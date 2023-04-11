// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"outgrow/ent/organizationaccount"
	"outgrow/ent/organizationaccountcategory"
	"outgrow/ent/transaction"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrganizationAccountCreate is the builder for creating a OrganizationAccount entity.
type OrganizationAccountCreate struct {
	config
	mutation *OrganizationAccountMutation
	hooks    []Hook
}

// SetCategoryID sets the "category_id" field.
func (oac *OrganizationAccountCreate) SetCategoryID(i int) *OrganizationAccountCreate {
	oac.mutation.SetCategoryID(i)
	return oac
}

// SetName sets the "name" field.
func (oac *OrganizationAccountCreate) SetName(s string) *OrganizationAccountCreate {
	oac.mutation.SetName(s)
	return oac
}

// SetCode sets the "code" field.
func (oac *OrganizationAccountCreate) SetCode(s string) *OrganizationAccountCreate {
	oac.mutation.SetCode(s)
	return oac
}

// SetBalance sets the "balance" field.
func (oac *OrganizationAccountCreate) SetBalance(f float64) *OrganizationAccountCreate {
	oac.mutation.SetBalance(f)
	return oac
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (oac *OrganizationAccountCreate) SetNillableBalance(f *float64) *OrganizationAccountCreate {
	if f != nil {
		oac.SetBalance(*f)
	}
	return oac
}

// SetCreatedAt sets the "created_at" field.
func (oac *OrganizationAccountCreate) SetCreatedAt(t time.Time) *OrganizationAccountCreate {
	oac.mutation.SetCreatedAt(t)
	return oac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oac *OrganizationAccountCreate) SetNillableCreatedAt(t *time.Time) *OrganizationAccountCreate {
	if t != nil {
		oac.SetCreatedAt(*t)
	}
	return oac
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (oac *OrganizationAccountCreate) AddTransactionIDs(ids ...uuid.UUID) *OrganizationAccountCreate {
	oac.mutation.AddTransactionIDs(ids...)
	return oac
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (oac *OrganizationAccountCreate) AddTransactions(t ...*Transaction) *OrganizationAccountCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return oac.AddTransactionIDs(ids...)
}

// SetAccCategoryID sets the "accCategory" edge to the OrganizationAccountCategory entity by ID.
func (oac *OrganizationAccountCreate) SetAccCategoryID(id int) *OrganizationAccountCreate {
	oac.mutation.SetAccCategoryID(id)
	return oac
}

// SetAccCategory sets the "accCategory" edge to the OrganizationAccountCategory entity.
func (oac *OrganizationAccountCreate) SetAccCategory(o *OrganizationAccountCategory) *OrganizationAccountCreate {
	return oac.SetAccCategoryID(o.ID)
}

// Mutation returns the OrganizationAccountMutation object of the builder.
func (oac *OrganizationAccountCreate) Mutation() *OrganizationAccountMutation {
	return oac.mutation
}

// Save creates the OrganizationAccount in the database.
func (oac *OrganizationAccountCreate) Save(ctx context.Context) (*OrganizationAccount, error) {
	oac.defaults()
	return withHooks[*OrganizationAccount, OrganizationAccountMutation](ctx, oac.sqlSave, oac.mutation, oac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oac *OrganizationAccountCreate) SaveX(ctx context.Context) *OrganizationAccount {
	v, err := oac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oac *OrganizationAccountCreate) Exec(ctx context.Context) error {
	_, err := oac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oac *OrganizationAccountCreate) ExecX(ctx context.Context) {
	if err := oac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oac *OrganizationAccountCreate) defaults() {
	if _, ok := oac.mutation.Balance(); !ok {
		v := organizationaccount.DefaultBalance
		oac.mutation.SetBalance(v)
	}
	if _, ok := oac.mutation.CreatedAt(); !ok {
		v := organizationaccount.DefaultCreatedAt()
		oac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oac *OrganizationAccountCreate) check() error {
	if _, ok := oac.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "OrganizationAccount.category_id"`)}
	}
	if v, ok := oac.mutation.CategoryID(); ok {
		if err := organizationaccount.CategoryIDValidator(v); err != nil {
			return &ValidationError{Name: "category_id", err: fmt.Errorf(`ent: validator failed for field "OrganizationAccount.category_id": %w`, err)}
		}
	}
	if _, ok := oac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OrganizationAccount.name"`)}
	}
	if v, ok := oac.mutation.Name(); ok {
		if err := organizationaccount.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrganizationAccount.name": %w`, err)}
		}
	}
	if _, ok := oac.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "OrganizationAccount.code"`)}
	}
	if v, ok := oac.mutation.Code(); ok {
		if err := organizationaccount.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "OrganizationAccount.code": %w`, err)}
		}
	}
	if _, ok := oac.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "OrganizationAccount.balance"`)}
	}
	if _, ok := oac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrganizationAccount.created_at"`)}
	}
	if _, ok := oac.mutation.AccCategoryID(); !ok {
		return &ValidationError{Name: "accCategory", err: errors.New(`ent: missing required edge "OrganizationAccount.accCategory"`)}
	}
	return nil
}

func (oac *OrganizationAccountCreate) sqlSave(ctx context.Context) (*OrganizationAccount, error) {
	if err := oac.check(); err != nil {
		return nil, err
	}
	_node, _spec := oac.createSpec()
	if err := sqlgraph.CreateNode(ctx, oac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oac.mutation.id = &_node.ID
	oac.mutation.done = true
	return _node, nil
}

func (oac *OrganizationAccountCreate) createSpec() (*OrganizationAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationAccount{config: oac.config}
		_spec = sqlgraph.NewCreateSpec(organizationaccount.Table, sqlgraph.NewFieldSpec(organizationaccount.FieldID, field.TypeInt))
	)
	if value, ok := oac.mutation.Name(); ok {
		_spec.SetField(organizationaccount.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := oac.mutation.Code(); ok {
		_spec.SetField(organizationaccount.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := oac.mutation.Balance(); ok {
		_spec.SetField(organizationaccount.FieldBalance, field.TypeFloat64, value)
		_node.Balance = value
	}
	if value, ok := oac.mutation.CreatedAt(); ok {
		_spec.SetField(organizationaccount.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := oac.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccount.TransactionsTable,
			Columns: []string{organizationaccount.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oac.mutation.AccCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccount.AccCategoryTable,
			Columns: []string{organizationaccount.AccCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationAccountCreateBulk is the builder for creating many OrganizationAccount entities in bulk.
type OrganizationAccountCreateBulk struct {
	config
	builders []*OrganizationAccountCreate
}

// Save creates the OrganizationAccount entities in the database.
func (oacb *OrganizationAccountCreateBulk) Save(ctx context.Context) ([]*OrganizationAccount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oacb.builders))
	nodes := make([]*OrganizationAccount, len(oacb.builders))
	mutators := make([]Mutator, len(oacb.builders))
	for i := range oacb.builders {
		func(i int, root context.Context) {
			builder := oacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, oacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oacb *OrganizationAccountCreateBulk) SaveX(ctx context.Context) []*OrganizationAccount {
	v, err := oacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oacb *OrganizationAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := oacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oacb *OrganizationAccountCreateBulk) ExecX(ctx context.Context) {
	if err := oacb.Exec(ctx); err != nil {
		panic(err)
	}
}
