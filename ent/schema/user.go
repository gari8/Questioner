package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().NotEmpty(),
		field.String("username").NotEmpty(),
		field.String("icon").Default(""),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.String("description").Default(""),
		field.Time("crated_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
		edge.To("answers", Answer.Type),
		edge.To("choiceanswers", ChoiceAnswer.Type),
	}
}
