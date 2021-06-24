// Code generated by entc, DO NOT EDIT.

package ent

import (
	"faves4/ent/answer"
	"faves4/ent/question"
	"faves4/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Answer is the model entity for the Answer schema.
type Answer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AnswerType holds the value of the "answer_type" field.
	AnswerType string `json:"answer_type,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CratedAt holds the value of the "crated_at" field.
	CratedAt time.Time `json:"crated_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnswerQuery when eager-loading is set.
	Edges            AnswerEdges `json:"edges"`
	question_answers *string
	user_answers     *string
}

// AnswerEdges holds the relations/edges for other nodes in the graph.
type AnswerEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Question `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerEdges) ParentOrErr() (*Question, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: question.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Answer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case answer.FieldID:
			values[i] = new(sql.NullInt64)
		case answer.FieldAnswerType, answer.FieldContent:
			values[i] = new(sql.NullString)
		case answer.FieldCratedAt, answer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case answer.ForeignKeys[0]: // question_answers
			values[i] = new(sql.NullString)
		case answer.ForeignKeys[1]: // user_answers
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Answer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Answer fields.
func (a *Answer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case answer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case answer.FieldAnswerType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field answer_type", values[i])
			} else if value.Valid {
				a.AnswerType = value.String
			}
		case answer.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				a.Content = value.String
			}
		case answer.FieldCratedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field crated_at", values[i])
			} else if value.Valid {
				a.CratedAt = value.Time
			}
		case answer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case answer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field question_answers", values[i])
			} else if value.Valid {
				a.question_answers = new(string)
				*a.question_answers = value.String
			}
		case answer.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_answers", values[i])
			} else if value.Valid {
				a.user_answers = new(string)
				*a.user_answers = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Answer entity.
func (a *Answer) QueryOwner() *UserQuery {
	return (&AnswerClient{config: a.config}).QueryOwner(a)
}

// QueryParent queries the "parent" edge of the Answer entity.
func (a *Answer) QueryParent() *QuestionQuery {
	return (&AnswerClient{config: a.config}).QueryParent(a)
}

// Update returns a builder for updating this Answer.
// Note that you need to call Answer.Unwrap() before calling this method if this Answer
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Answer) Update() *AnswerUpdateOne {
	return (&AnswerClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Answer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Answer) Unwrap() *Answer {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Answer is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Answer) String() string {
	var builder strings.Builder
	builder.WriteString("Answer(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", answer_type=")
	builder.WriteString(a.AnswerType)
	builder.WriteString(", content=")
	builder.WriteString(a.Content)
	builder.WriteString(", crated_at=")
	builder.WriteString(a.CratedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Answers is a parsable slice of Answer.
type Answers []*Answer

func (a Answers) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
