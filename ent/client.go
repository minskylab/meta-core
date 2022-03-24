// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/minskylab/meta-core/ent/migrate"
	uuid "github.com/satori/go.uuid"

	"github.com/minskylab/meta-core/ent/credential"
	"github.com/minskylab/meta-core/ent/deployment"
	"github.com/minskylab/meta-core/ent/process"
	"github.com/minskylab/meta-core/ent/provider"
	"github.com/minskylab/meta-core/ent/stack"
	"github.com/minskylab/meta-core/ent/task"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Credential is the client for interacting with the Credential builders.
	Credential *CredentialClient
	// Deployment is the client for interacting with the Deployment builders.
	Deployment *DeploymentClient
	// Process is the client for interacting with the Process builders.
	Process *ProcessClient
	// Provider is the client for interacting with the Provider builders.
	Provider *ProviderClient
	// Stack is the client for interacting with the Stack builders.
	Stack *StackClient
	// Task is the client for interacting with the Task builders.
	Task *TaskClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Credential = NewCredentialClient(c.config)
	c.Deployment = NewDeploymentClient(c.config)
	c.Process = NewProcessClient(c.config)
	c.Provider = NewProviderClient(c.config)
	c.Stack = NewStackClient(c.config)
	c.Task = NewTaskClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Credential: NewCredentialClient(cfg),
		Deployment: NewDeploymentClient(cfg),
		Process:    NewProcessClient(cfg),
		Provider:   NewProviderClient(cfg),
		Stack:      NewStackClient(cfg),
		Task:       NewTaskClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:     cfg,
		Credential: NewCredentialClient(cfg),
		Deployment: NewDeploymentClient(cfg),
		Process:    NewProcessClient(cfg),
		Provider:   NewProviderClient(cfg),
		Stack:      NewStackClient(cfg),
		Task:       NewTaskClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Credential.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Credential.Use(hooks...)
	c.Deployment.Use(hooks...)
	c.Process.Use(hooks...)
	c.Provider.Use(hooks...)
	c.Stack.Use(hooks...)
	c.Task.Use(hooks...)
}

// CredentialClient is a client for the Credential schema.
type CredentialClient struct {
	config
}

