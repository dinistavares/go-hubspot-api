package hubspot

import (
	"fmt"
)

// Tickets service
type TicketsService struct {
	service
}

type Ticket struct {
	Archived   bool        `json:"archived,omitempty"`
	CreatedAt  string      `json:"createdAt,omitempty"`
	ID         string      `json:"id,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
	UpdatedAt  string      `json:"updatedAt,omitempty"`
}

// Get ticket by id.
func (service *TicketsService) Get(id string, opts *QueryValues) (*Ticket, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/tickets/%s", *service.revision, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(Ticket)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Update an existing deal by ID
func (service *TicketsService) Update(id string, properties *Properties) (*Ticket, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/tickets/%s", *service.revision, id)

	body := Ticket{Properties: properties}

	req, _ := service.client.NewRequest("PATCH", _url, nil, body)

	data := new(Ticket)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
