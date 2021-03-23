package nanoleaf

import "net"

const (
	PANEL_AURORA = 0
	PANEL_CANVAS = 2
)
const (
	NANOLEAF_AURORA = "NL22"
	NANOLEAF_CANVAS = "NL29"
	NANOLEAF_SHAPES = "NL42"
)

type SwitchState struct {
	Value bool `json:"value"`
}

type RangeState struct {
	Value int `json:"value"`
	Min   int `json:"min"`
	Max   int `json:"max"`
}

type Info struct {
	Name            string                 `json:"name"`
	SerialNo        string                 `json:"serialNo"`
	Manufacturer    string                 `json:"manufacturer"`
	Model           string                 `json:"model"`
	FirmwareVersion string                 `json:"firmwareVersion"`
	Discovery       map[string]interface{} `json:"discovery"`
	State           *State                 `json:"state"`
	Effects         *Effects               `json:"effects"`
	PanelLayout     *PanelLayout           `json:"panelLayout"`
	Rhytm           map[string]interface{} `json:"rhythm"`
}
type State struct {
	On         *SwitchState `json:"state"`
	Brightness *RangeState  `json:"brightness"`
	Hue        *RangeState  `json:"hue"`
	Sat        *RangeState  `json:"sat"`
	Ct         *RangeState  `json:"ct"`
	ColorMode  string       `json:"colorMode"`
}

type Effects struct {
	List     []string `json:"effectsList"`
	Selected string   `json:"selected"`
}

type PanelLayout struct {
	GlobalOrientation *RangeState `json:"globalOrientation"`
	Layout            *Layout     `json:"layout"`
}
type Layout struct {
	NumPanels  uint16   `json:"numPanels"`
	SideLength int      `json:"sideLength"`
	Panels     []*Panel `json:"positionData"`
}

type NetworkInterface struct {
	Host string
	Port int
	IPv4 []net.IP
	IPv6 []net.IP
}

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	White uint8
}

func (c *Color) Bytes() []byte {
	return []byte{c.Red, c.Green, c.Blue, c.White}
}
