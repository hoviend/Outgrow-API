package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrganizationAccount holds the schema definition for the OrganizationAccount entity.
type OrganizationAccount struct {
	ent.Schema
}

// Fields of the OrganizationAccount.
func (OrganizationAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").Positive(),
		field.String("name").NotEmpty(),
		field.String("code").NotEmpty().Unique(),
		field.Float("balance").Default(0),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the OrganizationAccount.
func (OrganizationAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type),
		edge.From("accCategory", OrganizationAccountCategory.Type).
			Ref("accounts").
			Field("category_id").
			Required().
			Unique(),
	}
}
