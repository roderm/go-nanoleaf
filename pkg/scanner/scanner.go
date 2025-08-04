package scanner

import (
	"context"

	"github.com/roderm/go-nanoleaf/pkg/device"
)

type Scanner interface {
	Scan(context.Context) (<-chan *device.Device, error)
}
