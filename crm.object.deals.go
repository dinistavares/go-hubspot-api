package hubspot

import (
	"fmt"
)

// Deals service
type DealsService struct {
	service
}

type Deal struct {
	Archived   bool                   `json:"archived,omitempty"`
	CreatedAt  string                 `json:"createdAt,omitempty"`
	ID         string                 `json:"id,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	UpdatedAt  string                 `json:"updatedAt,omitempty"`
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
