package hubspot

import (
	"fmt"
)

// Owners service
type OwnersService struct {
	service
}

type ListOwnersResponse struct {
	Results *[]Owner `json:"results,omitempty"`
}

type Owner struct {
	Archived                bool   `json:"archived,omitempty"`
	CreatedAt               string `json:"createdAt,omitempty"`
	ID                      string `json:"id,omitempty"`
	UpdatedAt               string `json:"updatedAt,omitempty"`
	Email                   string `json:"email,omitempty"`
	FirstName               string `json:"firstName,omitempty"`
	LastName                string `json:"lastName,omitempty"`
	Type                    string `json:"type,omitempty"`
	UserID                  int    `json:"userId,omitempty"`
	UserIDIncludingInactive int    `json:"userIdIncludingInactive,omitempty"`
}

// List owners.
func (service *OwnersService) List(opts *QueryValues) (*ListOwnersResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/owners", *service.revision)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(ListOwnersResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Get owner by id.
func (service *OwnersService) Get(id string, opts *QueryValues) (*Owner, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/owners/%s", *service.revision, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(Owner)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
