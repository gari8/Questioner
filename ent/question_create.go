// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"faves4/ent/answer"
	"faves4/ent/choice"
	"faves4/ent/question"
	"faves4/ent/user"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QuestionCreate is the builder for creating a Question entity.
type QuestionCreate struct {
	config
	mutation *QuestionMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (qc *QuestionCreate) SetTitle(s string) *QuestionCreate {
	qc.mutation.SetTitle(s)
	return qc
}

// SetContent sets the "content" field.
func (qc *QuestionCreate) SetContent(s string) *QuestionCreate {
	qc.mutation.SetContent(s)
	return qc
}

// SetTextAfterAnswered sets the "text_after_answered" field.
func (qc *QuestionCreate) SetTextAfterAnswered(s string) *QuestionCreate {
	qc.mutation.SetTextAfterAnswered(s)
	return qc
}

// SetNillableTextAfterAnswered sets the "text_after_answered" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableTextAfterAnswered(s *string) *QuestionCreate {
	if s != nil {
		qc.SetTextAfterAnswered(*s)
	}
	return qc
}

// SetEnabled sets the "enabled" field.
func (qc *QuestionCreate) SetEnabled(b bool) *QuestionCreate {
	qc.mutation.SetEnabled(b)
	return qc
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableEnabled(b *bool) *QuestionCreate {
	if b != nil {
		qc.SetEnabled(*b)
	}
	return qc
}

// SetAnswerType sets the "answer_type" field.
func (qc *QuestionCreate) SetAnswerType(s string) *QuestionCreate {
	qc.mutation.SetAnswerType(s)
	return qc
}

// SetPublishedAt sets the "published_at" field.
func (qc *QuestionCreate) SetPublishedAt(t time.Time) *QuestionCreate {
	qc.mutation.SetPublishedAt(t)
	return qc
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (qc *QuestionCreate) SetNillablePublishedAt(t *time.Time) *QuestionCreate {
	if t != nil {
		qc.SetPublishedAt(*t)
	}
	return qc
}

// SetFinishedAt sets the "finished_at" field.
func (qc *QuestionCreate) SetFinishedAt(t time.Time) *QuestionCreate {
	qc.mutation.SetFinishedAt(t)
	return qc
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableFinishedAt(t *time.Time) *QuestionCreate {
	if t != nil {
		qc.SetFinishedAt(*t)
	}
	return qc
}

// SetCratedAt sets the "crated_at" field.
func (qc *QuestionCreate) SetCratedAt(t time.Time) *QuestionCreate {
	qc.mutation.SetCratedAt(t)
	return qc
}

// SetNillableCratedAt sets the "crated_at" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableCratedAt(t *time.Time) *QuestionCreate {
	if t != nil {
		qc.SetCratedAt(*t)
	}
	return qc
}

// SetUpdatedAt sets the "updated_at" field.
func (qc *QuestionCreate) SetUpdatedAt(t time.Time) *QuestionCreate {
	qc.mutation.SetUpdatedAt(t)
	return qc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableUpdatedAt(t *time.Time) *QuestionCreate {
	if t != nil {
		qc.SetUpdatedAt(*t)
	}
	return qc
}

// SetID sets the "id" field.
func (qc *QuestionCreate) SetID(s string) *QuestionCreate {
	qc.mutation.SetID(s)
	return qc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (qc *QuestionCreate) SetOwnerID(id string) *QuestionCreate {
	qc.mutation.SetOwnerID(id)
	return qc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (qc *QuestionCreate) SetNillableOwnerID(id *string) *QuestionCreate {
	if id != nil {
		qc = qc.SetOwnerID(*id)
	}
	return qc
}

// SetOwner sets the "owner" edge to the User entity.
func (qc *QuestionCreate) SetOwner(u *User) *QuestionCreate {
	return qc.SetOwnerID(u.ID)
}

// AddAnswerIDs adds the "answers" edge to the Answer entity by IDs.
func (qc *QuestionCreate) AddAnswerIDs(ids ...int) *QuestionCreate {
	qc.mutation.AddAnswerIDs(ids...)
	return qc
}

// AddAnswers adds the "answers" edges to the Answer entity.
func (qc *QuestionCreate) AddAnswers(a ...*Answer) *QuestionCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qc.AddAnswerIDs(ids...)
}

// AddChoiceIDs adds the "choices" edge to the Choice entity by IDs.
func (qc *QuestionCreate) AddChoiceIDs(ids ...int) *QuestionCreate {
	qc.mutation.AddChoiceIDs(ids...)
	return qc
}

// AddChoices adds the "choices" edges to the Choice entity.
func (qc *QuestionCreate) AddChoices(c ...*Choice) *QuestionCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return qc.AddChoiceIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (qc *QuestionCreate) Mutation() *QuestionMutation {
	return qc.mutation
}

// Save creates the Question in the database.
func (qc *QuestionCreate) Save(ctx context.Context) (*Question, error) {
	var (
		err  error
		node *Question
	)
	qc.defaults()
	if len(qc.hooks) == 0 {
		if err = qc.check(); err != nil {
			return nil, err
		}
		node, err = qc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qc.check(); err != nil {
				return nil, err
			}
			qc.mutation = mutation
			node, err = qc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qc.hooks) - 1; i >= 0; i-- {
			mut = qc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QuestionCreate) SaveX(ctx context.Context) *Question {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (qc *QuestionCreate) defaults() {
	if _, ok := qc.mutation.TextAfterAnswered(); !ok {
		v := question.DefaultTextAfterAnswered
		qc.mutation.SetTextAfterAnswered(v)
	}
	if _, ok := qc.mutation.Enabled(); !ok {
		v := question.DefaultEnabled
		qc.mutation.SetEnabled(v)
	}
	if _, ok := qc.mutation.CratedAt(); !ok {
		v := question.DefaultCratedAt()
		qc.mutation.SetCratedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qc *QuestionCreate) check() error {
	if _, ok := qc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if v, ok := qc.mutation.Title(); ok {
		if err := question.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if _, ok := qc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New("ent: missing required field \"content\"")}
	}
	if v, ok := qc.mutation.Content(); ok {
		if err := question.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf("ent: validator failed for field \"content\": %w", err)}
		}
	}
	if _, ok := qc.mutation.TextAfterAnswered(); !ok {
		return &ValidationError{Name: "text_after_answered", err: errors.New("ent: missing required field \"text_after_answered\"")}
	}
	if _, ok := qc.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New("ent: missing required field \"enabled\"")}
	}
	if _, ok := qc.mutation.AnswerType(); !ok {
		return &ValidationError{Name: "answer_type", err: errors.New("ent: missing required field \"answer_type\"")}
	}
	if v, ok := qc.mutation.AnswerType(); ok {
		if err := question.AnswerTypeValidator(v); err != nil {
			return &ValidationError{Name: "answer_type", err: fmt.Errorf("ent: validator failed for field \"answer_type\": %w", err)}
		}
	}
	if _, ok := qc.mutation.CratedAt(); !ok {
		return &ValidationError{Name: "crated_at", err: errors.New("ent: missing required field \"crated_at\"")}
	}
	if v, ok := qc.mutation.ID(); ok {
		if err := question.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (qc *QuestionCreate) sqlSave(ctx context.Context) (*Question, error) {
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (qc *QuestionCreate) createSpec() (*Question, *sqlgraph.CreateSpec) {
	var (
		_node = &Question{config: qc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: question.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: question.FieldID,
			},
		}
	)
	if id, ok := qc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := qc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := qc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := qc.mutation.TextAfterAnswered(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldTextAfterAnswered,
		})
		_node.TextAfterAnswered = value
	}
	if value, ok := qc.mutation.Enabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: question.FieldEnabled,
		})
		_node.Enabled = value
	}
	if value, ok := qc.mutation.AnswerType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldAnswerType,
		})
		_node.AnswerType = value
	}
	if value, ok := qc.mutation.PublishedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: question.FieldPublishedAt,
		})
		_node.PublishedAt = value
	}
	if value, ok := qc.mutation.FinishedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: question.FieldFinishedAt,
		})
		_node.FinishedAt = value
	}
	if value, ok := qc.mutation.CratedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: question.FieldCratedAt,
		})
		_node.CratedAt = value
	}
	if value, ok := qc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: question.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := qc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.OwnerTable,
			Columns: []string{question.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_questions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.ChoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.ChoicesTable,
			Columns: []string{question.ChoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: choice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// QuestionCreateBulk is the builder for creating many Question entities in bulk.
type QuestionCreateBulk struct {
	config
	builders []*QuestionCreate
}

// Save creates the Question entities in the database.
func (qcb *QuestionCreateBulk) Save(ctx context.Context) ([]*Question, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Question, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QuestionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QuestionCreateBulk) SaveX(ctx context.Context) []*Question {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
