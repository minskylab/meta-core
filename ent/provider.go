// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/minskylab/meta-core/ent/provider"

	uuid "github.com/satori/go.uuid"
)

// Provider is the model entity for the Provider schema.
type Provider struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Hostname holds the value of the "hostname" field.
	Hostname string `json:"hostname,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProviderQuery when eager-loading is set.
	Edges ProviderEdges `json:"edges"`
}

// ProviderEdges holds the relations/edges for other nodes in the graph.
type ProviderEdges struct {
	// Deployment holds the value of the deployment edge.
	Deployment []*Deployment `json:"deployment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DeploymentOrErr returns the Deployment value or an error if the edge
// was not loaded in eager-loading.
func (e ProviderEdges) DeploymentOrErr() ([]*Deployment, error) {
	if e.loadedTypes[0] {
		return e.Deployment, nil
	}
	return nil, &NotLoadedError{edge: "deployment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Provider) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case provider.FieldHostname, provider.FieldToken:
			values[i] = new(sql.NullString)
		case provider.FieldCreatedAt, provider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case provider.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Provider", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Provider fields.
func (pr *Provider) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provider.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case provider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case provider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case provider.FieldHostname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hostname", values[i])
			} else if value.Valid {
				pr.Hostname = value.String
			}
		case provider.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				pr.Token = value.String
			}
		}
	}
	return nil
}

// QueryDeployment queries the "deployment" edge of the Provider entity.
func (pr *Provider) QueryDeployment() *DeploymentQuery {
	return (&ProviderClient{config: pr.config}).QueryDeployment(pr)
}

// Update returns a builder for updating this Provider.
// Note that you need to call Provider.Unwrap() before calling this method if this Provider
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Provider) Update() *ProviderUpdateOne {
	return (&ProviderClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Provider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Provider) Unwrap() *Provider {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Provider is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Provider) String() string {
	var builder strings.Builder
	builder.WriteString("Provider(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", hostname=")
	builder.WriteString(pr.Hostname)
	builder.WriteString(", token=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// Providers is a parsable slice of Provider.
type Providers []*Provider

func (pr Providers) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
