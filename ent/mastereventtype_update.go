// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"outgrow/ent/mastereventtype"
	"outgrow/ent/predicate"
	"outgrow/ent/schema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// MasterEventTypeUpdate is the builder for updating MasterEventType entities.
type MasterEventTypeUpdate struct {
	config
	hooks    []Hook
	mutation *MasterEventTypeMutation
}

// Where appends a list predicates to the MasterEventTypeUpdate builder.
func (metu *MasterEventTypeUpdate) Where(ps ...predicate.MasterEventType) *MasterEventTypeUpdate {
	metu.mutation.Where(ps...)
	return metu
}

// SetName sets the "name" field.
func (metu *MasterEventTypeUpdate) SetName(s string) *MasterEventTypeUpdate {
	metu.mutation.SetName(s)
	return metu
}

// SetDescription sets the "description" field.
func (metu *MasterEventTypeUpdate) SetDescription(s string) *MasterEventTypeUpdate {
	metu.mutation.SetDescription(s)
	return metu
}

// SetRules sets the "rules" field.
func (metu *MasterEventTypeUpdate) SetRules(sr []schema.EventRules) *MasterEventTypeUpdate {
	metu.mutation.SetRules(sr)
	return metu
}

// AppendRules appends sr to the "rules" field.
func (metu *MasterEventTypeUpdate) AppendRules(sr []schema.EventRules) *MasterEventTypeUpdate {
	metu.mutation.AppendRules(sr)
	return metu
}

// Mutation returns the MasterEventTypeMutation object of the builder.
func (metu *MasterEventTypeUpdate) Mutation() *MasterEventTypeMutation {
	return metu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (metu *MasterEventTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MasterEventTypeMutation](ctx, metu.sqlSave, metu.mutation, metu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (metu *MasterEventTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := metu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (metu *MasterEventTypeUpdate) Exec(ctx context.Context) error {
	_, err := metu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (metu *MasterEventTypeUpdate) ExecX(ctx context.Context) {
	if err := metu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (metu *MasterEventTypeUpdate) check() error {
	if v, ok := metu.mutation.Name(); ok {
		if err := mastereventtype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MasterEventType.name": %w`, err)}
		}
	}
	return nil
}

func (metu *MasterEventTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := metu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(mastereventtype.Table, mastereventtype.Columns, sqlgraph.NewFieldSpec(mastereventtype.FieldID, field.TypeInt))
	if ps := metu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := metu.mutation.Name(); ok {
		_spec.SetField(mastereventtype.FieldName, field.TypeString, value)
	}
	if value, ok := metu.mutation.Description(); ok {
		_spec.SetField(mastereventtype.FieldDescription, field.TypeString, value)
	}
	if value, ok := metu.mutation.Rules(); ok {
		_spec.SetField(mastereventtype.FieldRules, field.TypeJSON, value)
	}
	if value, ok := metu.mutation.AppendedRules(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, mastereventtype.FieldRules, value)
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, metu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mastereventtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	metu.mutation.done = true
	return n, nil
}

// MasterEventTypeUpdateOne is the builder for updating a single MasterEventType entity.
type MasterEventTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MasterEventTypeMutation
}

// SetName sets the "name" field.
func (metuo *MasterEventTypeUpdateOne) SetName(s string) *MasterEventTypeUpdateOne {
	metuo.mutation.SetName(s)
	return metuo
}

// SetDescription sets the "description" field.
func (metuo *MasterEventTypeUpdateOne) SetDescription(s string) *MasterEventTypeUpdateOne {
	metuo.mutation.SetDescription(s)
	return metuo
}

// SetRules sets the "rules" field.
func (metuo *MasterEventTypeUpdateOne) SetRules(sr []schema.EventRules) *MasterEventTypeUpdateOne {
	metuo.mutation.SetRules(sr)
	return metuo
}

// AppendRules appends sr to the "rules" field.
func (metuo *MasterEventTypeUpdateOne) AppendRules(sr []schema.EventRules) *MasterEventTypeUpdateOne {
	metuo.mutation.AppendRules(sr)
	return metuo
}

// Mutation returns the MasterEventTypeMutation object of the builder.
func (metuo *MasterEventTypeUpdateOne) Mutation() *MasterEventTypeMutation {
	return metuo.mutation
}

// Where appends a list predicates to the MasterEventTypeUpdate builder.
func (metuo *MasterEventTypeUpdateOne) Where(ps ...predicate.MasterEventType) *MasterEventTypeUpdateOne {
	metuo.mutation.Where(ps...)
	return metuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (metuo *MasterEventTypeUpdateOne) Select(field string, fields ...string) *MasterEventTypeUpdateOne {
	metuo.fields = append([]string{field}, fields...)
	return metuo
}

// Save executes the query and returns the updated MasterEventType entity.
func (metuo *MasterEventTypeUpdateOne) Save(ctx context.Context) (*MasterEventType, error) {
	return withHooks[*MasterEventType, MasterEventTypeMutation](ctx, metuo.sqlSave, metuo.mutation, metuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (metuo *MasterEventTypeUpdateOne) SaveX(ctx context.Context) *MasterEventType {
	node, err := metuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (metuo *MasterEventTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := metuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (metuo *MasterEventTypeUpdateOne) ExecX(ctx context.Context) {
	if err := metuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (metuo *MasterEventTypeUpdateOne) check() error {
	if v, ok := metuo.mutation.Name(); ok {
		if err := mastereventtype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MasterEventType.name": %w`, err)}
		}
	}
	return nil
}

func (metuo *MasterEventTypeUpdateOne) sqlSave(ctx context.Context) (_node *MasterEventType, err error) {
	if err := metuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(mastereventtype.Table, mastereventtype.Columns, sqlgraph.NewFieldSpec(mastereventtype.FieldID, field.TypeInt))
	id, ok := metuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MasterEventType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := metuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mastereventtype.FieldID)
		for _, f := range fields {
			if !mastereventtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mastereventtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := metuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := metuo.mutation.Name(); ok {
		_spec.SetField(mastereventtype.FieldName, field.TypeString, value)
	}
	if value, ok := metuo.mutation.Description(); ok {
		_spec.SetField(mastereventtype.FieldDescription, field.TypeString, value)
	}
	if value, ok := metuo.mutation.Rules(); ok {
		_spec.SetField(mastereventtype.FieldRules, field.TypeJSON, value)
	}
	if value, ok := metuo.mutation.AppendedRules(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, mastereventtype.FieldRules, value)
		})
	}
	_node = &MasterEventType{config: metuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, metuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mastereventtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	metuo.mutation.done = true
	return _node, nil
}
