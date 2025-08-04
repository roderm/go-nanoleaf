package main

import (
	"fmt"
	"net"

	"github.com/roderm/go-nanoleaf/example"
	"github.com/roderm/go-nanoleaf/pkg/device"
)

func main() {
	ipAddr, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", example.IP))
	if err != nil {
		panic(err)
	}
	d := device.NewDevice(device.WithIP(ipAddr), device.WithPort(example.Port))
	err = d.Authorization()
	if err != nil {
		err := fmt.Errorf("Authorization failed: make sure your device in pairing mode. \n%s", err.Error())
		panic(err)
	}
	fmt.Println("Your new auth-token is:", d.AuthorizationToken())
}
