package scanner

import (
	"github.com/roderm/go-nanoleaf/pkg/device"
)

type Scanner interface {
	Scan(chan<- *device.Device) error
	Stop() error
}
