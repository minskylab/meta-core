// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/meta-core/ent/credential"
	"github.com/minskylab/meta-core/ent/deployment"
	"github.com/minskylab/meta-core/ent/predicate"
	"github.com/minskylab/meta-core/ent/provider"
	"github.com/minskylab/meta-core/ent/stack"
	"github.com/minskylab/meta-core/ent/task"
	uuid "github.com/satori/go.uuid"
)

// DeploymentUpdate is the builder for updating Deployment entities.
type DeploymentUpdate struct {
	config
	hooks    []Hook
	mutation *DeploymentMutation
}

// Where appends a list predicates to the DeploymentUpdate builder.
func (du *DeploymentUpdate) Where(ps ...predicate.Deployment) *DeploymentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DeploymentUpdate) SetCreatedAt(t time.Time) *DeploymentUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableCreatedAt(t *time.Time) *DeploymentUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DeploymentUpdate) SetUpdatedAt(t time.Time) *DeploymentUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetState sets the "state" field.
func (du *DeploymentUpdate) SetState(d deployment.State) *DeploymentUpdate {
	du.mutation.SetState(d)
	return du
}

// SetNillableState sets the "state" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableState(d *deployment.State) *DeploymentUpdate {
	if d != nil {
		du.SetState(*d)
	}
	return du
}

// SetName sets the "name" field.
func (du *DeploymentUpdate) SetName(s string) *DeploymentUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableName(s *string) *DeploymentUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// ClearName clears the value of the "name" field.
func (du *DeploymentUpdate) ClearName() *DeploymentUpdate {
	du.mutation.ClearName()
	return du
}

// SetTimeout sets the "timeout" field.
func (du *DeploymentUpdate) SetTimeout(i int) *DeploymentUpdate {
	du.mutation.ResetTimeout()
	du.mutation.SetTimeout(i)
	return du
}

// AddTimeout adds i to the "timeout" field.
func (du *DeploymentUpdate) AddTimeout(i int) *DeploymentUpdate {
	du.mutation.AddTimeout(i)
	return du
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (du *DeploymentUpdate) AddTaskIDs(ids ...uuid.UUID) *DeploymentUpdate {
	du.mutation.AddTaskIDs(ids...)
	return du
}

// AddTasks adds the "tasks" edges to the Task entity.
func (du *DeploymentUpdate) AddTasks(t ...*Task) *DeploymentUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return du.AddTaskIDs(ids...)
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (du *DeploymentUpdate) SetProviderID(id uuid.UUID) *DeploymentUpdate {
	du.mutation.SetProviderID(id)
	return du
}

// SetProvider sets the "provider" edge to the Provider entity.
func (du *DeploymentUpdate) SetProvider(p *Provider) *DeploymentUpdate {
	return du.SetProviderID(p.ID)
}

// SetStackID sets the "stack" edge to the Stack entity by ID.
func (du *DeploymentUpdate) SetStackID(id string) *DeploymentUpdate {
	du.mutation.SetStackID(id)
	return du
}

// SetNillableStackID sets the "stack" edge to the Stack entity by ID if the given value is not nil.
func (du *DeploymentUpdate) SetNillableStackID(id *string) *DeploymentUpdate {
	if id != nil {
		du = du.SetStackID(*id)
	}
	return du
}

// SetStack sets the "stack" edge to the Stack entity.
func (du *DeploymentUpdate) SetStack(s *Stack) *DeploymentUpdate {
	return du.SetStackID(s.ID)
}

// AddCredentialIDs adds the "credentials" edge to the Credential entity by IDs.
func (du *DeploymentUpdate) AddCredentialIDs(ids ...uuid.UUID) *DeploymentUpdate {
	du.mutation.AddCredentialIDs(ids...)
	return du
}

// AddCredentials adds the "credentials" edges to the Credential entity.
func (du *DeploymentUpdate) AddCredentials(c ...*Credential) *DeploymentUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.AddCredentialIDs(ids...)
}

// Mutation returns the DeploymentMutation object of the builder.
func (du *DeploymentUpdate) Mutation() *DeploymentMutation {
	return du.mutation
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (du *DeploymentUpdate) ClearTasks() *DeploymentUpdate {
	du.mutation.ClearTasks()
	return du
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (du *DeploymentUpdate) RemoveTaskIDs(ids ...uuid.UUID) *DeploymentUpdate {
	du.mutation.RemoveTaskIDs(ids...)
	return du
}

// RemoveTasks removes "tasks" edges to Task entities.
func (du *DeploymentUpdate) RemoveTasks(t ...*Task) *DeploymentUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return du.RemoveTaskIDs(ids...)
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (du *DeploymentUpdate) ClearProvider() *DeploymentUpdate {
	du.mutation.ClearProvider()
	return du
}

// ClearStack clears the "stack" edge to the Stack entity.
func (du *DeploymentUpdate) ClearStack() *DeploymentUpdate {
	du.mutation.ClearStack()
	return du
}

// ClearCredentials clears all "credentials" edges to the Credential entity.
func (du *DeploymentUpdate) ClearCredentials() *DeploymentUpdate {
	du.mutation.ClearCredentials()
	return du
}

// RemoveCredentialIDs removes the "credentials" edge to Credential entities by IDs.
func (du *DeploymentUpdate) RemoveCredentialIDs(ids ...uuid.UUID) *DeploymentUpdate {
	du.mutation.RemoveCredentialIDs(ids...)
	return du
}

// RemoveCredentials removes "credentials" edges to Credential entities.
func (du *DeploymentUpdate) RemoveCredentials(c ...*Credential) *DeploymentUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.RemoveCredentialIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeploymentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeploymentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeploymentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeploymentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DeploymentUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := deployment.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeploymentUpdate) check() error {
	if v, ok := du.mutation.State(); ok {
		if err := deployment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if v, ok := du.mutation.Timeout(); ok {
		if err := deployment.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	if _, ok := du.mutation.ProviderID(); du.mutation.ProviderCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"provider\"")
	}
	return nil
}

func (du *DeploymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deployment.Table,
			Columns: deployment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deployment.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: deployment.FieldState,
		})
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldName,
		})
	}
	if du.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deployment.FieldName,
		})
	}
	if value, ok := du.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deployment.FieldTimeout,
		})
	}
	if value, ok := du.mutation.AddedTimeout(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deployment.FieldTimeout,
		})
	}
	if du.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedTasksIDs(); len(nodes) > 0 && !du.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.TasksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.ProviderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ProviderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.StackCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.StackIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.CredentialsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedCredentialsIDs(); len(nodes) > 0 && !du.mutation.CredentialsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.CredentialsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deployment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DeploymentUpdateOne is the builder for updating a single Deployment entity.
type DeploymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeploymentMutation
}

