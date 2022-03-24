package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	uuid "github.com/satori/go.uuid"
)

// Deployment holds the schema definition for the Deployment entity.
type Deployment struct {
	ent.Schema
}

// Fields of the Deployment.
func (Deployment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.NewV4).Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.Enum("state").Values("pending", "provisioning", "running", "failed", "succeeded").Default("pending"),
		field.String("name").Default("unknown").Optional(),
		field.Int("timeout").NonNegative(),
		// field.Strings("logs").Optional(),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type).Required(),
		edge.To("provider", Provider.Type).Unique().Required(),

		edge.To("stack", Stack.Type).Unique(),
		edge.To("credentials", Credential.Type),

		// TODO: For the future, we need a Log entity.
		// edge.To("logs", Log.Type),
	}
}
