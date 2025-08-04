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
		panic(fmt.Errorf("Invalid IP-Address: %v", err))
	}
	d := device.NewDevice(
		device.WithIP(ipAddr),
		device.WithPort(example.Port),
		device.WithAuthKey(example.AuthKey),
	)
	d.On()
	d.SetBrightness(100, 10)
	effects, err := d.GetEffects()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got effects on Device:\n %v", effects)

	err = d.SelectEffect("Be Productive")
	if err != nil {
		panic(err)
	}
}
