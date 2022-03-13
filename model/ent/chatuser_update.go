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
	"github.com/gitploy-io/gitploy/model/ent/chatuser"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
	"github.com/gitploy-io/gitploy/model/ent/user"
)

// ChatUserUpdate is the builder for updating ChatUser entities.
type ChatUserUpdate struct {
	config
	hooks    []Hook
	mutation *ChatUserMutation
}

// Where appends a list predicates to the ChatUserUpdate builder.
func (cuu *ChatUserUpdate) Where(ps ...predicate.ChatUser) *ChatUserUpdate {
	cuu.mutation.Where(ps...)
	return cuu
}

// SetToken sets the "token" field.
func (cuu *ChatUserUpdate) SetToken(s string) *ChatUserUpdate {
	cuu.mutation.SetToken(s)
	return cuu
}

// SetRefresh sets the "refresh" field.
func (cuu *ChatUserUpdate) SetRefresh(s string) *ChatUserUpdate {
	cuu.mutation.SetRefresh(s)
	return cuu
}

// SetExpiry sets the "expiry" field.
func (cuu *ChatUserUpdate) SetExpiry(t time.Time) *ChatUserUpdate {
	cuu.mutation.SetExpiry(t)
	return cuu
}

// SetBotToken sets the "bot_token" field.
func (cuu *ChatUserUpdate) SetBotToken(s string) *ChatUserUpdate {
	cuu.mutation.SetBotToken(s)
	return cuu
}

// SetCreatedAt sets the "created_at" field.
func (cuu *ChatUserUpdate) SetCreatedAt(t time.Time) *ChatUserUpdate {
	cuu.mutation.SetCreatedAt(t)
	return cuu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuu *ChatUserUpdate) SetNillableCreatedAt(t *time.Time) *ChatUserUpdate {
	if t != nil {
		cuu.SetCreatedAt(*t)
	}
	return cuu
}

// SetUpdatedAt sets the "updated_at" field.
func (cuu *ChatUserUpdate) SetUpdatedAt(t time.Time) *ChatUserUpdate {
	cuu.mutation.SetUpdatedAt(t)
	return cuu
}

// SetUserID sets the "user_id" field.
func (cuu *ChatUserUpdate) SetUserID(i int64) *ChatUserUpdate {
	cuu.mutation.SetUserID(i)
	return cuu
}

// SetUser sets the "user" edge to the User entity.
func (cuu *ChatUserUpdate) SetUser(u *User) *ChatUserUpdate {
	return cuu.SetUserID(u.ID)
}

