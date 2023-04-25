// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProviders)
	t.Run("Instances", testInstances)
}

func TestSoftDelete(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersSoftDelete)
	t.Run("Instances", testInstancesSoftDelete)
}

func TestQuerySoftDeleteAll(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersQuerySoftDeleteAll)
	t.Run("Instances", testInstancesQuerySoftDeleteAll)
}

func TestSliceSoftDeleteAll(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersSliceSoftDeleteAll)
	t.Run("Instances", testInstancesSliceSoftDeleteAll)
}

func TestDelete(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersDelete)
	t.Run("Instances", testInstancesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersQueryDeleteAll)
	t.Run("Instances", testInstancesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersSliceDeleteAll)
	t.Run("Instances", testInstancesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersExists)
	t.Run("Instances", testInstancesExists)
}

func TestFind(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersFind)
	t.Run("Instances", testInstancesFind)
}

func TestBind(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersBind)
	t.Run("Instances", testInstancesBind)
}

func TestOne(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersOne)
	t.Run("Instances", testInstancesOne)
}

func TestAll(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersAll)
	t.Run("Instances", testInstancesAll)
}

func TestCount(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersCount)
	t.Run("Instances", testInstancesCount)
}

func TestHooks(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersHooks)
	t.Run("Instances", testInstancesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersInsert)
	t.Run("InstanceProviders", testInstanceProvidersInsertWhitelist)
	t.Run("Instances", testInstancesInsert)
	t.Run("Instances", testInstancesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("InstanceToInstanceProviderUsingInstanceProvider", testInstanceToOneInstanceProviderUsingInstanceProvider)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("InstanceProviderToInstances", testInstanceProviderToManyInstances)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("InstanceToInstanceProviderUsingInstances", testInstanceToOneSetOpInstanceProviderUsingInstanceProvider)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("InstanceProviderToInstances", testInstanceProviderToManyAddOpInstances)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersReload)
	t.Run("Instances", testInstancesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersReloadAll)
	t.Run("Instances", testInstancesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersSelect)
	t.Run("Instances", testInstancesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersUpdate)
	t.Run("Instances", testInstancesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("InstanceProviders", testInstanceProvidersSliceUpdateAll)
	t.Run("Instances", testInstancesSliceUpdateAll)
}
