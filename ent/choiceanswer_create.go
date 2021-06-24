// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"faves4/ent/choice"
	"faves4/ent/choiceanswer"
	"faves4/ent/user"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChoiceAnswerCreate is the builder for creating a ChoiceAnswer entity.
type ChoiceAnswerCreate struct {
	config
	mutation *ChoiceAnswerMutation
	hooks    []Hook
}

// SetCratedAt sets the "crated_at" field.
func (cac *ChoiceAnswerCreate) SetCratedAt(t time.Time) *ChoiceAnswerCreate {
	cac.mutation.SetCratedAt(t)
	return cac
}

// SetNillableCratedAt sets the "crated_at" field if the given value is not nil.
func (cac *ChoiceAnswerCreate) SetNillableCratedAt(t *time.Time) *ChoiceAnswerCreate {
	if t != nil {
		cac.SetCratedAt(*t)
	}
	return cac
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cac *ChoiceAnswerCreate) SetOwnerID(id string) *ChoiceAnswerCreate {
	cac.mutation.SetOwnerID(id)
	return cac
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cac *ChoiceAnswerCreate) SetNillableOwnerID(id *string) *ChoiceAnswerCreate {
	if id != nil {
		cac = cac.SetOwnerID(*id)
	}
	return cac
}

// SetOwner sets the "owner" edge to the User entity.
func (cac *ChoiceAnswerCreate) SetOwner(u *User) *ChoiceAnswerCreate {
	return cac.SetOwnerID(u.ID)
}

// SetParentID sets the "parent" edge to the Choice entity by ID.
func (cac *ChoiceAnswerCreate) SetParentID(id int) *ChoiceAnswerCreate {
	cac.mutation.SetParentID(id)
	return cac
}

// SetNillableParentID sets the "parent" edge to the Choice entity by ID if the given value is not nil.
func (cac *ChoiceAnswerCreate) SetNillableParentID(id *int) *ChoiceAnswerCreate {
	if id != nil {
		cac = cac.SetParentID(*id)
	}
	return cac
}

// SetParent sets the "parent" edge to the Choice entity.
func (cac *ChoiceAnswerCreate) SetParent(c *Choice) *ChoiceAnswerCreate {
	return cac.SetParentID(c.ID)
}

// Mutation returns the ChoiceAnswerMutation object of the builder.
func (cac *ChoiceAnswerCreate) Mutation() *ChoiceAnswerMutation {
	return cac.mutation
}

// Save creates the ChoiceAnswer in the database.
func (cac *ChoiceAnswerCreate) Save(ctx context.Context) (*ChoiceAnswer, error) {
	var (
		err  error
		node *ChoiceAnswer
	)
	cac.defaults()
	if len(cac.hooks) == 0 {
		if err = cac.check(); err != nil {
			return nil, err
		}
		node, err = cac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChoiceAnswerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cac.check(); err != nil {
				return nil, err
			}
			cac.mutation = mutation
			node, err = cac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cac.hooks) - 1; i >= 0; i-- {
			mut = cac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cac *ChoiceAnswerCreate) SaveX(ctx context.Context) *ChoiceAnswer {
	v, err := cac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cac *ChoiceAnswerCreate) defaults() {
	if _, ok := cac.mutation.CratedAt(); !ok {
		v := choiceanswer.DefaultCratedAt()
		cac.mutation.SetCratedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cac *ChoiceAnswerCreate) check() error {
	if _, ok := cac.mutation.CratedAt(); !ok {
		return &ValidationError{Name: "crated_at", err: errors.New("ent: missing required field \"crated_at\"")}
	}
	return nil
}

func (cac *ChoiceAnswerCreate) sqlSave(ctx context.Context) (*ChoiceAnswer, error) {
	_node, _spec := cac.createSpec()
	if err := sqlgraph.CreateNode(ctx, cac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cac *ChoiceAnswerCreate) createSpec() (*ChoiceAnswer, *sqlgraph.CreateSpec) {
	var (
		_node = &ChoiceAnswer{config: cac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: choiceanswer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: choiceanswer.FieldID,
			},
		}
	)
	if value, ok := cac.mutation.CratedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: choiceanswer.FieldCratedAt,
		})
		_node.CratedAt = value
	}
	if nodes := cac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choiceanswer.OwnerTable,
			Columns: []string{choiceanswer.OwnerColumn},
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
		_node.user_choiceanswers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cac.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choiceanswer.ParentTable,
			Columns: []string{choiceanswer.ParentColumn},
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
		_node.choice_choiceanswers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChoiceAnswerCreateBulk is the builder for creating many ChoiceAnswer entities in bulk.
type ChoiceAnswerCreateBulk struct {
	config
	builders []*ChoiceAnswerCreate
}

// Save creates the ChoiceAnswer entities in the database.
func (cacb *ChoiceAnswerCreateBulk) Save(ctx context.Context) ([]*ChoiceAnswer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cacb.builders))
	nodes := make([]*ChoiceAnswer, len(cacb.builders))
	mutators := make([]Mutator, len(cacb.builders))
	for i := range cacb.builders {
		func(i int, root context.Context) {
			builder := cacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChoiceAnswerMutation)
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
					_, err = mutators[i+1].Mutate(root, cacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cacb *ChoiceAnswerCreateBulk) SaveX(ctx context.Context) []*ChoiceAnswer {
	v, err := cacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
