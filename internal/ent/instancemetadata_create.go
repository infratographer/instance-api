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
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"go.infratographer.com/instance-api/internal/ent/instance"
	"go.infratographer.com/instance-api/internal/ent/instancemetadata"
)

// InstanceMetadataCreate is the builder for creating a InstanceMetadata entity.
type InstanceMetadataCreate struct {
	config
	mutation *InstanceMetadataMutation
	hooks    []Hook
}

// SetNamespace sets the "namespace" field.
func (imc *InstanceMetadataCreate) SetNamespace(s string) *InstanceMetadataCreate {
	imc.mutation.SetNamespace(s)
	return imc
}

// SetData sets the "data" field.
func (imc *InstanceMetadataCreate) SetData(jm json.RawMessage) *InstanceMetadataCreate {
	imc.mutation.SetData(jm)
	return imc
}

// SetCreatedAt sets the "created_at" field.
func (imc *InstanceMetadataCreate) SetCreatedAt(t time.Time) *InstanceMetadataCreate {
	imc.mutation.SetCreatedAt(t)
	return imc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (imc *InstanceMetadataCreate) SetNillableCreatedAt(t *time.Time) *InstanceMetadataCreate {
	if t != nil {
		imc.SetCreatedAt(*t)
	}
	return imc
}

// SetUpdatedAt sets the "updated_at" field.
func (imc *InstanceMetadataCreate) SetUpdatedAt(t time.Time) *InstanceMetadataCreate {
	imc.mutation.SetUpdatedAt(t)
	return imc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (imc *InstanceMetadataCreate) SetNillableUpdatedAt(t *time.Time) *InstanceMetadataCreate {
	if t != nil {
		imc.SetUpdatedAt(*t)
	}
	return imc
}

// SetInstanceID sets the "instance_id" field.
func (imc *InstanceMetadataCreate) SetInstanceID(u uuid.UUID) *InstanceMetadataCreate {
	imc.mutation.SetInstanceID(u)
	return imc
}

// SetID sets the "id" field.
func (imc *InstanceMetadataCreate) SetID(u uuid.UUID) *InstanceMetadataCreate {
	imc.mutation.SetID(u)
	return imc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (imc *InstanceMetadataCreate) SetNillableID(u *uuid.UUID) *InstanceMetadataCreate {
	if u != nil {
		imc.SetID(*u)
	}
	return imc
}

// SetInstance sets the "instance" edge to the Instance entity.
func (imc *InstanceMetadataCreate) SetInstance(i *Instance) *InstanceMetadataCreate {
	return imc.SetInstanceID(i.ID)
}

// Mutation returns the InstanceMetadataMutation object of the builder.
func (imc *InstanceMetadataCreate) Mutation() *InstanceMetadataMutation {
	return imc.mutation
}

// Save creates the InstanceMetadata in the database.
func (imc *InstanceMetadataCreate) Save(ctx context.Context) (*InstanceMetadata, error) {
	imc.defaults()
	return withHooks[*InstanceMetadata, InstanceMetadataMutation](ctx, imc.sqlSave, imc.mutation, imc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (imc *InstanceMetadataCreate) SaveX(ctx context.Context) *InstanceMetadata {
	v, err := imc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (imc *InstanceMetadataCreate) Exec(ctx context.Context) error {
	_, err := imc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (imc *InstanceMetadataCreate) ExecX(ctx context.Context) {
	if err := imc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (imc *InstanceMetadataCreate) defaults() {
	if _, ok := imc.mutation.CreatedAt(); !ok {
		v := instancemetadata.DefaultCreatedAt()
		imc.mutation.SetCreatedAt(v)
	}
	if _, ok := imc.mutation.UpdatedAt(); !ok {
		v := instancemetadata.DefaultUpdatedAt()
		imc.mutation.SetUpdatedAt(v)
	}
	if _, ok := imc.mutation.ID(); !ok {
		v := instancemetadata.DefaultID()
		imc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (imc *InstanceMetadataCreate) check() error {
	if _, ok := imc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required field "InstanceMetadata.namespace"`)}
	}
	if v, ok := imc.mutation.Namespace(); ok {
		if err := instancemetadata.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`ent: validator failed for field "InstanceMetadata.namespace": %w`, err)}
		}
	}
	if _, ok := imc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "InstanceMetadata.data"`)}
	}
	if _, ok := imc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "InstanceMetadata.created_at"`)}
	}
	if _, ok := imc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "InstanceMetadata.updated_at"`)}
	}
	if _, ok := imc.mutation.InstanceID(); !ok {
		return &ValidationError{Name: "instance_id", err: errors.New(`ent: missing required field "InstanceMetadata.instance_id"`)}
	}
	if _, ok := imc.mutation.InstanceID(); !ok {
		return &ValidationError{Name: "instance", err: errors.New(`ent: missing required edge "InstanceMetadata.instance"`)}
	}
	return nil
}

func (imc *InstanceMetadataCreate) sqlSave(ctx context.Context) (*InstanceMetadata, error) {
	if err := imc.check(); err != nil {
		return nil, err
	}
	_node, _spec := imc.createSpec()
	if err := sqlgraph.CreateNode(ctx, imc.driver, _spec); err != nil {
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
	imc.mutation.id = &_node.ID
	imc.mutation.done = true
	return _node, nil
}

func (imc *InstanceMetadataCreate) createSpec() (*InstanceMetadata, *sqlgraph.CreateSpec) {
	var (
		_node = &InstanceMetadata{config: imc.config}
		_spec = sqlgraph.NewCreateSpec(instancemetadata.Table, sqlgraph.NewFieldSpec(instancemetadata.FieldID, field.TypeUUID))
	)
	if id, ok := imc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := imc.mutation.Namespace(); ok {
		_spec.SetField(instancemetadata.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := imc.mutation.Data(); ok {
		_spec.SetField(instancemetadata.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	if value, ok := imc.mutation.CreatedAt(); ok {
		_spec.SetField(instancemetadata.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := imc.mutation.UpdatedAt(); ok {
		_spec.SetField(instancemetadata.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := imc.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   instancemetadata.InstanceTable,
			Columns: []string{instancemetadata.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.InstanceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InstanceMetadataCreateBulk is the builder for creating many InstanceMetadata entities in bulk.
type InstanceMetadataCreateBulk struct {
	config
	builders []*InstanceMetadataCreate
}

// Save creates the InstanceMetadata entities in the database.
func (imcb *InstanceMetadataCreateBulk) Save(ctx context.Context) ([]*InstanceMetadata, error) {
	specs := make([]*sqlgraph.CreateSpec, len(imcb.builders))
	nodes := make([]*InstanceMetadata, len(imcb.builders))
	mutators := make([]Mutator, len(imcb.builders))
	for i := range imcb.builders {
		func(i int, root context.Context) {
			builder := imcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InstanceMetadataMutation)
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
					_, err = mutators[i+1].Mutate(root, imcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, imcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, imcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (imcb *InstanceMetadataCreateBulk) SaveX(ctx context.Context) []*InstanceMetadata {
	v, err := imcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (imcb *InstanceMetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := imcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (imcb *InstanceMetadataCreateBulk) ExecX(ctx context.Context) {
	if err := imcb.Exec(ctx); err != nil {
		panic(err)
	}
}
