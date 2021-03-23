package main

import (
	"fmt"

	"github.com/roderm/go-nanoleaf"
	"github.com/roderm/go-nanoleaf/example"
)

func main() {
	d := nanoleaf.NewDevice(
		nanoleaf.WithIP(example.GetIp()),
		nanoleaf.WithPort(example.Port),
		nanoleaf.WithAuthKey(example.AuthKey),
	)

	listener := nanoleaf.NewListener(d)

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
