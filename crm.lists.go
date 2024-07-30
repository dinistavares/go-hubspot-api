package hubspot

import (
	"fmt"
)

// Lists service
type ListsService struct {
	service
}

type ListResponse struct {
	List *List `json:"list,omitempty"`
}

type ListsResponse struct {
	Lists   *[]List `json:"lists,omitempty"`
	Total   int     `json:"total,omitempty"`
	Offset  int     `json:"offset,omitempty"`
	HasMore bool    `json:"hasMore,omitempty"`
}

type List struct {
	ProcessingType   string            `json:"processingType,omitempty"`
	ObjectTypeID     string            `json:"objectTypeId,omitempty"`
	UpdatedByID      string            `json:"updatedById,omitempty"`
	FiltersUpdatedAt string            `json:"filtersUpdatedAt,omitempty"`
	ListID           string            `json:"listId,omitempty"`
	CreatedAt        string            `json:"createdAt,omitempty"`
	ProcessingStatus string            `json:"processingStatus,omitempty"`
	DeletedAt        string            `json:"deletedAt,omitempty"`
	ListVersion      int               `json:"listVersion,omitempty"`
	Size             int               `json:"size,omitempty"`
	Name             string            `json:"name,omitempty"`
	CreatedByID      string            `json:"createdById,omitempty"`
	UpdatedAt        string            `json:"updatedAt,omitempty"`
	FilterBranch     *ListFilterBranch `json:"filterBranch,omitempty"`
}

type ListFilterBranch struct {
	FilterBranchType     string         `json:"filterBranchType,omitempty"`
	FilterBranchOperator string         `json:"filterBranchOperator,omitempty"`
	FilterBranches       *[]interface{} `json:"filterBranches,omitempty"`
	Filters              *[]interface{} `json:"filters,omitempty"`
}

type ListMembershipOfResults struct {
	Results *[]ListMembershipOfResult `json:"results,omitempty"`
}

type ListMembershipOfResult struct {
	FirstAddedTimestamp string `json:"firstAddedTimestamp,omitempty"`
	LastAddedTimestamp  string `json:"lastAddedTimestamp,omitempty"`
	ListID              string `json:"listId,omitempty"`
	ListVersion         int    `json:"listVersion,omitempty"`
}

type SearchListsBody struct {
	Count                int                  `json:"count,omitempty"`
	ProcessingTypes      []ListProcessingType `json:"processingTypes,omitempty"`
	AdditionalProperties []string             `json:"additionalProperties,omitempty"`
	Sort                 string               `json:"sort,omitempty"`
	ListIds              []string             `json:"listIds,omitempty"`
	Offset               int                  `json:"offset,omitempty"`
	Query                string               `json:"query,omitempty"`
}

type PutRecordMembershipResponse struct {
	RecordIdsAdded   []string `json:"recordsIdsAdded,omitempty"`
	RecordIdsRemoved []string `json:"recordIdsRemoved,omitempty"`
}

type ListProcessingType string

const (
	ListProcessingTypeDynamic  ListProcessingType = "DYNAMIC"
	ListProcessingTypeManual   ListProcessingType = "MANUAL"
	ListProcessingTypeSnapshop ListProcessingType = "SNAPSHOT"
)

// Get list by id
func (service *ListsService) Get(id string, opts QueryValues) (*ListResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/lists/%s", *service.revision, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(ListResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Get multiple lists by ID
func (service *ListsService) GetMulitple(ids []string, opts QueryValues) (*ListsResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/lists", *service.revision)

	for i, id := range ids {
		if i == 0 {
			_url += "?listIds=" + id
		} else {
			_url += "&listIds=" + id
		}
	}

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(ListsResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Search lists
func (service *ListsService) SearchLists(body *SearchListsBody) (*ListsResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/lists/search", *service.revision)

	req, _ := service.client.NewRequest("POST", _url, nil, body)

	data := new(ListsResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Get lists record is member of
func (service *ListsService) GetListMembershipOf(objectTypeID ObjectTypeID, id string) (*ListMembershipOfResults, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/lists/records/%s/%s/memberships", *service.revision, objectTypeID, id)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	data := new(ListMembershipOfResults)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Add record to list
func (service *ListsService) AddRecordsToList(listID string, recordIDs []string) (*PutRecordMembershipResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/lists/%s/memberships/add", *service.revision, listID)

	req, _ := service.client.NewRequest("PUT", _url, nil, recordIDs)

	data := new(PutRecordMembershipResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}

// Remove record from list
func (service *ListsService) RemoveRecordsFromList(listID string, recordIDs []string) (*PutRecordMembershipResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/lists/%s/memberships/remove", *service.revision, listID)

	req, _ := service.client.NewRequest("PUT", _url, nil, recordIDs)

	data := new(PutRecordMembershipResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
