// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/hanjunlee/gitploy/ent/migrate"

	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/chatuser"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/notification"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ChatCallback is the client for interacting with the ChatCallback builders.
	ChatCallback *ChatCallbackClient
	// ChatUser is the client for interacting with the ChatUser builders.
	ChatUser *ChatUserClient
	// Deployment is the client for interacting with the Deployment builders.
	Deployment *DeploymentClient
	// Notification is the client for interacting with the Notification builders.
	Notification *NotificationClient
	// Perm is the client for interacting with the Perm builders.
	Perm *PermClient
	// Repo is the client for interacting with the Repo builders.
	Repo *RepoClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.ChatCallback = NewChatCallbackClient(c.config)
	c.ChatUser = NewChatUserClient(c.config)
	c.Deployment = NewDeploymentClient(c.config)
	c.Notification = NewNotificationClient(c.config)
	c.Perm = NewPermClient(c.config)
	c.Repo = NewRepoClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:          ctx,
		config:       cfg,
		ChatCallback: NewChatCallbackClient(cfg),
		ChatUser:     NewChatUserClient(cfg),
		Deployment:   NewDeploymentClient(cfg),
		Notification: NewNotificationClient(cfg),
		Perm:         NewPermClient(cfg),
		Repo:         NewRepoClient(cfg),
		User:         NewUserClient(cfg),
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
		config:       cfg,
		ChatCallback: NewChatCallbackClient(cfg),
		ChatUser:     NewChatUserClient(cfg),
		Deployment:   NewDeploymentClient(cfg),
		Notification: NewNotificationClient(cfg),
		Perm:         NewPermClient(cfg),
		Repo:         NewRepoClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ChatCallback.
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
	c.ChatCallback.Use(hooks...)
	c.ChatUser.Use(hooks...)
	c.Deployment.Use(hooks...)
	c.Notification.Use(hooks...)
	c.Perm.Use(hooks...)
	c.Repo.Use(hooks...)
	c.User.Use(hooks...)
}

// ChatCallbackClient is a client for the ChatCallback schema.
type ChatCallbackClient struct {
	config
}

