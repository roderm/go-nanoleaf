package main

import (
	"fmt"
	"net"
	"time"

	"github.com/roderm/go-nanoleaf/example"
	"github.com/roderm/go-nanoleaf/pkg/device"
)

func main() {
	ipAddr, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", example.IP))
	if err != nil {
		panic(err)
	}
	d := device.NewDevice(
		device.WithIP(ipAddr),
		device.WithPort(example.Port),
		device.WithAuthKey(example.AuthKey),
	)

	if err != nil {
		panic(err)
	}
	if err := d.On(); err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		d.SetBrightness(95, 0)
		time.Sleep(time.Second)
		d.SetBrightness(5, 0)
		time.Sleep(time.Second)
	}
	d.Off()
}
