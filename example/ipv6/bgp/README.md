# Configuring BGP

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

## Generate Go structs

I'm running the following from `$GOPATH/src/github.com/nleiva/xroc/pkg/network-instance`

```console
$ $GOPATH/src/github.com/openconfig/ygot/generator/generator -path ../../yang/ -exclude_modules=ietf-interfaces -compress_paths -package_name=ocni -generate_fakeroot -fakeroot_name=device -output_file=ocni.go ../../yang/network-instance/*.yang
```

## Running it

```console
$ go run main.go

BGP config applied on 2001:420:2cff:1204::5502:2 (Request ID: 6605)
```

## JSON payload generated

```json
{
  "openconfig-network-instance:network-instances": {
    "network-instance": [
      {
        "config": {
          "name": "default"
        },
        "name": "default",
        "protocols": {
          "protocol": [
            {
              "bgp": {
                "global": {
                  "afi-safis": {
                    "afi-safi": [
                      {
                        "afi-safi-name": "openconfig-bgp-types:IPV6_UNICAST",
                        "config": {
                          "afi-safi-name": "openconfig-bgp-types:IPV6_UNICAST",
                          "enabled": true
                        }
                      }
                    ]
                  },
                  "config": {
                    "as": 64512,
                    "router-id": "203.0.113.22"
                  }
                },
                "neighbors": {
                  "neighbor": [
                    {
                      "afi-safis": {
                        "afi-safi": [
                          {
                            "afi-safi-name": "openconfig-bgp-types:IPV6_UNICAST",
                            "config": {
                              "afi-safi-name": "openconfig-bgp-types:IPV6_UNICAST",
                              "enabled": true
                            }
                          }
                        ]
                      },
                      "config": {
                        "description": "iBGP session",
                        "neighbor-address": "2001:db8::11",
                        "peer-as": 64512
                      },
                      "neighbor-address": "2001:db8::11"
                    }
                  ]
                }
              },
              "config": {
                "identifier": "openconfig-policy-types:BGP",
                "name": "default"
              },
              "identifier": "openconfig-policy-types:BGP",
              "name": "default"
            }
          ]
        }
      }
    ]
  }
}
```

## Final result on the router (target)

```console
RP/0/RP0/CPU0:mrstn-5502-2.cisco.com#sh run router bgp
router bgp 64512
 bgp router-id 203.0.113.22
 address-family ipv6 unicast
 !
 neighbor 2001:db8::11
  remote-as 64512
  description iBGP session
  address-family ipv6 unicast
  !
 !
!
```