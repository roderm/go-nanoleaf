package types

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	White uint8
}

func (c *Color) Bytes() []byte {
	return []byte{c.Red, c.Green, c.Blue, c.White}
}
