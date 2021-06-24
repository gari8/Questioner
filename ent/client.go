// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"server/ent/migrate"

	"server/ent/answer"
	"server/ent/choice"
	"server/ent/choiceanswer"
	"server/ent/question"
	"server/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Answer is the client for interacting with the Answer builders.
	Answer *AnswerClient
	// Choice is the client for interacting with the Choice builders.
	Choice *ChoiceClient
	// ChoiceAnswer is the client for interacting with the ChoiceAnswer builders.
	ChoiceAnswer *ChoiceAnswerClient
	// Question is the client for interacting with the Question builders.
	Question *QuestionClient
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
	c.Answer = NewAnswerClient(c.config)
	c.Choice = NewChoiceClient(c.config)
	c.ChoiceAnswer = NewChoiceAnswerClient(c.config)
	c.Question = NewQuestionClient(c.config)
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
		Answer:       NewAnswerClient(cfg),
		Choice:       NewChoiceClient(cfg),
		ChoiceAnswer: NewChoiceAnswerClient(cfg),
		Question:     NewQuestionClient(cfg),
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
		Answer:       NewAnswerClient(cfg),
		Choice:       NewChoiceClient(cfg),
		ChoiceAnswer: NewChoiceAnswerClient(cfg),
		Question:     NewQuestionClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Answer.
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
	c.Answer.Use(hooks...)
	c.Choice.Use(hooks...)
	c.ChoiceAnswer.Use(hooks...)
	c.Question.Use(hooks...)
	c.User.Use(hooks...)
}

// AnswerClient is a client for the Answer schema.
type AnswerClient struct {
	config
}

