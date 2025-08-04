package main

import (
	"context"
	"fmt"
	"time"

	"github.com/roderm/go-nanoleaf/pkg/scanner"
)

func main() {
	// search for 5 minutes

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	scanner := scanner.NewMdns()
	leafs, err := scanner.Scan(ctx)
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
		case <-ctx.Done():
			return
		}
	}
}
