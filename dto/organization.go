package dto

import (
	"outgrow/ent"
	"outgrow/enum"
	"time"

	"github.com/google/uuid"
)

type OrganizationEventTypeSelectionResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetOrganizationOption struct {
	ID         uuid.UUID
	WithEvents bool
}

type GetOrganizationTransactionResponse struct {
	ID              uuid.UUID            `json:"id"`
	EventID         uuid.UUID            `json:"event_id"`
	EventName       string               `json:"event_name"`
	TransactionDate time.Time            `json:"transaction_date"`
	TransactionType enum.TransactionType `json:"transaction_type"`
	AccountID       int                  `json:"account_id"`
	AccountName     string               `json:"account_name"`
	Amount          float64              `json:"amount"`
	Notes           string               `json:"notes"`
}

func (r *GetOrganizationTransactionResponse) BuildResponse(trx *ent.Transaction) {
	r.ID = trx.ID
	r.EventID = trx.EventID
	r.TransactionDate = trx.TransactionDate
	r.TransactionType = trx.TransactionType
	r.AccountID = trx.AccountID
	r.Amount = trx.Amount
	r.Notes = trx.Notes

	r.EventName = trx.Edges.Event.Edges.Type.Name
	r.AccountName = trx.Edges.Account.Name
}