// NewAnswerClient returns a client for the Answer from the given config.
func NewAnswerClient(c config) *AnswerClient {
	return &AnswerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `answer.Hooks(f(g(h())))`.
func (c *AnswerClient) Use(hooks ...Hook) {
	c.hooks.Answer = append(c.hooks.Answer, hooks...)
}

// Create returns a create builder for Answer.
func (c *AnswerClient) Create() *AnswerCreate {
	mutation := newAnswerMutation(c.config, OpCreate)
	return &AnswerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Answer entities.
func (c *AnswerClient) CreateBulk(builders ...*AnswerCreate) *AnswerCreateBulk {
	return &AnswerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Answer.
func (c *AnswerClient) Update() *AnswerUpdate {
	mutation := newAnswerMutation(c.config, OpUpdate)
	return &AnswerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnswerClient) UpdateOne(a *Answer) *AnswerUpdateOne {
	mutation := newAnswerMutation(c.config, OpUpdateOne, withAnswer(a))
	return &AnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnswerClient) UpdateOneID(id int) *AnswerUpdateOne {
	mutation := newAnswerMutation(c.config, OpUpdateOne, withAnswerID(id))
	return &AnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Answer.
func (c *AnswerClient) Delete() *AnswerDelete {
	mutation := newAnswerMutation(c.config, OpDelete)
	return &AnswerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AnswerClient) DeleteOne(a *Answer) *AnswerDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AnswerClient) DeleteOneID(id int) *AnswerDeleteOne {
	builder := c.Delete().Where(answer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnswerDeleteOne{builder}
}

// Query returns a query builder for Answer.
func (c *AnswerClient) Query() *AnswerQuery {
	return &AnswerQuery{
		config: c.config,
	}
}

// Get returns a Answer entity by its id.
func (c *AnswerClient) Get(ctx context.Context, id int) (*Answer, error) {
	return c.Query().Where(answer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnswerClient) GetX(ctx context.Context, id int) *Answer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Answer.
func (c *AnswerClient) QueryOwner(a *Answer) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(answer.Table, answer.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, answer.OwnerTable, answer.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Answer.
func (c *AnswerClient) QueryParent(a *Answer) *QuestionQuery {
	query := &QuestionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(answer.Table, answer.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, answer.ParentTable, answer.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AnswerClient) Hooks() []Hook {
	return c.hooks.Answer
}

// ChoiceClient is a client for the Choice schema.
type ChoiceClient struct {
	config
}

// NewChoiceClient returns a client for the Choice from the given config.
func NewChoiceClient(c config) *ChoiceClient {
	return &ChoiceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `choice.Hooks(f(g(h())))`.
func (c *ChoiceClient) Use(hooks ...Hook) {
	c.hooks.Choice = append(c.hooks.Choice, hooks...)
}

// Create returns a create builder for Choice.
func (c *ChoiceClient) Create() *ChoiceCreate {
	mutation := newChoiceMutation(c.config, OpCreate)
	return &ChoiceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Choice entities.
func (c *ChoiceClient) CreateBulk(builders ...*ChoiceCreate) *ChoiceCreateBulk {
	return &ChoiceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Choice.
func (c *ChoiceClient) Update() *ChoiceUpdate {
	mutation := newChoiceMutation(c.config, OpUpdate)
	return &ChoiceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChoiceClient) UpdateOne(ch *Choice) *ChoiceUpdateOne {
	mutation := newChoiceMutation(c.config, OpUpdateOne, withChoice(ch))
	return &ChoiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChoiceClient) UpdateOneID(id int) *ChoiceUpdateOne {
	mutation := newChoiceMutation(c.config, OpUpdateOne, withChoiceID(id))
	return &ChoiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Choice.
func (c *ChoiceClient) Delete() *ChoiceDelete {
	mutation := newChoiceMutation(c.config, OpDelete)
	return &ChoiceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ChoiceClient) DeleteOne(ch *Choice) *ChoiceDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ChoiceClient) DeleteOneID(id int) *ChoiceDeleteOne {
	builder := c.Delete().Where(choice.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChoiceDeleteOne{builder}
}

// Query returns a query builder for Choice.
func (c *ChoiceClient) Query() *ChoiceQuery {
	return &ChoiceQuery{
		config: c.config,
	}
}

// Get returns a Choice entity by its id.
func (c *ChoiceClient) Get(ctx context.Context, id int) (*Choice, error) {
	return c.Query().Where(choice.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChoiceClient) GetX(ctx context.Context, id int) *Choice {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Choice.
func (c *ChoiceClient) QueryParent(ch *Choice) *QuestionQuery {
	query := &QuestionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(choice.Table, choice.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, choice.ParentTable, choice.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChoiceanswers queries the choiceanswers edge of a Choice.
func (c *ChoiceClient) QueryChoiceanswers(ch *Choice) *ChoiceAnswerQuery {
	query := &ChoiceAnswerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(choice.Table, choice.FieldID, id),
			sqlgraph.To(choiceanswer.Table, choiceanswer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, choice.ChoiceanswersTable, choice.ChoiceanswersColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChoiceClient) Hooks() []Hook {
	return c.hooks.Choice
}

// ChoiceAnswerClient is a client for the ChoiceAnswer schema.
type ChoiceAnswerClient struct {
	config
}

// NewChoiceAnswerClient returns a client for the ChoiceAnswer from the given config.
func NewChoiceAnswerClient(c config) *ChoiceAnswerClient {
	return &ChoiceAnswerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `choiceanswer.Hooks(f(g(h())))`.
func (c *ChoiceAnswerClient) Use(hooks ...Hook) {
	c.hooks.ChoiceAnswer = append(c.hooks.ChoiceAnswer, hooks...)
}

// Create returns a create builder for ChoiceAnswer.
func (c *ChoiceAnswerClient) Create() *ChoiceAnswerCreate {
	mutation := newChoiceAnswerMutation(c.config, OpCreate)
	return &ChoiceAnswerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChoiceAnswer entities.
func (c *ChoiceAnswerClient) CreateBulk(builders ...*ChoiceAnswerCreate) *ChoiceAnswerCreateBulk {
	return &ChoiceAnswerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChoiceAnswer.
func (c *ChoiceAnswerClient) Update() *ChoiceAnswerUpdate {
	mutation := newChoiceAnswerMutation(c.config, OpUpdate)
	return &ChoiceAnswerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChoiceAnswerClient) UpdateOne(ca *ChoiceAnswer) *ChoiceAnswerUpdateOne {
	mutation := newChoiceAnswerMutation(c.config, OpUpdateOne, withChoiceAnswer(ca))
	return &ChoiceAnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChoiceAnswerClient) UpdateOneID(id int) *ChoiceAnswerUpdateOne {
	mutation := newChoiceAnswerMutation(c.config, OpUpdateOne, withChoiceAnswerID(id))
	return &ChoiceAnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChoiceAnswer.
func (c *ChoiceAnswerClient) Delete() *ChoiceAnswerDelete {
	mutation := newChoiceAnswerMutation(c.config, OpDelete)
	return &ChoiceAnswerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ChoiceAnswerClient) DeleteOne(ca *ChoiceAnswer) *ChoiceAnswerDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ChoiceAnswerClient) DeleteOneID(id int) *ChoiceAnswerDeleteOne {
	builder := c.Delete().Where(choiceanswer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChoiceAnswerDeleteOne{builder}
}

// Query returns a query builder for ChoiceAnswer.
func (c *ChoiceAnswerClient) Query() *ChoiceAnswerQuery {
	return &ChoiceAnswerQuery{
		config: c.config,
	}
}

// Get returns a ChoiceAnswer entity by its id.
func (c *ChoiceAnswerClient) Get(ctx context.Context, id int) (*ChoiceAnswer, error) {
	return c.Query().Where(choiceanswer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChoiceAnswerClient) GetX(ctx context.Context, id int) *ChoiceAnswer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a ChoiceAnswer.
func (c *ChoiceAnswerClient) QueryOwner(ca *ChoiceAnswer) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(choiceanswer.Table, choiceanswer.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, choiceanswer.OwnerTable, choiceanswer.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a ChoiceAnswer.
func (c *ChoiceAnswerClient) QueryParent(ca *ChoiceAnswer) *ChoiceQuery {
	query := &ChoiceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(choiceanswer.Table, choiceanswer.FieldID, id),
			sqlgraph.To(choice.Table, choice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, choiceanswer.ParentTable, choiceanswer.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChoiceAnswerClient) Hooks() []Hook {
	return c.hooks.ChoiceAnswer
}

// QuestionClient is a client for the Question schema.
type QuestionClient struct {
	config
}

// NewQuestionClient returns a client for the Question from the given config.
func NewQuestionClient(c config) *QuestionClient {
	return &QuestionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `question.Hooks(f(g(h())))`.
func (c *QuestionClient) Use(hooks ...Hook) {
	c.hooks.Question = append(c.hooks.Question, hooks...)
}

// Create returns a create builder for Question.
func (c *QuestionClient) Create() *QuestionCreate {
	mutation := newQuestionMutation(c.config, OpCreate)
	return &QuestionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Question entities.
func (c *QuestionClient) CreateBulk(builders ...*QuestionCreate) *QuestionCreateBulk {
	return &QuestionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Question.
func (c *QuestionClient) Update() *QuestionUpdate {
	mutation := newQuestionMutation(c.config, OpUpdate)
	return &QuestionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *QuestionClient) UpdateOne(q *Question) *QuestionUpdateOne {
	mutation := newQuestionMutation(c.config, OpUpdateOne, withQuestion(q))
	return &QuestionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *QuestionClient) UpdateOneID(id string) *QuestionUpdateOne {
	mutation := newQuestionMutation(c.config, OpUpdateOne, withQuestionID(id))
	return &QuestionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Question.
func (c *QuestionClient) Delete() *QuestionDelete {
	mutation := newQuestionMutation(c.config, OpDelete)
	return &QuestionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *QuestionClient) DeleteOne(q *Question) *QuestionDeleteOne {
	return c.DeleteOneID(q.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *QuestionClient) DeleteOneID(id string) *QuestionDeleteOne {
	builder := c.Delete().Where(question.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &QuestionDeleteOne{builder}
}

// Query returns a query builder for Question.
func (c *QuestionClient) Query() *QuestionQuery {
	return &QuestionQuery{
		config: c.config,
	}
}

// Get returns a Question entity by its id.
func (c *QuestionClient) Get(ctx context.Context, id string) (*Question, error) {
	return c.Query().Where(question.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *QuestionClient) GetX(ctx context.Context, id string) *Question {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Question.
func (c *QuestionClient) QueryOwner(q *Question) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, question.OwnerTable, question.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAnswers queries the answers edge of a Question.
func (c *QuestionClient) QueryAnswers(q *Question) *AnswerQuery {
	query := &AnswerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, question.AnswersTable, question.AnswersColumn),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChoices queries the choices edge of a Question.
func (c *QuestionClient) QueryChoices(q *Question) *ChoiceQuery {
	query := &ChoiceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(choice.Table, choice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, question.ChoicesTable, question.ChoicesColumn),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *QuestionClient) Hooks() []Hook {
	return c.hooks.Question
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

// QueryQuestions queries the questions edge of a User.
func (c *UserClient) QueryQuestions(u *User) *QuestionQuery {
	query := &QuestionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.QuestionsTable, user.QuestionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAnswers queries the answers edge of a User.
func (c *UserClient) QueryAnswers(u *User) *AnswerQuery {
	query := &AnswerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.AnswersTable, user.AnswersColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChoiceanswers queries the choiceanswers edge of a User.
func (c *UserClient) QueryChoiceanswers(u *User) *ChoiceAnswerQuery {
	query := &ChoiceAnswerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(choiceanswer.Table, choiceanswer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ChoiceanswersTable, user.ChoiceanswersColumn),
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
