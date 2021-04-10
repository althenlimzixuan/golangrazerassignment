// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hello/ent/comprimiseemail"
	"hello/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ComprimiseEmailUpdate is the builder for updating ComprimiseEmail entities.
type ComprimiseEmailUpdate struct {
	config
	hooks    []Hook
	mutation *ComprimiseEmailMutation
}

// Where adds a new predicate for the ComprimiseEmailUpdate builder.
func (ceu *ComprimiseEmailUpdate) Where(ps ...predicate.ComprimiseEmail) *ComprimiseEmailUpdate {
	ceu.mutation.predicates = append(ceu.mutation.predicates, ps...)
	return ceu
}

// SetEmail sets the "Email" field.
func (ceu *ComprimiseEmailUpdate) SetEmail(s string) *ComprimiseEmailUpdate {
	ceu.mutation.SetEmail(s)
	return ceu
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (ceu *ComprimiseEmailUpdate) SetNillableEmail(s *string) *ComprimiseEmailUpdate {
	if s != nil {
		ceu.SetEmail(*s)
	}
	return ceu
}

// Mutation returns the ComprimiseEmailMutation object of the builder.
func (ceu *ComprimiseEmailUpdate) Mutation() *ComprimiseEmailMutation {
	return ceu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ceu *ComprimiseEmailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ceu.hooks) == 0 {
		affected, err = ceu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ComprimiseEmailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceu.mutation = mutation
			affected, err = ceu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ceu.hooks) - 1; i >= 0; i-- {
			mut = ceu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceu *ComprimiseEmailUpdate) SaveX(ctx context.Context) int {
	affected, err := ceu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ceu *ComprimiseEmailUpdate) Exec(ctx context.Context) error {
	_, err := ceu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceu *ComprimiseEmailUpdate) ExecX(ctx context.Context) {
	if err := ceu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ceu *ComprimiseEmailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comprimiseemail.Table,
			Columns: comprimiseemail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comprimiseemail.FieldID,
			},
		},
	}
	if ps := ceu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comprimiseemail.FieldEmail,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ceu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comprimiseemail.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ComprimiseEmailUpdateOne is the builder for updating a single ComprimiseEmail entity.
type ComprimiseEmailUpdateOne struct {
	config
	hooks    []Hook
	mutation *ComprimiseEmailMutation
}

// SetEmail sets the "Email" field.
func (ceuo *ComprimiseEmailUpdateOne) SetEmail(s string) *ComprimiseEmailUpdateOne {
	ceuo.mutation.SetEmail(s)
	return ceuo
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (ceuo *ComprimiseEmailUpdateOne) SetNillableEmail(s *string) *ComprimiseEmailUpdateOne {
	if s != nil {
		ceuo.SetEmail(*s)
	}
	return ceuo
}

// Mutation returns the ComprimiseEmailMutation object of the builder.
func (ceuo *ComprimiseEmailUpdateOne) Mutation() *ComprimiseEmailMutation {
	return ceuo.mutation
}

// Save executes the query and returns the updated ComprimiseEmail entity.
func (ceuo *ComprimiseEmailUpdateOne) Save(ctx context.Context) (*ComprimiseEmail, error) {
	var (
		err  error
		node *ComprimiseEmail
	)
	if len(ceuo.hooks) == 0 {
		node, err = ceuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ComprimiseEmailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceuo.mutation = mutation
			node, err = ceuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ceuo.hooks) - 1; i >= 0; i-- {
			mut = ceuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceuo *ComprimiseEmailUpdateOne) SaveX(ctx context.Context) *ComprimiseEmail {
	node, err := ceuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ceuo *ComprimiseEmailUpdateOne) Exec(ctx context.Context) error {
	_, err := ceuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceuo *ComprimiseEmailUpdateOne) ExecX(ctx context.Context) {
	if err := ceuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ceuo *ComprimiseEmailUpdateOne) sqlSave(ctx context.Context) (_node *ComprimiseEmail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comprimiseemail.Table,
			Columns: comprimiseemail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comprimiseemail.FieldID,
			},
		},
	}
	id, ok := ceuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ComprimiseEmail.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ceuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comprimiseemail.FieldEmail,
		})
	}
	_node = &ComprimiseEmail{config: ceuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ceuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comprimiseemail.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
