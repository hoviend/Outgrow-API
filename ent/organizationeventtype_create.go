// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"outgrow/ent/event"
	"outgrow/ent/organization"
	"outgrow/ent/organizationeventtype"
	"outgrow/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrganizationEventTypeCreate is the builder for creating a OrganizationEventType entity.
type OrganizationEventTypeCreate struct {
	config
	mutation *OrganizationEventTypeMutation
	hooks    []Hook
}

// SetOrganizationID sets the "organization_id" field.
func (oetc *OrganizationEventTypeCreate) SetOrganizationID(u uuid.UUID) *OrganizationEventTypeCreate {
	oetc.mutation.SetOrganizationID(u)
	return oetc
}

// SetName sets the "name" field.
func (oetc *OrganizationEventTypeCreate) SetName(s string) *OrganizationEventTypeCreate {
	oetc.mutation.SetName(s)
	return oetc
}

// SetDescription sets the "description" field.
func (oetc *OrganizationEventTypeCreate) SetDescription(s string) *OrganizationEventTypeCreate {
	oetc.mutation.SetDescription(s)
	return oetc
}

// SetRules sets the "rules" field.
func (oetc *OrganizationEventTypeCreate) SetRules(sr []schema.EventRules) *OrganizationEventTypeCreate {
	oetc.mutation.SetRules(sr)
	return oetc
}

// SetCreatedAt sets the "created_at" field.
func (oetc *OrganizationEventTypeCreate) SetCreatedAt(t time.Time) *OrganizationEventTypeCreate {
	oetc.mutation.SetCreatedAt(t)
	return oetc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oetc *OrganizationEventTypeCreate) SetNillableCreatedAt(t *time.Time) *OrganizationEventTypeCreate {
	if t != nil {
		oetc.SetCreatedAt(*t)
	}
	return oetc
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (oetc *OrganizationEventTypeCreate) AddEventIDs(ids ...uuid.UUID) *OrganizationEventTypeCreate {
	oetc.mutation.AddEventIDs(ids...)
	return oetc
}

// AddEvents adds the "events" edges to the Event entity.
func (oetc *OrganizationEventTypeCreate) AddEvents(e ...*Event) *OrganizationEventTypeCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return oetc.AddEventIDs(ids...)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (oetc *OrganizationEventTypeCreate) SetOrganization(o *Organization) *OrganizationEventTypeCreate {
	return oetc.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationEventTypeMutation object of the builder.
func (oetc *OrganizationEventTypeCreate) Mutation() *OrganizationEventTypeMutation {
	return oetc.mutation
}

// Save creates the OrganizationEventType in the database.
func (oetc *OrganizationEventTypeCreate) Save(ctx context.Context) (*OrganizationEventType, error) {
	oetc.defaults()
	return withHooks[*OrganizationEventType, OrganizationEventTypeMutation](ctx, oetc.sqlSave, oetc.mutation, oetc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oetc *OrganizationEventTypeCreate) SaveX(ctx context.Context) *OrganizationEventType {
	v, err := oetc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oetc *OrganizationEventTypeCreate) Exec(ctx context.Context) error {
	_, err := oetc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oetc *OrganizationEventTypeCreate) ExecX(ctx context.Context) {
	if err := oetc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oetc *OrganizationEventTypeCreate) defaults() {
	if _, ok := oetc.mutation.CreatedAt(); !ok {
		v := organizationeventtype.DefaultCreatedAt()
		oetc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oetc *OrganizationEventTypeCreate) check() error {
	if _, ok := oetc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "OrganizationEventType.organization_id"`)}
	}
	if _, ok := oetc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OrganizationEventType.name"`)}
	}
	if v, ok := oetc.mutation.Name(); ok {
		if err := organizationeventtype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrganizationEventType.name": %w`, err)}
		}
	}
	if _, ok := oetc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "OrganizationEventType.description"`)}
	}
	if _, ok := oetc.mutation.Rules(); !ok {
		return &ValidationError{Name: "rules", err: errors.New(`ent: missing required field "OrganizationEventType.rules"`)}
	}
	if _, ok := oetc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrganizationEventType.created_at"`)}
	}
	if _, ok := oetc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "OrganizationEventType.organization"`)}
	}
	return nil
}

func (oetc *OrganizationEventTypeCreate) sqlSave(ctx context.Context) (*OrganizationEventType, error) {
	if err := oetc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oetc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oetc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oetc.mutation.id = &_node.ID
	oetc.mutation.done = true
	return _node, nil
}

func (oetc *OrganizationEventTypeCreate) createSpec() (*OrganizationEventType, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationEventType{config: oetc.config}
		_spec = sqlgraph.NewCreateSpec(organizationeventtype.Table, sqlgraph.NewFieldSpec(organizationeventtype.FieldID, field.TypeInt))
	)
	if value, ok := oetc.mutation.Name(); ok {
		_spec.SetField(organizationeventtype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := oetc.mutation.Description(); ok {
		_spec.SetField(organizationeventtype.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := oetc.mutation.Rules(); ok {
		_spec.SetField(organizationeventtype.FieldRules, field.TypeJSON, value)
		_node.Rules = value
	}
	if value, ok := oetc.mutation.CreatedAt(); ok {
		_spec.SetField(organizationeventtype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := oetc.mutation.EventsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oetc.mutation.OrganizationIDs(); len(nodes) > 0 {
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
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationEventTypeCreateBulk is the builder for creating many OrganizationEventType entities in bulk.
type OrganizationEventTypeCreateBulk struct {
	config
	builders []*OrganizationEventTypeCreate
}

// Save creates the OrganizationEventType entities in the database.
func (oetcb *OrganizationEventTypeCreateBulk) Save(ctx context.Context) ([]*OrganizationEventType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oetcb.builders))
	nodes := make([]*OrganizationEventType, len(oetcb.builders))
	mutators := make([]Mutator, len(oetcb.builders))
	for i := range oetcb.builders {
		func(i int, root context.Context) {
			builder := oetcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationEventTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, oetcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oetcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oetcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oetcb *OrganizationEventTypeCreateBulk) SaveX(ctx context.Context) []*OrganizationEventType {
	v, err := oetcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oetcb *OrganizationEventTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := oetcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oetcb *OrganizationEventTypeCreateBulk) ExecX(ctx context.Context) {
	if err := oetcb.Exec(ctx); err != nil {
		panic(err)
	}
}
