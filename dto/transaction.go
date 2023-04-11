package dto

import (
	"outgrow/ent/schema"

	"github.com/google/uuid"
)

type GetTransactionsOption struct {
	Paginate       *PaginateParam
	OrganizationID uuid.UUID
}

type CreateOrganizationEventPayload struct {
	EventTypeID    int     `json:"event_type_id" validate:"required"`
	Amount         float64 `json:"amount" validate:"required"`
	OrganizationID uuid.UUID
}

type CreateOrganizationEventTypePayload struct {
	Name           string              `json:"name" validate:"required"`
	Rules          []schema.EventRules `json:"rules" validate:"required, dive"`
	Description    string              `json:"description"`
	OrganizationID uuid.UUID
}