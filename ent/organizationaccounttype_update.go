// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"outgrow/ent/organization"
	"outgrow/ent/organizationaccountcategory"
	"outgrow/ent/organizationaccounttype"
	"outgrow/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrganizationAccountTypeUpdate is the builder for updating OrganizationAccountType entities.
type OrganizationAccountTypeUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationAccountTypeMutation
}

// Where appends a list predicates to the OrganizationAccountTypeUpdate builder.
func (oatu *OrganizationAccountTypeUpdate) Where(ps ...predicate.OrganizationAccountType) *OrganizationAccountTypeUpdate {
	oatu.mutation.Where(ps...)
	return oatu
}

// SetOrganizationID sets the "organization_id" field.
func (oatu *OrganizationAccountTypeUpdate) SetOrganizationID(u uuid.UUID) *OrganizationAccountTypeUpdate {
	oatu.mutation.SetOrganizationID(u)
	return oatu
}

// SetName sets the "name" field.
func (oatu *OrganizationAccountTypeUpdate) SetName(s string) *OrganizationAccountTypeUpdate {
	oatu.mutation.SetName(s)
	return oatu
}

// AddCategoryIDs adds the "categories" edge to the OrganizationAccountCategory entity by IDs.
func (oatu *OrganizationAccountTypeUpdate) AddCategoryIDs(ids ...int) *OrganizationAccountTypeUpdate {
	oatu.mutation.AddCategoryIDs(ids...)
	return oatu
}

// AddCategories adds the "categories" edges to the OrganizationAccountCategory entity.
func (oatu *OrganizationAccountTypeUpdate) AddCategories(o ...*OrganizationAccountCategory) *OrganizationAccountTypeUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oatu.AddCategoryIDs(ids...)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (oatu *OrganizationAccountTypeUpdate) SetOrganization(o *Organization) *OrganizationAccountTypeUpdate {
	return oatu.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationAccountTypeMutation object of the builder.
func (oatu *OrganizationAccountTypeUpdate) Mutation() *OrganizationAccountTypeMutation {
	return oatu.mutation
}

// ClearCategories clears all "categories" edges to the OrganizationAccountCategory entity.
func (oatu *OrganizationAccountTypeUpdate) ClearCategories() *OrganizationAccountTypeUpdate {
	oatu.mutation.ClearCategories()
	return oatu
}

// RemoveCategoryIDs removes the "categories" edge to OrganizationAccountCategory entities by IDs.
func (oatu *OrganizationAccountTypeUpdate) RemoveCategoryIDs(ids ...int) *OrganizationAccountTypeUpdate {
	oatu.mutation.RemoveCategoryIDs(ids...)
	return oatu
}

// RemoveCategories removes "categories" edges to OrganizationAccountCategory entities.
func (oatu *OrganizationAccountTypeUpdate) RemoveCategories(o ...*OrganizationAccountCategory) *OrganizationAccountTypeUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oatu.RemoveCategoryIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (oatu *OrganizationAccountTypeUpdate) ClearOrganization() *OrganizationAccountTypeUpdate {
	oatu.mutation.ClearOrganization()
	return oatu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oatu *OrganizationAccountTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OrganizationAccountTypeMutation](ctx, oatu.sqlSave, oatu.mutation, oatu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oatu *OrganizationAccountTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := oatu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oatu *OrganizationAccountTypeUpdate) Exec(ctx context.Context) error {
	_, err := oatu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oatu *OrganizationAccountTypeUpdate) ExecX(ctx context.Context) {
	if err := oatu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oatu *OrganizationAccountTypeUpdate) check() error {
	if v, ok := oatu.mutation.Name(); ok {
		if err := organizationaccounttype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrganizationAccountType.name": %w`, err)}
		}
	}
	if _, ok := oatu.mutation.OrganizationID(); oatu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrganizationAccountType.organization"`)
	}
	return nil
}

