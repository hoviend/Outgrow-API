// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"outgrow/ent/organizationaccount"
	"outgrow/ent/organizationaccountcategory"
	"outgrow/ent/organizationaccounttype"
	"outgrow/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationAccountCategoryUpdate is the builder for updating OrganizationAccountCategory entities.
type OrganizationAccountCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationAccountCategoryMutation
}

// Where appends a list predicates to the OrganizationAccountCategoryUpdate builder.
func (oacu *OrganizationAccountCategoryUpdate) Where(ps ...predicate.OrganizationAccountCategory) *OrganizationAccountCategoryUpdate {
	oacu.mutation.Where(ps...)
	return oacu
}

// SetAccountTypeID sets the "account_type_id" field.
func (oacu *OrganizationAccountCategoryUpdate) SetAccountTypeID(i int) *OrganizationAccountCategoryUpdate {
	oacu.mutation.SetAccountTypeID(i)
	return oacu
}

// SetName sets the "name" field.
func (oacu *OrganizationAccountCategoryUpdate) SetName(s string) *OrganizationAccountCategoryUpdate {
	oacu.mutation.SetName(s)
	return oacu
}

// SetDescription sets the "description" field.
func (oacu *OrganizationAccountCategoryUpdate) SetDescription(s string) *OrganizationAccountCategoryUpdate {
	oacu.mutation.SetDescription(s)
	return oacu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (oacu *OrganizationAccountCategoryUpdate) SetNillableDescription(s *string) *OrganizationAccountCategoryUpdate {
	if s != nil {
		oacu.SetDescription(*s)
	}
	return oacu
}

// ClearDescription clears the value of the "description" field.
func (oacu *OrganizationAccountCategoryUpdate) ClearDescription() *OrganizationAccountCategoryUpdate {
	oacu.mutation.ClearDescription()
	return oacu
}

// AddAccountIDs adds the "accounts" edge to the OrganizationAccount entity by IDs.
func (oacu *OrganizationAccountCategoryUpdate) AddAccountIDs(ids ...int) *OrganizationAccountCategoryUpdate {
	oacu.mutation.AddAccountIDs(ids...)
	return oacu
}

// AddAccounts adds the "accounts" edges to the OrganizationAccount entity.
func (oacu *OrganizationAccountCategoryUpdate) AddAccounts(o ...*OrganizationAccount) *OrganizationAccountCategoryUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oacu.AddAccountIDs(ids...)
}

// SetTypeID sets the "type" edge to the OrganizationAccountType entity by ID.
func (oacu *OrganizationAccountCategoryUpdate) SetTypeID(id int) *OrganizationAccountCategoryUpdate {
	oacu.mutation.SetTypeID(id)
	return oacu
}

// SetType sets the "type" edge to the OrganizationAccountType entity.
func (oacu *OrganizationAccountCategoryUpdate) SetType(o *OrganizationAccountType) *OrganizationAccountCategoryUpdate {
	return oacu.SetTypeID(o.ID)
}

// Mutation returns the OrganizationAccountCategoryMutation object of the builder.
func (oacu *OrganizationAccountCategoryUpdate) Mutation() *OrganizationAccountCategoryMutation {
	return oacu.mutation
}

// ClearAccounts clears all "accounts" edges to the OrganizationAccount entity.
func (oacu *OrganizationAccountCategoryUpdate) ClearAccounts() *OrganizationAccountCategoryUpdate {
	oacu.mutation.ClearAccounts()
	return oacu
}

// RemoveAccountIDs removes the "accounts" edge to OrganizationAccount entities by IDs.
func (oacu *OrganizationAccountCategoryUpdate) RemoveAccountIDs(ids ...int) *OrganizationAccountCategoryUpdate {
	oacu.mutation.RemoveAccountIDs(ids...)
	return oacu
}

// RemoveAccounts removes "accounts" edges to OrganizationAccount entities.
func (oacu *OrganizationAccountCategoryUpdate) RemoveAccounts(o ...*OrganizationAccount) *OrganizationAccountCategoryUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oacu.RemoveAccountIDs(ids...)
}

