package main

import "C"
import (
	"github.com/roderm/go-nanoleaf"
	"github.com/roderm/go-nanoleaf/scanner"
)

//export Hello
func Hello() *C.char {
	return C.CString("Hello world!")
}

var nanoScan scanner.Scanner
var foundDevices chan *nanoleaf.Device

//export ScannerStart
func ScannerStart() {
	nanoScan = scanner.NewMdns()
	foundDevices = make(chan *nanoleaf.Device)
	go nanoScan.Scan(foundDevices)
}

//export ScannerFound
func ScannerFound() *C.char {
	dev := <-foundDevices
	return C.CString(dev.Name)
}

//export ScannerStop
func ScannerStop() {
	nanoScan.Stop()
}

// // export Device
// func Device(ip string, port int, authKey string) (*nanoleaf.Device, error) {
// 	ipAddr, _, err := net.ParseCIDR(ip)
// 	if err != nil {
// 		return nil, fmt.Errorf("Invalid IP")
// 	}
// 	if port == 0 {
// 		port = 16021
// 	}
// 	opts := []nanoleaf.Option{
// 		nanoleaf.WithIP(ipAddr),
// 		nanoleaf.WithPort(port),
// 	}
// 	if authKey != "" {
// 		opts = append(opts, nanoleaf.WithAuthKey(authKey))
// 	}
// 	return nanoleaf.NewDevice(opts...), nil
// }

func main() {}
