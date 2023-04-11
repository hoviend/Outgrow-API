package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrganizationAccountCategory holds the schema definition for the OrganizationAccountCategory entity.
type OrganizationAccountCategory struct {
	ent.Schema
}

// Fields of the OrganizationAccountCategory.
func (OrganizationAccountCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_type_id"),
		field.String("name").NotEmpty(),
		field.Text("description").Optional(),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the OrganizationAccountCategory.
func (OrganizationAccountCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", OrganizationAccount.Type),
		edge.From("type", OrganizationAccountType.Type).
			Ref("categories").
			Field("account_type_id").
			Required().
			Unique(),
	}
}
