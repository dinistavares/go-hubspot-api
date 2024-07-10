package hubspot

import "fmt"

// Properties service
type PropertiesService struct {
	service
}

type ObjectPropertiesResponse struct {
	Results *[]ObjectProperty `json:"results,omitempty"`
}

type ObjectProperty struct {
	UpdatedAt            string      `json:"updatedAt,omitempty"`
	CreatedAt            string      `json:"createdAt,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Label                string      `json:"label,omitempty"`
	Type                 string      `json:"type,omitempty"`
	FieldType            string      `json:"fieldType,omitempty"`
	Description          string      `json:"description,omitempty"`
	GroupName            string      `json:"groupName,omitempty"`
	Options              []any       `json:"options,omitempty"`
	DisplayOrder         int         `json:"displayOrder,omitempty"`
	Calculated           bool        `json:"calculated,omitempty"`
	ExternalOptions      bool        `json:"externalOptions,omitempty"`
	HasUniqueValue       bool        `json:"hasUniqueValue,omitempty"`
	Hidden               bool        `json:"hidden,omitempty"`
	HubspotDefined       bool        `json:"hubspotDefined,omitempty"`
	ShowCurrencySymbol   bool        `json:"showCurrencySymbol,omitempty"`
	ModificationMetadata interface{} `json:"modificationMetadata,omitempty"`
	FormField            bool        `json:"formField,omitempty"`
	CalculationFormula   string      `json:"calculationFormula,omitempty"`
	DataSensitivity      string      `json:"dataSensitivity,omitempty"`
	ReferencedObjectType string      `json:"referencedObjectType,omitempty"`
	Archived             bool        `json:"archived,omitempty"`
	CreatedUserID        string      `json:"createdUserId,omitempty"`
	UpdatedUserID        string      `json:"updatedUserId,omitempty"`
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
