package nanoleaf

import (
	"net"
	"net/http"
)

type Option func(*Device)

func WithAuthKey(apiKey string) Option {
	return func(d *Device) {
		d.apiKey = apiKey
	}
}

func WithIP(ip net.IP) Option {
	return func(d *Device) {
		d.Network.IPv4 = []net.IP{ip}
	}
}

func WithPort(port int) Option {
	return func(d *Device) {
		d.Network.Port = port
	}
}

func WithClient(client *http.Client) Option {
	return func(d *Device) {
		d.client = client
	}
}
