// Copyright 2023 The Infratographer Authors
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

package generated

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/instance-api/internal/ent/generated/instancemetadata"
	"go.infratographer.com/instance-api/internal/ent/generated/predicate"
)

// InstanceMetadataUpdate is the builder for updating InstanceMetadata entities.
type InstanceMetadataUpdate struct {
	config
	hooks    []Hook
	mutation *InstanceMetadataMutation
}

// Where appends a list predicates to the InstanceMetadataUpdate builder.
func (imu *InstanceMetadataUpdate) Where(ps ...predicate.InstanceMetadata) *InstanceMetadataUpdate {
	imu.mutation.Where(ps...)
	return imu
}

// SetNamespace sets the "namespace" field.
func (imu *InstanceMetadataUpdate) SetNamespace(s string) *InstanceMetadataUpdate {
	imu.mutation.SetNamespace(s)
	return imu
}

// SetData sets the "data" field.
func (imu *InstanceMetadataUpdate) SetData(jm json.RawMessage) *InstanceMetadataUpdate {
	imu.mutation.SetData(jm)
	return imu
}

// AppendData appends jm to the "data" field.
func (imu *InstanceMetadataUpdate) AppendData(jm json.RawMessage) *InstanceMetadataUpdate {
	imu.mutation.AppendData(jm)
	return imu
}

// Mutation returns the InstanceMetadataMutation object of the builder.
func (imu *InstanceMetadataUpdate) Mutation() *InstanceMetadataMutation {
	return imu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (imu *InstanceMetadataUpdate) Save(ctx context.Context) (int, error) {
	imu.defaults()
	return withHooks[int, InstanceMetadataMutation](ctx, imu.sqlSave, imu.mutation, imu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (imu *InstanceMetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := imu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (imu *InstanceMetadataUpdate) Exec(ctx context.Context) error {
	_, err := imu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (imu *InstanceMetadataUpdate) ExecX(ctx context.Context) {
	if err := imu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (imu *InstanceMetadataUpdate) defaults() {
	if _, ok := imu.mutation.UpdatedAt(); !ok {
		v := instancemetadata.UpdateDefaultUpdatedAt()
		imu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (imu *InstanceMetadataUpdate) check() error {
	if v, ok := imu.mutation.Namespace(); ok {
		if err := instancemetadata.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`generated: validator failed for field "InstanceMetadata.namespace": %w`, err)}
		}
	}
	if _, ok := imu.mutation.InstanceID(); imu.mutation.InstanceCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "InstanceMetadata.instance"`)
	}
	return nil
}

func (imu *InstanceMetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := imu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(instancemetadata.Table, instancemetadata.Columns, sqlgraph.NewFieldSpec(instancemetadata.FieldID, field.TypeString))
	if ps := imu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := imu.mutation.Namespace(); ok {
		_spec.SetField(instancemetadata.FieldNamespace, field.TypeString, value)
	}
	if value, ok := imu.mutation.Data(); ok {
		_spec.SetField(instancemetadata.FieldData, field.TypeJSON, value)
	}
	if value, ok := imu.mutation.AppendedData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, instancemetadata.FieldData, value)
		})
	}
	if value, ok := imu.mutation.UpdatedAt(); ok {
		_spec.SetField(instancemetadata.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, imu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instancemetadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	imu.mutation.done = true
	return n, nil
}

// InstanceMetadataUpdateOne is the builder for updating a single InstanceMetadata entity.
type InstanceMetadataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InstanceMetadataMutation
}

// SetNamespace sets the "namespace" field.
func (imuo *InstanceMetadataUpdateOne) SetNamespace(s string) *InstanceMetadataUpdateOne {
	imuo.mutation.SetNamespace(s)
	return imuo
}

// SetData sets the "data" field.
func (imuo *InstanceMetadataUpdateOne) SetData(jm json.RawMessage) *InstanceMetadataUpdateOne {
	imuo.mutation.SetData(jm)
	return imuo
}

// AppendData appends jm to the "data" field.
func (imuo *InstanceMetadataUpdateOne) AppendData(jm json.RawMessage) *InstanceMetadataUpdateOne {
	imuo.mutation.AppendData(jm)
	return imuo
}

// Mutation returns the InstanceMetadataMutation object of the builder.
func (imuo *InstanceMetadataUpdateOne) Mutation() *InstanceMetadataMutation {
	return imuo.mutation
}

// Where appends a list predicates to the InstanceMetadataUpdate builder.
func (imuo *InstanceMetadataUpdateOne) Where(ps ...predicate.InstanceMetadata) *InstanceMetadataUpdateOne {
	imuo.mutation.Where(ps...)
	return imuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (imuo *InstanceMetadataUpdateOne) Select(field string, fields ...string) *InstanceMetadataUpdateOne {
	imuo.fields = append([]string{field}, fields...)
	return imuo
}

// Save executes the query and returns the updated InstanceMetadata entity.
func (imuo *InstanceMetadataUpdateOne) Save(ctx context.Context) (*InstanceMetadata, error) {
	imuo.defaults()
	return withHooks[*InstanceMetadata, InstanceMetadataMutation](ctx, imuo.sqlSave, imuo.mutation, imuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (imuo *InstanceMetadataUpdateOne) SaveX(ctx context.Context) *InstanceMetadata {
	node, err := imuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (imuo *InstanceMetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := imuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (imuo *InstanceMetadataUpdateOne) ExecX(ctx context.Context) {
	if err := imuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (imuo *InstanceMetadataUpdateOne) defaults() {
	if _, ok := imuo.mutation.UpdatedAt(); !ok {
		v := instancemetadata.UpdateDefaultUpdatedAt()
		imuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (imuo *InstanceMetadataUpdateOne) check() error {
	if v, ok := imuo.mutation.Namespace(); ok {
		if err := instancemetadata.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`generated: validator failed for field "InstanceMetadata.namespace": %w`, err)}
		}
	}
	if _, ok := imuo.mutation.InstanceID(); imuo.mutation.InstanceCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "InstanceMetadata.instance"`)
	}
	return nil
}

func (imuo *InstanceMetadataUpdateOne) sqlSave(ctx context.Context) (_node *InstanceMetadata, err error) {
	if err := imuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(instancemetadata.Table, instancemetadata.Columns, sqlgraph.NewFieldSpec(instancemetadata.FieldID, field.TypeString))
	id, ok := imuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "InstanceMetadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := imuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instancemetadata.FieldID)
		for _, f := range fields {
			if !instancemetadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != instancemetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := imuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := imuo.mutation.Namespace(); ok {
		_spec.SetField(instancemetadata.FieldNamespace, field.TypeString, value)
	}
	if value, ok := imuo.mutation.Data(); ok {
		_spec.SetField(instancemetadata.FieldData, field.TypeJSON, value)
	}
	if value, ok := imuo.mutation.AppendedData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, instancemetadata.FieldData, value)
		})
	}
	if value, ok := imuo.mutation.UpdatedAt(); ok {
		_spec.SetField(instancemetadata.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &InstanceMetadata{config: imuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, imuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instancemetadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	imuo.mutation.done = true
	return _node, nil
}
