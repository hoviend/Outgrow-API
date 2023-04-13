package schema

import (
	"outgrow/enum"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

type EventRules struct {
	AccountID       int                  `json:"account_id" validate:"required"`
	TransactionType enum.TransactionType `json:"transaction_type" validate:"oneof=DEBIT CREDIT"`
	Amount          float64              `json:"amount" validate:"required,er_amount"`
	RuleType        enum.EventRuleType   `json:"rule_type" validate:"oneof=PERCENTAGE FLAT"`
}

// MasterEventType holds the schema definition for the MasterEventType entity.
type MasterEventType struct {
	ent.Schema
}

// Fields of the MasterEventType.
func (MasterEventType) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.String("description"),
		field.JSON("rules", []EventRules{}),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the MasterEventType.
func (MasterEventType) Edges() []ent.Edge {
	return nil
}
