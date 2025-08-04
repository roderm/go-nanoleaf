package main

import (
	"context"
	"fmt"
	"net"

	"github.com/roderm/go-nanoleaf/example"
	"github.com/roderm/go-nanoleaf/pkg/device"
	"github.com/roderm/go-nanoleaf/pkg/stream"
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

	ctrl, err := stream.NewControll(d)
	if err != nil {
		panic(err)
	}
	ctrl.Stream(context.TODO())

	// ctrl.SetPanel(types.Color{
	// 	Red:   255,
	// 	White: 0,
	// }, 1, d.Info.PanelLayout.Layout.Panels...)
	// time.Sleep(time.Second)
	// for i := 0; i < 6; i++ {
	// 	for _, p := range d.Info.PanelLayout.Layout.Panels {
	// 		ctrl.SetPanel(types.Color{
	// 			Blue:  uint8(((i + 1) % 2) * 255),
	// 			Green: uint8((i % 2) * 255),
	// 		}, 0,
	// 			p)
	// 		time.Sleep(time.Millisecond * 200)
	// 		ctrl.SetPanel(types.Color{}, 1, d.Info.PanelLayout.Layout.Panels...)
	// 	}
	// }
	// ctrl.SetPanel(types.Color{
	// 	Red: 255,
	// }, 1, d.Info.PanelLayout.Layout.Panels...)
}
