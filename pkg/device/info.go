package device

import (
	"net"

	"github.com/roderm/go-nanoleaf/pkg/types"
)

type NetworkInterface struct {
	Host string
	Port int
	IPv4 []net.IP
	IPv6 []net.IP
}

type Info struct {
	Name            string                 `json:"name"`
	SerialNo        string                 `json:"serialNo"`
	Manufacturer    string                 `json:"manufacturer"`
	Model           string                 `json:"model"`
	FirmwareVersion string                 `json:"firmwareVersion"`
	Discovery       map[string]interface{} `json:"discovery"`
	State           *types.State           `json:"state"`
	Effects         *types.Effects         `json:"effects"`
	PanelLayout     *PanelLayout           `json:"panelLayout"`
	Rhytm           map[string]interface{} `json:"rhythm"`
}

type PanelLayout struct {
	GlobalOrientation *types.RangeState `json:"globalOrientation"`
	Layout            *Layout           `json:"layout"`
}
type Layout struct {
	NumPanels  uint16   `json:"numPanels"`
	SideLength int      `json:"sideLength"`
	Panels     []*Panel `json:"positionData"`
}
