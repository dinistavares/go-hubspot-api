package hubspot

import (
	"fmt"
)

// Pipelines service
type PipelinesService struct {
	service
}

type ListPipelinesResponse struct {
	Results *[]Pipeline `json:"results,omitempty"`
}

type Pipeline struct {
	Label        string          `json:"label,omitempty"`
	DisplayOrder int             `json:"displayOrder,omitempty"`
	ID           string          `json:"id,omitempty"`
	Stages       *[]PipelineStage `json:"stages,omitempty"`
	CreatedAt    string          `json:"createdAt,omitempty"`
	UpdatedAt    string          `json:"updatedAt,omitempty"`
	Archived     bool            `json:"archived,omitempty"`
}

type PipelineStage struct {
	Label            string                `json:"label,omitempty"`
	DisplayOrder     int                   `json:"displayOrder,omitempty"`
	Metadata        *PipelineStageMetadata `json:"metadata,omitempty"`
	ID               string                `json:"id,omitempty"`
	CreatedAt        string                `json:"createdAt,omitempty"`
	UpdatedAt        string                `json:"updatedAt,omitempty"`
	Archived         bool                  `json:"archived,omitempty"`
	WritePermissions string                `json:"writePermissions,omitempty"`
}

type PipelineStageMetadata struct {
	IsClosed    string `json:"isClosed,omitempty"`
	Probability string `json:"probability,omitempty"`
}

// List pipelines by ID.
func (service *PipelinesService) List(objectType ObjectType) (*ListPipelinesResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/pipelines/%s", *service.revision, objectType)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	data := new(ListPipelinesResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
