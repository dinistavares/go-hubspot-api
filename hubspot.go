package hubspot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion               = "1.1"
	defaultAuthHeaderName        = "Authorization"
	defaultOAuthPrefix           = "Bearer"
	defaultRestEndpointURL       = "https://api.hubapi.com"
	defaultRestAPIRevision       = "v3"
	acceptedContentType          = "application/json"
	userAgent                    = "go-hubspot-api/" + libraryVersion
	clientRequestRetryAttempts   = 2
	clientRequestRetryHoldMillis = 1000
)

var (
	errorDoAllAttemptsExhausted = errors.New("all request attempts were exhausted")
	errorDoAttemptNilRequest    = errors.New("request could not be constructed")
)

type ApiType string

type ClientConfig struct {
	HttpClient           *http.Client
	RestEndpointURL      string
	RestEndpointRevision string
}

type auth struct {
	Available       bool
	AccessToken     string
	HeaderName      string
	DeveloperApiKey string
}

type Client struct {
	config  *ClientConfig
	client  *http.Client
	auth    *auth
	baseURL *url.URL

	// CRM
	AccountInformation *AccountInformationService
	Associations       *AssociationsService
	Contacts           *ContactsService
	Deals              *DealsService
	Events             *EventsService
	EventTemplates     *EventTemplatesService
	Lists              *ListsService
	Pipelines          *PipelinesService
	Properties         *PropertiesService
	Tickets            *TicketsService

	// Marketing
	SubscriptionPreferences *SubscriptionPreferencesService
}

type service struct {
	client   *Client
	revision *string
}

type GenericResponse struct {
	Response *http.Response

	Category      string   `json:"category,omitempty"`
	CorrelationID string   `json:"correlationId,omitempty"`
	Errors        *[]Error `json:"errors,omitempty"`
	Message       string   `json:"message,omitempty"`
	Status        string   `json:"status,omitempty"`
}

type ErrorSource struct {
	Pointer string `json:"pointer,omitempty"`
}

type Response struct {
	*http.Response
}

type Error struct {
	Message string      `json:"message,omitempty"`
	In      string      `json:"in,omitempty"`
	Context interface{} `json:"context,omitempty"`
}

func (response *GenericResponse) Error() string {
	errorString := fmt.Sprintf("%v %v: %d",
		response.Response.Request.Method, response.Response.Request.URL,
		response.Response.StatusCode)

	if response.Category != "" {
		errorString += fmt.Sprintf(" Category: %s", response.Category)
	}

	if response.Message != "" {
		errorString += fmt.Sprintf(" Message: %s", response.Message)
	}

	if response.Errors != nil && len(*response.Errors) > 0 {
		firstError := (*response.Errors)[0]

		if firstError.Message != "" {
			errorString += fmt.Sprintf(" Error: %s", firstError.Message)
		}

		if firstError.Context != "" {
			errorString += fmt.Sprintf(" Context: %+v", firstError.Context)
		}
	}

	if response.CorrelationID != "" {
		errorString += fmt.Sprintf(" (correlation_id): %s", response.CorrelationID)
	}

	return errorString
}

func NewWithConfig(config ClientConfig) *Client {
	if config.HttpClient == nil {
		config.HttpClient = http.DefaultClient
	}

	if config.RestEndpointURL == "" {
		config.RestEndpointURL = defaultRestEndpointURL
	}

	if config.RestEndpointRevision == "" {
		config.RestEndpointRevision = defaultRestAPIRevision
	}

	// Create client
	baseURL, _ := url.Parse(config.RestEndpointURL)

	client := &Client{config: &config, client: config.HttpClient, auth: &auth{}, baseURL: baseURL}

	// Map services
	client.AccountInformation = &AccountInformationService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.Associations = &AssociationsService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.Contacts = &ContactsService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.Deals = &DealsService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.Events = &EventsService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.EventTemplates = &EventTemplatesService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.Lists = &ListsService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.Pipelines = &PipelinesService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.Properties = &PropertiesService{service{client: client, revision: &client.config.RestEndpointRevision}}
	client.Tickets = &TicketsService{service{client: client, revision: &client.config.RestEndpointRevision}}

	client.SubscriptionPreferences = &SubscriptionPreferencesService{service{client: client, revision: &client.config.RestEndpointRevision}}

	return client
}

func New() *Client {
	return NewWithConfig(ClientConfig{})
}

func (client *Client) Authenticate(accessToken string) {
	client.auth.HeaderName = defaultAuthHeaderName
	client.auth.AccessToken = accessToken
	client.auth.Available = true
}

// Set authentication key for developer routes (eg. EventTemplates service routes)
func (client *Client) DeveloperAuthenticate(apiKey string) {
	client.auth.DeveloperApiKey = apiKey
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, opts interface{}, body interface{}) (*http.Request, error) {
	// Append Query Params to URL
	if opts, ok := isPointerWithQueryValues(opts); ok {
		if v, ok := opts.(QueryValues); ok {
			urlStr += v.getQueryValues().encode()
		}
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := client.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)

		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	// Only append access token for routes without a HUBSPOT_DEVELOPER_API_KEY.
	if hapikey := url.Query().Get("hapikey"); hapikey == "" && client.auth.Available {
		req.Header.Add(client.auth.HeaderName, fmt.Sprintf("Bearer %s", client.auth.AccessToken))
	}

	req.Header.Add("Accept", acceptedContentType)
	req.Header.Add("Content-type", acceptedContentType)
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	var lastErr error

	attempts := 0

	for attempts < clientRequestRetryAttempts {
		// Hold before this attempt? (ie. not first attempt)
		if attempts > 0 {
			time.Sleep(clientRequestRetryHoldMillis * time.Millisecond)
		}

		// Dispatch request attempt
		attempts++
		resp, shouldRetry, err := client.doAttempt(req, v)

		// Return response straight away? (we are done)
		if !shouldRetry {
			return resp, err
		}

		// Should retry: store last error (we are not done)
		lastErr = err
	}

	// Set default error? (all attempts failed, but no error is set)
	if lastErr == nil {
		lastErr = errorDoAllAttemptsExhausted
	}

	// All attempts failed, return last attempt error
	return nil, lastErr
}

func (client *Client) doAttempt(req *http.Request, v interface{}) (*Response, bool, error) {
	if req == nil {
		return nil, false, errorDoAttemptNilRequest
	}

	resp, err := client.client.Do(req)

	if checkRequestRetry(resp, err) {
		return nil, true, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		return response, false, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return response, false, err
}

func newResponse(httpResponse *http.Response) *Response {
	response := Response{Response: httpResponse}

	return &response
}

// checkRequestRetry checks if should retry request
func checkRequestRetry(response *http.Response, err error) bool {
	// Low-level error, or response status is a server error? (HTTP 5xx)
	if err != nil || response.StatusCode >= 500 {
		return true
	}

	// No low-level error (should not retry)
	return false
}

// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
	// No error in response? (HTTP 2xx)
	if code := response.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	// Map response error data (eg. HTTP 4xx)
	errorResponse := &GenericResponse{Response: response}

	data, err := io.ReadAll(response.Body)

	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

func (client *Client) isDeveloperApiKeySet() error {
	if client.auth.DeveloperApiKey == "" {
		return errors.New("please set developer api key using the 'DeveloperAuthenticate' function on the hubspot client")
	}

	return nil
}

func (client *Client) setDeveloperAuthenticationParams() QueryValues {
	params := QueryValues{}
	params.setDeveloperAPIKey(client.auth.DeveloperApiKey)

	return params
}
