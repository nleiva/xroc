# A collection of OpenConfig and Cisco IOS XR examples

[![BSD License](https://img.shields.io/pypi/l/Django.svg)](LICENSE)
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/nleiva/xroc)

The goal of this repo is to provide a set of examples that illustrate how to program network elements using [OpenConfig](http://www.openconfig.net/) and [gRPC](https://grpc.io/). We will leverage different Open Source projects:

- [Go](https://github.com/golang/go) as the main programming language.
- [xrgrpc](https://nleiva.github.io/xrgrpc/) (gRPC library for Cisco IOS XR) to interact with IOS XR via [gRPC](https://grpc.io/).
- [ygot](https://github.com/openconfig/ygot) (**Y**ANG **Go** **T**ools) to generate the [OpenConfig](http://www.openconfig.net/) data structures required for the gRPC calls.

## Examples

### Interface Configuration

- [OpenConfig interface](example/ipv6/interface/README.md)

### Telemetry Stream Configuration

- [OpenConfig telemetry](example/ipv6/telemetry/README.md)

### BGP Neighbor Configuration

- [OpenConfig Network Instance: BGP](example/ipv6/bgp/README.md)

### Configuring and Validating a BGP Peer session

The objective is to configure an interface (link) and bring up a BGP session, while setting up a BGP Neighbor Telemetry stream to track the current state of the BGP Finite State Machine. Check out the code [here](example/ipv6/closeloop/main.go)

#### Steps

1. Configure a Streaming Telemetry subscription using [ygot](https://github.com/openconfig/ygot) and [xrgrpc](https://nleiva.github.io/xrgrpc/).

2. Configure the Peer link (interface) using [ygot](https://github.com/openconfig/ygot) and [xrgrpc](https://nleiva.github.io/xrgrpc/).

3. Configure a BGP neighbor using [ygot](https://github.com/openconfig/ygot) and [xrgrpc](https://nleiva.github.io/xrgrpc/).

4. Subscribe to a Telemetry stream to learn about BGP Neighbor status with [xrgrpc](https://nleiva.github.io/xrgrpc/).

   ![oc-config-validate](static/images/closeloop.svg)

## Tutorials

- [Programming IOS-XR with gRPC and Go](https://xrdocs.github.io/programmability/tutorials/2017-08-04-programming-ios-xr-with-grpc-and-go/).
- [Validate the intent of network config changes](https://xrdocs.github.io/programmability/tutorials/2017-08-14-validate-the-intent-of-network-config-changes/).
- [Programming Network Devices with gRPC and OpenConfig](https://goo.gl/dMZxd2)
