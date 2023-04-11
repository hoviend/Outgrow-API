package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MasterAccountCategory holds the schema definition for the MasterAccountCategory entity.
type MasterAccountCategory struct {
	ent.Schema
}

// Fields of the MasterAccountCategory.
func (MasterAccountCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_type_id"),
		field.String("name").NotEmpty(),
		field.Text("description").Optional(),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the MasterAccountCategory.
func (MasterAccountCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", MasterAccount.Type),
		edge.From("accType", MasterAccountType.Type).
			Ref("categories").
			Field("account_type_id").
			Required().
			Unique(),
	}
}
