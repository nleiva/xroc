/*
Configures an interface on an IOS XR device using ygot.

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
	ocint "github.com/nleiva/xroc/pkg/interface"
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
	// Generate the Interface config
	////////////////////////////////////////////////////////////
	name := "HundredGigE0/0/0/16"

	// Device is our fake-root for "openconfig-interfaces:interfaces"
	d := &ocint.Device{}
	d.Interface = make(map[string]*ocint.Interface)

	d.Interface[name] = &ocint.Interface{
		Name:        ygot.String(name),
		Description: ygot.String("ygot test"),
		Enabled:     ygot.Bool(true),
		Type:        ocint.IETFInterfaces_InterfaceType_ethernetCsmacd,
		Mtu:         ygot.Uint16(9192),
		Ethernet:    &ocint.Interface_Ethernet{AutoNegotiate: ygot.Bool(false)},
	}

	s, err := d.Interface[name].NewSubinterface(0)
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
	// Apply Interface config
	////////////////////////////////////////////////////////////
	_, err = xr.MergeConfig(ctx, conn, json, id)
	if err != nil {
		log.Fatalf("failed to config %s: %v\n", r.IP, err)
	} else {
		fmt.Printf("\n%sInterface%s config applied on %s (Request ID: %v)\n", blue, white, r.IP, id)
	}
}
