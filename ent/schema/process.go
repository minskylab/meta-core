package schema

import (
	"time"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Process holds the schema definition for the Process entity.
type Process struct {
	ent.Schema
}

// Fields of the Process.
func (Process) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.String("name").Optional(),
		field.String("token"),
		field.String("instance_type"),
		field.String("ami_id"),
		field.String("resource_prefix"),
		field.Int("timeout").NonNegative().Optional(),
		field.Int("expiration").NonNegative(),
		field.Enum("state").Values("PENDING", "PROVISIONING", "DEPLOYING", "RUNNING", "IDLE", "STOPPING", "EXIT_TIMEOUT", "EXIT_SUCCESS").Default("PENDING"),
	}
}

// Edges of the Process.
func (Process) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stack", Stack.Type).Unique(),
		edge.To("credentials", Credential.Type),
		edge.To("tasks", Task.Type),
	}
}
