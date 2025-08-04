package event

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/donovanhide/eventsource"
	"github.com/roderm/go-nanoleaf/pkg/device"
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
	dev *device.Device
}

func NewListener(dev *device.Device) *Listener {
	l := &Listener{
		dev: dev,
	}
	return l
}

func (l *Listener) reqUrl(eventType int) *url.URL {
	uri, _ := url.Parse(fmt.Sprintf("http://%s:%d/api/v1/%s/events", l.dev.Network.IPv4[0], l.dev.Network.Port, l.dev.GetApiKey()))
	query := uri.Query()
	query.Add("id", fmt.Sprint(eventType))
	uri.RawQuery = query.Encode()
	return uri
}

func (l *Listener) States() (<-chan []EventState, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := l.StatesContext(ctx)
	return stream, cancel, err
}

func (l *Listener) StatesContext(ctx context.Context) (<-chan []EventState, error) {
	stream, err := eventsource.Subscribe(l.reqUrl(EVENT_TYPE_STATE).String(), "")
	if err != nil {
		return nil, err
	}
	states := make(chan []EventState)
	type eventState struct {
		Events []EventState `json:"events"`
		Error  error        `json:"error"`
	}
	go func(ctx context.Context, stream *eventsource.Stream, states chan<- []EventState) {
		for {
			var state eventState
			select {
			case <-ctx.Done():
				stream.Close()
				close(states)
				return
			case event := <-stream.Events:
				err := json.Unmarshal([]byte(event.Data()), &state)
				state.Error = err
			case err := <-stream.Errors:
				state.Error = err
			}
			states <- state.Events
		}
	}(ctx, stream, states)
	return states, err
}
func (l *Listener) Touches() (<-chan []EventTouch, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := l.TouchesContext(ctx)
	return stream, cancel, err
}

func (l *Listener) TouchesContext(ctx context.Context) (<-chan []EventTouch, error) {
	stream, err := eventsource.Subscribe(l.reqUrl(EVENT_TYPE_TOUCH).String(), "")
	if err != nil {
		return nil, err
	}
	touches := make(chan []EventTouch)
	type eventState struct {
		Events []EventTouch `json:"events"`
		Error  error        `json:"error"`
	}
	go func(ctx context.Context, stream *eventsource.Stream, touches chan<- []EventTouch) {
		for {
			var touche eventState
			select {
			case <-ctx.Done():
				stream.Close()
				close(touches)
				return
			case event := <-stream.Events:
				err := json.Unmarshal([]byte(event.Data()), &touche)
				touche.Error = err
			case err := <-stream.Errors:
				touche.Error = err
			}
			touches <- touche.Events
		}
	}(ctx, stream, touches)
	return touches, err
}
