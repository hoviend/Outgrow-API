package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MasterAccountType holds the schema definition for the MasterAccountType entity.
type MasterAccountType struct {
	ent.Schema
}

// Fields of the MasterAccountType.
func (MasterAccountType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")).Immutable(),
	}
}

// Edges of the MasterAccountType.
func (MasterAccountType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("categories", MasterAccountCategory.Type),
	}
}
