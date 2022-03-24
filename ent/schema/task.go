package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	uuid "github.com/satori/go.uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.NewV4).Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.String("image"),
		field.Int("timeout").Optional(),
		field.String("name").Optional(),
		field.String("cmd").Optional(),
		field.Bool("detached").Default(false).Optional(),
		field.JSON("environment", map[string]interface{}{}).Optional(),
		field.JSON("ports", []string{}).Optional(),
		field.JSON("volumes", []string{}).Optional(),
		field.String("restart").Optional(),
		field.JSON("security_opt", []string{}).Optional(),
		field.JSON("cap_add", []string{}).Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment", Deployment.Type).Ref("tasks").Unique(),
	}
}