// Mutation returns the ChatUserMutation object of the builder.
func (cuu *ChatUserUpdate) Mutation() *ChatUserMutation {
	return cuu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuu *ChatUserUpdate) ClearUser() *ChatUserUpdate {
	cuu.mutation.ClearUser()
	return cuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cuu *ChatUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cuu.defaults()
	if len(cuu.hooks) == 0 {
		if err = cuu.check(); err != nil {
			return 0, err
		}
		affected, err = cuu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuu.check(); err != nil {
				return 0, err
			}
			cuu.mutation = mutation
			affected, err = cuu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cuu.hooks) - 1; i >= 0; i-- {
			if cuu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuu *ChatUserUpdate) SaveX(ctx context.Context) int {
	affected, err := cuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cuu *ChatUserUpdate) Exec(ctx context.Context) error {
	_, err := cuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuu *ChatUserUpdate) ExecX(ctx context.Context) {
	if err := cuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuu *ChatUserUpdate) defaults() {
	if _, ok := cuu.mutation.UpdatedAt(); !ok {
		v := chatuser.UpdateDefaultUpdatedAt()
		cuu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuu *ChatUserUpdate) check() error {
	if _, ok := cuu.mutation.UserID(); cuu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChatUser.user"`)
	}
	return nil
}

func (cuu *ChatUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chatuser.Table,
			Columns: chatuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chatuser.FieldID,
			},
		},
	}
	if ps := cuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuu.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldToken,
		})
	}
	if value, ok := cuu.mutation.Refresh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldRefresh,
		})
	}
	if value, ok := cuu.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldExpiry,
		})
	}
	if value, ok := cuu.mutation.BotToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldBotToken,
		})
	}
	if value, ok := cuu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldCreatedAt,
		})
	}
	if value, ok := cuu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldUpdatedAt,
		})
	}
	if cuu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chatuser.UserTable,
			Columns: []string{chatuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chatuser.UserTable,
			Columns: []string{chatuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChatUserUpdateOne is the builder for updating a single ChatUser entity.
type ChatUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatUserMutation
}

// SetToken sets the "token" field.
func (cuuo *ChatUserUpdateOne) SetToken(s string) *ChatUserUpdateOne {
	cuuo.mutation.SetToken(s)
	return cuuo
}

// SetRefresh sets the "refresh" field.
func (cuuo *ChatUserUpdateOne) SetRefresh(s string) *ChatUserUpdateOne {
	cuuo.mutation.SetRefresh(s)
	return cuuo
}

// SetExpiry sets the "expiry" field.
func (cuuo *ChatUserUpdateOne) SetExpiry(t time.Time) *ChatUserUpdateOne {
	cuuo.mutation.SetExpiry(t)
	return cuuo
}

// SetBotToken sets the "bot_token" field.
func (cuuo *ChatUserUpdateOne) SetBotToken(s string) *ChatUserUpdateOne {
	cuuo.mutation.SetBotToken(s)
	return cuuo
}

// SetCreatedAt sets the "created_at" field.
func (cuuo *ChatUserUpdateOne) SetCreatedAt(t time.Time) *ChatUserUpdateOne {
	cuuo.mutation.SetCreatedAt(t)
	return cuuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuuo *ChatUserUpdateOne) SetNillableCreatedAt(t *time.Time) *ChatUserUpdateOne {
	if t != nil {
		cuuo.SetCreatedAt(*t)
	}
	return cuuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuuo *ChatUserUpdateOne) SetUpdatedAt(t time.Time) *ChatUserUpdateOne {
	cuuo.mutation.SetUpdatedAt(t)
	return cuuo
}

// SetUserID sets the "user_id" field.
func (cuuo *ChatUserUpdateOne) SetUserID(i int64) *ChatUserUpdateOne {
	cuuo.mutation.SetUserID(i)
	return cuuo
}

// SetUser sets the "user" edge to the User entity.
func (cuuo *ChatUserUpdateOne) SetUser(u *User) *ChatUserUpdateOne {
	return cuuo.SetUserID(u.ID)
}

// Mutation returns the ChatUserMutation object of the builder.
func (cuuo *ChatUserUpdateOne) Mutation() *ChatUserMutation {
	return cuuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuuo *ChatUserUpdateOne) ClearUser() *ChatUserUpdateOne {
	cuuo.mutation.ClearUser()
	return cuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuuo *ChatUserUpdateOne) Select(field string, fields ...string) *ChatUserUpdateOne {
	cuuo.fields = append([]string{field}, fields...)
	return cuuo
}

// Save executes the query and returns the updated ChatUser entity.
func (cuuo *ChatUserUpdateOne) Save(ctx context.Context) (*ChatUser, error) {
	var (
		err  error
		node *ChatUser
	)
	cuuo.defaults()
	if len(cuuo.hooks) == 0 {
		if err = cuuo.check(); err != nil {
			return nil, err
		}
		node, err = cuuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuuo.check(); err != nil {
				return nil, err
			}
			cuuo.mutation = mutation
			node, err = cuuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuuo.hooks) - 1; i >= 0; i-- {
			if cuuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuuo *ChatUserUpdateOne) SaveX(ctx context.Context) *ChatUser {
	node, err := cuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuuo *ChatUserUpdateOne) Exec(ctx context.Context) error {
	_, err := cuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuuo *ChatUserUpdateOne) ExecX(ctx context.Context) {
	if err := cuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuuo *ChatUserUpdateOne) defaults() {
	if _, ok := cuuo.mutation.UpdatedAt(); !ok {
		v := chatuser.UpdateDefaultUpdatedAt()
		cuuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuuo *ChatUserUpdateOne) check() error {
	if _, ok := cuuo.mutation.UserID(); cuuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChatUser.user"`)
	}
	return nil
}

func (cuuo *ChatUserUpdateOne) sqlSave(ctx context.Context) (_node *ChatUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chatuser.Table,
			Columns: chatuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chatuser.FieldID,
			},
		},
	}
	id, ok := cuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChatUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatuser.FieldID)
		for _, f := range fields {
			if !chatuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chatuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuuo.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldToken,
		})
	}
	if value, ok := cuuo.mutation.Refresh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldRefresh,
		})
	}
	if value, ok := cuuo.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldExpiry,
		})
	}
	if value, ok := cuuo.mutation.BotToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldBotToken,
		})
	}
	if value, ok := cuuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldCreatedAt,
		})
	}
	if value, ok := cuuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldUpdatedAt,
		})
	}
	if cuuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chatuser.UserTable,
			Columns: []string{chatuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chatuser.UserTable,
			Columns: []string{chatuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ChatUser{config: cuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
