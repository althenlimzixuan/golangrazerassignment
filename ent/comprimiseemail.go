// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hello/ent/comprimiseemail"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ComprimiseEmail is the model entity for the ComprimiseEmail schema.
type ComprimiseEmail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ComprimiseEmail) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comprimiseemail.FieldID:
			values[i] = &sql.NullInt64{}
		case comprimiseemail.FieldEmail:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ComprimiseEmail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ComprimiseEmail fields.
func (ce *ComprimiseEmail) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comprimiseemail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ce.ID = int(value.Int64)
		case comprimiseemail.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Email", values[i])
			} else if value.Valid {
				ce.Email = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ComprimiseEmail.
// Note that you need to call ComprimiseEmail.Unwrap() before calling this method if this ComprimiseEmail
// was returned from a transaction, and the transaction was committed or rolled back.
func (ce *ComprimiseEmail) Update() *ComprimiseEmailUpdateOne {
	return (&ComprimiseEmailClient{config: ce.config}).UpdateOne(ce)
}

// Unwrap unwraps the ComprimiseEmail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ce *ComprimiseEmail) Unwrap() *ComprimiseEmail {
	tx, ok := ce.config.driver.(*txDriver)
	if !ok {
		panic("ent: ComprimiseEmail is not a transactional entity")
	}
	ce.config.driver = tx.drv
	return ce
}

// String implements the fmt.Stringer.
func (ce *ComprimiseEmail) String() string {
	var builder strings.Builder
	builder.WriteString("ComprimiseEmail(")
	builder.WriteString(fmt.Sprintf("id=%v", ce.ID))
	builder.WriteString(", Email=")
	builder.WriteString(ce.Email)
	builder.WriteByte(')')
	return builder.String()
}

// ComprimiseEmails is a parsable slice of ComprimiseEmail.
type ComprimiseEmails []*ComprimiseEmail

func (ce ComprimiseEmails) config(cfg config) {
	for _i := range ce {
		ce[_i].config = cfg
	}
}
