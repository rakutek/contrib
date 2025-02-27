// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithenum"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithEnumCreate is the builder for creating a MessageWithEnum entity.
type MessageWithEnumCreate struct {
	config
	mutation *MessageWithEnumMutation
	hooks    []Hook
}

// SetEnumType sets the "enum_type" field.
func (mwec *MessageWithEnumCreate) SetEnumType(mt messagewithenum.EnumType) *MessageWithEnumCreate {
	mwec.mutation.SetEnumType(mt)
	return mwec
}

// SetNillableEnumType sets the "enum_type" field if the given value is not nil.
func (mwec *MessageWithEnumCreate) SetNillableEnumType(mt *messagewithenum.EnumType) *MessageWithEnumCreate {
	if mt != nil {
		mwec.SetEnumType(*mt)
	}
	return mwec
}

// SetEnumWithoutDefault sets the "enum_without_default" field.
func (mwec *MessageWithEnumCreate) SetEnumWithoutDefault(mwd messagewithenum.EnumWithoutDefault) *MessageWithEnumCreate {
	mwec.mutation.SetEnumWithoutDefault(mwd)
	return mwec
}

// SetEnumWithSpecialCharacters sets the "enum_with_special_characters" field.
func (mwec *MessageWithEnumCreate) SetEnumWithSpecialCharacters(mwsc messagewithenum.EnumWithSpecialCharacters) *MessageWithEnumCreate {
	mwec.mutation.SetEnumWithSpecialCharacters(mwsc)
	return mwec
}

// Mutation returns the MessageWithEnumMutation object of the builder.
func (mwec *MessageWithEnumCreate) Mutation() *MessageWithEnumMutation {
	return mwec.mutation
}

// Save creates the MessageWithEnum in the database.
func (mwec *MessageWithEnumCreate) Save(ctx context.Context) (*MessageWithEnum, error) {
	mwec.defaults()
	return withHooks[*MessageWithEnum, MessageWithEnumMutation](ctx, mwec.sqlSave, mwec.mutation, mwec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mwec *MessageWithEnumCreate) SaveX(ctx context.Context) *MessageWithEnum {
	v, err := mwec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwec *MessageWithEnumCreate) Exec(ctx context.Context) error {
	_, err := mwec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwec *MessageWithEnumCreate) ExecX(ctx context.Context) {
	if err := mwec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mwec *MessageWithEnumCreate) defaults() {
	if _, ok := mwec.mutation.EnumType(); !ok {
		v := messagewithenum.DefaultEnumType
		mwec.mutation.SetEnumType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mwec *MessageWithEnumCreate) check() error {
	if _, ok := mwec.mutation.EnumType(); !ok {
		return &ValidationError{Name: "enum_type", err: errors.New(`ent: missing required field "MessageWithEnum.enum_type"`)}
	}
	if v, ok := mwec.mutation.EnumType(); ok {
		if err := messagewithenum.EnumTypeValidator(v); err != nil {
			return &ValidationError{Name: "enum_type", err: fmt.Errorf(`ent: validator failed for field "MessageWithEnum.enum_type": %w`, err)}
		}
	}
	if _, ok := mwec.mutation.EnumWithoutDefault(); !ok {
		return &ValidationError{Name: "enum_without_default", err: errors.New(`ent: missing required field "MessageWithEnum.enum_without_default"`)}
	}
	if v, ok := mwec.mutation.EnumWithoutDefault(); ok {
		if err := messagewithenum.EnumWithoutDefaultValidator(v); err != nil {
			return &ValidationError{Name: "enum_without_default", err: fmt.Errorf(`ent: validator failed for field "MessageWithEnum.enum_without_default": %w`, err)}
		}
	}
	if _, ok := mwec.mutation.EnumWithSpecialCharacters(); !ok {
		return &ValidationError{Name: "enum_with_special_characters", err: errors.New(`ent: missing required field "MessageWithEnum.enum_with_special_characters"`)}
	}
	if v, ok := mwec.mutation.EnumWithSpecialCharacters(); ok {
		if err := messagewithenum.EnumWithSpecialCharactersValidator(v); err != nil {
			return &ValidationError{Name: "enum_with_special_characters", err: fmt.Errorf(`ent: validator failed for field "MessageWithEnum.enum_with_special_characters": %w`, err)}
		}
	}
	return nil
}

func (mwec *MessageWithEnumCreate) sqlSave(ctx context.Context) (*MessageWithEnum, error) {
	if err := mwec.check(); err != nil {
		return nil, err
	}
	_node, _spec := mwec.createSpec()
	if err := sqlgraph.CreateNode(ctx, mwec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mwec.mutation.id = &_node.ID
	mwec.mutation.done = true
	return _node, nil
}

func (mwec *MessageWithEnumCreate) createSpec() (*MessageWithEnum, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageWithEnum{config: mwec.config}
		_spec = sqlgraph.NewCreateSpec(messagewithenum.Table, sqlgraph.NewFieldSpec(messagewithenum.FieldID, field.TypeInt))
	)
	if value, ok := mwec.mutation.EnumType(); ok {
		_spec.SetField(messagewithenum.FieldEnumType, field.TypeEnum, value)
		_node.EnumType = value
	}
	if value, ok := mwec.mutation.EnumWithoutDefault(); ok {
		_spec.SetField(messagewithenum.FieldEnumWithoutDefault, field.TypeEnum, value)
		_node.EnumWithoutDefault = value
	}
	if value, ok := mwec.mutation.EnumWithSpecialCharacters(); ok {
		_spec.SetField(messagewithenum.FieldEnumWithSpecialCharacters, field.TypeEnum, value)
		_node.EnumWithSpecialCharacters = value
	}
	return _node, _spec
}

// MessageWithEnumCreateBulk is the builder for creating many MessageWithEnum entities in bulk.
type MessageWithEnumCreateBulk struct {
	config
	builders []*MessageWithEnumCreate
}

// Save creates the MessageWithEnum entities in the database.
func (mwecb *MessageWithEnumCreateBulk) Save(ctx context.Context) ([]*MessageWithEnum, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mwecb.builders))
	nodes := make([]*MessageWithEnum, len(mwecb.builders))
	mutators := make([]Mutator, len(mwecb.builders))
	for i := range mwecb.builders {
		func(i int, root context.Context) {
			builder := mwecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageWithEnumMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mwecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mwecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, mwecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mwecb *MessageWithEnumCreateBulk) SaveX(ctx context.Context) []*MessageWithEnum {
	v, err := mwecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwecb *MessageWithEnumCreateBulk) Exec(ctx context.Context) error {
	_, err := mwecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwecb *MessageWithEnumCreateBulk) ExecX(ctx context.Context) {
	if err := mwecb.Exec(ctx); err != nil {
		panic(err)
	}
}
