package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Device struct {
	client          *http.Client
	Id              string
	Name            string
	Type            string
	FirmwareVersion string
	Network         *NetworkInterface
	apiKey          string
	Info            *Info
}

func NewDevice(opts ...Option) *Device {
	dev := &Device{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		Network: &NetworkInterface{
			Port: 16021,
		},
	}
	for _, o := range opts {
		o(dev)
	}
	return dev
}

func (d *Device) handleApiError(resp *http.Response, result interface{}) error {
	if resp.StatusCode == 204 {
		return nil
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("error in HTTP-Request (%s)", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(result)
}

func (d *Device) getStream(reqBody interface{}) io.ReadCloser {
	r, w := io.Pipe()
	go func(w *io.PipeWriter) {
		defer w.Close()
		if err := json.NewEncoder(w).Encode(reqBody); err != nil {
			panic(err)
		}
	}(w)
	return r
}

func (d *Device) GetApiKey() string {
	return d.apiKey
}

// PUT
func (d *Device) set(path string, reqBody interface{}, result interface{}) error {
	uri, err := url.Parse(fmt.Sprintf("http://%s:%d/api/v1/%s%s", d.Network.IPv4[0], d.Network.Port, d.apiKey, path))
	if err != nil {
		return err
	}
	resp, err := d.client.Do(&http.Request{
		Method: "PUT",
		URL:    uri,
		Body:   d.getStream(reqBody),
	})
	if err != nil {
		return err
	}
	return d.handleApiError(resp, result)
}

// GET
func (d *Device) get(path string, result interface{}) error {
	resp, err := d.client.Get(fmt.Sprintf("http://%s:%d/api/v1/%s%s", d.Network.IPv4[0], d.Network.Port, d.apiKey, path))
	if err != nil {
		return err
	}
	return d.handleApiError(resp, result)
}
