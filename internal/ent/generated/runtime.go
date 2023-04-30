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
	"time"

	"go.infratographer.com/instance-api/internal/ent/generated/instance"
	"go.infratographer.com/instance-api/internal/ent/generated/instancemetadata"
	"go.infratographer.com/instance-api/internal/ent/generated/instanceprovider"
	"go.infratographer.com/instance-api/internal/ent/schema"
	"go.infratographer.com/x/gidx"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	instanceMixin := schema.Instance{}.Mixin()
	instanceMixinFields0 := instanceMixin[0].Fields()
	_ = instanceMixinFields0
	instanceFields := schema.Instance{}.Fields()
	_ = instanceFields
	// instanceDescCreatedAt is the schema descriptor for created_at field.
	instanceDescCreatedAt := instanceMixinFields0[0].Descriptor()
	// instance.DefaultCreatedAt holds the default value on creation for the created_at field.
	instance.DefaultCreatedAt = instanceDescCreatedAt.Default.(func() time.Time)
	// instanceDescUpdatedAt is the schema descriptor for updated_at field.
	instanceDescUpdatedAt := instanceMixinFields0[1].Descriptor()
	// instance.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	instance.DefaultUpdatedAt = instanceDescUpdatedAt.Default.(func() time.Time)
	// instance.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	instance.UpdateDefaultUpdatedAt = instanceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// instanceDescName is the schema descriptor for name field.
	instanceDescName := instanceFields[1].Descriptor()
	// instance.NameValidator is a validator for the "name" field. It is called by the builders before save.
	instance.NameValidator = instanceDescName.Validators[0].(func(string) error)
	// instanceDescID is the schema descriptor for id field.
	instanceDescID := instanceFields[0].Descriptor()
	// instance.DefaultID holds the default value on creation for the id field.
	instance.DefaultID = instanceDescID.Default.(func() gidx.PrefixedID)
	instancemetadataMixin := schema.InstanceMetadata{}.Mixin()
	instancemetadataMixinFields0 := instancemetadataMixin[0].Fields()
	_ = instancemetadataMixinFields0
	instancemetadataMixinFields1 := instancemetadataMixin[1].Fields()
	_ = instancemetadataMixinFields1
	instancemetadataFields := schema.InstanceMetadata{}.Fields()
	_ = instancemetadataFields
	// instancemetadataDescNamespace is the schema descriptor for namespace field.
	instancemetadataDescNamespace := instancemetadataMixinFields0[0].Descriptor()
	// instancemetadata.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	instancemetadata.NamespaceValidator = func() func(string) error {
		validators := instancemetadataDescNamespace.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(namespace string) error {
			for _, fn := range fns {
				if err := fn(namespace); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// instancemetadataDescCreatedAt is the schema descriptor for created_at field.
	instancemetadataDescCreatedAt := instancemetadataMixinFields1[0].Descriptor()
	// instancemetadata.DefaultCreatedAt holds the default value on creation for the created_at field.
	instancemetadata.DefaultCreatedAt = instancemetadataDescCreatedAt.Default.(func() time.Time)
	// instancemetadataDescUpdatedAt is the schema descriptor for updated_at field.
	instancemetadataDescUpdatedAt := instancemetadataMixinFields1[1].Descriptor()
	// instancemetadata.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	instancemetadata.DefaultUpdatedAt = instancemetadataDescUpdatedAt.Default.(func() time.Time)
	// instancemetadata.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	instancemetadata.UpdateDefaultUpdatedAt = instancemetadataDescUpdatedAt.UpdateDefault.(func() time.Time)
	// instancemetadataDescID is the schema descriptor for id field.
	instancemetadataDescID := instancemetadataFields[0].Descriptor()
	// instancemetadata.DefaultID holds the default value on creation for the id field.
	instancemetadata.DefaultID = instancemetadataDescID.Default.(func() gidx.PrefixedID)
	instanceproviderMixin := schema.InstanceProvider{}.Mixin()
	instanceproviderMixinFields0 := instanceproviderMixin[0].Fields()
	_ = instanceproviderMixinFields0
	instanceproviderFields := schema.InstanceProvider{}.Fields()
	_ = instanceproviderFields
	// instanceproviderDescCreatedAt is the schema descriptor for created_at field.
	instanceproviderDescCreatedAt := instanceproviderMixinFields0[0].Descriptor()
	// instanceprovider.DefaultCreatedAt holds the default value on creation for the created_at field.
	instanceprovider.DefaultCreatedAt = instanceproviderDescCreatedAt.Default.(func() time.Time)
	// instanceproviderDescUpdatedAt is the schema descriptor for updated_at field.
	instanceproviderDescUpdatedAt := instanceproviderMixinFields0[1].Descriptor()
	// instanceprovider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	instanceprovider.DefaultUpdatedAt = instanceproviderDescUpdatedAt.Default.(func() time.Time)
	// instanceprovider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	instanceprovider.UpdateDefaultUpdatedAt = instanceproviderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// instanceproviderDescName is the schema descriptor for name field.
	instanceproviderDescName := instanceproviderFields[1].Descriptor()
	// instanceprovider.NameValidator is a validator for the "name" field. It is called by the builders before save.
	instanceprovider.NameValidator = instanceproviderDescName.Validators[0].(func(string) error)
	// instanceproviderDescID is the schema descriptor for id field.
	instanceproviderDescID := instanceproviderFields[0].Descriptor()
	// instanceprovider.DefaultID holds the default value on creation for the id field.
	instanceprovider.DefaultID = instanceproviderDescID.Default.(func() gidx.PrefixedID)
}
