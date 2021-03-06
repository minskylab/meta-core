// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/meta-core/ent/credential"
	"github.com/minskylab/meta-core/ent/deployment"
	"github.com/minskylab/meta-core/ent/predicate"
	uuid "github.com/satori/go.uuid"
)

// CredentialUpdate is the builder for updating Credential entities.
type CredentialUpdate struct {
	config
	hooks    []Hook
	mutation *CredentialMutation
}

// Where appends a list predicates to the CredentialUpdate builder.
func (cu *CredentialUpdate) Where(ps ...predicate.Credential) *CredentialUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CredentialUpdate) SetCreatedAt(t time.Time) *CredentialUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CredentialUpdate) SetNillableCreatedAt(t *time.Time) *CredentialUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CredentialUpdate) SetUpdatedAt(t time.Time) *CredentialUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetRegistry sets the "registry" field.
func (cu *CredentialUpdate) SetRegistry(s string) *CredentialUpdate {
	cu.mutation.SetRegistry(s)
	return cu
}

// SetUsername sets the "username" field.
func (cu *CredentialUpdate) SetUsername(s string) *CredentialUpdate {
	cu.mutation.SetUsername(s)
	return cu
}

// SetPassword sets the "password" field.
func (cu *CredentialUpdate) SetPassword(s string) *CredentialUpdate {
	cu.mutation.SetPassword(s)
	return cu
}

// SetDeploymentID sets the "deployment" edge to the Deployment entity by ID.
func (cu *CredentialUpdate) SetDeploymentID(id uuid.UUID) *CredentialUpdate {
	cu.mutation.SetDeploymentID(id)
	return cu
}

// SetNillableDeploymentID sets the "deployment" edge to the Deployment entity by ID if the given value is not nil.
func (cu *CredentialUpdate) SetNillableDeploymentID(id *uuid.UUID) *CredentialUpdate {
	if id != nil {
		cu = cu.SetDeploymentID(*id)
	}
	return cu
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (cu *CredentialUpdate) SetDeployment(d *Deployment) *CredentialUpdate {
	return cu.SetDeploymentID(d.ID)
}

// Mutation returns the CredentialMutation object of the builder.
func (cu *CredentialUpdate) Mutation() *CredentialMutation {
	return cu.mutation
}

// ClearDeployment clears the "deployment" edge to the Deployment entity.
func (cu *CredentialUpdate) ClearDeployment() *CredentialUpdate {
	cu.mutation.ClearDeployment()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CredentialUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CredentialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CredentialUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CredentialUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CredentialUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := credential.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   credential.Table,
			Columns: credential.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: credential.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Registry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldRegistry,
		})
	}
	if value, ok := cu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldUsername,
		})
	}
	if value, ok := cu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPassword,
		})
	}
	if cu.mutation.DeploymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.DeploymentTable,
			Columns: []string{credential.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deployment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.DeploymentTable,
			Columns: []string{credential.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CredentialUpdateOne is the builder for updating a single Credential entity.
type CredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CredentialMutation
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CredentialUpdateOne) SetCreatedAt(t time.Time) *CredentialUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillableCreatedAt(t *time.Time) *CredentialUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CredentialUpdateOne) SetUpdatedAt(t time.Time) *CredentialUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetRegistry sets the "registry" field.
func (cuo *CredentialUpdateOne) SetRegistry(s string) *CredentialUpdateOne {
	cuo.mutation.SetRegistry(s)
	return cuo
}

// SetUsername sets the "username" field.
func (cuo *CredentialUpdateOne) SetUsername(s string) *CredentialUpdateOne {
	cuo.mutation.SetUsername(s)
	return cuo
}

// SetPassword sets the "password" field.
func (cuo *CredentialUpdateOne) SetPassword(s string) *CredentialUpdateOne {
	cuo.mutation.SetPassword(s)
	return cuo
}

// SetDeploymentID sets the "deployment" edge to the Deployment entity by ID.
func (cuo *CredentialUpdateOne) SetDeploymentID(id uuid.UUID) *CredentialUpdateOne {
	cuo.mutation.SetDeploymentID(id)
	return cuo
}

// SetNillableDeploymentID sets the "deployment" edge to the Deployment entity by ID if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillableDeploymentID(id *uuid.UUID) *CredentialUpdateOne {
	if id != nil {
		cuo = cuo.SetDeploymentID(*id)
	}
	return cuo
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (cuo *CredentialUpdateOne) SetDeployment(d *Deployment) *CredentialUpdateOne {
	return cuo.SetDeploymentID(d.ID)
}

// Mutation returns the CredentialMutation object of the builder.
func (cuo *CredentialUpdateOne) Mutation() *CredentialMutation {
	return cuo.mutation
}

// ClearDeployment clears the "deployment" edge to the Deployment entity.
func (cuo *CredentialUpdateOne) ClearDeployment() *CredentialUpdateOne {
	cuo.mutation.ClearDeployment()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CredentialUpdateOne) Select(field string, fields ...string) *CredentialUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Credential entity.
func (cuo *CredentialUpdateOne) Save(ctx context.Context) (*Credential, error) {
	var (
		err  error
		node *Credential
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CredentialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CredentialUpdateOne) SaveX(ctx context.Context) *Credential {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CredentialUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CredentialUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := credential.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CredentialUpdateOne) sqlSave(ctx context.Context) (_node *Credential, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   credential.Table,
			Columns: credential.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: credential.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Credential.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, credential.FieldID)
		for _, f := range fields {
			if !credential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != credential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Registry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldRegistry,
		})
	}
	if value, ok := cuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldUsername,
		})
	}
	if value, ok := cuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPassword,
		})
	}
	if cuo.mutation.DeploymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.DeploymentTable,
			Columns: []string{credential.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deployment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.DeploymentTable,
			Columns: []string{credential.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Credential{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