// SetCreatedAt sets the "created_at" field.
func (duo *DeploymentUpdateOne) SetCreatedAt(t time.Time) *DeploymentUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableCreatedAt(t *time.Time) *DeploymentUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DeploymentUpdateOne) SetUpdatedAt(t time.Time) *DeploymentUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetState sets the "state" field.
func (duo *DeploymentUpdateOne) SetState(d deployment.State) *DeploymentUpdateOne {
	duo.mutation.SetState(d)
	return duo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableState(d *deployment.State) *DeploymentUpdateOne {
	if d != nil {
		duo.SetState(*d)
	}
	return duo
}

// SetName sets the "name" field.
func (duo *DeploymentUpdateOne) SetName(s string) *DeploymentUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableName(s *string) *DeploymentUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// ClearName clears the value of the "name" field.
func (duo *DeploymentUpdateOne) ClearName() *DeploymentUpdateOne {
	duo.mutation.ClearName()
	return duo
}

// SetTimeout sets the "timeout" field.
func (duo *DeploymentUpdateOne) SetTimeout(i int) *DeploymentUpdateOne {
	duo.mutation.ResetTimeout()
	duo.mutation.SetTimeout(i)
	return duo
}

// AddTimeout adds i to the "timeout" field.
func (duo *DeploymentUpdateOne) AddTimeout(i int) *DeploymentUpdateOne {
	duo.mutation.AddTimeout(i)
	return duo
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (duo *DeploymentUpdateOne) AddTaskIDs(ids ...uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.AddTaskIDs(ids...)
	return duo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (duo *DeploymentUpdateOne) AddTasks(t ...*Task) *DeploymentUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return duo.AddTaskIDs(ids...)
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (duo *DeploymentUpdateOne) SetProviderID(id uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.SetProviderID(id)
	return duo
}

// SetProvider sets the "provider" edge to the Provider entity.
func (duo *DeploymentUpdateOne) SetProvider(p *Provider) *DeploymentUpdateOne {
	return duo.SetProviderID(p.ID)
}

// SetStackID sets the "stack" edge to the Stack entity by ID.
func (duo *DeploymentUpdateOne) SetStackID(id string) *DeploymentUpdateOne {
	duo.mutation.SetStackID(id)
	return duo
}

// SetNillableStackID sets the "stack" edge to the Stack entity by ID if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableStackID(id *string) *DeploymentUpdateOne {
	if id != nil {
		duo = duo.SetStackID(*id)
	}
	return duo
}

// SetStack sets the "stack" edge to the Stack entity.
func (duo *DeploymentUpdateOne) SetStack(s *Stack) *DeploymentUpdateOne {
	return duo.SetStackID(s.ID)
}

// AddCredentialIDs adds the "credentials" edge to the Credential entity by IDs.
func (duo *DeploymentUpdateOne) AddCredentialIDs(ids ...uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.AddCredentialIDs(ids...)
	return duo
}

// AddCredentials adds the "credentials" edges to the Credential entity.
func (duo *DeploymentUpdateOne) AddCredentials(c ...*Credential) *DeploymentUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.AddCredentialIDs(ids...)
}

// Mutation returns the DeploymentMutation object of the builder.
func (duo *DeploymentUpdateOne) Mutation() *DeploymentMutation {
	return duo.mutation
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (duo *DeploymentUpdateOne) ClearTasks() *DeploymentUpdateOne {
	duo.mutation.ClearTasks()
	return duo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (duo *DeploymentUpdateOne) RemoveTaskIDs(ids ...uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.RemoveTaskIDs(ids...)
	return duo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (duo *DeploymentUpdateOne) RemoveTasks(t ...*Task) *DeploymentUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return duo.RemoveTaskIDs(ids...)
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (duo *DeploymentUpdateOne) ClearProvider() *DeploymentUpdateOne {
	duo.mutation.ClearProvider()
	return duo
}

// ClearStack clears the "stack" edge to the Stack entity.
func (duo *DeploymentUpdateOne) ClearStack() *DeploymentUpdateOne {
	duo.mutation.ClearStack()
	return duo
}

// ClearCredentials clears all "credentials" edges to the Credential entity.
func (duo *DeploymentUpdateOne) ClearCredentials() *DeploymentUpdateOne {
	duo.mutation.ClearCredentials()
	return duo
}

// RemoveCredentialIDs removes the "credentials" edge to Credential entities by IDs.
func (duo *DeploymentUpdateOne) RemoveCredentialIDs(ids ...uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.RemoveCredentialIDs(ids...)
	return duo
}

// RemoveCredentials removes "credentials" edges to Credential entities.
func (duo *DeploymentUpdateOne) RemoveCredentials(c ...*Credential) *DeploymentUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.RemoveCredentialIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeploymentUpdateOne) Select(field string, fields ...string) *DeploymentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Deployment entity.
func (duo *DeploymentUpdateOne) Save(ctx context.Context) (*Deployment, error) {
	var (
		err  error
		node *Deployment
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeploymentUpdateOne) SaveX(ctx context.Context) *Deployment {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeploymentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeploymentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DeploymentUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := deployment.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeploymentUpdateOne) check() error {
	if v, ok := duo.mutation.State(); ok {
		if err := deployment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Timeout(); ok {
		if err := deployment.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	if _, ok := duo.mutation.ProviderID(); duo.mutation.ProviderCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"provider\"")
	}
	return nil
}

func (duo *DeploymentUpdateOne) sqlSave(ctx context.Context) (_node *Deployment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deployment.Table,
			Columns: deployment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deployment.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Deployment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deployment.FieldID)
		for _, f := range fields {
			if !deployment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deployment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: deployment.FieldState,
		})
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldName,
		})
	}
	if duo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deployment.FieldName,
		})
	}
	if value, ok := duo.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deployment.FieldTimeout,
		})
	}
	if value, ok := duo.mutation.AddedTimeout(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: deployment.FieldTimeout,
		})
	}
	if duo.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !duo.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.TasksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.ProviderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ProviderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.StackCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.StackIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.CredentialsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedCredentialsIDs(); len(nodes) > 0 && !duo.mutation.CredentialsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.CredentialsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Deployment{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deployment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
