// Code generated by codegen. DO NOT EDIT.
// Generated by sql/dmlgen. DO NOT EDIT.
package store

import (
	"testing"

	"github.com/corestoreio/errors"
	"github.com/corestoreio/pkg/util/assert"
	"github.com/corestoreio/pkg/util/pseudo"
)

func TestNewTablesNonDB_8fbf75a91e4e6bd670de701be5c9ec85(t *testing.T) {
	ps := pseudo.MustNewService(0, &pseudo.Options{Lang: "de", MaxFloatDecimals: 6})
	_ = ps
	t.Run("Store_Empty", func(t *testing.T) {
		e := new(Store)
		assert.NoError(t, ps.FakeData(e))
		e.Empty()
		assert.Exactly(t, *e, Store{})
	})
	t.Run("Store_Copy", func(t *testing.T) {
		e := new(Store)
		assert.NoError(t, ps.FakeData(e))
		e2 := e.Copy()
		assert.Exactly(t, e, e2)
		assert.NoError(t, ps.FakeData(e))
		assert.NotEqual(t, e, e2)
	})
	t.Run("Stores_Validate", func(t *testing.T) {
		c := Stores{Data: []*Store{nil}}
		assert.True(t, errors.NotValid.Match(c.Validate()))
	})
	t.Run("StoreGroup_Empty", func(t *testing.T) {
		e := new(StoreGroup)
		assert.NoError(t, ps.FakeData(e))
		e.Empty()
		assert.Exactly(t, *e, StoreGroup{})
	})
	t.Run("StoreGroup_Copy", func(t *testing.T) {
		e := new(StoreGroup)
		assert.NoError(t, ps.FakeData(e))
		e2 := e.Copy()
		assert.Exactly(t, e, e2)
		assert.NoError(t, ps.FakeData(e))
		assert.NotEqual(t, e, e2)
	})
	t.Run("StoreGroups_Validate", func(t *testing.T) {
		c := StoreGroups{Data: []*StoreGroup{nil}}
		assert.True(t, errors.NotValid.Match(c.Validate()))
	})
	t.Run("StoreWebsite_Empty", func(t *testing.T) {
		e := new(StoreWebsite)
		assert.NoError(t, ps.FakeData(e))
		e.Empty()
		assert.Exactly(t, *e, StoreWebsite{})
	})
	t.Run("StoreWebsite_Copy", func(t *testing.T) {
		e := new(StoreWebsite)
		assert.NoError(t, ps.FakeData(e))
		e2 := e.Copy()
		assert.Exactly(t, e, e2)
		assert.NoError(t, ps.FakeData(e))
		assert.NotEqual(t, e, e2)
	})
	t.Run("StoreWebsites_Validate", func(t *testing.T) {
		c := StoreWebsites{Data: []*StoreWebsite{nil}}
		assert.True(t, errors.NotValid.Match(c.Validate()))
	})
}
