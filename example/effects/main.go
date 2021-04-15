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
		panic(fmt.Errorf("Invalid IP-Address: %v", err))
	}
	d := nanoleaf.NewDevice(
		nanoleaf.WithIP(ipAddr),
		nanoleaf.WithPort(example.Port),
		nanoleaf.WithAuthKey(example.AuthKey),
	)
	d.On()
	effects, err := d.GetEffects()
	if err != nil {
		panic(err)
	}
	fmt.Println("Got effects on Device:", effects)

	err = d.SelectEffect("Starlight")
	if err != nil {
		panic(err)
	}
}
