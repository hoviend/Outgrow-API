// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"outgrow/ent/masteraccountcategory"
	"outgrow/ent/masteraccounttype"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MasterAccountTypeCreate is the builder for creating a MasterAccountType entity.
type MasterAccountTypeCreate struct {
	config
	mutation *MasterAccountTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (matc *MasterAccountTypeCreate) SetName(s string) *MasterAccountTypeCreate {
	matc.mutation.SetName(s)
	return matc
}

// SetCreatedAt sets the "created_at" field.
func (matc *MasterAccountTypeCreate) SetCreatedAt(t time.Time) *MasterAccountTypeCreate {
	matc.mutation.SetCreatedAt(t)
	return matc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (matc *MasterAccountTypeCreate) SetNillableCreatedAt(t *time.Time) *MasterAccountTypeCreate {
	if t != nil {
		matc.SetCreatedAt(*t)
	}
	return matc
}

// AddCategoryIDs adds the "categories" edge to the MasterAccountCategory entity by IDs.
func (matc *MasterAccountTypeCreate) AddCategoryIDs(ids ...int) *MasterAccountTypeCreate {
	matc.mutation.AddCategoryIDs(ids...)
	return matc
}

// AddCategories adds the "categories" edges to the MasterAccountCategory entity.
func (matc *MasterAccountTypeCreate) AddCategories(m ...*MasterAccountCategory) *MasterAccountTypeCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return matc.AddCategoryIDs(ids...)
}

// Mutation returns the MasterAccountTypeMutation object of the builder.
func (matc *MasterAccountTypeCreate) Mutation() *MasterAccountTypeMutation {
	return matc.mutation
}

// Save creates the MasterAccountType in the database.
func (matc *MasterAccountTypeCreate) Save(ctx context.Context) (*MasterAccountType, error) {
	matc.defaults()
	return withHooks[*MasterAccountType, MasterAccountTypeMutation](ctx, matc.sqlSave, matc.mutation, matc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (matc *MasterAccountTypeCreate) SaveX(ctx context.Context) *MasterAccountType {
	v, err := matc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (matc *MasterAccountTypeCreate) Exec(ctx context.Context) error {
	_, err := matc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (matc *MasterAccountTypeCreate) ExecX(ctx context.Context) {
	if err := matc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (matc *MasterAccountTypeCreate) defaults() {
	if _, ok := matc.mutation.CreatedAt(); !ok {
		v := masteraccounttype.DefaultCreatedAt()
		matc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (matc *MasterAccountTypeCreate) check() error {
	if _, ok := matc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MasterAccountType.name"`)}
	}
	if v, ok := matc.mutation.Name(); ok {
		if err := masteraccounttype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MasterAccountType.name": %w`, err)}
		}
	}
	if _, ok := matc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MasterAccountType.created_at"`)}
	}
	return nil
}

func (matc *MasterAccountTypeCreate) sqlSave(ctx context.Context) (*MasterAccountType, error) {
	if err := matc.check(); err != nil {
		return nil, err
	}
	_node, _spec := matc.createSpec()
	if err := sqlgraph.CreateNode(ctx, matc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	matc.mutation.id = &_node.ID
	matc.mutation.done = true
	return _node, nil
}

func (matc *MasterAccountTypeCreate) createSpec() (*MasterAccountType, *sqlgraph.CreateSpec) {
	var (
		_node = &MasterAccountType{config: matc.config}
		_spec = sqlgraph.NewCreateSpec(masteraccounttype.Table, sqlgraph.NewFieldSpec(masteraccounttype.FieldID, field.TypeInt))
	)
	if value, ok := matc.mutation.Name(); ok {
		_spec.SetField(masteraccounttype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := matc.mutation.CreatedAt(); ok {
		_spec.SetField(masteraccounttype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := matc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   masteraccounttype.CategoriesTable,
			Columns: []string{masteraccounttype.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(masteraccountcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MasterAccountTypeCreateBulk is the builder for creating many MasterAccountType entities in bulk.
type MasterAccountTypeCreateBulk struct {
	config
	builders []*MasterAccountTypeCreate
}

// Save creates the MasterAccountType entities in the database.
func (matcb *MasterAccountTypeCreateBulk) Save(ctx context.Context) ([]*MasterAccountType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(matcb.builders))
	nodes := make([]*MasterAccountType, len(matcb.builders))
	mutators := make([]Mutator, len(matcb.builders))
	for i := range matcb.builders {
		func(i int, root context.Context) {
			builder := matcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MasterAccountTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, matcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, matcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, matcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (matcb *MasterAccountTypeCreateBulk) SaveX(ctx context.Context) []*MasterAccountType {
	v, err := matcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (matcb *MasterAccountTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := matcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (matcb *MasterAccountTypeCreateBulk) ExecX(ctx context.Context) {
	if err := matcb.Exec(ctx); err != nil {
		panic(err)
	}
}