package hubspot

import "fmt"

// Properties service
type PropertiesService struct {
	service
}

type Properties map[string]interface{}

type ObjectPropertiesResponse struct {
	Results *[]ObjectProperty `json:"results,omitempty"`
}

type ObjectProperty struct {
	UpdatedAt            string                `json:"updatedAt,omitempty"`
	CreatedAt            string                `json:"createdAt,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Label                string                `json:"label,omitempty"`
	Type                 string                `json:"type,omitempty"`
	FieldType            string                `json:"fieldType,omitempty"`
	Description          string                `json:"description,omitempty"`
	GroupName            string                `json:"groupName,omitempty"`
	Options              []interface{}         `json:"options,omitempty"`
	DisplayOrder         int                   `json:"displayOrder,omitempty"`
	Calculated           bool                  `json:"calculated,omitempty"`
	ExternalOptions      bool                  `json:"externalOptions,omitempty"`
	HasUniqueValue       bool                  `json:"hasUniqueValue,omitempty"`
	Hidden               bool                  `json:"hidden,omitempty"`
	HubspotDefined       bool                  `json:"hubspotDefined,omitempty"`
	ShowCurrencySymbol   bool                  `json:"showCurrencySymbol,omitempty"`
	FormField            bool                  `json:"formField,omitempty"`
	CalculationFormula   string                `json:"calculationFormula,omitempty"`
	DataSensitivity      string                `json:"dataSensitivity,omitempty"`
	ReferencedObjectType string                `json:"referencedObjectType,omitempty"`
	Archived             bool                  `json:"archived,omitempty"`
	CreatedUserID        string                `json:"createdUserId,omitempty"`
	UpdatedUserID        string                `json:"updatedUserId,omitempty"`
	ModificationMetadata *ModificationMetadata `json:"modificationMetadata,omitempty"`
}

type ModificationMetadata struct {
	Archivable         bool `json:"archivable"`
	ReadOnlyDefinition bool `json:"readOnlyDefinition"`
	ReadOnlyValue      bool `json:"readOnlyValue"`
}

type CreateProperty struct {
	FormField            bool              `json:"formField,omitempty"`
	ReferencedObjectType string            `json:"referencedObjectType,omitempty"`
	CalculationFormula   string            `json:"calculationFormula,omitempty"`
	HasUniqueValue       bool              `json:"hasUniqueValue,omitempty"`
	ExternalOptions      bool              `json:"externalOptions,omitempty"`
	Hidden               bool              `json:"hidden,omitempty"`
	DisplayOrder         int               `json:"displayOrder,omitempty"`
	Description          string            `json:"description,omitempty"`
	Label                string            `json:"label,omitempty"`
	Type                 string            `json:"type,omitempty"`
	GroupName            string            `json:"groupName,omitempty"`
	Name                 string            `json:"name,omitempty"`
	FieldType            string            `json:"fieldType,omitempty"`
	Options              *[]PropertyOption `json:"options,omitempty"`
}

type PropertyOption struct {
	Hidden      bool   `json:"hidden,omitempty"`
	Description string `json:"description,omitempty"`
	Label       string `json:"label,omitempty"`
	Value       string `json:"value,omitempty"`
}

// Get a proptery by property type and name.
func (service *PropertiesService) Get(objectType ObjectType, name string, opts *QueryValues) (*ObjectProperty, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/properties/%s/%s", *service.revision, objectType, name)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	accounts := new(ObjectProperty)
	response, err := service.client.Do(req, accounts)

	if err != nil {
		return nil, response, err
	}

	return accounts, response, nil
}

// Create property for object type.
func (service *PropertiesService) Create(objectType ObjectType, property *CreateProperty) (*ObjectProperty, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/properties/%s", *service.revision, objectType)

	req, _ := service.client.NewRequest("POST", _url, nil, property)

	data := new(ObjectProperty)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// List properties by property type.
func (service *PropertiesService) List(objectType ObjectType, opts *QueryValues) (*ObjectPropertiesResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/properties/%s", *service.revision, objectType)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	accounts := new(ObjectPropertiesResponse)
	response, err := service.client.Do(req, accounts)

	if err != nil {
		return nil, response, err
	}

	return accounts, response, nil
}

func (p *Properties) Add(key string, value interface{}) {
	(*p)[key] = value
}

func (p *Properties) Get(key string) interface{} {
	if property, ok := (*p)[key]; ok {
		return property
	}

	return nil
}