// ClearType clears the "type" edge to the OrganizationAccountType entity.
func (oacu *OrganizationAccountCategoryUpdate) ClearType() *OrganizationAccountCategoryUpdate {
	oacu.mutation.ClearType()
	return oacu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oacu *OrganizationAccountCategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OrganizationAccountCategoryMutation](ctx, oacu.sqlSave, oacu.mutation, oacu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oacu *OrganizationAccountCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := oacu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oacu *OrganizationAccountCategoryUpdate) Exec(ctx context.Context) error {
	_, err := oacu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oacu *OrganizationAccountCategoryUpdate) ExecX(ctx context.Context) {
	if err := oacu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oacu *OrganizationAccountCategoryUpdate) check() error {
	if v, ok := oacu.mutation.Name(); ok {
		if err := organizationaccountcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrganizationAccountCategory.name": %w`, err)}
		}
	}
	if _, ok := oacu.mutation.TypeID(); oacu.mutation.TypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrganizationAccountCategory.type"`)
	}
	return nil
}

func (oacu *OrganizationAccountCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oacu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationaccountcategory.Table, organizationaccountcategory.Columns, sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt))
	if ps := oacu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oacu.mutation.Name(); ok {
		_spec.SetField(organizationaccountcategory.FieldName, field.TypeString, value)
	}
	if value, ok := oacu.mutation.Description(); ok {
		_spec.SetField(organizationaccountcategory.FieldDescription, field.TypeString, value)
	}
	if oacu.mutation.DescriptionCleared() {
		_spec.ClearField(organizationaccountcategory.FieldDescription, field.TypeString)
	}
	if oacu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccountcategory.AccountsTable,
			Columns: []string{organizationaccountcategory.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccount.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oacu.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !oacu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccountcategory.AccountsTable,
			Columns: []string{organizationaccountcategory.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oacu.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccountcategory.AccountsTable,
			Columns: []string{organizationaccountcategory.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oacu.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccountcategory.TypeTable,
			Columns: []string{organizationaccountcategory.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oacu.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccountcategory.TypeTable,
			Columns: []string{organizationaccountcategory.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oacu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationaccountcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oacu.mutation.done = true
	return n, nil
}

// OrganizationAccountCategoryUpdateOne is the builder for updating a single OrganizationAccountCategory entity.
type OrganizationAccountCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationAccountCategoryMutation
}

// SetAccountTypeID sets the "account_type_id" field.
func (oacuo *OrganizationAccountCategoryUpdateOne) SetAccountTypeID(i int) *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.SetAccountTypeID(i)
	return oacuo
}

// SetName sets the "name" field.
func (oacuo *OrganizationAccountCategoryUpdateOne) SetName(s string) *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.SetName(s)
	return oacuo
}

// SetDescription sets the "description" field.
func (oacuo *OrganizationAccountCategoryUpdateOne) SetDescription(s string) *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.SetDescription(s)
	return oacuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (oacuo *OrganizationAccountCategoryUpdateOne) SetNillableDescription(s *string) *OrganizationAccountCategoryUpdateOne {
	if s != nil {
		oacuo.SetDescription(*s)
	}
	return oacuo
}

// ClearDescription clears the value of the "description" field.
func (oacuo *OrganizationAccountCategoryUpdateOne) ClearDescription() *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.ClearDescription()
	return oacuo
}

// AddAccountIDs adds the "accounts" edge to the OrganizationAccount entity by IDs.
func (oacuo *OrganizationAccountCategoryUpdateOne) AddAccountIDs(ids ...int) *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.AddAccountIDs(ids...)
	return oacuo
}

// AddAccounts adds the "accounts" edges to the OrganizationAccount entity.
func (oacuo *OrganizationAccountCategoryUpdateOne) AddAccounts(o ...*OrganizationAccount) *OrganizationAccountCategoryUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oacuo.AddAccountIDs(ids...)
}

// SetTypeID sets the "type" edge to the OrganizationAccountType entity by ID.
func (oacuo *OrganizationAccountCategoryUpdateOne) SetTypeID(id int) *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.SetTypeID(id)
	return oacuo
}

// SetType sets the "type" edge to the OrganizationAccountType entity.
func (oacuo *OrganizationAccountCategoryUpdateOne) SetType(o *OrganizationAccountType) *OrganizationAccountCategoryUpdateOne {
	return oacuo.SetTypeID(o.ID)
}

