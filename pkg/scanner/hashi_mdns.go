package scanner

import (
	"context"
	"net"
	"strings"

	"github.com/hashicorp/mdns"
	"github.com/roderm/go-nanoleaf/pkg/device"
)

type mdnsHashiScanner bool

func NewHashiMdns() Scanner {
	return new(mdnsHashiScanner)
}

func (s *mdnsHashiScanner) toDevice(se *mdns.ServiceEntry) *device.Device {
	device := &device.Device{
		Name: se.Name,
		Network: &device.NetworkInterface{
			Port: se.Port,
			Host: se.Host,
			IPv4: []net.IP{se.AddrV4},
			IPv6: []net.IP{se.AddrV6},
		},
	}
	for _, v := range se.InfoFields {
		value := strings.Split(v, "=")
		if len(value) == 0 {
			continue
		}
		switch value[0] {
		case "id":
			device.Id = strings.Join(value[1:], "=")
		case "md":
			device.Type = strings.Join(value[1:], "=")
		case "srcvers":
			device.FirmwareVersion = strings.Join(value[1:], "=")
		}
	}
	return device
}
func (s *mdnsHashiScanner) Scan(ctx context.Context) (<-chan *device.Device, error) {
	ses := make(chan *mdns.ServiceEntry)
	devs := make(chan *device.Device)
	go func(ctx context.Context, ses chan *mdns.ServiceEntry, devs chan<- *device.Device) {
		for {
			select {
			case <-ctx.Done():
				close(ses)
				close(devs)
				return
			case se := <-ses:
				devs <- s.toDevice(se)
			}
		}
	}(ctx, ses, devs)
	err := mdns.Lookup("_nanoleafapi._tcp", ses)
	return devs, err
}
