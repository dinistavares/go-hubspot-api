package hubspot

import "fmt"

// Events service
type EventsService struct {
	service
}

type Event struct {
	EventTemplateID string               `json:"eventTemplateId,omitempty"`
	Domain          string               `json:"domain,omitempty"`
	ID              string               `json:"id,omitempty"`
	Utk             string               `json:"utk,omitempty"`
	Email           string               `json:"email,omitempty"`
	ObjectID        string               `json:"objectId,omitempty"`
	Timestamp       string               `json:"timestamp,omitempty"`
	ObjectType      ObjectType           `json:"objectType,omitempty"`
	ExtraData       interface{}          `json:"extraData,omitempty"`
	Tokens          interface{}          `json:"tokens,omitempty"`
	TimelineIFrame  *EventTimelineIFrame `json:"timelineIFrame,omitempty"`
}

type EventTimelineIFrame struct {
	LinkLabel   string `json:"linkLabel,omitempty"`
	HeaderLabel string `json:"headerLabel,omitempty"`
	Width       int    `json:"width,omitempty"`
	URL         string `json:"url,omitempty"`
	Height      int    `json:"height,omitempty"`
}

// Creat event.
func (service *EventsService) Create(event *Event) (*Event, *Response, error) {
	_url := fmt.Sprintf("/integrators/timeline/%s/events", *service.revision)

	req, _ := service.client.NewRequest("POST", _url, nil, event)

	data := new(Event)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
