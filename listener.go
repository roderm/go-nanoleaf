package nanoleaf

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/donovanhide/eventsource"
)

const (
	EVENT_ATTRIBUTE_ON        = 1
	EVENT_ATTRIBUTE_BRIGHT    = 2
	EVENT_ATTRIBUTE_HUE       = 2
	EVENT_ATTRIBUTE_SAT       = 4
	EVENT_ATTRIBUTE_CCT       = 5
	EVENT_ATTRIBUTE_COLORMODE = 6
)

const (
	EVENT_TYPE_STATE  = 1
	EVENT_TYPE_LAYOUT = 2
	EVENT_TYPE_EFFECT = 3
	EVENT_TYPE_TOUCH  = 4
)

type EventState struct {
	Attribute uint8       `json:"attr"`
	Value     interface{} `json:"value"`
}
type EventTouch struct {
	Gesture uint8  `json:"gesture"`
	PanelId uint16 `json:"panelId"`
}

type Listener struct {
	dev  *Device
	stop chan struct{}
}

func NewListener(dev *Device) *Listener {
	l := &Listener{
		dev: dev,
	}
	return l
}

func (l *Listener) reqUrl(eventType int) *url.URL {
	uri, _ := url.Parse(fmt.Sprintf("http://%s:%d/api/v1/%s/events", l.dev.Network.IPv4[0], l.dev.Network.Port, l.dev.apiKey))
	query := uri.Query()
	query.Add("id", fmt.Sprint(eventType))
	uri.RawQuery = query.Encode()
	return uri
}

func (l *Listener) States() (chan []EventState, func(), error) {
	states := make(chan []EventState)
	cancel := func() {
		close(states)
	}

	stream, err := eventsource.Subscribe(l.reqUrl(EVENT_TYPE_STATE).String(), "")
	if err != nil {
		return states, cancel, err
	}
	cancel = func() {
		stream.Close()
		close(states)
	}
	type eventState struct {
		Events []EventState `json:"events"`
	}
	go func(stream *eventsource.Stream, states chan []EventState) {
		defer stream.Close()
		for event := range stream.Events {
			var state eventState
			err := json.Unmarshal([]byte(event.Data()), &state)
			if err != nil {
				continue
			}
			states <- state.Events
		}
	}(stream, states)
	return states, cancel, err
}

func (l *Listener) Touches() (chan []EventTouch, func(), error) {
	touches := make(chan []EventTouch)
	cancel := func() {
		close(touches)
	}

	stream, err := eventsource.Subscribe(l.reqUrl(EVENT_TYPE_TOUCH).String(), "")
	if err != nil {
		return touches, cancel, err
	}
	cancel = func() {
		stream.Close()
		close(touches)
	}
	type eventState struct {
		Events []EventTouch `json:"events"`
	}
	go func(stream *eventsource.Stream, touches chan []EventTouch) {
		defer stream.Close()
		for event := range stream.Events {
			var touche eventState
			err := json.Unmarshal([]byte(event.Data()), &touche)
			if err != nil {
				continue
			}
			touches <- touche.Events
		}
	}(stream, touches)
	return touches, cancel, err
}
