package hubspot

// Account Information Service
type AccountInformationService struct {
	service
}

type AccountDetails struct {
	PortalID              int      `json:"portalId,omitempty"`
	AccountType           string   `json:"accountType,omitempty"`
	TimeZone              string   `json:"timeZone,omitempty"`
	CompanyCurrency       string   `json:"companyCurrency,omitempty"`
	AdditionalCurrencies  []string `json:"additionalCurrencies,omitempty"`
	UtcOffset             string   `json:"utcOffset,omitempty"`
	UtcOffsetMilliseconds int      `json:"utcOffsetMilliseconds,omitempty"`
	UIDomain              string   `json:"uiDomain,omitempty"`
	DataHostingLocation   string   `json:"dataHostingLocation,omitempty"`
}

// List associations by ID.
func (service *AccountInformationService) Get() (*AccountDetails, *Response, error) {
	_url := "/account-info/v3/details"

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	data := new(AccountDetails)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}