// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"outgrow/ent/event"
	"outgrow/ent/organization"
	"outgrow/ent/organizationeventtype"
	"outgrow/ent/predicate"
	"outgrow/ent/schema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrganizationEventTypeUpdate is the builder for updating OrganizationEventType entities.
type OrganizationEventTypeUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationEventTypeMutation
}

// Where appends a list predicates to the OrganizationEventTypeUpdate builder.
func (oetu *OrganizationEventTypeUpdate) Where(ps ...predicate.OrganizationEventType) *OrganizationEventTypeUpdate {
	oetu.mutation.Where(ps...)
	return oetu
}

// SetOrganizationID sets the "organization_id" field.
func (oetu *OrganizationEventTypeUpdate) SetOrganizationID(u uuid.UUID) *OrganizationEventTypeUpdate {
	oetu.mutation.SetOrganizationID(u)
	return oetu
}

// SetName sets the "name" field.
func (oetu *OrganizationEventTypeUpdate) SetName(s string) *OrganizationEventTypeUpdate {
	oetu.mutation.SetName(s)
	return oetu
}

// SetDescription sets the "description" field.
func (oetu *OrganizationEventTypeUpdate) SetDescription(s string) *OrganizationEventTypeUpdate {
	oetu.mutation.SetDescription(s)
	return oetu
}

// SetRules sets the "rules" field.
func (oetu *OrganizationEventTypeUpdate) SetRules(sr []schema.EventRules) *OrganizationEventTypeUpdate {
	oetu.mutation.SetRules(sr)
	return oetu
}

