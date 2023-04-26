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
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.infratographer.com/instance-api/internal/ent/instance"
	"go.infratographer.com/instance-api/internal/ent/instancemetadata"
	"go.infratographer.com/instance-api/internal/ent/instanceprovider"
	"go.infratographer.com/instance-api/internal/ent/predicate"
)

// InstanceWhereInput represents a where input for filtering Instance queries.
type InstanceWhereInput struct {
	Predicates []predicate.Instance  `json:"-"`
	Not        *InstanceWhereInput   `json:"not,omitempty"`
	Or         []*InstanceWhereInput `json:"or,omitempty"`
	And        []*InstanceWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "instance_provider_id" field predicates.
	InstanceProviderID      *uuid.UUID  `json:"instanceProviderID,omitempty"`
	InstanceProviderIDNEQ   *uuid.UUID  `json:"instanceProviderIDNEQ,omitempty"`
	InstanceProviderIDIn    []uuid.UUID `json:"instanceProviderIDIn,omitempty"`
	InstanceProviderIDNotIn []uuid.UUID `json:"instanceProviderIDNotIn,omitempty"`

	// "tenant_id" field predicates.
	TenantID      *uuid.UUID  `json:"tenantID,omitempty"`
	TenantIDNEQ   *uuid.UUID  `json:"tenantIDNEQ,omitempty"`
	TenantIDIn    []uuid.UUID `json:"tenantIDIn,omitempty"`
	TenantIDNotIn []uuid.UUID `json:"tenantIDNotIn,omitempty"`
	TenantIDGT    *uuid.UUID  `json:"tenantIDGT,omitempty"`
	TenantIDGTE   *uuid.UUID  `json:"tenantIDGTE,omitempty"`
	TenantIDLT    *uuid.UUID  `json:"tenantIDLT,omitempty"`
	TenantIDLTE   *uuid.UUID  `json:"tenantIDLTE,omitempty"`

	// "instance_provider" edge predicates.
	HasInstanceProvider     *bool                         `json:"hasInstanceProvider,omitempty"`
	HasInstanceProviderWith []*InstanceProviderWhereInput `json:"hasInstanceProviderWith,omitempty"`

	// "metadata" edge predicates.
	HasMetadata     *bool                         `json:"hasMetadata,omitempty"`
	HasMetadataWith []*InstanceMetadataWhereInput `json:"hasMetadataWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *InstanceWhereInput) AddPredicates(predicates ...predicate.Instance) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the InstanceWhereInput filter on the InstanceQuery builder.
func (i *InstanceWhereInput) Filter(q *InstanceQuery) (*InstanceQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyInstanceWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyInstanceWhereInput is returned in case the InstanceWhereInput is empty.
var ErrEmptyInstanceWhereInput = errors.New("ent: empty predicate InstanceWhereInput")

// P returns a predicate for filtering instances.
// An error is returned if the input is empty or invalid.
func (i *InstanceWhereInput) P() (predicate.Instance, error) {
	var predicates []predicate.Instance
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, instance.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Instance, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, instance.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Instance, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, instance.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, instance.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, instance.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, instance.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, instance.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, instance.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, instance.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, instance.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, instance.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, instance.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, instance.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, instance.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, instance.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, instance.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, instance.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, instance.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, instance.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, instance.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, instance.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, instance.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, instance.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, instance.NameContainsFold(*i.NameContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, instance.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, instance.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, instance.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, instance.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, instance.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, instance.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, instance.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, instance.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, instance.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, instance.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, instance.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, instance.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, instance.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, instance.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, instance.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, instance.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.InstanceProviderID != nil {
		predicates = append(predicates, instance.InstanceProviderIDEQ(*i.InstanceProviderID))
	}
	if i.InstanceProviderIDNEQ != nil {
		predicates = append(predicates, instance.InstanceProviderIDNEQ(*i.InstanceProviderIDNEQ))
	}
	if len(i.InstanceProviderIDIn) > 0 {
		predicates = append(predicates, instance.InstanceProviderIDIn(i.InstanceProviderIDIn...))
	}
	if len(i.InstanceProviderIDNotIn) > 0 {
		predicates = append(predicates, instance.InstanceProviderIDNotIn(i.InstanceProviderIDNotIn...))
	}
	if i.TenantID != nil {
		predicates = append(predicates, instance.TenantIDEQ(*i.TenantID))
	}
	if i.TenantIDNEQ != nil {
		predicates = append(predicates, instance.TenantIDNEQ(*i.TenantIDNEQ))
	}
	if len(i.TenantIDIn) > 0 {
		predicates = append(predicates, instance.TenantIDIn(i.TenantIDIn...))
	}
	if len(i.TenantIDNotIn) > 0 {
		predicates = append(predicates, instance.TenantIDNotIn(i.TenantIDNotIn...))
	}
	if i.TenantIDGT != nil {
		predicates = append(predicates, instance.TenantIDGT(*i.TenantIDGT))
	}
	if i.TenantIDGTE != nil {
		predicates = append(predicates, instance.TenantIDGTE(*i.TenantIDGTE))
	}
	if i.TenantIDLT != nil {
		predicates = append(predicates, instance.TenantIDLT(*i.TenantIDLT))
	}
	if i.TenantIDLTE != nil {
		predicates = append(predicates, instance.TenantIDLTE(*i.TenantIDLTE))
	}

	if i.HasInstanceProvider != nil {
		p := instance.HasInstanceProvider()
		if !*i.HasInstanceProvider {
			p = instance.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasInstanceProviderWith) > 0 {
		with := make([]predicate.InstanceProvider, 0, len(i.HasInstanceProviderWith))
		for _, w := range i.HasInstanceProviderWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasInstanceProviderWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, instance.HasInstanceProviderWith(with...))
	}
	if i.HasMetadata != nil {
		p := instance.HasMetadata()
		if !*i.HasMetadata {
			p = instance.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasMetadataWith) > 0 {
		with := make([]predicate.InstanceMetadata, 0, len(i.HasMetadataWith))
		for _, w := range i.HasMetadataWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasMetadataWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, instance.HasMetadataWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyInstanceWhereInput
	case 1:
		return predicates[0], nil
	default:
		return instance.And(predicates...), nil
	}
}

// InstanceMetadataWhereInput represents a where input for filtering InstanceMetadata queries.
type InstanceMetadataWhereInput struct {
	Predicates []predicate.InstanceMetadata  `json:"-"`
	Not        *InstanceMetadataWhereInput   `json:"not,omitempty"`
	Or         []*InstanceMetadataWhereInput `json:"or,omitempty"`
	And        []*InstanceMetadataWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "namespace" field predicates.
	Namespace             *string  `json:"namespace,omitempty"`
	NamespaceNEQ          *string  `json:"namespaceNEQ,omitempty"`
	NamespaceIn           []string `json:"namespaceIn,omitempty"`
	NamespaceNotIn        []string `json:"namespaceNotIn,omitempty"`
	NamespaceGT           *string  `json:"namespaceGT,omitempty"`
	NamespaceGTE          *string  `json:"namespaceGTE,omitempty"`
	NamespaceLT           *string  `json:"namespaceLT,omitempty"`
	NamespaceLTE          *string  `json:"namespaceLTE,omitempty"`
	NamespaceContains     *string  `json:"namespaceContains,omitempty"`
	NamespaceHasPrefix    *string  `json:"namespaceHasPrefix,omitempty"`
	NamespaceHasSuffix    *string  `json:"namespaceHasSuffix,omitempty"`
	NamespaceEqualFold    *string  `json:"namespaceEqualFold,omitempty"`
	NamespaceContainsFold *string  `json:"namespaceContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "instance_id" field predicates.
	InstanceID      *uuid.UUID  `json:"instanceID,omitempty"`
	InstanceIDNEQ   *uuid.UUID  `json:"instanceIDNEQ,omitempty"`
	InstanceIDIn    []uuid.UUID `json:"instanceIDIn,omitempty"`
	InstanceIDNotIn []uuid.UUID `json:"instanceIDNotIn,omitempty"`

	// "instance" edge predicates.
	HasInstance     *bool                 `json:"hasInstance,omitempty"`
	HasInstanceWith []*InstanceWhereInput `json:"hasInstanceWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *InstanceMetadataWhereInput) AddPredicates(predicates ...predicate.InstanceMetadata) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the InstanceMetadataWhereInput filter on the InstanceMetadataQuery builder.
func (i *InstanceMetadataWhereInput) Filter(q *InstanceMetadataQuery) (*InstanceMetadataQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyInstanceMetadataWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyInstanceMetadataWhereInput is returned in case the InstanceMetadataWhereInput is empty.
var ErrEmptyInstanceMetadataWhereInput = errors.New("ent: empty predicate InstanceMetadataWhereInput")

// P returns a predicate for filtering instancemetadataslice.
// An error is returned if the input is empty or invalid.
func (i *InstanceMetadataWhereInput) P() (predicate.InstanceMetadata, error) {
	var predicates []predicate.InstanceMetadata
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, instancemetadata.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.InstanceMetadata, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, instancemetadata.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.InstanceMetadata, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, instancemetadata.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, instancemetadata.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, instancemetadata.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, instancemetadata.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, instancemetadata.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, instancemetadata.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, instancemetadata.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, instancemetadata.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, instancemetadata.IDLTE(*i.IDLTE))
	}
	if i.Namespace != nil {
		predicates = append(predicates, instancemetadata.NamespaceEQ(*i.Namespace))
	}
	if i.NamespaceNEQ != nil {
		predicates = append(predicates, instancemetadata.NamespaceNEQ(*i.NamespaceNEQ))
	}
	if len(i.NamespaceIn) > 0 {
		predicates = append(predicates, instancemetadata.NamespaceIn(i.NamespaceIn...))
	}
	if len(i.NamespaceNotIn) > 0 {
		predicates = append(predicates, instancemetadata.NamespaceNotIn(i.NamespaceNotIn...))
	}
	if i.NamespaceGT != nil {
		predicates = append(predicates, instancemetadata.NamespaceGT(*i.NamespaceGT))
	}
	if i.NamespaceGTE != nil {
		predicates = append(predicates, instancemetadata.NamespaceGTE(*i.NamespaceGTE))
	}
	if i.NamespaceLT != nil {
		predicates = append(predicates, instancemetadata.NamespaceLT(*i.NamespaceLT))
	}
	if i.NamespaceLTE != nil {
		predicates = append(predicates, instancemetadata.NamespaceLTE(*i.NamespaceLTE))
	}
	if i.NamespaceContains != nil {
		predicates = append(predicates, instancemetadata.NamespaceContains(*i.NamespaceContains))
	}
	if i.NamespaceHasPrefix != nil {
		predicates = append(predicates, instancemetadata.NamespaceHasPrefix(*i.NamespaceHasPrefix))
	}
	if i.NamespaceHasSuffix != nil {
		predicates = append(predicates, instancemetadata.NamespaceHasSuffix(*i.NamespaceHasSuffix))
	}
	if i.NamespaceEqualFold != nil {
		predicates = append(predicates, instancemetadata.NamespaceEqualFold(*i.NamespaceEqualFold))
	}
	if i.NamespaceContainsFold != nil {
		predicates = append(predicates, instancemetadata.NamespaceContainsFold(*i.NamespaceContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, instancemetadata.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, instancemetadata.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, instancemetadata.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, instancemetadata.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, instancemetadata.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, instancemetadata.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, instancemetadata.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, instancemetadata.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, instancemetadata.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, instancemetadata.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, instancemetadata.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, instancemetadata.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, instancemetadata.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, instancemetadata.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, instancemetadata.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, instancemetadata.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.InstanceID != nil {
		predicates = append(predicates, instancemetadata.InstanceIDEQ(*i.InstanceID))
	}
	if i.InstanceIDNEQ != nil {
		predicates = append(predicates, instancemetadata.InstanceIDNEQ(*i.InstanceIDNEQ))
	}
	if len(i.InstanceIDIn) > 0 {
		predicates = append(predicates, instancemetadata.InstanceIDIn(i.InstanceIDIn...))
	}
	if len(i.InstanceIDNotIn) > 0 {
		predicates = append(predicates, instancemetadata.InstanceIDNotIn(i.InstanceIDNotIn...))
	}

	if i.HasInstance != nil {
		p := instancemetadata.HasInstance()
		if !*i.HasInstance {
			p = instancemetadata.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasInstanceWith) > 0 {
		with := make([]predicate.Instance, 0, len(i.HasInstanceWith))
		for _, w := range i.HasInstanceWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasInstanceWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, instancemetadata.HasInstanceWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyInstanceMetadataWhereInput
	case 1:
		return predicates[0], nil
	default:
		return instancemetadata.And(predicates...), nil
	}
}

// InstanceProviderWhereInput represents a where input for filtering InstanceProvider queries.
type InstanceProviderWhereInput struct {
	Predicates []predicate.InstanceProvider  `json:"-"`
	Not        *InstanceProviderWhereInput   `json:"not,omitempty"`
	Or         []*InstanceProviderWhereInput `json:"or,omitempty"`
	And        []*InstanceProviderWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "instances" edge predicates.
	HasInstances     *bool                 `json:"hasInstances,omitempty"`
	HasInstancesWith []*InstanceWhereInput `json:"hasInstancesWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *InstanceProviderWhereInput) AddPredicates(predicates ...predicate.InstanceProvider) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the InstanceProviderWhereInput filter on the InstanceProviderQuery builder.
func (i *InstanceProviderWhereInput) Filter(q *InstanceProviderQuery) (*InstanceProviderQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyInstanceProviderWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyInstanceProviderWhereInput is returned in case the InstanceProviderWhereInput is empty.
var ErrEmptyInstanceProviderWhereInput = errors.New("ent: empty predicate InstanceProviderWhereInput")

// P returns a predicate for filtering instanceproviders.
// An error is returned if the input is empty or invalid.
func (i *InstanceProviderWhereInput) P() (predicate.InstanceProvider, error) {
	var predicates []predicate.InstanceProvider
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, instanceprovider.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.InstanceProvider, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, instanceprovider.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.InstanceProvider, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, instanceprovider.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, instanceprovider.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, instanceprovider.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, instanceprovider.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, instanceprovider.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, instanceprovider.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, instanceprovider.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, instanceprovider.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, instanceprovider.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, instanceprovider.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, instanceprovider.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, instanceprovider.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, instanceprovider.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, instanceprovider.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, instanceprovider.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, instanceprovider.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, instanceprovider.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, instanceprovider.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, instanceprovider.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, instanceprovider.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, instanceprovider.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, instanceprovider.NameContainsFold(*i.NameContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, instanceprovider.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, instanceprovider.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, instanceprovider.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, instanceprovider.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, instanceprovider.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, instanceprovider.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, instanceprovider.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, instanceprovider.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, instanceprovider.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, instanceprovider.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, instanceprovider.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, instanceprovider.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, instanceprovider.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, instanceprovider.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, instanceprovider.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, instanceprovider.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	if i.HasInstances != nil {
		p := instanceprovider.HasInstances()
		if !*i.HasInstances {
			p = instanceprovider.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasInstancesWith) > 0 {
		with := make([]predicate.Instance, 0, len(i.HasInstancesWith))
		for _, w := range i.HasInstancesWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasInstancesWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, instanceprovider.HasInstancesWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyInstanceProviderWhereInput
	case 1:
		return predicates[0], nil
	default:
		return instanceprovider.And(predicates...), nil
	}
}