// NewChatCallbackClient returns a client for the ChatCallback from the given config.
func NewChatCallbackClient(c config) *ChatCallbackClient {
	return &ChatCallbackClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chatcallback.Hooks(f(g(h())))`.
func (c *ChatCallbackClient) Use(hooks ...Hook) {
	c.hooks.ChatCallback = append(c.hooks.ChatCallback, hooks...)
}

// Create returns a create builder for ChatCallback.
func (c *ChatCallbackClient) Create() *ChatCallbackCreate {
	mutation := newChatCallbackMutation(c.config, OpCreate)
	return &ChatCallbackCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChatCallback entities.
func (c *ChatCallbackClient) CreateBulk(builders ...*ChatCallbackCreate) *ChatCallbackCreateBulk {
	return &ChatCallbackCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChatCallback.
func (c *ChatCallbackClient) Update() *ChatCallbackUpdate {
	mutation := newChatCallbackMutation(c.config, OpUpdate)
	return &ChatCallbackUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatCallbackClient) UpdateOne(cc *ChatCallback) *ChatCallbackUpdateOne {
	mutation := newChatCallbackMutation(c.config, OpUpdateOne, withChatCallback(cc))
	return &ChatCallbackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatCallbackClient) UpdateOneID(id int) *ChatCallbackUpdateOne {
	mutation := newChatCallbackMutation(c.config, OpUpdateOne, withChatCallbackID(id))
	return &ChatCallbackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChatCallback.
func (c *ChatCallbackClient) Delete() *ChatCallbackDelete {
	mutation := newChatCallbackMutation(c.config, OpDelete)
	return &ChatCallbackDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ChatCallbackClient) DeleteOne(cc *ChatCallback) *ChatCallbackDeleteOne {
	return c.DeleteOneID(cc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ChatCallbackClient) DeleteOneID(id int) *ChatCallbackDeleteOne {
	builder := c.Delete().Where(chatcallback.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChatCallbackDeleteOne{builder}
}

// Query returns a query builder for ChatCallback.
func (c *ChatCallbackClient) Query() *ChatCallbackQuery {
	return &ChatCallbackQuery{
		config: c.config,
	}
}

// Get returns a ChatCallback entity by its id.
func (c *ChatCallbackClient) Get(ctx context.Context, id int) (*ChatCallback, error) {
	return c.Query().Where(chatcallback.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatCallbackClient) GetX(ctx context.Context, id int) *ChatCallback {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChatUser queries the chat_user edge of a ChatCallback.
func (c *ChatCallbackClient) QueryChatUser(cc *ChatCallback) *ChatUserQuery {
	query := &ChatUserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatcallback.Table, chatcallback.FieldID, id),
			sqlgraph.To(chatuser.Table, chatuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chatcallback.ChatUserTable, chatcallback.ChatUserColumn),
		)
		fromV = sqlgraph.Neighbors(cc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRepo queries the repo edge of a ChatCallback.
func (c *ChatCallbackClient) QueryRepo(cc *ChatCallback) *RepoQuery {
	query := &RepoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatcallback.Table, chatcallback.FieldID, id),
			sqlgraph.To(repo.Table, repo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chatcallback.RepoTable, chatcallback.RepoColumn),
		)
		fromV = sqlgraph.Neighbors(cc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChatCallbackClient) Hooks() []Hook {
	return c.hooks.ChatCallback
}

// ChatUserClient is a client for the ChatUser schema.
type ChatUserClient struct {
	config
}

// NewChatUserClient returns a client for the ChatUser from the given config.
func NewChatUserClient(c config) *ChatUserClient {
	return &ChatUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chatuser.Hooks(f(g(h())))`.
func (c *ChatUserClient) Use(hooks ...Hook) {
	c.hooks.ChatUser = append(c.hooks.ChatUser, hooks...)
}

// Create returns a create builder for ChatUser.
func (c *ChatUserClient) Create() *ChatUserCreate {
	mutation := newChatUserMutation(c.config, OpCreate)
	return &ChatUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChatUser entities.
func (c *ChatUserClient) CreateBulk(builders ...*ChatUserCreate) *ChatUserCreateBulk {
	return &ChatUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChatUser.
func (c *ChatUserClient) Update() *ChatUserUpdate {
	mutation := newChatUserMutation(c.config, OpUpdate)
	return &ChatUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatUserClient) UpdateOne(cu *ChatUser) *ChatUserUpdateOne {
	mutation := newChatUserMutation(c.config, OpUpdateOne, withChatUser(cu))
	return &ChatUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatUserClient) UpdateOneID(id string) *ChatUserUpdateOne {
	mutation := newChatUserMutation(c.config, OpUpdateOne, withChatUserID(id))
	return &ChatUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChatUser.
func (c *ChatUserClient) Delete() *ChatUserDelete {
	mutation := newChatUserMutation(c.config, OpDelete)
	return &ChatUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ChatUserClient) DeleteOne(cu *ChatUser) *ChatUserDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ChatUserClient) DeleteOneID(id string) *ChatUserDeleteOne {
	builder := c.Delete().Where(chatuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChatUserDeleteOne{builder}
}

// Query returns a query builder for ChatUser.
func (c *ChatUserClient) Query() *ChatUserQuery {
	return &ChatUserQuery{
		config: c.config,
	}
}

// Get returns a ChatUser entity by its id.
func (c *ChatUserClient) Get(ctx context.Context, id string) (*ChatUser, error) {
	return c.Query().Where(chatuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatUserClient) GetX(ctx context.Context, id string) *ChatUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChatCallback queries the chat_callback edge of a ChatUser.
func (c *ChatUserClient) QueryChatCallback(cu *ChatUser) *ChatCallbackQuery {
	query := &ChatCallbackQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatuser.Table, chatuser.FieldID, id),
			sqlgraph.To(chatcallback.Table, chatcallback.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, chatuser.ChatCallbackTable, chatuser.ChatCallbackColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a ChatUser.
func (c *ChatUserClient) QueryUser(cu *ChatUser) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chatuser.Table, chatuser.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, chatuser.UserTable, chatuser.UserColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChatUserClient) Hooks() []Hook {
	return c.hooks.ChatUser
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
func (c *DeploymentClient) UpdateOneID(id int) *DeploymentUpdateOne {
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
func (c *DeploymentClient) DeleteOneID(id int) *DeploymentDeleteOne {
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
func (c *DeploymentClient) Get(ctx context.Context, id int) (*Deployment, error) {
	return c.Query().Where(deployment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeploymentClient) GetX(ctx context.Context, id int) *Deployment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Deployment.
func (c *DeploymentClient) QueryUser(d *Deployment) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deployment.UserTable, deployment.UserColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRepo queries the repo edge of a Deployment.
func (c *DeploymentClient) QueryRepo(d *Deployment) *RepoQuery {
	query := &RepoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(repo.Table, repo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deployment.RepoTable, deployment.RepoColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNotifications queries the notifications edge of a Deployment.
func (c *DeploymentClient) QueryNotifications(d *Deployment) *NotificationQuery {
	query := &NotificationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, deployment.NotificationsTable, deployment.NotificationsColumn),
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

// NotificationClient is a client for the Notification schema.
type NotificationClient struct {
	config
}

// NewNotificationClient returns a client for the Notification from the given config.
func NewNotificationClient(c config) *NotificationClient {
	return &NotificationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notification.Hooks(f(g(h())))`.
func (c *NotificationClient) Use(hooks ...Hook) {
	c.hooks.Notification = append(c.hooks.Notification, hooks...)
}

// Create returns a create builder for Notification.
func (c *NotificationClient) Create() *NotificationCreate {
	mutation := newNotificationMutation(c.config, OpCreate)
	return &NotificationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Notification entities.
func (c *NotificationClient) CreateBulk(builders ...*NotificationCreate) *NotificationCreateBulk {
	return &NotificationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Notification.
func (c *NotificationClient) Update() *NotificationUpdate {
	mutation := newNotificationMutation(c.config, OpUpdate)
	return &NotificationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NotificationClient) UpdateOne(n *Notification) *NotificationUpdateOne {
	mutation := newNotificationMutation(c.config, OpUpdateOne, withNotification(n))
	return &NotificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NotificationClient) UpdateOneID(id int) *NotificationUpdateOne {
	mutation := newNotificationMutation(c.config, OpUpdateOne, withNotificationID(id))
	return &NotificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Notification.
func (c *NotificationClient) Delete() *NotificationDelete {
	mutation := newNotificationMutation(c.config, OpDelete)
	return &NotificationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NotificationClient) DeleteOne(n *Notification) *NotificationDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NotificationClient) DeleteOneID(id int) *NotificationDeleteOne {
	builder := c.Delete().Where(notification.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NotificationDeleteOne{builder}
}

// Query returns a query builder for Notification.
func (c *NotificationClient) Query() *NotificationQuery {
	return &NotificationQuery{
		config: c.config,
	}
}

// Get returns a Notification entity by its id.
func (c *NotificationClient) Get(ctx context.Context, id int) (*Notification, error) {
	return c.Query().Where(notification.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NotificationClient) GetX(ctx context.Context, id int) *Notification {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Notification.
func (c *NotificationClient) QueryUser(n *Notification) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notification.UserTable, notification.UserColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRepo queries the repo edge of a Notification.
func (c *NotificationClient) QueryRepo(n *Notification) *RepoQuery {
	query := &RepoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, id),
			sqlgraph.To(repo.Table, repo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notification.RepoTable, notification.RepoColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDeployment queries the deployment edge of a Notification.
func (c *NotificationClient) QueryDeployment(n *Notification) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notification.DeploymentTable, notification.DeploymentColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NotificationClient) Hooks() []Hook {
	return c.hooks.Notification
}

// PermClient is a client for the Perm schema.
type PermClient struct {
	config
}

// NewPermClient returns a client for the Perm from the given config.
func NewPermClient(c config) *PermClient {
	return &PermClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `perm.Hooks(f(g(h())))`.
func (c *PermClient) Use(hooks ...Hook) {
	c.hooks.Perm = append(c.hooks.Perm, hooks...)
}

// Create returns a create builder for Perm.
func (c *PermClient) Create() *PermCreate {
	mutation := newPermMutation(c.config, OpCreate)
	return &PermCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Perm entities.
func (c *PermClient) CreateBulk(builders ...*PermCreate) *PermCreateBulk {
	return &PermCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Perm.
func (c *PermClient) Update() *PermUpdate {
	mutation := newPermMutation(c.config, OpUpdate)
	return &PermUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PermClient) UpdateOne(pe *Perm) *PermUpdateOne {
	mutation := newPermMutation(c.config, OpUpdateOne, withPerm(pe))
	return &PermUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PermClient) UpdateOneID(id int) *PermUpdateOne {
	mutation := newPermMutation(c.config, OpUpdateOne, withPermID(id))
	return &PermUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Perm.
func (c *PermClient) Delete() *PermDelete {
	mutation := newPermMutation(c.config, OpDelete)
	return &PermDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PermClient) DeleteOne(pe *Perm) *PermDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PermClient) DeleteOneID(id int) *PermDeleteOne {
	builder := c.Delete().Where(perm.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PermDeleteOne{builder}
}

// Query returns a query builder for Perm.
func (c *PermClient) Query() *PermQuery {
	return &PermQuery{
		config: c.config,
	}
}

// Get returns a Perm entity by its id.
func (c *PermClient) Get(ctx context.Context, id int) (*Perm, error) {
	return c.Query().Where(perm.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PermClient) GetX(ctx context.Context, id int) *Perm {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Perm.
func (c *PermClient) QueryUser(pe *Perm) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(perm.Table, perm.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, perm.UserTable, perm.UserColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRepo queries the repo edge of a Perm.
func (c *PermClient) QueryRepo(pe *Perm) *RepoQuery {
	query := &RepoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(perm.Table, perm.FieldID, id),
			sqlgraph.To(repo.Table, repo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, perm.RepoTable, perm.RepoColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PermClient) Hooks() []Hook {
	return c.hooks.Perm
}

// RepoClient is a client for the Repo schema.
type RepoClient struct {
	config
}

// NewRepoClient returns a client for the Repo from the given config.
func NewRepoClient(c config) *RepoClient {
	return &RepoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `repo.Hooks(f(g(h())))`.
func (c *RepoClient) Use(hooks ...Hook) {
	c.hooks.Repo = append(c.hooks.Repo, hooks...)
}

// Create returns a create builder for Repo.
func (c *RepoClient) Create() *RepoCreate {
	mutation := newRepoMutation(c.config, OpCreate)
	return &RepoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Repo entities.
func (c *RepoClient) CreateBulk(builders ...*RepoCreate) *RepoCreateBulk {
	return &RepoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Repo.
func (c *RepoClient) Update() *RepoUpdate {
	mutation := newRepoMutation(c.config, OpUpdate)
	return &RepoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RepoClient) UpdateOne(r *Repo) *RepoUpdateOne {
	mutation := newRepoMutation(c.config, OpUpdateOne, withRepo(r))
	return &RepoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RepoClient) UpdateOneID(id string) *RepoUpdateOne {
	mutation := newRepoMutation(c.config, OpUpdateOne, withRepoID(id))
	return &RepoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Repo.
func (c *RepoClient) Delete() *RepoDelete {
	mutation := newRepoMutation(c.config, OpDelete)
	return &RepoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RepoClient) DeleteOne(r *Repo) *RepoDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RepoClient) DeleteOneID(id string) *RepoDeleteOne {
	builder := c.Delete().Where(repo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RepoDeleteOne{builder}
}

// Query returns a query builder for Repo.
func (c *RepoClient) Query() *RepoQuery {
	return &RepoQuery{
		config: c.config,
	}
}

// Get returns a Repo entity by its id.
func (c *RepoClient) Get(ctx context.Context, id string) (*Repo, error) {
	return c.Query().Where(repo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RepoClient) GetX(ctx context.Context, id string) *Repo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPerms queries the perms edge of a Repo.
func (c *RepoClient) QueryPerms(r *Repo) *PermQuery {
	query := &PermQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(repo.Table, repo.FieldID, id),
			sqlgraph.To(perm.Table, perm.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, repo.PermsTable, repo.PermsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDeployments queries the deployments edge of a Repo.
func (c *RepoClient) QueryDeployments(r *Repo) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(repo.Table, repo.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, repo.DeploymentsTable, repo.DeploymentsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChatCallback queries the chat_callback edge of a Repo.
func (c *RepoClient) QueryChatCallback(r *Repo) *ChatCallbackQuery {
	query := &ChatCallbackQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(repo.Table, repo.FieldID, id),
			sqlgraph.To(chatcallback.Table, chatcallback.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, repo.ChatCallbackTable, repo.ChatCallbackColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNotifications queries the notifications edge of a Repo.
func (c *RepoClient) QueryNotifications(r *Repo) *NotificationQuery {
	query := &NotificationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(repo.Table, repo.FieldID, id),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, repo.NotificationsTable, repo.NotificationsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RepoClient) Hooks() []Hook {
	return c.hooks.Repo
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChatUser queries the chat_user edge of a User.
func (c *UserClient) QueryChatUser(u *User) *ChatUserQuery {
	query := &ChatUserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(chatuser.Table, chatuser.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.ChatUserTable, user.ChatUserColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPerms queries the perms edge of a User.
func (c *UserClient) QueryPerms(u *User) *PermQuery {
	query := &PermQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(perm.Table, perm.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PermsTable, user.PermsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDeployments queries the deployments edge of a User.
func (c *UserClient) QueryDeployments(u *User) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DeploymentsTable, user.DeploymentsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNotification queries the notification edge of a User.
func (c *UserClient) QueryNotification(u *User) *NotificationQuery {
	query := &NotificationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.NotificationTable, user.NotificationColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
