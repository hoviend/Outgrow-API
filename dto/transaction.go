package dto

import (
	"outgrow/ent/schema"

	"github.com/google/uuid"
)

type GetTransactionsOption struct {
	Paginate        *PaginateParam
	OrganizationID  uuid.UUID
	JoinAccount     bool
	JoinEvent       bool
	AccountIDFilter int
}

type CreateOrganizationEventPayload struct {
	EventTypeID    int     `json:"event_type_id" validate:"required"`
	Amount         float64 `json:"amount" validate:"required"`
	Notes          string  `json:"notes"`
	OrganizationID uuid.UUID
}

type CreateOrganizationEventTypePayload struct {
	Name           string              `json:"name" validate:"required"`
	Rules          []schema.EventRules `json:"rules" validate:"gt=0,dive"`
	Description    string              `json:"description"`
	OrganizationID uuid.UUID
}
