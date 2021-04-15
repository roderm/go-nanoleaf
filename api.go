package nanoleaf

import (
	"fmt"
	"net"
)

type increment struct {
	Increment int `json:"increment"`
	Duration  int `json:"duration"`
}
type intValue struct {
	Value    int `json:"value"`
	Duration int `json:"duration"`
}

func (d *Device) Get() (*Info, error) {
	err := d.get("/", &d.Info)
	if err == nil {
		for _, p := range d.Info.PanelLayout.Layout.Panels {
			p.device = d
		}
	}
	return d.Info, err
}

func (d *Device) SetState(state interface{}) (err error) {
	type onOff struct {
		Value SwitchState `json:"on"`
	}
	err = d.set("/state", state, nil)
	return
}

func (d *Device) SetOn(on bool) error {
	type onOff struct {
		Value SwitchState `json:"on"`
	}
	return d.SetState(onOff{
		SwitchState{
			Value: on,
		}})
}

func (d *Device) On() error {
	return d.SetOn(true)
}
func (d *Device) Off() error {
	return d.SetOn(false)
}

func (d *Device) SetHue(hue int, duration int) error {
	if hue > d.Info.State.Hue.Max || hue < d.Info.State.Brightness.Min {
		return fmt.Errorf("Value for 'hue' must be between %d and %d", d.Info.State.Brightness.Min, d.Info.State.Hue.Max)
	}
	type hueVal struct {
		Hue intValue `json:"hue"`
	}
	return d.SetState(hueVal{
		Hue: intValue{
			Value:    hue,
			Duration: duration,
		},
	})
}

func (d *Device) IncrementHue(inc int, duration int) error {
	type hueVal struct {
		Hue increment `json:"hue"`
	}
	return d.SetState(hueVal{
		Hue: increment{
			Increment: inc,
			Duration:  duration,
		},
	})
}

func (d *Device) SetSaturation(saturation int, duration int) error {
	if saturation > d.Info.State.Sat.Max || saturation < d.Info.State.Sat.Min {
		return fmt.Errorf("Value for 'saturation' must be between %d and %d", d.Info.State.Sat.Min, d.Info.State.Sat.Max)
	}
	type saturationVal struct {
		Saturation intValue `json:"saturation"`
	}
	return d.SetState(saturationVal{
		Saturation: intValue{
			Value:    saturation,
			Duration: duration,
		},
	})
}

func (d *Device) IncrementSaturation(inc int, duration int) error {
	type saturationVal struct {
		Saturation increment `json:"Saturation"`
	}
	return d.SetState(saturationVal{
		Saturation: increment{
			Increment: inc,
			Duration:  duration,
		},
	})
}

func (d *Device) SetBrightness(bright int, duration int) error {
	if bright > d.Info.State.Sat.Max || bright < d.Info.State.Sat.Min {
		return fmt.Errorf("Value for 'bright' must be between %d and %d", d.Info.State.Sat.Min, d.Info.State.Sat.Max)
	}
	type brightVal struct {
		Brightness intValue `json:"brightness"`
	}
	return d.SetState(brightVal{
		Brightness: intValue{
			Value:    bright,
			Duration: duration,
		},
	})
}

func (d *Device) IncrementBrightness(inc int, duration int) error {
	type brightVal struct {
		Brightness increment `json:"Brightness"`
	}
	return d.SetState(brightVal{
		Brightness: increment{
			Increment: inc,
			Duration:  duration,
		},
	})
}

type StreamResponse struct {
	StreamControlIpAddr   net.IP `json:"streamControlIpAddr"`
	StreamControlPort     int    `json:"streamControlPort"`
	StreamControlProtocol string `json:"streamControlProtocol"`
}

type Effect struct {
	Write EffectBody `json:"write"`
}

type EffectBody struct {
	Command           string   `json:"command"`
	AnimType          string   `json:"animType"`
	AnimData          string   `json:"animData"`
	ExtControlVersion string   `json:"extControlVersion"`
	Palette           []string `json:"palette"`
	ColorType         string   `json:"colorType"`
}

func (d *Device) SetExternalStream() (resp StreamResponse, err error) {
	err = d.set("/effects", Effect{
		Write: EffectBody{
			Command:  "display",
			AnimType: "extControl",
			ExtControlVersion: func(d *Device) string {
				switch d.Info.Model {
				case NANOLEAF_AURORA:
					return "v1"
				case NANOLEAF_CANVAS, NANOLEAF_SHAPES:
					return "v2"
				}
				return "v1"
			}(d),
		}}, &resp)
	return
}

func (d *Device) GetEffects() (resp []string, err error) {
	err = d.get("/effects/effectsList", &resp)
	return
}

func (d *Device) SelectEffect(effect string) (err error) {
	type request struct {
		Select string `json:"select"`
	}
	return d.set("/effects", request{
		Select: effect,
	}, nil)
}

func (d *Device) Effect(effect Effect) (err error) {
	return d.set("/effects", effect, nil)
}
