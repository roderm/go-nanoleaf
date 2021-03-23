package scanner

import "github.com/roderm/go-nanoleaf"

type Scanner interface {
	Scan(chan<- *nanoleaf.Device) error
	Stop() error
}
