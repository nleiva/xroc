# A collection of OpenConfig examples with Cisco IOS XR

[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

The goal of this repo is to provide a set of examples that illustrate how to program network elements using OpenConfig and gRPC. We will leverage different Open Source projects:

- [Go](https://github.com/golang/go) as the main programming language.
- [xrgrpc](https://github.com/nleiva/xrgrpc) (gRPC library for Cisco IOS XR) to interact with IOS XR via [gRPC](https://grpc.io/).
- [ygot](https://github.com/openconfig/ygot) (**Y**ANG **Go** **T**ools) to generate the [OpenConfig](http://www.openconfig.net/) data required for the transactions.

## Examples

### Configuring and Validating BGP Peer session

The objective is to configure an interface (link) and bring up a BGP session, while setting up a BGP Neighbor Telemetry stream to track the current state of the BGP Finite State Machine. Check the code [here](example/ipv6/bgp/main.go)

__Steps__
1. Configure a Streaming Telemetry subscription using using [ygot](https://github.com/openconfig/ygot) and [xrgrpc](https://github.com/nleiva/xrgrpc).
2. Configure the Peer link (Interface) using using using [ygot](https://github.com/openconfig/ygot) and [xrgrpc](https://github.com/nleiva/xrgrpc).
3. Configure a BGP neighbor using using using [ygot](https://github.com/openconfig/ygot) and [xrgrpc](https://github.com/nleiva/xrgrpc).
4. Subscribe to a Telemetry stream to learn about BGP Neighbor status with [xrgrpc](https://github.com/nleiva/xrgrpc).

  ![oc-config-validate](https://github.com/nleiva/xroc/blob/gh-pages/ygot-bgp3.gif)

## Tutorials
- [Programming IOS-XR with gRPC and Go](https://xrdocs.github.io/programmability/tutorials/2017-08-04-programming-ios-xr-with-grpc-and-go/).
- [Validate the intent of network config changes](https://xrdocs.github.io/programmability/tutorials/2017-08-14-validate-the-intent-of-network-config-changes/).
