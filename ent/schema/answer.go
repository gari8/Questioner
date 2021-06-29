package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return []ent.Field{
		field.String("answer_type").NotEmpty(),
		field.String("content").NotEmpty(),
		field.Time("crated_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("answers").Unique(),
		edge.From("parent", Question.Type).Ref("answers").Unique(),
	}
}
