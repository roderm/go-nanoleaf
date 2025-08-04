package stream

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/roderm/go-nanoleaf/pkg/device"
	"github.com/roderm/go-nanoleaf/pkg/types"
)

type externalController struct {
	config device.StreamResponse
	dev    *device.Device
	conn   net.Conn
}

type Option func(*externalController)

func NewControll(dev *device.Device, opts ...Option) (ctrl *externalController, err error) {
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
	ctrl.conn, err = net.Dial(ctrl.config.StreamControlProtocol, fmt.Sprintf("%s:%d", ctrl.config.StreamControlIpAddr.String(), ctrl.config.StreamControlPort))
	return
}

func (ctrl *externalController) SetPanel(color types.Color, transitionTime uint16, panels ...*device.Panel) error {
	var bs []byte
	var timeBuff []byte
	switch ctrl.dev.Info.Model {
	case types.NANOLEAF_AURORA:
		bs = []byte{uint8(len(panels))}
		timeBuff = []byte{uint8(transitionTime)}
	case types.NANOLEAF_CANVAS, types.NANOLEAF_SHAPES:
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

func (ctrl *externalController) Stream(ctx context.Context) {
	for {
		buff := make([]byte, 1024)
		if _, err := ctrl.conn.Read(buff); err == io.EOF {
			return
		}
		if len(buff) > 0 {
			fmt.Println(string(buff))
		}
	}
}
