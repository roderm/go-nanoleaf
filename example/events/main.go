package main

import (
	"fmt"
	"net"

	"github.com/roderm/go-nanoleaf/example"
	"github.com/roderm/go-nanoleaf/pkg/device"
	"github.com/roderm/go-nanoleaf/pkg/event"
)

func main() {
	ipAddr, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", example.IP))
	if err != nil {
		panic(fmt.Errorf("invalid IP-Address: %v", err))
	}
	d := device.NewDevice(
		device.WithIP(ipAddr),
		device.WithPort(example.Port),
		device.WithAuthKey(example.AuthKey),
	)

	listener := event.NewListener(d)

	states, cancelStates, err := listener.States()
	if err != nil {
		panic(err)
	}
	defer cancelStates()
	touches, cancelTouch, err := listener.Touches()
	if err != nil {
		panic(err)
	}
	defer cancelTouch()
	fmt.Println("Watching changes")
	for {
		select {
		case state := <-states:
			fmt.Println("Has new state", state)
		case touch := <-touches:
			fmt.Println("Got Touched", touch)
		}
	}
}
