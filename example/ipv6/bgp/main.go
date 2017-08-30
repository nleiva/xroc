/*
1. Configures a Streaming Telemetry subscription using an OpenConfig model template.
2. Configures the Peer link (Interface) using ygot.
3. Configures a BGP neighbor using ygot.
4. Subscribes to the Telemetry stream to learn about BGP neighbor status.

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
	ocbgp "github.com/nleiva/xroc/pkg/bgp"
	ocint "github.com/nleiva/xroc/pkg/interface"
	octele "github.com/nleiva/xroc/pkg/telemetry"
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
	// Determine the ID for first the transaction.
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := ra.Int63n(10000)

	////////////////////////////////////////////////////////////
	// Manually specify target parameters
	////////////////////////////////////////////////////////////
	router, err := xr.BuildRouter(
		xr.WithUsername("cisco"),
		xr.WithPassword("cisco"),
		xr.WithHost("[2001:420:2cff:1204::5502:1]:57344"),
		xr.WithCert("../../input/certificate/ems5502-1.pem"),
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
	subscriptionID := "BGP-OC"
	sensorGroupID := "BGPNeighbor-OC"
	path := "openconfig-bgp:bgp/neighbors/neighbor/state"
	var sint uint64 = 1000

	t := &octele.OpenconfigTelemetry_TelemetrySystem{}
	ygot.BuildEmptyTree(t)
	sg, err := t.SensorGroups.NewSensorGroup(sensorGroupID)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", sensorGroupID, err)
	}
	ygot.BuildEmptyTree(sg)
	sg.Config.SensorGroupId = &sensorGroupID

	ygot.BuildEmptyTree(sg)
	sp, err := sg.SensorPaths.NewSensorPath(path)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", path, err)
	}
	ygot.BuildEmptyTree(sp)
	sp.Config.Path = &path

	sb, err := t.Subscriptions.Persistent.NewSubscription(subscriptionID)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", subscriptionID, err)
	}
	ygot.BuildEmptyTree(sb)
	sb.Config.SubscriptionId = &subscriptionID
	spf, err := sb.SensorProfiles.NewSensorProfile(sensorGroupID)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", sensorGroupID, err)
	}
	ygot.BuildEmptyTree(spf)
	spf.SensorGroup = &sensorGroupID
	spf.Config.SensorGroup = ygot.String(sensorGroupID)
	spf.Config.SampleInterval = ygot.Uint64(sint)

	// EmitJSON will do this anyways, here just for demo purposes
	if err := t.Validate(); err != nil {
		log.Fatalf("telemetry config validation failed: %v", err)
	}
	json, err := ygot.EmitJSON(t, &ygot.EmitJSONConfig{
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
	// Apply Telemetry Config
	////////////////////////////////////////////////////////////
	_, err = xr.MergeConfig(ctx, conn, json, id)
	if err != nil {
		log.Fatalf("failed to config %s: %v\n", r.IP, err)
	} else {
		fmt.Printf("\n1)\n%sTelemetry%s config applied on %s (Request ID: %v)\n", blue, white, r.IP, id)
	}
	id++

	// Pause
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	////////////////////////////////////////////////////////////
	// Generate the Interface config
	////////////////////////////////////////////////////////////
	intf := "HundredGigE0/0/0/1"
	idesc := "Peer Link"
	var mtu uint16 = 9192

	d := &ocint.Device{}
	i, err := d.NewInterface(intf)
	i.Name = &intf
	i.Type = ocint.IETFInterfaces_InterfaceType_ethernetCsmacd

	i.Mtu = &mtu
	i.Description = &idesc
	i.Enabled = ygot.Bool(true)
	i.Ethernet = &ocint.Interface_Ethernet{AutoNegotiate: ygot.Bool(false)}

	s, err := d.Interface[intf].NewSubinterface(0)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", intf, err)
	}
	ygot.BuildEmptyTree(s)
	addresses := []struct {
		address string
		mask    uint8
	}{{
		address: "2001:db8:cafe::1",
		mask:    64,
	}}

	for _, addr := range addresses {
		a, err := s.Ipv6.NewAddress(addr.address)
		if err != nil {
			panic(err)
		}
		a.PrefixLength = ygot.Uint8(addr.mask)
	}
	// EmitJSON will do this anyways, here just for demo purposes
	if err := d.Validate(); err != nil {
		log.Fatalf("interface config validation failed: %v", err)
	}
	json, err = ygot.EmitJSON(d, &ygot.EmitJSONConfig{
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
		fmt.Printf("\n2)\n%sInterface%s config applied on %s (Request ID: %v)\n", blue, white, r.IP, id)
	}
	id++

	// Pause
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	////////////////////////////////////////////////////////////
	// Generate the BGP config
	////////////////////////////////////////////////////////////
	rid := "162.151.250.1"
	nei := "2001:db8:cafe::2"
	desc := "iBGP session"
	var asn uint32 = 64512

	b := &ocbgp.Bgp{
		Global: &ocbgp.Bgp_Global{
			As:       ygot.Uint32(asn),
			RouterId: ygot.String(rid),
		},
	}

	af6 := ocbgp.OpenconfigBgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST
	safi, err := b.Global.NewAfiSafi(af6)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", "BGP Config", err)
	}
	safi.Enabled = ygot.Bool(true)
	b.Global.AfiSafi[ocbgp.OpenconfigBgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST] = safi
	p, err := b.NewNeighbor(nei)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", nei, err)
	}
	p.PeerAs = ygot.Uint32(asn)
	p.Description = ygot.String(desc)

	psafi, err := p.NewAfiSafi(af6)
	// err
	psafi.Enabled = ygot.Bool(true)
	p.AfiSafi[ocbgp.OpenconfigBgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST] = psafi

	// EmitJSON will do this anyways, here just for demo purposes
	if err := d.Validate(); err != nil {
		log.Fatalf("interface config validation failed: %v", err)
	}
	json, err = ygot.EmitJSON(b, &ygot.EmitJSONConfig{
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
		fmt.Printf("\n3)\n%sBGP%s Config applied on %s (Request ID: %v)\n", blue, white, r.IP, id)
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
	ch, ech, err := xr.GetSubscription(ctx, conn, subscriptionID, id, e)
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
		case "connection-state":
			if *ok {
				state := f.GetStringValue()
				switch state {
				case "bgp-st-active", "bgp-st-idle":
					color = red
				case "bgp-st-opensent", "bgp-st-connect", "bgp-st-openconfirm":
					color = yellow
				case "bgp-st-estab":
					color = green
				default:
					color = white
				}
				t := time.Now()
				fmt.Printf("\rNeighbor: %s, Time: %v, State: %s%s%s     ", peer, t.Format("15:04:05"), color, state, white)
				if state == "bgp-st-estab" {
					fmt.Printf("\n\n\n                        Session \u2705 \n\n\n")
					os.Exit(0)
				}
			}
		default:
		}
	default:
	}
}
