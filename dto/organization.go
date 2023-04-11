package dto

import "github.com/google/uuid"

type OrganizationEventTypeSelectionResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetOrganizationOption struct {
	ID         uuid.UUID
	WithEvents bool
}
