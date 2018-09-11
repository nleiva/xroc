/*
Configures a Telemetry Subscription on an IOS XR device using ygot.

Libraries:
	xrgrpc -> https://nleiva.github.io/xrgrpc/
	ygot -> https://github.com/openconfig/ygot/
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	xr "github.com/nleiva/xrgrpc"
	ocni "github.com/nleiva/xroc/pkg/network-instance"
	"github.com/openconfig/ygot/ygot"
)

// Colors, just for fun.
const (
	blue  = "\x1b[34;1m"
	white = "\x1b[0m"
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
	d := &ocni.Device{}

	// Create a new BGP instance
	d.NewNetworkInstance("default")
	ni, err := d.NetworkInstance["default"].NewProtocol(bgpID, "default")
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
	if err := d.Validate(); err != nil {
		log.Fatalf("interface config validation failed: %v", err)
	}
	// Generate the json payload for our message
	json, err := ygot.EmitJSON(d, &ygot.EmitJSONConfig{
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

}
