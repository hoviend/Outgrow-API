package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// OrganizationEventType holds the schema definition for the OrganizationEventType entity.
type OrganizationEventType struct {
	ent.Schema
}

// Fields of the OrganizationEventType.
func (OrganizationEventType) Fields() []ent.Field {

	return []ent.Field{
		field.UUID("organization_id", uuid.UUID{}),
		field.String("name").NotEmpty(),
		field.String("description"),
		field.JSON("rules", []EventRules{}),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the OrganizationEventType.
func (OrganizationEventType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type),
		edge.From("organization", Organization.Type).
			Ref("eventTypes").
			Field("organization_id").
			Required().
			Unique(),
	}
}

// Indexes of the OrganizationEventType.
func (OrganizationEventType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("organization_id", "name").Unique(),
	}
}
