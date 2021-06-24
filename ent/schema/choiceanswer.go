package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// ChoiceAnswer holds the schema definition for the ChoiceAnswer entity.
type ChoiceAnswer struct {
	ent.Schema
}

// Fields of the ChoiceAnswer.
func (ChoiceAnswer) Fields() []ent.Field {
	return []ent.Field {
		field.Time("crated_at").Default(time.Now).Immutable(),
	}
}

// Edges of the ChoiceAnswer.
func (ChoiceAnswer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("choiceanswers").Unique(),
		edge.From("parent", Choice.Type).Ref("choiceanswers").Unique(),
	}
}