func (oatu *OrganizationAccountTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oatu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationaccounttype.Table, organizationaccounttype.Columns, sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt))
	if ps := oatu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oatu.mutation.Name(); ok {
		_spec.SetField(organizationaccounttype.FieldName, field.TypeString, value)
	}
	if oatu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccounttype.CategoriesTable,
			Columns: []string{organizationaccounttype.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oatu.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !oatu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccounttype.CategoriesTable,
			Columns: []string{organizationaccounttype.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oatu.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccounttype.CategoriesTable,
			Columns: []string{organizationaccounttype.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oatu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccounttype.OrganizationTable,
			Columns: []string{organizationaccounttype.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oatu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccounttype.OrganizationTable,
			Columns: []string{organizationaccounttype.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oatu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationaccounttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oatu.mutation.done = true
	return n, nil
}

// OrganizationAccountTypeUpdateOne is the builder for updating a single OrganizationAccountType entity.
type OrganizationAccountTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationAccountTypeMutation
}

// SetOrganizationID sets the "organization_id" field.
func (oatuo *OrganizationAccountTypeUpdateOne) SetOrganizationID(u uuid.UUID) *OrganizationAccountTypeUpdateOne {
	oatuo.mutation.SetOrganizationID(u)
	return oatuo
}

// SetName sets the "name" field.
func (oatuo *OrganizationAccountTypeUpdateOne) SetName(s string) *OrganizationAccountTypeUpdateOne {
	oatuo.mutation.SetName(s)
	return oatuo
}

// AddCategoryIDs adds the "categories" edge to the OrganizationAccountCategory entity by IDs.
func (oatuo *OrganizationAccountTypeUpdateOne) AddCategoryIDs(ids ...int) *OrganizationAccountTypeUpdateOne {
	oatuo.mutation.AddCategoryIDs(ids...)
	return oatuo
}

// AddCategories adds the "categories" edges to the OrganizationAccountCategory entity.
func (oatuo *OrganizationAccountTypeUpdateOne) AddCategories(o ...*OrganizationAccountCategory) *OrganizationAccountTypeUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oatuo.AddCategoryIDs(ids...)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (oatuo *OrganizationAccountTypeUpdateOne) SetOrganization(o *Organization) *OrganizationAccountTypeUpdateOne {
	return oatuo.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationAccountTypeMutation object of the builder.
func (oatuo *OrganizationAccountTypeUpdateOne) Mutation() *OrganizationAccountTypeMutation {
	return oatuo.mutation
}

// ClearCategories clears all "categories" edges to the OrganizationAccountCategory entity.
func (oatuo *OrganizationAccountTypeUpdateOne) ClearCategories() *OrganizationAccountTypeUpdateOne {
	oatuo.mutation.ClearCategories()
	return oatuo
}

// RemoveCategoryIDs removes the "categories" edge to OrganizationAccountCategory entities by IDs.
func (oatuo *OrganizationAccountTypeUpdateOne) RemoveCategoryIDs(ids ...int) *OrganizationAccountTypeUpdateOne {
	oatuo.mutation.RemoveCategoryIDs(ids...)
	return oatuo
}

// RemoveCategories removes "categories" edges to OrganizationAccountCategory entities.
func (oatuo *OrganizationAccountTypeUpdateOne) RemoveCategories(o ...*OrganizationAccountCategory) *OrganizationAccountTypeUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oatuo.RemoveCategoryIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (oatuo *OrganizationAccountTypeUpdateOne) ClearOrganization() *OrganizationAccountTypeUpdateOne {
	oatuo.mutation.ClearOrganization()
	return oatuo
}

// Where appends a list predicates to the OrganizationAccountTypeUpdate builder.
func (oatuo *OrganizationAccountTypeUpdateOne) Where(ps ...predicate.OrganizationAccountType) *OrganizationAccountTypeUpdateOne {
	oatuo.mutation.Where(ps...)
	return oatuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oatuo *OrganizationAccountTypeUpdateOne) Select(field string, fields ...string) *OrganizationAccountTypeUpdateOne {
	oatuo.fields = append([]string{field}, fields...)
	return oatuo
}

// Save executes the query and returns the updated OrganizationAccountType entity.
func (oatuo *OrganizationAccountTypeUpdateOne) Save(ctx context.Context) (*OrganizationAccountType, error) {
	return withHooks[*OrganizationAccountType, OrganizationAccountTypeMutation](ctx, oatuo.sqlSave, oatuo.mutation, oatuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oatuo *OrganizationAccountTypeUpdateOne) SaveX(ctx context.Context) *OrganizationAccountType {
	node, err := oatuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oatuo *OrganizationAccountTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := oatuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oatuo *OrganizationAccountTypeUpdateOne) ExecX(ctx context.Context) {
	if err := oatuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oatuo *OrganizationAccountTypeUpdateOne) check() error {
	if v, ok := oatuo.mutation.Name(); ok {
		if err := organizationaccounttype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrganizationAccountType.name": %w`, err)}
		}
	}
	if _, ok := oatuo.mutation.OrganizationID(); oatuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrganizationAccountType.organization"`)
	}
	return nil
}

func (oatuo *OrganizationAccountTypeUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationAccountType, err error) {
	if err := oatuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationaccounttype.Table, organizationaccounttype.Columns, sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt))
	id, ok := oatuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrganizationAccountType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oatuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationaccounttype.FieldID)
		for _, f := range fields {
			if !organizationaccounttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organizationaccounttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oatuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oatuo.mutation.Name(); ok {
		_spec.SetField(organizationaccounttype.FieldName, field.TypeString, value)
	}
	if oatuo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccounttype.CategoriesTable,
			Columns: []string{organizationaccounttype.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oatuo.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !oatuo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccounttype.CategoriesTable,
			Columns: []string{organizationaccounttype.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oatuo.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationaccounttype.CategoriesTable,
			Columns: []string{organizationaccounttype.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccountcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oatuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccounttype.OrganizationTable,
			Columns: []string{organizationaccounttype.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oatuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationaccounttype.OrganizationTable,
			Columns: []string{organizationaccounttype.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrganizationAccountType{config: oatuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oatuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationaccounttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oatuo.mutation.done = true
	return _node, nil
}
