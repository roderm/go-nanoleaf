package main

import "C"
import (
	"fmt"
	"net"

	"github.com/roderm/go-nanoleaf"
)

// export Device
func Device(ip string, port int, authKey string) (*nanoleaf.Device, error) {
	ipAddr, _, err := net.ParseCIDR(ip)
	if err != nil {
		return nil, fmt.Errorf("Invalid IP")
	}
	if port == 0 {
		port = 16021
	}
	opts := []nanoleaf.Option{
		nanoleaf.WithIP(ipAddr),
		nanoleaf.WithPort(port),
	}
	if authKey != "" {
		opts = append(opts, nanoleaf.WithAuthKey(authKey))
	}
	return nanoleaf.NewDevice(opts...), nil
}
func main() {}
