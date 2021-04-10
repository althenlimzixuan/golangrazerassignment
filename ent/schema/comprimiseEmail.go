package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type ComprimiseEmail struct {
	ent.Schema
}

// Fields of the User.
func (ComprimiseEmail) Fields() []ent.Field {
	return []ent.Field{
		field.String("Email").
			Default("unknown"),
	}
}
