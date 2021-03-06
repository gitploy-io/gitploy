// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/model/ent/chatuser"
	"github.com/gitploy-io/gitploy/model/ent/user"
)

// ChatUserCreate is the builder for creating a ChatUser entity.
type ChatUserCreate struct {
	config
	mutation *ChatUserMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (cuc *ChatUserCreate) SetToken(s string) *ChatUserCreate {
	cuc.mutation.SetToken(s)
	return cuc
}

// SetRefresh sets the "refresh" field.
func (cuc *ChatUserCreate) SetRefresh(s string) *ChatUserCreate {
	cuc.mutation.SetRefresh(s)
	return cuc
}

// SetExpiry sets the "expiry" field.
func (cuc *ChatUserCreate) SetExpiry(t time.Time) *ChatUserCreate {
	cuc.mutation.SetExpiry(t)
	return cuc
}

// SetBotToken sets the "bot_token" field.
func (cuc *ChatUserCreate) SetBotToken(s string) *ChatUserCreate {
	cuc.mutation.SetBotToken(s)
	return cuc
}

// SetCreatedAt sets the "created_at" field.
func (cuc *ChatUserCreate) SetCreatedAt(t time.Time) *ChatUserCreate {
	cuc.mutation.SetCreatedAt(t)
	return cuc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuc *ChatUserCreate) SetNillableCreatedAt(t *time.Time) *ChatUserCreate {
	if t != nil {
		cuc.SetCreatedAt(*t)
	}
	return cuc
}

// SetUpdatedAt sets the "updated_at" field.
func (cuc *ChatUserCreate) SetUpdatedAt(t time.Time) *ChatUserCreate {
	cuc.mutation.SetUpdatedAt(t)
	return cuc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuc *ChatUserCreate) SetNillableUpdatedAt(t *time.Time) *ChatUserCreate {
	if t != nil {
		cuc.SetUpdatedAt(*t)
	}
	return cuc
}

// SetUserID sets the "user_id" field.
func (cuc *ChatUserCreate) SetUserID(i int64) *ChatUserCreate {
	cuc.mutation.SetUserID(i)
	return cuc
}

// SetID sets the "id" field.
func (cuc *ChatUserCreate) SetID(s string) *ChatUserCreate {
	cuc.mutation.SetID(s)
	return cuc
}

// SetUser sets the "user" edge to the User entity.
func (cuc *ChatUserCreate) SetUser(u *User) *ChatUserCreate {
	return cuc.SetUserID(u.ID)
}

// Mutation returns the ChatUserMutation object of the builder.
func (cuc *ChatUserCreate) Mutation() *ChatUserMutation {
	return cuc.mutation
}

// Save creates the ChatUser in the database.
func (cuc *ChatUserCreate) Save(ctx context.Context) (*ChatUser, error) {
	var (
		err  error
		node *ChatUser
	)
	cuc.defaults()
	if len(cuc.hooks) == 0 {
		if err = cuc.check(); err != nil {
			return nil, err
		}
		node, err = cuc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuc.check(); err != nil {
				return nil, err
			}
			cuc.mutation = mutation
			if node, err = cuc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cuc.hooks) - 1; i >= 0; i-- {
			if cuc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cuc *ChatUserCreate) SaveX(ctx context.Context) *ChatUser {
	v, err := cuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cuc *ChatUserCreate) Exec(ctx context.Context) error {
	_, err := cuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuc *ChatUserCreate) ExecX(ctx context.Context) {
	if err := cuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuc *ChatUserCreate) defaults() {
	if _, ok := cuc.mutation.CreatedAt(); !ok {
		v := chatuser.DefaultCreatedAt()
		cuc.mutation.SetCreatedAt(v)
	}
	if _, ok := cuc.mutation.UpdatedAt(); !ok {
		v := chatuser.DefaultUpdatedAt()
		cuc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuc *ChatUserCreate) check() error {
	if _, ok := cuc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "ChatUser.token"`)}
	}
	if _, ok := cuc.mutation.Refresh(); !ok {
		return &ValidationError{Name: "refresh", err: errors.New(`ent: missing required field "ChatUser.refresh"`)}
	}
	if _, ok := cuc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`ent: missing required field "ChatUser.expiry"`)}
	}
	if _, ok := cuc.mutation.BotToken(); !ok {
		return &ValidationError{Name: "bot_token", err: errors.New(`ent: missing required field "ChatUser.bot_token"`)}
	}
	if _, ok := cuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ChatUser.created_at"`)}
	}
	if _, ok := cuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ChatUser.updated_at"`)}
	}
	if _, ok := cuc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ChatUser.user_id"`)}
	}
	if _, ok := cuc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "ChatUser.user"`)}
	}
	return nil
}

func (cuc *ChatUserCreate) sqlSave(ctx context.Context) (*ChatUser, error) {
	_node, _spec := cuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ChatUser.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (cuc *ChatUserCreate) createSpec() (*ChatUser, *sqlgraph.CreateSpec) {
	var (
		_node = &ChatUser{config: cuc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: chatuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chatuser.FieldID,
			},
		}
	)
	if id, ok := cuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cuc.mutation.Token(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldToken,
		})
		_node.Token = value
	}
	if value, ok := cuc.mutation.Refresh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldRefresh,
		})
		_node.Refresh = value
	}
	if value, ok := cuc.mutation.Expiry(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldExpiry,
		})
		_node.Expiry = value
	}
	if value, ok := cuc.mutation.BotToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatuser.FieldBotToken,
		})
		_node.BotToken = value
	}
	if value, ok := cuc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cuc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := cuc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChatUserCreateBulk is the builder for creating many ChatUser entities in bulk.
type ChatUserCreateBulk struct {
	config
	builders []*ChatUserCreate
}

// Save creates the ChatUser entities in the database.
func (cucb *ChatUserCreateBulk) Save(ctx context.Context) ([]*ChatUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cucb.builders))
	nodes := make([]*ChatUser, len(cucb.builders))
	mutators := make([]Mutator, len(cucb.builders))
	for i := range cucb.builders {
		func(i int, root context.Context) {
			builder := cucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChatUserMutation)
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
					_, err = mutators[i+1].Mutate(root, cucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cucb *ChatUserCreateBulk) SaveX(ctx context.Context) []*ChatUser {
	v, err := cucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cucb *ChatUserCreateBulk) Exec(ctx context.Context) error {
	_, err := cucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cucb *ChatUserCreateBulk) ExecX(ctx context.Context) {
	if err := cucb.Exec(ctx); err != nil {
		panic(err)
	}
}
