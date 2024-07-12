package hubspot

type GenericCreateBody struct {
	Associations *[]Association `json:"associations,omitempty"`
	Properties   *Properties    `json:"properties,omitempty"`
}
