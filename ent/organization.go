// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"outgrow/ent/organization"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// PublicID holds the value of the "public_id" field.
	PublicID string `json:"public_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges        OrganizationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// AccountTypes holds the value of the accountTypes edge.
	AccountTypes []*OrganizationAccountType `json:"accountTypes,omitempty"`
	// EventTypes holds the value of the eventTypes edge.
	EventTypes []*OrganizationEventType `json:"eventTypes,omitempty"`
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// AccountTypesOrErr returns the AccountTypes value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) AccountTypesOrErr() ([]*OrganizationAccountType, error) {
	if e.loadedTypes[1] {
		return e.AccountTypes, nil
	}
	return nil, &NotLoadedError{edge: "accountTypes"}
}

// EventTypesOrErr returns the EventTypes value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) EventTypesOrErr() ([]*OrganizationEventType, error) {
	if e.loadedTypes[2] {
		return e.EventTypes, nil
	}
	return nil, &NotLoadedError{edge: "eventTypes"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[3] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldName, organization.FieldPublicID:
			values[i] = new(sql.NullString)
		case organization.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case organization.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		case organization.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case organization.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case organization.FieldPublicID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_id", values[i])
			} else if value.Valid {
				o.PublicID = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Organization.
// This includes values selected through modifiers, order, etc.
func (o *Organization) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Organization entity.
func (o *Organization) QueryUsers() *UserQuery {
	return NewOrganizationClient(o.config).QueryUsers(o)
}

// QueryAccountTypes queries the "accountTypes" edge of the Organization entity.
func (o *Organization) QueryAccountTypes() *OrganizationAccountTypeQuery {
	return NewOrganizationClient(o.config).QueryAccountTypes(o)
}

// QueryEventTypes queries the "eventTypes" edge of the Organization entity.
func (o *Organization) QueryEventTypes() *OrganizationEventTypeQuery {
	return NewOrganizationClient(o.config).QueryEventTypes(o)
}

// QueryEvents queries the "events" edge of the Organization entity.
func (o *Organization) QueryEvents() *EventQuery {
	return NewOrganizationClient(o.config).QueryEvents(o)
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return NewOrganizationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("public_id=")
	builder.WriteString(o.PublicID)
	builder.WriteByte(')')
	return builder.String()
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization
