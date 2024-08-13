package hubspot

import (
	"fmt"
)

// Event Templates service
type EventTemplatesService struct {
	service
}

type EventTemplateResults struct {
	EventTemplates *[]EventTemplate `json:"results,omitempty"`
}

type EventTemplate struct {
	ID             string                `json:"id,omitempty"`
	Name           string                `json:"name,omitempty"`
	ObjectType     string                `json:"objectType,omitempty"`
	DetailTemplate string                `json:"detailTemplate,omitempty"`
	HeaderTemplate string                `json:"headerTemplate,omitempty"`
	CreatedAt      string                `json:"createdAt,omitempty"`
	UpdatedAt      string                `json:"updatedAt,omitempty"`
	Tokens         *[]EventTemplateToken `json:"tokens,omitempty"`
}

type EventTemplateToken struct {
	Name               string              `json:"name,omitempty"`
	Type               string              `json:"type,omitempty"`
	Label              string              `json:"label,omitempty"`
	CreatedAt          string              `json:"createdAt,omitempty"`
	UpdatedAt          string              `json:"updatedAt,omitempty"`
	ObjectPropertyName string              `json:"objectPropertyName,omitempty"`
	Options            *[]EventTokenOption `json:"options,omitempty"`
}

type EventTokenOption struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

// Get event template by ID
func (service *EventTemplatesService) Get(appID string, templateID string) (*EventTemplate, *Response, error) {
	_url := fmt.Sprintf("/integrators/timeline/%s/%s/event-templates/%s", *service.revision, appID, templateID)

	if err := service.client.isDeveloperApiKeySet(); err != nil {
		return nil, nil, err
	}

	req, _ := service.client.NewRequest("GET", _url, service.client.setDeveloperAuthenticationParams(), nil)

	data := new(EventTemplate)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// List event all app event templates
func (service *EventTemplatesService) List(appID string) (*EventTemplateResults, *Response, error) {
	_url := fmt.Sprintf("/integrators/timeline/%s/%s/event-templates", *service.revision, appID)

	if err := service.client.isDeveloperApiKeySet(); err != nil {
		return nil, nil, err
	}

	req, _ := service.client.NewRequest("GET", _url, service.client.setDeveloperAuthenticationParams(), nil)

	data := new(EventTemplateResults)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Create event template
func (service *EventTemplatesService) Create(appID string, template *EventTemplate) (*EventTemplate, *Response, error) {
	_url := fmt.Sprintf("/integrators/timeline/%s/%s/event-templates", *service.revision, appID)

	if err := service.client.isDeveloperApiKeySet(); err != nil {
		return nil, nil, err
	}

	req, _ := service.client.NewRequest("POST", _url, service.client.setDeveloperAuthenticationParams(), template)

	data := new(EventTemplate)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Update an existing event template
func (service *EventTemplatesService) Update(appID string, templateID string, template *EventTemplate) (*EventTemplate, *Response, error) {
	_url := fmt.Sprintf("/integrators/timeline/%s/%s/event-templates/%s", *service.revision, appID, templateID)

	if err := service.client.isDeveloperApiKeySet(); err != nil {
		return nil, nil, err
	}

	req, _ := service.client.NewRequest("PUT", _url, service.client.setDeveloperAuthenticationParams(), template)

	data := new(EventTemplate)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Delete an event template
func (service *EventTemplatesService) Delete(appID string, templateID string) (*Response, error) {
	_url := fmt.Sprintf("/integrators/timeline/%s/%s/event-templates/%s", *service.revision, appID, templateID)

	if err := service.client.isDeveloperApiKeySet(); err != nil {
		return nil, err
	}

	req, _ := service.client.NewRequest("DELETE", _url, service.client.setDeveloperAuthenticationParams(), nil)

	data := new(EventTemplate)
	response, err := service.client.Do(req, data)

	if err != nil {
		return response, err
	}

	return response, nil
}
