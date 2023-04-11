// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"outgrow/ent/event"
	"outgrow/ent/organization"
	"outgrow/ent/organizationeventtype"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organization_id,omitempty"`
	// EventTypeID holds the value of the "event_type_id" field.
	EventTypeID int `json:"event_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes string `json:"notes,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges EventEdges `json:"edges"`
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// Type holds the value of the type edge.
	Type *OrganizationEventType `json:"type,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[0] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// TypeOrErr returns the Type value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventEdges) TypeOrErr() (*OrganizationEventType, error) {
	if e.loadedTypes[1] {
		if e.Type == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organizationeventtype.Label}
		}
		return e.Type, nil
	}
	return nil, &NotLoadedError{edge: "type"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventEdges) OrganizationOrErr() (*Organization, error) {
	if e.loadedTypes[2] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case event.FieldEventTypeID:
			values[i] = new(sql.NullInt64)
		case event.FieldNotes:
			values[i] = new(sql.NullString)
		case event.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case event.FieldID, event.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Event", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case event.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				e.OrganizationID = *value
			}
		case event.FieldEventTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field event_type_id", values[i])
			} else if value.Valid {
				e.EventTypeID = int(value.Int64)
			}
		case event.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				e.Amount = value.Float64
			}
		case event.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				e.Notes = value.String
			}
		case event.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTransactions queries the "transactions" edge of the Event entity.
func (e *Event) QueryTransactions() *TransactionQuery {
	return NewEventClient(e.config).QueryTransactions(e)
}

// QueryType queries the "type" edge of the Event entity.
func (e *Event) QueryType() *OrganizationEventTypeQuery {
	return NewEventClient(e.config).QueryType(e)
}

// QueryOrganization queries the "organization" edge of the Event entity.
func (e *Event) QueryOrganization() *OrganizationQuery {
	return NewEventClient(e.config).QueryOrganization(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return NewEventClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", e.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("event_type_id=")
	builder.WriteString(fmt.Sprintf("%v", e.EventTypeID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", e.Amount))
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(e.Notes)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event