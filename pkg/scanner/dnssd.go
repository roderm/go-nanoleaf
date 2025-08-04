package scanner

import (
	"context"
	"fmt"

	"github.com/brutella/dnssd"
	"github.com/roderm/go-nanoleaf/pkg/device"
)

type dnssdScanner bool

func NewDnssd() Scanner {
	return new(dnssdScanner)
}

func (s *dnssdScanner) Scan(ctx context.Context) (<-chan *device.Device, error) {
	ch := make(chan *device.Device)
	go func(ctx context.Context) {
		dnssd.LookupType(ctx, "_nanoleafapi._tcp", func(be dnssd.BrowseEntry) {
			ch <- s.deviceFromBrowserEntry(be)
		}, func(be dnssd.BrowseEntry) {})
		close(ch)
	}(ctx)
	return ch, nil
}

func (s *dnssdScanner) deviceFromBrowserEntry(be dnssd.BrowseEntry) *device.Device {
	dev := &device.Device{
		Network: &device.NetworkInterface{
			Host: be.Host,
			Port: be.Port,
			IPv4: be.IPs,
		},
		// Id: string(be.IPs]),
	}
	for n, txt := range be.Text {
		fmt.Println(n, txt)
	}
	return dev
}
