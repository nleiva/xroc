# Configuring an interface

## Input

The script needs the username and password of the target device, along with the IP address and port number gRPC is listening to. To secure the connection you also need to provide a TLS certificate ([instructions to get it](https://github.com/nleiva/xrgrpc/blob/master/README.md#certificate-file)).

```go
router, err := xr.BuildRouter(
  xr.WithUsername("cisco"),
  xr.WithPassword("cisco"),
  xr.WithHost("[2001:420:2cff:1204::5502:2]:57344"),
  xr.WithCert("../../input/certificate/ems5502-2.pem"),
  xr.WithTimeout(45),
)
```

## A word about OC Telemetry model

IOS XR 6.5.1 (used for this example) supports openconfig-telemetry v0.4.1. On the other hand, OpenConfig just released v0.5.0.

You can see the changes made to adecuate to this by running `diff` or `git diff openconfig-telemetry.yang openconfig-telemetry.modified.yang` in `$GOPATH/src/github.com/nleiva/xroc/yang/telemetry`.

## Generate Go structs

I'm running the following from `$GOPATH/src/github.com/nleiva/xroc/pkg/telemetry`

```console
$ $GOPATH/src/github.com/openconfig/ygot/generator/generator -path ../../yang/ -compress_paths -package_name=octel -generate_fakeroot -fakeroot_name=device -output_file=octel.go ../../yang/telemetry/openconfig-telemetry.modified.yang openconfig-telemetry-types.yang
```

## Running it

```console
$ go run main.go

Telemetry config applied on 2001:420:2cff:1204::5502:2 (Request ID: 7319)
```

## JSON payload generated

```json
{
  "openconfig-telemetry:telemetry-system": {
    "sensor-groups": {
      "sensor-group": [
        {
          "config": {
            "sensor-group-id": "BGPNeighbor-OC"
          },
          "sensor-group-id": "BGPNeighbor-OC",
          "sensor-paths": {
            "sensor-path": [
              {
                "config": {
                  "path": "openconfig-network-instance:network-instances/network-instance/protocols/protocol/bgp/neighbors/neighbor/state"
                },
                "path": "openconfig-network-instance:network-instances/network-instance/protocols/protocol/bgp/neighbors/neighbor/state"
              }
            ]
          }
        }
      ]
    },
    "subscriptions": {
      "persistent": {
        "subscription": [
          {
            "config": {
              "subscription-id": "BGP-OC"
            },
            "sensor-profiles": {
              "sensor-profile": [
                {
                  "config": {
                    "sample-interval": "1000",
                    "sensor-group": "BGPNeighbor-OC"
                  },
                  "sensor-group": "BGPNeighbor-OC"
                }
              ]
            },
            "subscription-id": "BGP-OC"
          }
        ]
      }
    }
  }
}
```

## Final result on the router (target)

```console
RP/0/RP0/CPU0:mrstn-5502-2.cisco.com#show run telemetry model-driven
telemetry model-driven
 sensor-group BGPNeighbor-OC
  sensor-path openconfig-network-instance:network-instances/network-instance/protocols/protocol/bgp/neighbors/neighbor/state
 !
 subscription BGP-OC
  sensor-group-id BGPNeighbor-OC sample-interval 1000
 !
!
```