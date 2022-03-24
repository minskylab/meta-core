package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stack holds the schema definition for the Stack entity.
type Stack struct {
	ent.Schema
}

// Fields of the Stack.
func (Stack) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.String("instance"),
		field.String("vpc_id"),
		field.String("public_ip"),
		field.String("public_dns").Optional(),
		field.String("username"),

		field.String("security_group"),

		field.String("key_pair"),
		field.String("name"),
		field.String("private_key"),
		field.String("filepath").Optional(),
		// field.Enum("state").Values("pending", "running", "stopped", "deleted").Default("pending"),
		// field.Float("cpu").Positive(),
		// field.Float("memory").Positive(),
	}
}

// Edges of the Stack.
func (Stack) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment", Deployment.Type).Ref("stack").Unique(), //.Required(),
		// edge.From("provider", Provider.Type).Unique().Ref("stacks"),
	}
}
