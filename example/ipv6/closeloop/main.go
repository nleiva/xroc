/*
1. Configures a Streaming Telemetry subscription using an OpenConfig model template.
2. Configures the Peer link (Interface) using ygot.
3. Configures a BGP neighbor using ygot.
4. Subscribes to a Telemetry stream to learn about BGP neighbor status.

Libraries:
	xrgrpc -> https://nleiva.github.io/xrgrpc/
	ygot -> https://github.com/openconfig/ygot/
*/
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/golang/protobuf/proto"
	xr "github.com/nleiva/xrgrpc"
	"github.com/nleiva/xrgrpc/proto/telemetry"
	ocint "github.com/nleiva/xroc/pkg/interface"
	ocni "github.com/nleiva/xroc/pkg/network-instance"
	octel "github.com/nleiva/xroc/pkg/telemetry"
	"github.com/openconfig/ygot/ygot"
)

// Colors, just for fun.
const (
	blue   = "\x1b[34;1m"
	white  = "\x1b[0m"
	red    = "\x1b[31;1m"
	green  = "\x1b[32;1m"
	yellow = "\x1b[33;1m"
)

func main() {
	// Determine the ID for the first RPC call.
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := ra.Int63n(10000)

	////////////////////////////////////////////////////////////
	// Manually specify target parameters
	////////////////////////////////////////////////////////////
	router, err := xr.BuildRouter(
		xr.WithUsername("cisco"),
		xr.WithPassword("cisco"),
		xr.WithHost("[2001:420:2cff:1204::5502:2]:57344"),
		xr.WithCert("../../input/certificate/ems5502-2.pem"),
		xr.WithTimeout(45),
	)
	if err != nil {
		log.Fatalf("target parameters are incorrect: %s", err)
	}

	// Extract the IP address
	r, err := net.ResolveTCPAddr("tcp", router.Host)
	if err != nil {
		log.Fatalf("Incorrect IP address: %v", err)
	}

	// Connect to the router
	conn, ctx, err := xr.Connect(*router)
	if err != nil {
		log.Fatalf("could not setup a client connection to %s, %v", r.IP, err)
	}
	defer conn.Close()

	// Dealing with Cancellation
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	////////////////////////////////////////////////////////////
	// Generate the Telemetry config
	////////////////////////////////////////////////////////////
	//subscriptionID := "BGP-OC"
	subscriptionIDName := "BGP-OC"
	sensorGroupID := "BGPNeighbor-OC"
	path := "openconfig-network-instance:network-instances/network-instance/protocols/protocol/bgp/neighbors/neighbor/state"
	var sint uint64 = 1000

	// Device is our fake-root for "openconfig-telemetry:telemetry-system"
	dt := &octel.Device{}
	t := new(octel.TelemetrySystem)
	dt.TelemetrySystem = t

	//t := &octel.TelemetrySystem{}
	ygot.BuildEmptyTree(t)

	sg, err := t.NewSensorGroup(sensorGroupID)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", sensorGroupID, err)
	}
	ygot.BuildEmptyTree(sg)
	sg.SensorGroupId = &sensorGroupID

	ygot.BuildEmptyTree(sg)
	sp, err := sg.NewSensorPath(path)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", path, err)
	}
	ygot.BuildEmptyTree(sp)
	sp.Path = &path

	// Without Modified OpenConfig model
	// sb, err := t.Subscriptions.NewPersistentSubscription(subscriptionIDName)
	sb, err := t.Subscriptions.NewSubscription(subscriptionIDName)
	if err != nil {
		log.Fatalf("Failed to generate %vs: %v", subscriptionIDName, err)
	}
	ygot.BuildEmptyTree(sb)

	spf, err := sb.NewSensorProfile(sensorGroupID)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", sensorGroupID, err)
	}
	ygot.BuildEmptyTree(spf)
	spf.SensorGroup = &sensorGroupID
	spf.SensorGroup = ygot.String(sensorGroupID)
	spf.SampleInterval = ygot.Uint64(sint)

	// Validate the format. EmitJSON will do this anyways, here just for demo purposes.
	if err := dt.Validate(); err != nil {
		log.Fatalf("interface config validation failed: %v", err)
	}
	// Generate the json payload for our message
	json, err := ygot.EmitJSON(dt, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	if err != nil {
		log.Fatalf("JSON generation failed: %v", err)
	}

	////////////////////////////////////////////////////////////
	// Apply Telemetry config
	////////////////////////////////////////////////////////////

	_, err = xr.MergeConfig(ctx, conn, json, id)
	if err != nil {
		log.Fatalf("failed to config %s: %v\n", r.IP, err)
	} else {
		fmt.Printf("\n%sTelemetry%s config applied on %s (Request ID: %v)\n", blue, white, r.IP, id)
	}
	id++

	// Pause
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	////////////////////////////////////////////////////////////
	// Generate the Interface config
	////////////////////////////////////////////////////////////
	name := "HundredGigE0/0/0/16"

	// Device is our fake-root for "openconfig-interfaces:interfaces"
	di := &ocint.Device{}
	di.Interface = make(map[string]*ocint.Interface)

	di.Interface[name] = &ocint.Interface{
		Name:        ygot.String(name),
		Description: ygot.String("ygot test"),
		Enabled:     ygot.Bool(true),
		Type:        ocint.IETFInterfaces_InterfaceType_ethernetCsmacd,
		Mtu:         ygot.Uint16(9192),
		Ethernet:    &ocint.Interface_Ethernet{AutoNegotiate: ygot.Bool(false)},
	}

	s, err := di.Interface[name].NewSubinterface(0)
	if err != nil {
		log.Fatalf("Failed to generate sub-interface config for %s: %v", name, err)
	}

	ygot.BuildEmptyTree(s)
	// List of addresses to configure in this (sub)interface
	addresses := []struct {
		address string
		mask    uint8
	}{{
		address: "2001:db8:cafe::22",
		mask:    64,
	}}
	// Add the addresses to our Go generated struct
	for _, addr := range addresses {
		a, err := s.Ipv6.NewAddress(addr.address)
		if err != nil {
			panic(err)
		}
		a.PrefixLength = ygot.Uint8(addr.mask)
	}
	// Validate the format. EmitJSON will do this anyways, here just for demo purposes.
	if err := di.Validate(); err != nil {
		log.Fatalf("interface config validation failed: %v", err)
	}
	// Generate the json payload for our message
	json, err = ygot.EmitJSON(di, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	if err != nil {
		log.Fatalf("JSON generation failed: %v", err)
	}

	////////////////////////////////////////////////////////////
	// Apply Interface config
	////////////////////////////////////////////////////////////
	_, err = xr.MergeConfig(ctx, conn, json, id)
	if err != nil {
		log.Fatalf("failed to config %s: %v\n", r.IP, err)
	} else {
		fmt.Printf("\n%sInterface%s config applied on %s (Request ID: %v)\n", blue, white, r.IP, id)
	}
	id++

	// Pause
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	////////////////////////////////////////////////////////////
	// Generate the BGP config
	////////////////////////////////////////////////////////////
	rid := "203.0.113.22"
	nei := "2001:db8::11"
	desc := "iBGP session"
	bgpID := ocni.OpenconfigPolicyTypes_INSTALL_PROTOCOL_TYPE_BGP
	ipv6uniAF := ocni.OpenconfigBgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST
	var asn uint32 = 64512

	// Device is our fake-root for "network-instance:network-instance"
	db := &ocni.Device{}

	// Create a new BGP instance
	db.NewNetworkInstance("default")
	ni, err := db.NetworkInstance["default"].NewProtocol(bgpID, "default")
	if err != nil {
		log.Fatalf("Failed to generate a BGP instance %s: %v", "default", err)
	}

	// Specify Global BGP config
	ni.Bgp = &ocni.NetworkInstance_Protocol_Bgp{
		Global: &ocni.NetworkInstance_Protocol_Bgp_Global{
			As:       ygot.Uint32(asn),
			RouterId: ygot.String(rid),
		},
	}

	// Initialize the IPV6 Unicast address family.
	af6 := ipv6uniAF
	safi, err := ni.Bgp.Global.NewAfiSafi(af6)
	if err != nil {
		log.Fatalf("Failed to initialize the AF type %v: %v", af6, err)
	}
	safi.Enabled = ygot.Bool(true)
	ni.Bgp.Global.AfiSafi[ipv6uniAF] = safi

	// Configure a BGP Neighbor
	p, err := ni.Bgp.NewNeighbor(nei)
	if err != nil {
		log.Fatalf("Failed to create a new BGP neigbor IP: %s: %v", nei, err)
	}
	p.PeerAs = ygot.Uint32(asn)
	p.Description = ygot.String(desc)

	psafi, err := p.NewAfiSafi(af6)
	if err != nil {
		log.Fatalf("Failed to generate BGP AF type %v: %v", af6, err)
	}
	psafi.Enabled = ygot.Bool(true)
	p.AfiSafi[ipv6uniAF] = psafi

	// Validate the format. EmitJSON will do this anyways, here just for demo purposes.
	if err := db.Validate(); err != nil {
		log.Fatalf("interface config validation failed: %v", err)
	}
	// Generate the json payload for our message
	json, err = ygot.EmitJSON(db, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	if err != nil {
		log.Fatalf("JSON generation failed: %v", err)
	}

	////////////////////////////////////////////////////////////
	// Apply BGP config
	////////////////////////////////////////////////////////////

	_, err = xr.MergeConfig(ctx, conn, json, id)
	if err != nil {
		log.Fatalf("failed to config %s: %v\n", r.IP, err)
	} else {
		fmt.Printf("\n%sBGP%s config applied on %s (Request ID: %v)\n", blue, white, r.IP, id)
	}

	// Pause
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	////////////////////////////////////////////////////////////
	// Subscribe to Telemetry Stream
	////////////////////////////////////////////////////////////
	// Encoding GPBKV
	var e int64 = 3
	id++
	ch, ech, err := xr.GetSubscription(ctx, conn, subscriptionIDName, id, e)
	if err != nil {
		log.Fatalf("could not setup Telemetry Subscription: %v\n", err)
	}

	// Dealing with Cancellation (Telemetry Subscription)
	go func() {
		select {
		case <-c:
			fmt.Printf("\nmanually cancelled the session to %v\n\n", r.IP)
			cancel()
			return
		case <-ctx.Done():
			// Timeout: "context deadline exceeded"
			err = ctx.Err()
			fmt.Printf("\ngRPC session timed out after %v seconds: %v\n\n", router.Timeout, err.Error())
			return
		case err = <-ech:
			// Session canceled: "context canceled"
			fmt.Printf("\ngRPC session to %v failed: %v\n\n", r.IP, err.Error())
			return
		}
	}()
	fmt.Printf("\n4)\nReceiving %sTelemetry%s from %s ->\n\n", blue, white, r.IP)

	////////////////////////////////////////////////////////////
	// Process Telemetry messages
	////////////////////////////////////////////////////////////
	for tele := range ch {
		message := new(telemetry.Telemetry)
		err := proto.Unmarshal(tele, message)
		if err != nil {
			log.Fatalf("could not unmarshall the message: %v\n", err)
		}
		ok := false
		exploreFields(message.GetDataGpbkv(), "", nei, &ok)
	}
}

func exploreFields(f []*telemetry.TelemetryField, indent string, peer string, ok *bool) {
	for _, field := range f {
		switch field.GetFields() {
		case nil:
			decodeKV(field, indent, peer, ok)
		default:
			exploreFields(field.GetFields(), indent+" ", peer, ok)
		}
	}
}

func decodeKV(f *telemetry.TelemetryField, indent string, peer string, ok *bool) {
	// This is a very specific scenario, just for this example.
	color := white
	switch f.GetValueByType().(type) {
	case *telemetry.TelemetryField_StringValue:
		switch f.GetName() {
		case "neighbor-address":
			addr := f.GetStringValue()
			if addr == peer {
				*ok = true
			} else {
				*ok = false
			}
		case "session-state":
			if *ok {
				state := f.GetStringValue()
				switch state {
				case "ACTIVE", "IDLE":
					color = red
				case "OPENSENT", "CONNECT", "OPENCONFIRM":
					color = yellow
				case "ESTABLISHED":
					color = green
				default:
					color = white
				}
				t := time.Now()
				fmt.Printf("\rNeighbor: %s, Time: %v, State: %s%s%s     ", peer, t.Format("15:04:05"), color, state, white)
				if state == "ESTABLISHED" {
					fmt.Printf("\n\n\n                        Session \u2705 \n\n\n")
					os.Exit(0)
				}
			}
		default:
		}
	default:
	}
}
