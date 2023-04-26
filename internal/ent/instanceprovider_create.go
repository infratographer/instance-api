// Copyright 2022 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"go.infratographer.com/instance-api/internal/ent/instance"
	"go.infratographer.com/instance-api/internal/ent/instanceprovider"
)

// InstanceProviderCreate is the builder for creating a InstanceProvider entity.
type InstanceProviderCreate struct {
	config
	mutation *InstanceProviderMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ipc *InstanceProviderCreate) SetName(s string) *InstanceProviderCreate {
	ipc.mutation.SetName(s)
	return ipc
}

// SetCreatedAt sets the "created_at" field.
func (ipc *InstanceProviderCreate) SetCreatedAt(t time.Time) *InstanceProviderCreate {
	ipc.mutation.SetCreatedAt(t)
	return ipc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ipc *InstanceProviderCreate) SetNillableCreatedAt(t *time.Time) *InstanceProviderCreate {
	if t != nil {
		ipc.SetCreatedAt(*t)
	}
	return ipc
}

// SetUpdatedAt sets the "updated_at" field.
func (ipc *InstanceProviderCreate) SetUpdatedAt(t time.Time) *InstanceProviderCreate {
	ipc.mutation.SetUpdatedAt(t)
	return ipc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ipc *InstanceProviderCreate) SetNillableUpdatedAt(t *time.Time) *InstanceProviderCreate {
	if t != nil {
		ipc.SetUpdatedAt(*t)
	}
	return ipc
}

// SetID sets the "id" field.
func (ipc *InstanceProviderCreate) SetID(u uuid.UUID) *InstanceProviderCreate {
	ipc.mutation.SetID(u)
	return ipc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ipc *InstanceProviderCreate) SetNillableID(u *uuid.UUID) *InstanceProviderCreate {
	if u != nil {
		ipc.SetID(*u)
	}
	return ipc
}

// AddInstanceIDs adds the "instances" edge to the Instance entity by IDs.
func (ipc *InstanceProviderCreate) AddInstanceIDs(ids ...uuid.UUID) *InstanceProviderCreate {
	ipc.mutation.AddInstanceIDs(ids...)
	return ipc
}

// AddInstances adds the "instances" edges to the Instance entity.
func (ipc *InstanceProviderCreate) AddInstances(i ...*Instance) *InstanceProviderCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ipc.AddInstanceIDs(ids...)
}

// Mutation returns the InstanceProviderMutation object of the builder.
func (ipc *InstanceProviderCreate) Mutation() *InstanceProviderMutation {
	return ipc.mutation
}

// Save creates the InstanceProvider in the database.
func (ipc *InstanceProviderCreate) Save(ctx context.Context) (*InstanceProvider, error) {
	ipc.defaults()
	return withHooks[*InstanceProvider, InstanceProviderMutation](ctx, ipc.sqlSave, ipc.mutation, ipc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ipc *InstanceProviderCreate) SaveX(ctx context.Context) *InstanceProvider {
	v, err := ipc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ipc *InstanceProviderCreate) Exec(ctx context.Context) error {
	_, err := ipc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipc *InstanceProviderCreate) ExecX(ctx context.Context) {
	if err := ipc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ipc *InstanceProviderCreate) defaults() {
	if _, ok := ipc.mutation.CreatedAt(); !ok {
		v := instanceprovider.DefaultCreatedAt()
		ipc.mutation.SetCreatedAt(v)
	}
	if _, ok := ipc.mutation.UpdatedAt(); !ok {
		v := instanceprovider.DefaultUpdatedAt()
		ipc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ipc.mutation.ID(); !ok {
		v := instanceprovider.DefaultID()
		ipc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ipc *InstanceProviderCreate) check() error {
	if _, ok := ipc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "InstanceProvider.name"`)}
	}
	if v, ok := ipc.mutation.Name(); ok {
		if err := instanceprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "InstanceProvider.name": %w`, err)}
		}
	}
	if _, ok := ipc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "InstanceProvider.created_at"`)}
	}
	if _, ok := ipc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "InstanceProvider.updated_at"`)}
	}
	return nil
}

func (ipc *InstanceProviderCreate) sqlSave(ctx context.Context) (*InstanceProvider, error) {
	if err := ipc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ipc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ipc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ipc.mutation.id = &_node.ID
	ipc.mutation.done = true
	return _node, nil
}

func (ipc *InstanceProviderCreate) createSpec() (*InstanceProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &InstanceProvider{config: ipc.config}
		_spec = sqlgraph.NewCreateSpec(instanceprovider.Table, sqlgraph.NewFieldSpec(instanceprovider.FieldID, field.TypeUUID))
	)
	if id, ok := ipc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ipc.mutation.Name(); ok {
		_spec.SetField(instanceprovider.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ipc.mutation.CreatedAt(); ok {
		_spec.SetField(instanceprovider.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ipc.mutation.UpdatedAt(); ok {
		_spec.SetField(instanceprovider.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ipc.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   instanceprovider.InstancesTable,
			Columns: []string{instanceprovider.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InstanceProviderCreateBulk is the builder for creating many InstanceProvider entities in bulk.
type InstanceProviderCreateBulk struct {
	config
	builders []*InstanceProviderCreate
}

// Save creates the InstanceProvider entities in the database.
func (ipcb *InstanceProviderCreateBulk) Save(ctx context.Context) ([]*InstanceProvider, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ipcb.builders))
	nodes := make([]*InstanceProvider, len(ipcb.builders))
	mutators := make([]Mutator, len(ipcb.builders))
	for i := range ipcb.builders {
		func(i int, root context.Context) {
			builder := ipcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InstanceProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, ipcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ipcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, ipcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ipcb *InstanceProviderCreateBulk) SaveX(ctx context.Context) []*InstanceProvider {
	v, err := ipcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ipcb *InstanceProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := ipcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipcb *InstanceProviderCreateBulk) ExecX(ctx context.Context) {
	if err := ipcb.Exec(ctx); err != nil {
		panic(err)
	}
}
