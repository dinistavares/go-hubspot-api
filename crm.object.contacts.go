package hubspot

import (
	"fmt"
)

// Contacts service
type ContactsService struct {
	service
}

type ContactSearchResults struct {
	Contacts *[]Contact `json:"results,omitempty"`
	GenericSearchResults
}

type Contact struct {
	Archived              bool        `json:"archived,omitempty"`
	CreatedAt             string      `json:"createdAt,omitempty"`
	ID                    string      `json:"id,omitempty"`
	UpdatedAt             string      `json:"updatedAt,omitempty"`
	Properties            *Properties `json:"properties,omitempty"`
	PropertiesWithHistory *Properties `json:"propertiesWithHistory,omitempty"`
}

// Get contact by id.
func (service *ContactsService) Get(id string, opts *QueryValues) (*Contact, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/contacts/%s", *service.revision, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(Contact)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Creat contact.
func (service *ContactsService) Create(contact *GenericCreateBody) (*Contact, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/contacts", *service.revision)

	req, _ := service.client.NewRequest("POST", _url, nil, contact)

	data := new(Contact)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Update an existing deal by ID
func (service *ContactsService) Update(id string, properties *Properties) (*Contact, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/contacts/%s", *service.revision, id)

	body := Contact{Properties: properties}

	req, _ := service.client.NewRequest("PATCH", _url, nil, body)

	data := new(Contact)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Search contacts
func (service *ContactsService) Search(search *SearchRequest) (*ContactSearchResults, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/contacts/search", *service.revision)

	req, _ := service.client.NewRequest("POST", _url, nil, search)

	data := new(ContactSearchResults)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
