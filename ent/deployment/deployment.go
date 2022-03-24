// Code generated by entc, DO NOT EDIT.

package deployment

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	// Label holds the string label denoting the deployment type in the database.
	Label = "deployment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTimeout holds the string denoting the timeout field in the database.
	FieldTimeout = "timeout"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// EdgeProvider holds the string denoting the provider edge name in mutations.
	EdgeProvider = "provider"
	// EdgeStack holds the string denoting the stack edge name in mutations.
	EdgeStack = "stack"
	// EdgeCredentials holds the string denoting the credentials edge name in mutations.
	EdgeCredentials = "credentials"
	// Table holds the table name of the deployment in the database.
	Table = "deployments"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "deployment_tasks"
	// ProviderTable is the table that holds the provider relation/edge.
	ProviderTable = "deployments"
	// ProviderInverseTable is the table name for the Provider entity.
	// It exists in this package in order to avoid circular dependency with the "provider" package.
	ProviderInverseTable = "providers"
	// ProviderColumn is the table column denoting the provider relation/edge.
	ProviderColumn = "deployment_provider"
	// StackTable is the table that holds the stack relation/edge.
	StackTable = "stacks"
	// StackInverseTable is the table name for the Stack entity.
	// It exists in this package in order to avoid circular dependency with the "stack" package.
	StackInverseTable = "stacks"
	// StackColumn is the table column denoting the stack relation/edge.
	StackColumn = "deployment_stack"
	// CredentialsTable is the table that holds the credentials relation/edge.
	CredentialsTable = "credentials"
	// CredentialsInverseTable is the table name for the Credential entity.
	// It exists in this package in order to avoid circular dependency with the "credential" package.
	CredentialsInverseTable = "credentials"
	// CredentialsColumn is the table column denoting the credentials relation/edge.
	CredentialsColumn = "deployment_credentials"
)

// Columns holds all SQL columns for deployment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldState,
	FieldName,
	FieldTimeout,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "deployments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"deployment_provider",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// TimeoutValidator is a validator for the "timeout" field. It is called by the builders before save.
	TimeoutValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// State defines the type for the "state" enum field.
type State string

// StatePending is the default value of the State enum.
const DefaultState = StatePending

// State values.
const (
	StatePending      State = "pending"
	StateProvisioning State = "provisioning"
	StateRunning      State = "running"
	StateFailed       State = "failed"
	StateSucceeded    State = "succeeded"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StatePending, StateProvisioning, StateRunning, StateFailed, StateSucceeded:
		return nil
	default:
		return fmt.Errorf("deployment: invalid enum value for state field: %q", s)
	}
}