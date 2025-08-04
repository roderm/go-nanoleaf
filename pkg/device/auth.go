package device

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (d *Device) AuthorizationToken() string {
	return d.apiKey
}

func (d *Device) InvalidateKey(authToken string) error {
	uri, err := url.Parse(fmt.Sprintf("http://%s:%d/api/v1/%s", d.Network.IPv4[0], d.Network.Port, authToken))
	if err != nil {
		return err
	}
	resp, err := d.client.Do(&http.Request{
		Method: "DELETE",
		URL:    uri,
	})
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("error in HTTP-Request on authorize (%s)", resp.Status)
	}
	return nil
}
func (d *Device) Authorization() error {
	type responseBody struct {
		AuthToken string `json:"auth_token"`
	}

	var token responseBody
	response, err := d.client.Post(fmt.Sprintf("http://%s:%d/api/v1/new", d.Network.IPv4[0], d.Network.Port), "application/json", nil)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("error in HTTP-Request on authorize (%d - %s)", response.StatusCode, response.Status)
	}
	err = json.NewDecoder(response.Body).Decode(&token)
	if err != nil {
		return err
	}
	d.apiKey = token.AuthToken
	_, err = d.Get()
	return err
}
