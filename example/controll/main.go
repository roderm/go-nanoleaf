package main

import (
	"fmt"
	"net"
	"time"

	"github.com/roderm/go-nanoleaf"
	"github.com/roderm/go-nanoleaf/example"
)

func main() {
	ipAddr, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", example.IP))
	if err != nil {
		panic(err)
	}
	d := nanoleaf.NewDevice(
		nanoleaf.WithIP(ipAddr),
		nanoleaf.WithPort(example.Port),
		nanoleaf.WithAuthKey(example.AuthKey),
	)
	// Get initial data
	_, err = d.Get()
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
