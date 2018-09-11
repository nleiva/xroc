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
	octel "github.com/nleiva/xroc/pkg/telemetry"
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
	// Generate the Telemetry config
	////////////////////////////////////////////////////////////
	//subscriptionID := "BGP-OC"
	subscriptionIDName := "BGP-OC"
	// var subscriptionID uint64 = 1111
	sensorGroupID := "BGPNeighbor-OC"
	path := "openconfig-network-instance:network-instances/network-instance/protocols/protocol/bgp/neighbors/neighbor/state"
	var sint uint64 = 1000

	// Device is our fake-root for "openconfig-telemetry:telemetry-system"
	d := &octel.Device{}
	t := new(octel.TelemetrySystem)
	d.TelemetrySystem = t

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
	// used as string before. Is it a uint64 now?. Don't think so.

	// The follwoing would setup:
	// "state": { "id": "subscriptionID"}
	// sb.Id = &subscriptionID

	spf, err := sb.NewSensorProfile(sensorGroupID)
	if err != nil {
		log.Fatalf("Failed to generate %s: %v", sensorGroupID, err)
	}
	ygot.BuildEmptyTree(spf)
	spf.SensorGroup = &sensorGroupID
	spf.SensorGroup = ygot.String(sensorGroupID)
	spf.SampleInterval = ygot.Uint64(sint)

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
	// Apply Telemetry config
	////////////////////////////////////////////////////////////

	_, err = xr.MergeConfig(ctx, conn, json, id)
	if err != nil {
		log.Fatalf("failed to config %s: %v\n", r.IP, err)
	} else {
		fmt.Printf("\n%sTelemetry%s config applied on %s (Request ID: %v)\n", blue, white, r.IP, id)
	}

}
