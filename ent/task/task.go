// Code generated by entc, DO NOT EDIT.

package task

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldTimeout holds the string denoting the timeout field in the database.
	FieldTimeout = "timeout"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCmd holds the string denoting the cmd field in the database.
	FieldCmd = "cmd"
	// FieldDetached holds the string denoting the detached field in the database.
	FieldDetached = "detached"
	// FieldEnvironment holds the string denoting the environment field in the database.
	FieldEnvironment = "environment"
	// FieldPorts holds the string denoting the ports field in the database.
	FieldPorts = "ports"
	// FieldVolumes holds the string denoting the volumes field in the database.
	FieldVolumes = "volumes"
	// FieldRestart holds the string denoting the restart field in the database.
	FieldRestart = "restart"
	// FieldSecurityOpt holds the string denoting the security_opt field in the database.
	FieldSecurityOpt = "security_opt"
	// FieldCapAdd holds the string denoting the cap_add field in the database.
	FieldCapAdd = "cap_add"
	// EdgeDeployment holds the string denoting the deployment edge name in mutations.
	EdgeDeployment = "deployment"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// DeploymentTable is the table that holds the deployment relation/edge.
	DeploymentTable = "tasks"
	// DeploymentInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentInverseTable = "deployments"
	// DeploymentColumn is the table column denoting the deployment relation/edge.
	DeploymentColumn = "deployment_tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldImage,
	FieldTimeout,
	FieldName,
	FieldCmd,
	FieldDetached,
	FieldEnvironment,
	FieldPorts,
	FieldVolumes,
	FieldRestart,
	FieldSecurityOpt,
	FieldCapAdd,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"deployment_tasks",
	"process_tasks",
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
	// DefaultDetached holds the default value on creation for the "detached" field.
	DefaultDetached bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
