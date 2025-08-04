package types

type State struct {
	On         *SwitchState `json:"state"`
	Brightness *RangeState  `json:"brightness"`
	Hue        *RangeState  `json:"hue"`
	Sat        *RangeState  `json:"sat"`
	Ct         *RangeState  `json:"ct"`
	ColorMode  string       `json:"colorMode"`
}

type SwitchState struct {
	Value bool `json:"value"`
}

type RangeState struct {
	Value int `json:"value"`
	Min   int `json:"min"`
	Max   int `json:"max"`
}
