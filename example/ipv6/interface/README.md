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

## Generate Go structs

I'm running the following from `$GOPATH/src/github.com/nleiva/xroc/pkg/interface`

```console
$GOPATH/src/github.com/openconfig/ygot/generator/generator -path ../../yang/ -exclude_modules=ietf-interfaces -compress_paths -package_name=ocint -generate_fakeroot -fakeroot_name=device -output_file=ocint.go ../../yang/interfaces/*.yang
```

## Running it

```console
$ go run main.go

Interface config applied on 2001:420:2cff:1204::5502:2 (Request ID: 8974)
```

## JSON payload generated

```json
{
  "openconfig-interfaces:interfaces": {
    "interface": [
      {
        "config": {
          "description": "ygot test",
          "enabled": true,
          "mtu": 9192,
          "name": "HundredGigE0/0/0/16",
          "type": "iana-if-type:ethernetCsmacd"
        },
        "name": "HundredGigE0/0/0/16",
        "openconfig-if-ethernet:ethernet": {
          "config": {
            "auto-negotiate": false
          }
        },
        "subinterfaces": {
          "subinterface": [
            {
              "config": {
                "index": 0
              },
              "index": 0,
              "openconfig-if-ip:ipv6": {
                "addresses": {
                  "address": [
                    {
                      "config": {
                        "ip": "2001:db8:cafe::22",
                        "prefix-length": 64
                      },
                      "ip": "2001:db8:cafe::22"
                    }
                  ]
                }
              }
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
RP/0/RP0/CPU0:mrstn-5502-2.cisco.com# sh run int HundredGigE0/0/0/16
interface HundredGigE0/0/0/16
 description ygot test
 mtu 9192
 ipv6 address 2001:db8:cafe::22/64
!
```