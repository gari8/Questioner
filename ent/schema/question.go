package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field {
		field.String("id").Unique().NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("content").NotEmpty(),
		field.String("text_after_answered").Optional(),
		field.Bool("enabled").Default(true),
		field.String("answer_type").Optional(),
		field.Time("published_at").Optional(),
		field.Time("finished_at").Optional(),
		field.Time("crated_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("questions").Unique(),
		edge.To("answers", Answer.Type),
		edge.To("choices", Choice.Type),
	}
}
