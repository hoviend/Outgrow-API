package schema

import (
	"outgrow/enum"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("event_id", uuid.UUID{}),
		field.Time("transaction_date").SchemaType(map[string]string{dialect.Postgres: "DATE"}),
		field.Enum("transaction_type").GoType(enum.TransactionType("")),
		field.Int("account_id"),
		field.Float("amount").Default(0),
		field.Text("notes").Optional(),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", OrganizationAccount.Type).
			Ref("transactions").
			Field("account_id").
			Required().
			Unique(),
		edge.From("event", Event.Type).
			Ref("transactions").
			Field("event_id").
			Required().
			Unique(),
	}
}
