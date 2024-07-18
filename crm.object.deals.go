package hubspot

import (
	"fmt"
)

// Deals service
type DealsService struct {
	service
}

type Deal struct {
	Archived   bool        `json:"archived,omitempty"`
	CreatedAt  string      `json:"createdAt,omitempty"`
	ID         string      `json:"id,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
	UpdatedAt  string      `json:"updatedAt,omitempty"`
}

// Get deal by id.
func (service *DealsService) Get(id string, opts *QueryValues) (*Deal, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/deals/%s", *service.revision, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(Deal)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Create a deal
func (service *DealsService) Create(deal *GenericCreateBody) (*Deal, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/deals", *service.revision)

	req, _ := service.client.NewRequest("POST", _url, nil, deal)

	data := new(Deal)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Update an existing deal by ID
func (service *DealsService) Update(id string, properties *Properties) (*Deal, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/deals/%s", *service.revision, id)

	body := Deal{Properties: properties}

	req, _ := service.client.NewRequest("PATCH", _url, nil, body)

	data := new(Deal)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
