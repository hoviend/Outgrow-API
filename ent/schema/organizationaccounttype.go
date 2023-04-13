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

// OrganizationAccountType holds the schema definition for the OrganizationAccountType entity.
type OrganizationAccountType struct {
	ent.Schema
}

// Fields of the OrganizationAccountType.
func (OrganizationAccountType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("organization_id", uuid.UUID{}),
		field.String("name").NotEmpty(),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the OrganizationAccountType.
func (OrganizationAccountType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("categories", OrganizationAccountCategory.Type),
		edge.From("organization", Organization.Type).
			Ref("accountTypes").
			Field("organization_id").
			Required().
			Unique(),
	}
}

// Indexes of the OrganizationAccountType.
func (OrganizationAccountType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("organization_id", "name").Unique(),
	}
}
