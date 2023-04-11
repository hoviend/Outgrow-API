package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("organization_id", uuid.UUID{}),
		field.Int("event_type_id"),
		field.Float("amount").Default(0),
		field.Text("notes").Optional(),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type),
		edge.From("type", OrganizationEventType.Type).
			Ref("events").
			Field("event_type_id").
			Required().
			Unique(),
		edge.From("organization", Organization.Type).
			Ref("events").
			Field("organization_id").
			Required().
			Unique(),
	}
}
