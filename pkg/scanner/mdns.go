package scanner

import (
	"strings"

	"github.com/ideasynthesis/mdns"
	"github.com/roderm/go-nanoleaf/pkg/device"
)

type mdnsScanner struct {
	services chan *mdns.ServiceEntry
}

func NewMdns() Scanner {
	scanner := new(mdnsScanner)
	return scanner
}

func (s *mdnsScanner) Scan(entries chan<- *device.Device) error {
	s.services = make(chan *mdns.ServiceEntry)
	go func(entries chan<- *device.Device, services <-chan *mdns.ServiceEntry) {
		defer close(entries)
		for s := range services {
			device := &device.Device{
				Name: s.Name,
				Network: &device.NetworkInterface{
					Port: s.Port,
					Host: s.Host,
					IPv4: s.AddrV4,
					IPv6: s.AddrV6,
				},
			}
			for _, v := range s.InfoFields {
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
			entries <- device
		}
	}(entries, s.services)
	return mdns.Lookup("_nanoleafapi._tcp", s.services)
}

func (s *mdnsScanner) Stop() error {
	close(s.services)
	return nil
}
