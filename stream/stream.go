package stream

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/roderm/go-nanoleaf"
)

type externalController struct {
	config nanoleaf.StreamResponse
	dev    *nanoleaf.Device
	conn   net.Conn
}

type Option func(*externalController)

func NewControll(dev *nanoleaf.Device, opts ...Option) (ctrl *externalController, err error) {
	ctrl = &externalController{
		dev: dev,
	}
	ctrl.config, err = dev.SetExternalStream()
	if ctrl.config.StreamControlIpAddr == nil {
		ctrl.config.StreamControlIpAddr = dev.Network.IPv4[0]
	}
	if ctrl.config.StreamControlPort == 0 {
		ctrl.config.StreamControlPort = 60222
	}
	if ctrl.config.StreamControlProtocol == "" {
		ctrl.config.StreamControlProtocol = "udp"
	}

	if err != nil {
		return nil, err
	}
	ctrl.conn, err = net.Dial(ctrl.config.StreamControlProtocol, fmt.Sprintf("%s:%d", ctrl.config.StreamControlIpAddr, ctrl.config.StreamControlPort))
	return
}

func (ctrl *externalController) SetPanel(color nanoleaf.Color, transitionTime uint16, panels ...*nanoleaf.Panel) error {
	var bs []byte
	var timeBuff []byte
	switch ctrl.dev.Info.Model {
	case nanoleaf.NANOLEAF_AURORA:
		bs = []byte{uint8(len(panels))}
		timeBuff = []byte{uint8(transitionTime)}
	case nanoleaf.NANOLEAF_CANVAS, nanoleaf.NANOLEAF_SHAPES:
		bs = make([]byte, 2)
		timeBuff = make([]byte, 2)
		binary.BigEndian.PutUint16(bs, uint16(len(panels)))
		binary.BigEndian.PutUint16(timeBuff, transitionTime)
	}
	for _, p := range panels {
		bs = append(bs, p.IdBytes()...)
		bs = append(bs, color.Bytes()...)
		bs = append(bs, timeBuff...)
	}

	_, err := ctrl.conn.Write(bs)
	return err
}
