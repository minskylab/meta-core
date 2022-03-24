// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/meta-core/ent/deployment"

	"github.com/minskylab/meta-core/ent/credential"
	"github.com/minskylab/meta-core/ent/provider"
	"github.com/minskylab/meta-core/ent/stack"
	"github.com/minskylab/meta-core/ent/task"
	uuid "github.com/satori/go.uuid"
)

// DeploymentCreate is the builder for creating a Deployment entity.
type DeploymentCreate struct {
	config
	mutation *DeploymentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeploymentCreate) SetCreatedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableCreatedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DeploymentCreate) SetUpdatedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableUpdatedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetState sets the "state" field.
func (dc *DeploymentCreate) SetState(d deployment.State) *DeploymentCreate {
	dc.mutation.SetState(d)
	return dc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableState(d *deployment.State) *DeploymentCreate {
	if d != nil {
		dc.SetState(*d)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DeploymentCreate) SetName(s string) *DeploymentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableName(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetName(*s)
	}
	return dc
}

// SetTimeout sets the "timeout" field.
func (dc *DeploymentCreate) SetTimeout(i int) *DeploymentCreate {
	dc.mutation.SetTimeout(i)
	return dc
}

// SetID sets the "id" field.
func (dc *DeploymentCreate) SetID(u uuid.UUID) *DeploymentCreate {
	dc.mutation.SetID(u)
	return dc
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (dc *DeploymentCreate) AddTaskIDs(ids ...uuid.UUID) *DeploymentCreate {
	dc.mutation.AddTaskIDs(ids...)
	return dc
}

// AddTasks adds the "tasks" edges to the Task entity.
func (dc *DeploymentCreate) AddTasks(t ...*Task) *DeploymentCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return dc.AddTaskIDs(ids...)
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (dc *DeploymentCreate) SetProviderID(id uuid.UUID) *DeploymentCreate {
	dc.mutation.SetProviderID(id)
	return dc
}

// SetProvider sets the "provider" edge to the Provider entity.
func (dc *DeploymentCreate) SetProvider(p *Provider) *DeploymentCreate {
	return dc.SetProviderID(p.ID)
}

// SetStackID sets the "stack" edge to the Stack entity by ID.
func (dc *DeploymentCreate) SetStackID(id string) *DeploymentCreate {
	dc.mutation.SetStackID(id)
	return dc
}

// SetNillableStackID sets the "stack" edge to the Stack entity by ID if the given value is not nil.
func (dc *DeploymentCreate) SetNillableStackID(id *string) *DeploymentCreate {
	if id != nil {
		dc = dc.SetStackID(*id)
	}
	return dc
}

// SetStack sets the "stack" edge to the Stack entity.
func (dc *DeploymentCreate) SetStack(s *Stack) *DeploymentCreate {
	return dc.SetStackID(s.ID)
}

// AddCredentialIDs adds the "credentials" edge to the Credential entity by IDs.
func (dc *DeploymentCreate) AddCredentialIDs(ids ...uuid.UUID) *DeploymentCreate {
	dc.mutation.AddCredentialIDs(ids...)
	return dc
}

// AddCredentials adds the "credentials" edges to the Credential entity.
func (dc *DeploymentCreate) AddCredentials(c ...*Credential) *DeploymentCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dc.AddCredentialIDs(ids...)
}

// Mutation returns the DeploymentMutation object of the builder.
func (dc *DeploymentCreate) Mutation() *DeploymentMutation {
	return dc.mutation
}

// Save creates the Deployment in the database.
func (dc *DeploymentCreate) Save(ctx context.Context) (*Deployment, error) {
	var (
		err  error
		node *Deployment
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeploymentCreate) SaveX(ctx context.Context) *Deployment {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeploymentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeploymentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeploymentCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := deployment.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := deployment.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.State(); !ok {
		v := deployment.DefaultState
		dc.mutation.SetState(v)
	}
	if _, ok := dc.mutation.Name(); !ok {
		v := deployment.DefaultName
		dc.mutation.SetName(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := deployment.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeploymentCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := dc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := dc.mutation.State(); ok {
		if err := deployment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Timeout(); !ok {
		return &ValidationError{Name: "timeout", err: errors.New(`ent: missing required field "timeout"`)}
	}
	if v, ok := dc.mutation.Timeout(); ok {
		if err := deployment.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf(`ent: validator failed for field "timeout": %w`, err)}
		}
	}
	if len(dc.mutation.TasksIDs()) == 0 {
		return &ValidationError{Name: "tasks", err: errors.New("ent: missing required edge \"tasks\"")}
	}
	if _, ok := dc.mutation.ProviderID(); !ok {
		return &ValidationError{Name: "provider", err: errors.New("ent: missing required edge \"provider\"")}
	}
	return nil
}

func (dc *DeploymentCreate) sqlSave(ctx context.Context) (*Deployment, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (dc *DeploymentCreate) createSpec() (*Deployment, *sqlgraph.CreateSpec) {
	var (
		_node = &Deployment{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deployment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deployment.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: deployment.FieldState,
		})
		_node.State = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dc.mutation.Timeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deployment.FieldTimeout,
		})
		_node.Timeout = value
	}
	if nodes := dc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deployment.TasksTable,
			Columns: []string{deployment.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.ProviderTable,
			Columns: []string{deployment.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provider.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deployment_provider = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.StackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   deployment.StackTable,
			Columns: []string{deployment.StackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: stack.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.CredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deployment.CredentialsTable,
			Columns: []string{deployment.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: credential.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeploymentCreateBulk is the builder for creating many Deployment entities in bulk.
type DeploymentCreateBulk struct {
	config
	builders []*DeploymentCreate
}

// Save creates the Deployment entities in the database.
func (dcb *DeploymentCreateBulk) Save(ctx context.Context) ([]*Deployment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Deployment, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeploymentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) SaveX(ctx context.Context) []*Deployment {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeploymentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}