// AppendRules appends sr to the "rules" field.
func (oetu *OrganizationEventTypeUpdate) AppendRules(sr []schema.EventRules) *OrganizationEventTypeUpdate {
	oetu.mutation.AppendRules(sr)
	return oetu
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (oetu *OrganizationEventTypeUpdate) AddEventIDs(ids ...uuid.UUID) *OrganizationEventTypeUpdate {
	oetu.mutation.AddEventIDs(ids...)
	return oetu
}

// AddEvents adds the "events" edges to the Event entity.
func (oetu *OrganizationEventTypeUpdate) AddEvents(e ...*Event) *OrganizationEventTypeUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return oetu.AddEventIDs(ids...)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (oetu *OrganizationEventTypeUpdate) SetOrganization(o *Organization) *OrganizationEventTypeUpdate {
	return oetu.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationEventTypeMutation object of the builder.
func (oetu *OrganizationEventTypeUpdate) Mutation() *OrganizationEventTypeMutation {
	return oetu.mutation
}

// ClearEvents clears all "events" edges to the Event entity.
func (oetu *OrganizationEventTypeUpdate) ClearEvents() *OrganizationEventTypeUpdate {
	oetu.mutation.ClearEvents()
	return oetu
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (oetu *OrganizationEventTypeUpdate) RemoveEventIDs(ids ...uuid.UUID) *OrganizationEventTypeUpdate {
	oetu.mutation.RemoveEventIDs(ids...)
	return oetu
}

// RemoveEvents removes "events" edges to Event entities.
func (oetu *OrganizationEventTypeUpdate) RemoveEvents(e ...*Event) *OrganizationEventTypeUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return oetu.RemoveEventIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (oetu *OrganizationEventTypeUpdate) ClearOrganization() *OrganizationEventTypeUpdate {
	oetu.mutation.ClearOrganization()
	return oetu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oetu *OrganizationEventTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OrganizationEventTypeMutation](ctx, oetu.sqlSave, oetu.mutation, oetu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oetu *OrganizationEventTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := oetu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oetu *OrganizationEventTypeUpdate) Exec(ctx context.Context) error {
	_, err := oetu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oetu *OrganizationEventTypeUpdate) ExecX(ctx context.Context) {
	if err := oetu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oetu *OrganizationEventTypeUpdate) check() error {
	if v, ok := oetu.mutation.Name(); ok {
		if err := organizationeventtype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrganizationEventType.name": %w`, err)}
		}
	}
	if _, ok := oetu.mutation.OrganizationID(); oetu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrganizationEventType.organization"`)
	}
	return nil
}

func (oetu *OrganizationEventTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oetu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationeventtype.Table, organizationeventtype.Columns, sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt))
	if ps := oetu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oetu.mutation.Name(); ok {
		_spec.SetField(organizationeventtype.FieldName, field.TypeString, value)
	}
	if value, ok := oetu.mutation.Description(); ok {
		_spec.SetField(organizationeventtype.FieldDescription, field.TypeString, value)
	}
	if value, ok := oetu.mutation.Rules(); ok {
		_spec.SetField(organizationeventtype.FieldRules, field.TypeJSON, value)
	}
	if value, ok := oetu.mutation.AppendedRules(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationeventtype.FieldRules, value)
		})
	}
	if oetu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationeventtype.EventsTable,
			Columns: []string{organizationeventtype.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oetu.mutation.RemovedEventsIDs(); len(nodes) > 0 && !oetu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationeventtype.EventsTable,
			Columns: []string{organizationeventtype.EventsColumn},
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
	if nodes := oetu.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationeventtype.EventsTable,
			Columns: []string{organizationeventtype.EventsColumn},
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
	if oetu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationeventtype.OrganizationTable,
			Columns: []string{organizationeventtype.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oetu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationeventtype.OrganizationTable,
			Columns: []string{organizationeventtype.OrganizationColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, oetu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationeventtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oetu.mutation.done = true
	return n, nil
}

// OrganizationEventTypeUpdateOne is the builder for updating a single OrganizationEventType entity.
type OrganizationEventTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationEventTypeMutation
}

// SetOrganizationID sets the "organization_id" field.
func (oetuo *OrganizationEventTypeUpdateOne) SetOrganizationID(u uuid.UUID) *OrganizationEventTypeUpdateOne {
	oetuo.mutation.SetOrganizationID(u)
	return oetuo
}

// SetName sets the "name" field.
func (oetuo *OrganizationEventTypeUpdateOne) SetName(s string) *OrganizationEventTypeUpdateOne {
	oetuo.mutation.SetName(s)
	return oetuo
}

// SetDescription sets the "description" field.
func (oetuo *OrganizationEventTypeUpdateOne) SetDescription(s string) *OrganizationEventTypeUpdateOne {
	oetuo.mutation.SetDescription(s)
	return oetuo
}

// SetRules sets the "rules" field.
func (oetuo *OrganizationEventTypeUpdateOne) SetRules(sr []schema.EventRules) *OrganizationEventTypeUpdateOne {
	oetuo.mutation.SetRules(sr)
	return oetuo
}

// AppendRules appends sr to the "rules" field.
func (oetuo *OrganizationEventTypeUpdateOne) AppendRules(sr []schema.EventRules) *OrganizationEventTypeUpdateOne {
	oetuo.mutation.AppendRules(sr)
	return oetuo
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (oetuo *OrganizationEventTypeUpdateOne) AddEventIDs(ids ...uuid.UUID) *OrganizationEventTypeUpdateOne {
	oetuo.mutation.AddEventIDs(ids...)
	return oetuo
}

// AddEvents adds the "events" edges to the Event entity.
func (oetuo *OrganizationEventTypeUpdateOne) AddEvents(e ...*Event) *OrganizationEventTypeUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return oetuo.AddEventIDs(ids...)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (oetuo *OrganizationEventTypeUpdateOne) SetOrganization(o *Organization) *OrganizationEventTypeUpdateOne {
	return oetuo.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationEventTypeMutation object of the builder.
func (oetuo *OrganizationEventTypeUpdateOne) Mutation() *OrganizationEventTypeMutation {
	return oetuo.mutation
}

// ClearEvents clears all "events" edges to the Event entity.
func (oetuo *OrganizationEventTypeUpdateOne) ClearEvents() *OrganizationEventTypeUpdateOne {
	oetuo.mutation.ClearEvents()
	return oetuo
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (oetuo *OrganizationEventTypeUpdateOne) RemoveEventIDs(ids ...uuid.UUID) *OrganizationEventTypeUpdateOne {
	oetuo.mutation.RemoveEventIDs(ids...)
	return oetuo
}

// RemoveEvents removes "events" edges to Event entities.
func (oetuo *OrganizationEventTypeUpdateOne) RemoveEvents(e ...*Event) *OrganizationEventTypeUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return oetuo.RemoveEventIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (oetuo *OrganizationEventTypeUpdateOne) ClearOrganization() *OrganizationEventTypeUpdateOne {
	oetuo.mutation.ClearOrganization()
	return oetuo
}

// Where appends a list predicates to the OrganizationEventTypeUpdate builder.
func (oetuo *OrganizationEventTypeUpdateOne) Where(ps ...predicate.OrganizationEventType) *OrganizationEventTypeUpdateOne {
	oetuo.mutation.Where(ps...)
	return oetuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oetuo *OrganizationEventTypeUpdateOne) Select(field string, fields ...string) *OrganizationEventTypeUpdateOne {
	oetuo.fields = append([]string{field}, fields...)
	return oetuo
}

// Save executes the query and returns the updated OrganizationEventType entity.
func (oetuo *OrganizationEventTypeUpdateOne) Save(ctx context.Context) (*OrganizationEventType, error) {
	return withHooks[*OrganizationEventType, OrganizationEventTypeMutation](ctx, oetuo.sqlSave, oetuo.mutation, oetuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oetuo *OrganizationEventTypeUpdateOne) SaveX(ctx context.Context) *OrganizationEventType {
	node, err := oetuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oetuo *OrganizationEventTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := oetuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oetuo *OrganizationEventTypeUpdateOne) ExecX(ctx context.Context) {
	if err := oetuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oetuo *OrganizationEventTypeUpdateOne) check() error {
	if v, ok := oetuo.mutation.Name(); ok {
		if err := organizationeventtype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrganizationEventType.name": %w`, err)}
		}
	}
	if _, ok := oetuo.mutation.OrganizationID(); oetuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrganizationEventType.organization"`)
	}
	return nil
}

func (oetuo *OrganizationEventTypeUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationEventType, err error) {
	if err := oetuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationeventtype.Table, organizationeventtype.Columns, sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt))
	id, ok := oetuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrganizationEventType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oetuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationeventtype.FieldID)
		for _, f := range fields {
			if !organizationeventtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organizationeventtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oetuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oetuo.mutation.Name(); ok {
		_spec.SetField(organizationeventtype.FieldName, field.TypeString, value)
	}
	if value, ok := oetuo.mutation.Description(); ok {
		_spec.SetField(organizationeventtype.FieldDescription, field.TypeString, value)
	}
	if value, ok := oetuo.mutation.Rules(); ok {
		_spec.SetField(organizationeventtype.FieldRules, field.TypeJSON, value)
	}
	if value, ok := oetuo.mutation.AppendedRules(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationeventtype.FieldRules, value)
		})
	}
	if oetuo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationeventtype.EventsTable,
			Columns: []string{organizationeventtype.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oetuo.mutation.RemovedEventsIDs(); len(nodes) > 0 && !oetuo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationeventtype.EventsTable,
			Columns: []string{organizationeventtype.EventsColumn},
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
	if nodes := oetuo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationeventtype.EventsTable,
			Columns: []string{organizationeventtype.EventsColumn},
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
	if oetuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationeventtype.OrganizationTable,
			Columns: []string{organizationeventtype.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oetuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationeventtype.OrganizationTable,
			Columns: []string{organizationeventtype.OrganizationColumn},
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
	_node = &OrganizationEventType{config: oetuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oetuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationeventtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oetuo.mutation.done = true
	return _node, nil
}
