package main

import (
	"fmt"
	"net"

	"github.com/roderm/go-nanoleaf"
	"github.com/roderm/go-nanoleaf/example"
)

func main() {
	ipAddr, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", example.IP))
	if err != nil {
		panic(err)
	}
	d := nanoleaf.NewDevice(nanoleaf.WithIP(ipAddr), nanoleaf.WithPort(example.Port))
	err = d.Authorization()
	if err != nil {
		err := fmt.Errorf("Authorization failed: make sure your device in pairing mode. \n%s", err.Error())
		panic(err)
	}
	fmt.Println("Your new auth-token is:", d.AuthorizationToken())
}