// NewCredentialClient returns a client for the Credential from the given config.
func NewCredentialClient(c config) *CredentialClient {
	return &CredentialClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `credential.Hooks(f(g(h())))`.
func (c *CredentialClient) Use(hooks ...Hook) {
	c.hooks.Credential = append(c.hooks.Credential, hooks...)
}

// Create returns a create builder for Credential.
func (c *CredentialClient) Create() *CredentialCreate {
	mutation := newCredentialMutation(c.config, OpCreate)
	return &CredentialCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Credential entities.
func (c *CredentialClient) CreateBulk(builders ...*CredentialCreate) *CredentialCreateBulk {
	return &CredentialCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Credential.
func (c *CredentialClient) Update() *CredentialUpdate {
	mutation := newCredentialMutation(c.config, OpUpdate)
	return &CredentialUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CredentialClient) UpdateOne(cr *Credential) *CredentialUpdateOne {
	mutation := newCredentialMutation(c.config, OpUpdateOne, withCredential(cr))
	return &CredentialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CredentialClient) UpdateOneID(id uuid.UUID) *CredentialUpdateOne {
	mutation := newCredentialMutation(c.config, OpUpdateOne, withCredentialID(id))
	return &CredentialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Credential.
func (c *CredentialClient) Delete() *CredentialDelete {
	mutation := newCredentialMutation(c.config, OpDelete)
	return &CredentialDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CredentialClient) DeleteOne(cr *Credential) *CredentialDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CredentialClient) DeleteOneID(id uuid.UUID) *CredentialDeleteOne {
	builder := c.Delete().Where(credential.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CredentialDeleteOne{builder}
}

// Query returns a query builder for Credential.
func (c *CredentialClient) Query() *CredentialQuery {
	return &CredentialQuery{
		config: c.config,
	}
}

// Get returns a Credential entity by its id.
func (c *CredentialClient) Get(ctx context.Context, id uuid.UUID) (*Credential, error) {
	return c.Query().Where(credential.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CredentialClient) GetX(ctx context.Context, id uuid.UUID) *Credential {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDeployment queries the deployment edge of a Credential.
func (c *CredentialClient) QueryDeployment(cr *Credential) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(credential.Table, credential.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, credential.DeploymentTable, credential.DeploymentColumn),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CredentialClient) Hooks() []Hook {
	return c.hooks.Credential
}

// DeploymentClient is a client for the Deployment schema.
type DeploymentClient struct {
	config
}

// NewDeploymentClient returns a client for the Deployment from the given config.
func NewDeploymentClient(c config) *DeploymentClient {
	return &DeploymentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `deployment.Hooks(f(g(h())))`.
func (c *DeploymentClient) Use(hooks ...Hook) {
	c.hooks.Deployment = append(c.hooks.Deployment, hooks...)
}

// Create returns a create builder for Deployment.
func (c *DeploymentClient) Create() *DeploymentCreate {
	mutation := newDeploymentMutation(c.config, OpCreate)
	return &DeploymentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Deployment entities.
func (c *DeploymentClient) CreateBulk(builders ...*DeploymentCreate) *DeploymentCreateBulk {
	return &DeploymentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Deployment.
func (c *DeploymentClient) Update() *DeploymentUpdate {
	mutation := newDeploymentMutation(c.config, OpUpdate)
	return &DeploymentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeploymentClient) UpdateOne(d *Deployment) *DeploymentUpdateOne {
	mutation := newDeploymentMutation(c.config, OpUpdateOne, withDeployment(d))
	return &DeploymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeploymentClient) UpdateOneID(id uuid.UUID) *DeploymentUpdateOne {
	mutation := newDeploymentMutation(c.config, OpUpdateOne, withDeploymentID(id))
	return &DeploymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Deployment.
func (c *DeploymentClient) Delete() *DeploymentDelete {
	mutation := newDeploymentMutation(c.config, OpDelete)
	return &DeploymentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DeploymentClient) DeleteOne(d *Deployment) *DeploymentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DeploymentClient) DeleteOneID(id uuid.UUID) *DeploymentDeleteOne {
	builder := c.Delete().Where(deployment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeploymentDeleteOne{builder}
}

// Query returns a query builder for Deployment.
func (c *DeploymentClient) Query() *DeploymentQuery {
	return &DeploymentQuery{
		config: c.config,
	}
}

// Get returns a Deployment entity by its id.
func (c *DeploymentClient) Get(ctx context.Context, id uuid.UUID) (*Deployment, error) {
	return c.Query().Where(deployment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeploymentClient) GetX(ctx context.Context, id uuid.UUID) *Deployment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTasks queries the tasks edge of a Deployment.
func (c *DeploymentClient) QueryTasks(d *Deployment) *TaskQuery {
	query := &TaskQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, deployment.TasksTable, deployment.TasksColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProvider queries the provider edge of a Deployment.
func (c *DeploymentClient) QueryProvider(d *Deployment) *ProviderQuery {
	query := &ProviderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(provider.Table, provider.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, deployment.ProviderTable, deployment.ProviderColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStack queries the stack edge of a Deployment.
func (c *DeploymentClient) QueryStack(d *Deployment) *StackQuery {
	query := &StackQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(stack.Table, stack.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, deployment.StackTable, deployment.StackColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCredentials queries the credentials edge of a Deployment.
func (c *DeploymentClient) QueryCredentials(d *Deployment) *CredentialQuery {
	query := &CredentialQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(credential.Table, credential.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, deployment.CredentialsTable, deployment.CredentialsColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DeploymentClient) Hooks() []Hook {
	return c.hooks.Deployment
}

// ProcessClient is a client for the Process schema.
type ProcessClient struct {
	config
}

// NewProcessClient returns a client for the Process from the given config.
func NewProcessClient(c config) *ProcessClient {
	return &ProcessClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `process.Hooks(f(g(h())))`.
func (c *ProcessClient) Use(hooks ...Hook) {
	c.hooks.Process = append(c.hooks.Process, hooks...)
}

// Create returns a create builder for Process.
func (c *ProcessClient) Create() *ProcessCreate {
	mutation := newProcessMutation(c.config, OpCreate)
	return &ProcessCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Process entities.
func (c *ProcessClient) CreateBulk(builders ...*ProcessCreate) *ProcessCreateBulk {
	return &ProcessCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Process.
func (c *ProcessClient) Update() *ProcessUpdate {
	mutation := newProcessMutation(c.config, OpUpdate)
	return &ProcessUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProcessClient) UpdateOne(pr *Process) *ProcessUpdateOne {
	mutation := newProcessMutation(c.config, OpUpdateOne, withProcess(pr))
	return &ProcessUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProcessClient) UpdateOneID(id string) *ProcessUpdateOne {
	mutation := newProcessMutation(c.config, OpUpdateOne, withProcessID(id))
	return &ProcessUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Process.
func (c *ProcessClient) Delete() *ProcessDelete {
	mutation := newProcessMutation(c.config, OpDelete)
	return &ProcessDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProcessClient) DeleteOne(pr *Process) *ProcessDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProcessClient) DeleteOneID(id string) *ProcessDeleteOne {
	builder := c.Delete().Where(process.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProcessDeleteOne{builder}
}

// Query returns a query builder for Process.
func (c *ProcessClient) Query() *ProcessQuery {
	return &ProcessQuery{
		config: c.config,
	}
}

// Get returns a Process entity by its id.
func (c *ProcessClient) Get(ctx context.Context, id string) (*Process, error) {
	return c.Query().Where(process.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProcessClient) GetX(ctx context.Context, id string) *Process {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStack queries the stack edge of a Process.
func (c *ProcessClient) QueryStack(pr *Process) *StackQuery {
	query := &StackQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(process.Table, process.FieldID, id),
			sqlgraph.To(stack.Table, stack.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, process.StackTable, process.StackColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCredentials queries the credentials edge of a Process.
func (c *ProcessClient) QueryCredentials(pr *Process) *CredentialQuery {
	query := &CredentialQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(process.Table, process.FieldID, id),
			sqlgraph.To(credential.Table, credential.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, process.CredentialsTable, process.CredentialsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTasks queries the tasks edge of a Process.
func (c *ProcessClient) QueryTasks(pr *Process) *TaskQuery {
	query := &TaskQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(process.Table, process.FieldID, id),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, process.TasksTable, process.TasksColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProcessClient) Hooks() []Hook {
	return c.hooks.Process
}

// ProviderClient is a client for the Provider schema.
type ProviderClient struct {
	config
}

// NewProviderClient returns a client for the Provider from the given config.
func NewProviderClient(c config) *ProviderClient {
	return &ProviderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `provider.Hooks(f(g(h())))`.
func (c *ProviderClient) Use(hooks ...Hook) {
	c.hooks.Provider = append(c.hooks.Provider, hooks...)
}

// Create returns a create builder for Provider.
func (c *ProviderClient) Create() *ProviderCreate {
	mutation := newProviderMutation(c.config, OpCreate)
	return &ProviderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Provider entities.
func (c *ProviderClient) CreateBulk(builders ...*ProviderCreate) *ProviderCreateBulk {
	return &ProviderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Provider.
func (c *ProviderClient) Update() *ProviderUpdate {
	mutation := newProviderMutation(c.config, OpUpdate)
	return &ProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProviderClient) UpdateOne(pr *Provider) *ProviderUpdateOne {
	mutation := newProviderMutation(c.config, OpUpdateOne, withProvider(pr))
	return &ProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProviderClient) UpdateOneID(id uuid.UUID) *ProviderUpdateOne {
	mutation := newProviderMutation(c.config, OpUpdateOne, withProviderID(id))
	return &ProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Provider.
func (c *ProviderClient) Delete() *ProviderDelete {
	mutation := newProviderMutation(c.config, OpDelete)
	return &ProviderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProviderClient) DeleteOne(pr *Provider) *ProviderDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProviderClient) DeleteOneID(id uuid.UUID) *ProviderDeleteOne {
	builder := c.Delete().Where(provider.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProviderDeleteOne{builder}
}

// Query returns a query builder for Provider.
func (c *ProviderClient) Query() *ProviderQuery {
	return &ProviderQuery{
		config: c.config,
	}
}

// Get returns a Provider entity by its id.
func (c *ProviderClient) Get(ctx context.Context, id uuid.UUID) (*Provider, error) {
	return c.Query().Where(provider.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProviderClient) GetX(ctx context.Context, id uuid.UUID) *Provider {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDeployment queries the deployment edge of a Provider.
func (c *ProviderClient) QueryDeployment(pr *Provider) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(provider.Table, provider.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, provider.DeploymentTable, provider.DeploymentColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProviderClient) Hooks() []Hook {
	return c.hooks.Provider
}

// StackClient is a client for the Stack schema.
type StackClient struct {
	config
}

// NewStackClient returns a client for the Stack from the given config.
func NewStackClient(c config) *StackClient {
	return &StackClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `stack.Hooks(f(g(h())))`.
func (c *StackClient) Use(hooks ...Hook) {
	c.hooks.Stack = append(c.hooks.Stack, hooks...)
}

// Create returns a create builder for Stack.
func (c *StackClient) Create() *StackCreate {
	mutation := newStackMutation(c.config, OpCreate)
	return &StackCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Stack entities.
func (c *StackClient) CreateBulk(builders ...*StackCreate) *StackCreateBulk {
	return &StackCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Stack.
func (c *StackClient) Update() *StackUpdate {
	mutation := newStackMutation(c.config, OpUpdate)
	return &StackUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StackClient) UpdateOne(s *Stack) *StackUpdateOne {
	mutation := newStackMutation(c.config, OpUpdateOne, withStack(s))
	return &StackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StackClient) UpdateOneID(id string) *StackUpdateOne {
	mutation := newStackMutation(c.config, OpUpdateOne, withStackID(id))
	return &StackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Stack.
func (c *StackClient) Delete() *StackDelete {
	mutation := newStackMutation(c.config, OpDelete)
	return &StackDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StackClient) DeleteOne(s *Stack) *StackDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StackClient) DeleteOneID(id string) *StackDeleteOne {
	builder := c.Delete().Where(stack.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StackDeleteOne{builder}
}

// Query returns a query builder for Stack.
func (c *StackClient) Query() *StackQuery {
	return &StackQuery{
		config: c.config,
	}
}

// Get returns a Stack entity by its id.
func (c *StackClient) Get(ctx context.Context, id string) (*Stack, error) {
	return c.Query().Where(stack.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StackClient) GetX(ctx context.Context, id string) *Stack {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDeployment queries the deployment edge of a Stack.
func (c *StackClient) QueryDeployment(s *Stack) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(stack.Table, stack.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, stack.DeploymentTable, stack.DeploymentColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StackClient) Hooks() []Hook {
	return c.hooks.Stack
}

// TaskClient is a client for the Task schema.
type TaskClient struct {
	config
}

// NewTaskClient returns a client for the Task from the given config.
func NewTaskClient(c config) *TaskClient {
	return &TaskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `task.Hooks(f(g(h())))`.
func (c *TaskClient) Use(hooks ...Hook) {
	c.hooks.Task = append(c.hooks.Task, hooks...)
}

// Create returns a create builder for Task.
func (c *TaskClient) Create() *TaskCreate {
	mutation := newTaskMutation(c.config, OpCreate)
	return &TaskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Task entities.
func (c *TaskClient) CreateBulk(builders ...*TaskCreate) *TaskCreateBulk {
	return &TaskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Task.
func (c *TaskClient) Update() *TaskUpdate {
	mutation := newTaskMutation(c.config, OpUpdate)
	return &TaskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TaskClient) UpdateOne(t *Task) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTask(t))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TaskClient) UpdateOneID(id uuid.UUID) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTaskID(id))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Task.
func (c *TaskClient) Delete() *TaskDelete {
	mutation := newTaskMutation(c.config, OpDelete)
	return &TaskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TaskClient) DeleteOne(t *Task) *TaskDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TaskClient) DeleteOneID(id uuid.UUID) *TaskDeleteOne {
	builder := c.Delete().Where(task.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TaskDeleteOne{builder}
}

// Query returns a query builder for Task.
func (c *TaskClient) Query() *TaskQuery {
	return &TaskQuery{
		config: c.config,
	}
}

// Get returns a Task entity by its id.
func (c *TaskClient) Get(ctx context.Context, id uuid.UUID) (*Task, error) {
	return c.Query().Where(task.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TaskClient) GetX(ctx context.Context, id uuid.UUID) *Task {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDeployment queries the deployment edge of a Task.
func (c *TaskClient) QueryDeployment(t *Task) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, task.DeploymentTable, task.DeploymentColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TaskClient) Hooks() []Hook {
	return c.hooks.Task
}
