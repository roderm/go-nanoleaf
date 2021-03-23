package nanoleaf

import (
	"encoding/binary"
	"fmt"
)

const (
	SHAPETYPE_TRIANGLE              = 0
	SHAPETYPE_RHYTM                 = 1
	SHAPETYPE_SQUARE                = 2
	SHAPETYPE_CONTROL_SQUARE_MASTER = 3
	SHAPETYPE_CONTROL_SQUARE_PASSIV = 4
	SHAPETYPE_HEXAGON_SHAPE         = 7
	SHAPETYPE_TRIANGLE_SHAPE        = 8
	SHAPETYPE_TRIANGLE_MINI_SHAPE   = 9
	SHAPETYPE_CONTROL_SHAPE         = 12
)

type Panel struct {
	device    *Device
	Id        uint16 `json:"panelId"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	O         int    `json:"o"`
	ShapeType uint8  `json:"shapeType"`
}

func (p *Panel) GetDevice() *Device {
	return p.device
}
func (p *Panel) IdBytes() []byte {
	switch p.device.Info.Model {
	case NANOLEAF_AURORA:
		return []byte{uint8(p.Id)}
	case NANOLEAF_SHAPES, NANOLEAF_CANVAS:
		id := make([]byte, 2)
		binary.BigEndian.PutUint16(id, p.Id)
		return id
	}
	return []byte{}
}

func (p *Panel) GetColorMode(red uint8, green uint8, blue uint8, transitionTime uint16) string {
	return fmt.Sprintf(
		"1 %d 1 %d %d %d 0 %d",
		p.Id,
		red,
		green,
		blue,
		transitionTime,
	)
}

func (p *Panel) SetColor(red uint8, green uint8, blue uint8, transitionTime uint16) error {
	return p.device.set("/effects",
		Effect{
			EffectBody{
				Command:  "display",
				AnimType: "static",
				AnimData: p.GetColorMode(red, green, blue, transitionTime),
				Palette:  []string{},
			},
		},
		nil)
}
