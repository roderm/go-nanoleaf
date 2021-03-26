package main

import (
	"fmt"
	"net"
	"time"

	"github.com/roderm/go-nanoleaf"
	"github.com/roderm/go-nanoleaf/example"
	"github.com/roderm/go-nanoleaf/stream"
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
	d.Get()

	ctrl, err := stream.NewControll(d)
	if err != nil {
		panic(err)
	}
	ctrl.SetPanel(nanoleaf.Color{
		Red:   255,
		White: 0,
	}, 1, d.Info.PanelLayout.Layout.Panels...)
	time.Sleep(time.Second)
	for i := 0; i < 6; i++ {
		for _, p := range d.Info.PanelLayout.Layout.Panels {
			ctrl.SetPanel(nanoleaf.Color{
				Blue:  uint8(((i + 1) % 2) * 255),
				Green: uint8((i % 2) * 255),
			}, 0,
				p)
			time.Sleep(time.Millisecond * 200)
			ctrl.SetPanel(nanoleaf.Color{}, 1, d.Info.PanelLayout.Layout.Panels...)
		}
	}
	ctrl.SetPanel(nanoleaf.Color{
		Red: 255,
	}, 1, d.Info.PanelLayout.Layout.Panels...)
}
