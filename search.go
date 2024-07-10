package hubspot

type GenericSearchResults struct {
	Paging  *SearchPaging    `json:"paging,omitempty"`
	Total   int       `json:"total,omitempty"`
}

type SearchPaging struct {
	Next *SearchNext `json:"next,omitempty"`
}

type SearchNext struct {
	After string `json:"after,omitempty"`
}

type SearchRequest struct {
	Query        string               `json:"query,omitempty"`
	Limit        int                  `json:"limit,omitempty"`
	After        string               `json:"after,omitempty"`
	Sorts        *[]string            `json:"sorts,omitempty"`
	Properties   *[]string            `json:"properties,omitempty"`
	FilterGroups *[]SearchFilterGroup `json:"filterGroups,omitempty"`
}

type SearchFilterGroup struct {
	Filters *[]SearchFilter `json:"filters,omitempty"`
}

type SearchFilter struct {
	HighValue    string          `json:"highValue,omitempty"`
	PropertyName string          `json:"propertyName,omitempty"`
	Value        string          `json:"value,omitempty"`
	Operator     SearchOperator `json:"operator,omitempty"`
	Values       *[]string       `json:"values,omitempty"`
}

type SearchOperator string

const (
	SearchOperatorEquals              = "EQ"
	SearchOperatorNotEqual            = "NEQ"
	SearchOperatorLessThan            = "LT"
	SearchOperatorLessThanOrEqual     = "LTE"
	SearchOperatorGreaterThan         = "GT"
	SearchOperatorGreaterThanOrEqual  = "GTE"
	SearchOperatorBetween             = "BETWEEN"
	SearchOperatorIncluded            = "IN"
	SearchOperatorNotIncluded         = "NOT_IN"
	SearchOperatorHasProperty         = "HAS_PROPERTY"
	SearchOperatorDoesNotHaveProperty = "NOT_HAS_PROPERTY"
	SearchOperatorContainsToken       = "CONTAINS_TOKEN"
	SearchOperatorDoesNotContainToken = "NOT_CONTAINS_TOKEN"
)

func (r *SearchRequest) CreateNewFilterGroup() *SearchFilterGroup {
	filterGroup := SearchFilterGroup{}

	if r.FilterGroups == nil {
		r.FilterGroups = &[]SearchFilterGroup{}
	}

	*r.FilterGroups = append(*r.FilterGroups, filterGroup)

	return &(*r.FilterGroups)[len(*r.FilterGroups)-1]
}

func (f *SearchFilterGroup) AddFilter(filter SearchFilter) {
	if f.Filters == nil {
		f.Filters = &[]SearchFilter{}
	}

	*f.Filters = append(*f.Filters, filter)
}
