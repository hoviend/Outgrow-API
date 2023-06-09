// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"outgrow/ent/event"
	"outgrow/ent/organization"
	"outgrow/ent/organizationaccounttype"
	"outgrow/ent/organizationeventtype"
	"outgrow/ent/predicate"
	"outgrow/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrganizationUpdate is the builder for updating Organization entities.
type OrganizationUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationMutation
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ou *OrganizationUpdate) Where(ps ...predicate.Organization) *OrganizationUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetName sets the "name" field.
func (ou *OrganizationUpdate) SetName(s string) *OrganizationUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetPublicID sets the "public_id" field.
func (ou *OrganizationUpdate) SetPublicID(s string) *OrganizationUpdate {
	ou.mutation.SetPublicID(s)
	return ou
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ou *OrganizationUpdate) AddUserIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.AddUserIDs(ids...)
	return ou
}

// AddUsers adds the "users" edges to the User entity.
func (ou *OrganizationUpdate) AddUsers(u ...*User) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ou.AddUserIDs(ids...)
}

// AddAccountTypeIDs adds the "accountTypes" edge to the OrganizationAccountType entity by IDs.
func (ou *OrganizationUpdate) AddAccountTypeIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.AddAccountTypeIDs(ids...)
	return ou
}

// AddAccountTypes adds the "accountTypes" edges to the OrganizationAccountType entity.
func (ou *OrganizationUpdate) AddAccountTypes(o ...*OrganizationAccountType) *OrganizationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddAccountTypeIDs(ids...)
}

// AddEventTypeIDs adds the "eventTypes" edge to the OrganizationEventType entity by IDs.
func (ou *OrganizationUpdate) AddEventTypeIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.AddEventTypeIDs(ids...)
	return ou
}

// AddEventTypes adds the "eventTypes" edges to the OrganizationEventType entity.
func (ou *OrganizationUpdate) AddEventTypes(o ...*OrganizationEventType) *OrganizationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddEventTypeIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (ou *OrganizationUpdate) AddEventIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.AddEventIDs(ids...)
	return ou
}

// AddEvents adds the "events" edges to the Event entity.
func (ou *OrganizationUpdate) AddEvents(e ...*Event) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ou.AddEventIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ou *OrganizationUpdate) Mutation() *OrganizationMutation {
	return ou.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (ou *OrganizationUpdate) ClearUsers() *OrganizationUpdate {
	ou.mutation.ClearUsers()
	return ou
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ou *OrganizationUpdate) RemoveUserIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.RemoveUserIDs(ids...)
	return ou
}

// RemoveUsers removes "users" edges to User entities.
func (ou *OrganizationUpdate) RemoveUsers(u ...*User) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ou.RemoveUserIDs(ids...)
}

// ClearAccountTypes clears all "accountTypes" edges to the OrganizationAccountType entity.
func (ou *OrganizationUpdate) ClearAccountTypes() *OrganizationUpdate {
	ou.mutation.ClearAccountTypes()
	return ou
}

// RemoveAccountTypeIDs removes the "accountTypes" edge to OrganizationAccountType entities by IDs.
func (ou *OrganizationUpdate) RemoveAccountTypeIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.RemoveAccountTypeIDs(ids...)
	return ou
}

// RemoveAccountTypes removes "accountTypes" edges to OrganizationAccountType entities.
func (ou *OrganizationUpdate) RemoveAccountTypes(o ...*OrganizationAccountType) *OrganizationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveAccountTypeIDs(ids...)
}

// ClearEventTypes clears all "eventTypes" edges to the OrganizationEventType entity.
func (ou *OrganizationUpdate) ClearEventTypes() *OrganizationUpdate {
	ou.mutation.ClearEventTypes()
	return ou
}

// RemoveEventTypeIDs removes the "eventTypes" edge to OrganizationEventType entities by IDs.
func (ou *OrganizationUpdate) RemoveEventTypeIDs(ids ...int) *OrganizationUpdate {
	ou.mutation.RemoveEventTypeIDs(ids...)
	return ou
}

// RemoveEventTypes removes "eventTypes" edges to OrganizationEventType entities.
func (ou *OrganizationUpdate) RemoveEventTypes(o ...*OrganizationEventType) *OrganizationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveEventTypeIDs(ids...)
}

