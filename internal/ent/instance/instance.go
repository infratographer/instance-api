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

package instance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the instance type in the database.
	Label = "instance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldInstanceProviderID holds the string denoting the instance_provider_id field in the database.
	FieldInstanceProviderID = "instance_provider_id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// EdgeInstanceProvider holds the string denoting the instance_provider edge name in mutations.
	EdgeInstanceProvider = "instance_provider"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// Table holds the table name of the instance in the database.
	Table = "instances"
	// InstanceProviderTable is the table that holds the instance_provider relation/edge.
	InstanceProviderTable = "instances"
	// InstanceProviderInverseTable is the table name for the InstanceProvider entity.
	// It exists in this package in order to avoid circular dependency with the "instanceprovider" package.
	InstanceProviderInverseTable = "instance_providers"
	// InstanceProviderColumn is the table column denoting the instance_provider relation/edge.
	InstanceProviderColumn = "instance_provider_id"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "instance_metadata"
	// MetadataInverseTable is the table name for the InstanceMetadata entity.
	// It exists in this package in order to avoid circular dependency with the "instancemetadata" package.
	MetadataInverseTable = "instance_metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "instance_id"
)

// Columns holds all SQL columns for instance fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldInstanceProviderID,
	FieldTenantID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Instance queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByInstanceProviderID orders the results by the instance_provider_id field.
func ByInstanceProviderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceProviderID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByInstanceProviderField orders the results by instance_provider field.
func ByInstanceProviderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInstanceProviderStep(), sql.OrderByField(field, opts...))
	}
}

// ByMetadataCount orders the results by metadata count.
func ByMetadataCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMetadataStep(), opts...)
	}
}

// ByMetadata orders the results by metadata terms.
func ByMetadata(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetadataStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newInstanceProviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InstanceProviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, InstanceProviderTable, InstanceProviderColumn),
	)
}
func newMetadataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetadataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, MetadataTable, MetadataColumn),
	)
}
