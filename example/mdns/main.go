package main

import (
	"fmt"
	"time"

	"github.com/roderm/go-nanoleaf/pkg/device"
	"github.com/roderm/go-nanoleaf/pkg/scanner"
)

func main() {
	// search for 5 minutes
	deadline := time.NewTimer(time.Minute * 5)

	leafs := make(chan *device.Device)
	scanner := scanner.NewMdns()
	err := scanner.Scan(leafs)
	fmt.Println("scanner startet...")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case nl, ok := <-leafs:
			if !ok {
				fmt.Println("Channel already closed")
				return
			}
			fmt.Printf("Found a nanoleaf %s (%s): \n\t IPs: %v\n", nl.Name, nl.Id, nl.Network.IPv4)
		case <-deadline.C:
			return
		}
	}
}
