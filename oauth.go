package hubspot

import (
	"fmt"
)

// OAuth service
type OAuthService struct {
	service
}

type OAuthAccessToken struct {
	AppID             int                     `json:"app_id,omitempty"`
	ExpiresIn         int                     `json:"expires_in,omitempty"`
	HubDomain         string                  `json:"hub_domain,omitempty"`
	HubID             int                     `json:"hub_id,omitempty"`
	Token             string                  `json:"token,omitempty"`
	TokenType         string                  `json:"token_type,omitempty"`
	User              string                  `json:"user,omitempty"`
	UserID            int                     `json:"user_id,omitempty"`
	Scopes            *[]string               `json:"scopes,omitempty"`
	SignedAccessToken *OAuthSignedAccessToken `json:"signed_access_token,omitempty"`
}

type OAuthSignedAccessToken struct {
	AppID                     int    `json:"appId,omitempty"`
	ExpiresAt                 int64  `json:"expiresAt,omitempty"`
	HubID                     int    `json:"hubId,omitempty"`
	Hublet                    string `json:"hublet,omitempty"`
	IsUserLevel               bool   `json:"isUserLevel,omitempty"`
	NewSignature              string `json:"newSignature,omitempty"`
	ScopeToScopeGroupPks      string `json:"scopeToScopeGroupPks,omitempty"`
	Scopes                    string `json:"scopes,omitempty"`
	Signature                 string `json:"signature,omitempty"`
	TrialScopeToScopeGroupPks string `json:"trialScopeToScopeGroupPks,omitempty"`
	TrialScopes               string `json:"trialScopes,omitempty"`
	UserID                    int    `json:"userId,omitempty"`
}

// Get access token information.
func (service *OAuthService) GetAccessTokenInformation() (*OAuthAccessToken, *Response, error) {
	_url := fmt.Sprintf("/oauth/v1/access-tokens/%s", service.client.auth.AccessToken)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	data := new(OAuthAccessToken)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
