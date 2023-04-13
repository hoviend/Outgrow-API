package dto

import (
	"outgrow/enum"
)

type EventSimulationPayload struct {
	EventTypeID int     `json:"event_type_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
}

type EventSimulationResponse struct {
	AccountName     string               `json:"account_name"`
	TransactionType enum.TransactionType `json:"transaction_type"`
	Amount          float64              `json:"amount"`
}

type TotalEventRuleAmounts struct {
	TotalDebitAmount      float64
	TotalCreditAmount     float64
	TotalCreditFlatAmount float64
	TotalDebitFlatAmount  float64
	TotalDebitPercentage  float64
	TotalCreditPercentage float64
}
