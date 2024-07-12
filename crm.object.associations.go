package hubspot

import (
	"fmt"
)

// Associations service
type AssociationsService struct {
	service
}

type ListAssociationResponse struct {
	Associations *[]Association `json:"results,omitempty"`
}

type Association struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// List associations by ID.
func (service *AssociationsService) List(objectType ObjectTypeID, id string, objectToType ObjectTypeID, opts QueryValues) (*ListAssociationResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/%s/%s/associations/%s", *service.revision, objectType, id, objectToType)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(ListAssociationResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