// Mutation returns the OrganizationAccountCategoryMutation object of the builder.
func (oacuo *OrganizationAccountCategoryUpdateOne) Mutation() *OrganizationAccountCategoryMutation {
	return oacuo.mutation
}

// ClearAccounts clears all "accounts" edges to the OrganizationAccount entity.
func (oacuo *OrganizationAccountCategoryUpdateOne) ClearAccounts() *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.ClearAccounts()
	return oacuo
}

// RemoveAccountIDs removes the "accounts" edge to OrganizationAccount entities by IDs.
func (oacuo *OrganizationAccountCategoryUpdateOne) RemoveAccountIDs(ids ...int) *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.RemoveAccountIDs(ids...)
	return oacuo
}

// RemoveAccounts removes "accounts" edges to OrganizationAccount entities.
func (oacuo *OrganizationAccountCategoryUpdateOne) RemoveAccounts(o ...*OrganizationAccount) *OrganizationAccountCategoryUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oacuo.RemoveAccountIDs(ids...)
}

// ClearType clears the "type" edge to the OrganizationAccountType entity.
func (oacuo *OrganizationAccountCategoryUpdateOne) ClearType() *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.ClearType()
	return oacuo
}

// Where appends a list predicates to the OrganizationAccountCategoryUpdate builder.
func (oacuo *OrganizationAccountCategoryUpdateOne) Where(ps ...predicate.OrganizationAccountCategory) *OrganizationAccountCategoryUpdateOne {
	oacuo.mutation.Where(ps...)
	return oacuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oacuo *OrganizationAccountCategoryUpdateOne) Select(field string, fields ...string) *OrganizationAccountCategoryUpdateOne {
	oacuo.fields = append([]string{field}, fields...)
	return oacuo
}

// Save executes the query and returns the updated OrganizationAccountCategory entity.
func (oacuo *OrganizationAccountCategoryUpdateOne) Save(ctx context.Context) (*OrganizationAccountCategory, error) {
	return withHooks[*OrganizationAccountCategory, OrganizationAccountCategoryMutation](ctx, oacuo.sqlSave, oacuo.mutation, oacuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oacuo *OrganizationAccountCategoryUpdateOne) SaveX(ctx context.Context) *OrganizationAccountCategory {
	node, err := oacuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oacuo *OrganizationAccountCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := oacuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oacuo *OrganizationAccountCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := oacuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oacuo *OrganizationAccountCategoryUpdateOne) check() error {
	if v, ok := oacuo.mutation.Name(); ok {
		if err := organizationaccountcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrganizationAccountCategory.name": %w`, err)}
		}
	}
	if _, ok := oacuo.mutation.TypeID(); oacuo.mutation.TypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrganizationAccountCategory.type"`)
	}
	return nil
}

func (oacuo *OrganizationAccountCategoryUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationAccountCategory, err error) {
	if err := oacuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationaccountcategory.Table, organizationaccountcategory.Columns, sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt))
	id, ok := oacuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrganizationAccountCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oacuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationaccountcategory.FieldID)
		for _, f := range fields {
			if !organizationaccountcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organizationaccountcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oacuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oacuo.mutation.Name(); ok {
		_spec.SetField(organizationaccountcategory.FieldName, field.TypeString, value)
	}
	if value, ok := oacuo.mutation.Description(); ok {
		_spec.SetField(organizationaccountcategory.FieldDescription, field.TypeString, value)
	}
	if oacuo.mutation.DescriptionCleared() {
		_spec.ClearField(organizationaccountcategory.FieldDescription, field.TypeString)
	}
	if oacuo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccountcategory.AccountsTable,
			Columns: []string{organizationaccountcategory.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccount.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oacuo.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !oacuo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccountcategory.AccountsTable,
			Columns: []string{organizationaccountcategory.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oacuo.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccountcategory.AccountsTable,
			Columns: []string{organizationaccountcategory.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oacuo.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccountcategory.TypeTable,
			Columns: []string{organizationaccountcategory.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oacuo.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccountcategory.TypeTable,
			Columns: []string{organizationaccountcategory.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrganizationAccountCategory{config: oacuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oacuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationaccountcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oacuo.mutation.done = true
	return _node, nil
}