// ClearEvents clears all "events" edges to the Event entity.
func (ou *OrganizationUpdate) ClearEvents() *OrganizationUpdate {
	ou.mutation.ClearEvents()
	return ou
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (ou *OrganizationUpdate) RemoveEventIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.RemoveEventIDs(ids...)
	return ou
}

// RemoveEvents removes "events" edges to Event entities.
func (ou *OrganizationUpdate) RemoveEvents(e ...*Event) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ou.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrganizationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OrganizationMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrganizationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrganizationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrganizationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrganizationUpdate) check() error {
	if v, ok := ou.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.PublicID(); ok {
		if err := organization.PublicIDValidator(v); err != nil {
			return &ValidationError{Name: "public_id", err: fmt.Errorf(`ent: validator failed for field "Organization.public_id": %w`, err)}
		}
	}
	return nil
}

func (ou *OrganizationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.PublicID(); ok {
		_spec.SetField(organization.FieldPublicID, field.TypeString, value)
	}
	if ou.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.UsersTable,
			Columns: organization.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ou.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.UsersTable,
			Columns: organization.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.UsersTable,
			Columns: organization.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.AccountTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.AccountTypesTable,
			Columns: []string{organization.AccountTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedAccountTypesIDs(); len(nodes) > 0 && !ou.mutation.AccountTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.AccountTypesTable,
			Columns: []string{organization.AccountTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.AccountTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.AccountTypesTable,
			Columns: []string{organization.AccountTypesColumn},
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
	if ou.mutation.EventTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventTypesTable,
			Columns: []string{organization.EventTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedEventTypesIDs(); len(nodes) > 0 && !ou.mutation.EventTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventTypesTable,
			Columns: []string{organization.EventTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.EventTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventTypesTable,
			Columns: []string{organization.EventTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventsTable,
			Columns: []string{organization.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedEventsIDs(); len(nodes) > 0 && !ou.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventsTable,
			Columns: []string{organization.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventsTable,
			Columns: []string{organization.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrganizationUpdateOne is the builder for updating a single Organization entity.
type OrganizationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationMutation
}

// SetName sets the "name" field.
func (ouo *OrganizationUpdateOne) SetName(s string) *OrganizationUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetPublicID sets the "public_id" field.
func (ouo *OrganizationUpdateOne) SetPublicID(s string) *OrganizationUpdateOne {
	ouo.mutation.SetPublicID(s)
	return ouo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ouo *OrganizationUpdateOne) AddUserIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.AddUserIDs(ids...)
	return ouo
}

// AddUsers adds the "users" edges to the User entity.
func (ouo *OrganizationUpdateOne) AddUsers(u ...*User) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ouo.AddUserIDs(ids...)
}

// AddAccountTypeIDs adds the "accountTypes" edge to the OrganizationAccountType entity by IDs.
func (ouo *OrganizationUpdateOne) AddAccountTypeIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.AddAccountTypeIDs(ids...)
	return ouo
}

// AddAccountTypes adds the "accountTypes" edges to the OrganizationAccountType entity.
func (ouo *OrganizationUpdateOne) AddAccountTypes(o ...*OrganizationAccountType) *OrganizationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddAccountTypeIDs(ids...)
}

// AddEventTypeIDs adds the "eventTypes" edge to the OrganizationEventType entity by IDs.
func (ouo *OrganizationUpdateOne) AddEventTypeIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.AddEventTypeIDs(ids...)
	return ouo
}

// AddEventTypes adds the "eventTypes" edges to the OrganizationEventType entity.
func (ouo *OrganizationUpdateOne) AddEventTypes(o ...*OrganizationEventType) *OrganizationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddEventTypeIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (ouo *OrganizationUpdateOne) AddEventIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.AddEventIDs(ids...)
	return ouo
}

// AddEvents adds the "events" edges to the Event entity.
func (ouo *OrganizationUpdateOne) AddEvents(e ...*Event) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ouo.AddEventIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ouo *OrganizationUpdateOne) Mutation() *OrganizationMutation {
	return ouo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (ouo *OrganizationUpdateOne) ClearUsers() *OrganizationUpdateOne {
	ouo.mutation.ClearUsers()
	return ouo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.RemoveUserIDs(ids...)
	return ouo
}

// RemoveUsers removes "users" edges to User entities.
func (ouo *OrganizationUpdateOne) RemoveUsers(u ...*User) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ouo.RemoveUserIDs(ids...)
}

// ClearAccountTypes clears all "accountTypes" edges to the OrganizationAccountType entity.
func (ouo *OrganizationUpdateOne) ClearAccountTypes() *OrganizationUpdateOne {
	ouo.mutation.ClearAccountTypes()
	return ouo
}

// RemoveAccountTypeIDs removes the "accountTypes" edge to OrganizationAccountType entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveAccountTypeIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.RemoveAccountTypeIDs(ids...)
	return ouo
}

// RemoveAccountTypes removes "accountTypes" edges to OrganizationAccountType entities.
func (ouo *OrganizationUpdateOne) RemoveAccountTypes(o ...*OrganizationAccountType) *OrganizationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveAccountTypeIDs(ids...)
}

// ClearEventTypes clears all "eventTypes" edges to the OrganizationEventType entity.
func (ouo *OrganizationUpdateOne) ClearEventTypes() *OrganizationUpdateOne {
	ouo.mutation.ClearEventTypes()
	return ouo
}

// RemoveEventTypeIDs removes the "eventTypes" edge to OrganizationEventType entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveEventTypeIDs(ids ...int) *OrganizationUpdateOne {
	ouo.mutation.RemoveEventTypeIDs(ids...)
	return ouo
}

// RemoveEventTypes removes "eventTypes" edges to OrganizationEventType entities.
func (ouo *OrganizationUpdateOne) RemoveEventTypes(o ...*OrganizationEventType) *OrganizationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveEventTypeIDs(ids...)
}

// ClearEvents clears all "events" edges to the Event entity.
func (ouo *OrganizationUpdateOne) ClearEvents() *OrganizationUpdateOne {
	ouo.mutation.ClearEvents()
	return ouo
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveEventIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.RemoveEventIDs(ids...)
	return ouo
}

// RemoveEvents removes "events" edges to Event entities.
func (ouo *OrganizationUpdateOne) RemoveEvents(e ...*Event) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ouo.RemoveEventIDs(ids...)
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ouo *OrganizationUpdateOne) Where(ps ...predicate.Organization) *OrganizationUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrganizationUpdateOne) Select(field string, fields ...string) *OrganizationUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Organization entity.
func (ouo *OrganizationUpdateOne) Save(ctx context.Context) (*Organization, error) {
	return withHooks[*Organization, OrganizationMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) SaveX(ctx context.Context) *Organization {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrganizationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrganizationUpdateOne) check() error {
	if v, ok := ouo.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.PublicID(); ok {
		if err := organization.PublicIDValidator(v); err != nil {
			return &ValidationError{Name: "public_id", err: fmt.Errorf(`ent: validator failed for field "Organization.public_id": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrganizationUpdateOne) sqlSave(ctx context.Context) (_node *Organization, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Organization.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for _, f := range fields {
			if !organization.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.PublicID(); ok {
		_spec.SetField(organization.FieldPublicID, field.TypeString, value)
	}
	if ouo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.UsersTable,
			Columns: organization.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ouo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.UsersTable,
			Columns: organization.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.UsersTable,
			Columns: organization.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.AccountTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.AccountTypesTable,
			Columns: []string{organization.AccountTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedAccountTypesIDs(); len(nodes) > 0 && !ouo.mutation.AccountTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.AccountTypesTable,
			Columns: []string{organization.AccountTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationaccounttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.AccountTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.AccountTypesTable,
			Columns: []string{organization.AccountTypesColumn},
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
	if ouo.mutation.EventTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventTypesTable,
			Columns: []string{organization.EventTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedEventTypesIDs(); len(nodes) > 0 && !ouo.mutation.EventTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventTypesTable,
			Columns: []string{organization.EventTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.EventTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventTypesTable,
			Columns: []string{organization.EventTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventsTable,
			Columns: []string{organization.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedEventsIDs(); len(nodes) > 0 && !ouo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventsTable,
			Columns: []string{organization.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EventsTable,
			Columns: []string{organization.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Organization{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
