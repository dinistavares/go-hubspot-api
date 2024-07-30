package hubspot

import "fmt"

// Subscription Preferences Service
type SubscriptionPreferencesService struct {
	service
}

// Get subscription statuses for a contact by email.
func (service *SubscriptionPreferencesService) Get(email string) (*interface{}, *Response, error) {
	_url := fmt.Sprintf("/communication-preferences/%s/status/email/%s", *service.revision, email)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	data := new(interface{})
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}