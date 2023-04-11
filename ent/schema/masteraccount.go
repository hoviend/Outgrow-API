package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MasterAccount holds the schema definition for the MasterAccount entity.
type MasterAccount struct {
	ent.Schema
}

// Fields of the MasterAccount.
func (MasterAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").Positive(),
		field.String("name").NotEmpty(),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the MasterAccount.
func (MasterAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("accCategory", MasterAccountCategory.Type).
			Ref("accounts").
			Field("category_id").
			Required().
			Unique(),
	}
}
