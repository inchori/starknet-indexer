// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/inchori/starknet-indexer/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/inchori/starknet-indexer/ent/starknetdeclaretx"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// StarknetDeclareTx is the client for interacting with the StarknetDeclareTx builders.
	StarknetDeclareTx *StarknetDeclareTxClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.StarknetDeclareTx = NewStarknetDeclareTxClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		StarknetDeclareTx: NewStarknetDeclareTxClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:               ctx,
		config:            cfg,
		StarknetDeclareTx: NewStarknetDeclareTxClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		StarknetDeclareTx.
//		Query().
//		Count(ctx)
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
	c.StarknetDeclareTx.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.StarknetDeclareTx.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *StarknetDeclareTxMutation:
		return c.StarknetDeclareTx.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// StarknetDeclareTxClient is a client for the StarknetDeclareTx schema.
type StarknetDeclareTxClient struct {
	config
}

// NewStarknetDeclareTxClient returns a client for the StarknetDeclareTx from the given config.
func NewStarknetDeclareTxClient(c config) *StarknetDeclareTxClient {
	return &StarknetDeclareTxClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `starknetdeclaretx.Hooks(f(g(h())))`.
func (c *StarknetDeclareTxClient) Use(hooks ...Hook) {
	c.hooks.StarknetDeclareTx = append(c.hooks.StarknetDeclareTx, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `starknetdeclaretx.Intercept(f(g(h())))`.
func (c *StarknetDeclareTxClient) Intercept(interceptors ...Interceptor) {
	c.inters.StarknetDeclareTx = append(c.inters.StarknetDeclareTx, interceptors...)
}

// Create returns a builder for creating a StarknetDeclareTx entity.
func (c *StarknetDeclareTxClient) Create() *StarknetDeclareTxCreate {
	mutation := newStarknetDeclareTxMutation(c.config, OpCreate)
	return &StarknetDeclareTxCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StarknetDeclareTx entities.
func (c *StarknetDeclareTxClient) CreateBulk(builders ...*StarknetDeclareTxCreate) *StarknetDeclareTxCreateBulk {
	return &StarknetDeclareTxCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StarknetDeclareTxClient) MapCreateBulk(slice any, setFunc func(*StarknetDeclareTxCreate, int)) *StarknetDeclareTxCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StarknetDeclareTxCreateBulk{err: fmt.Errorf("calling to StarknetDeclareTxClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StarknetDeclareTxCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StarknetDeclareTxCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StarknetDeclareTx.
func (c *StarknetDeclareTxClient) Update() *StarknetDeclareTxUpdate {
	mutation := newStarknetDeclareTxMutation(c.config, OpUpdate)
	return &StarknetDeclareTxUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StarknetDeclareTxClient) UpdateOne(sdt *StarknetDeclareTx) *StarknetDeclareTxUpdateOne {
	mutation := newStarknetDeclareTxMutation(c.config, OpUpdateOne, withStarknetDeclareTx(sdt))
	return &StarknetDeclareTxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StarknetDeclareTxClient) UpdateOneID(id int) *StarknetDeclareTxUpdateOne {
	mutation := newStarknetDeclareTxMutation(c.config, OpUpdateOne, withStarknetDeclareTxID(id))
	return &StarknetDeclareTxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StarknetDeclareTx.
func (c *StarknetDeclareTxClient) Delete() *StarknetDeclareTxDelete {
	mutation := newStarknetDeclareTxMutation(c.config, OpDelete)
	return &StarknetDeclareTxDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StarknetDeclareTxClient) DeleteOne(sdt *StarknetDeclareTx) *StarknetDeclareTxDeleteOne {
	return c.DeleteOneID(sdt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StarknetDeclareTxClient) DeleteOneID(id int) *StarknetDeclareTxDeleteOne {
	builder := c.Delete().Where(starknetdeclaretx.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StarknetDeclareTxDeleteOne{builder}
}

// Query returns a query builder for StarknetDeclareTx.
func (c *StarknetDeclareTxClient) Query() *StarknetDeclareTxQuery {
	return &StarknetDeclareTxQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStarknetDeclareTx},
		inters: c.Interceptors(),
	}
}

// Get returns a StarknetDeclareTx entity by its id.
func (c *StarknetDeclareTxClient) Get(ctx context.Context, id int) (*StarknetDeclareTx, error) {
	return c.Query().Where(starknetdeclaretx.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StarknetDeclareTxClient) GetX(ctx context.Context, id int) *StarknetDeclareTx {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *StarknetDeclareTxClient) Hooks() []Hook {
	return c.hooks.StarknetDeclareTx
}

// Interceptors returns the client interceptors.
func (c *StarknetDeclareTxClient) Interceptors() []Interceptor {
	return c.inters.StarknetDeclareTx
}

func (c *StarknetDeclareTxClient) mutate(ctx context.Context, m *StarknetDeclareTxMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StarknetDeclareTxCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StarknetDeclareTxUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StarknetDeclareTxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StarknetDeclareTxDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown StarknetDeclareTx mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		StarknetDeclareTx []ent.Hook
	}
	inters struct {
		StarknetDeclareTx []ent.Interceptor
	}
)
