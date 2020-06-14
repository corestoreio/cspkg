// Code generated by corestoreio/pkg/util/codegen. DO NOT EDIT.
// Generated by sql/dmlgen. DO NOT EDIT.
package dmltestgenerated2

import (
	"testing"

	"github.com/corestoreio/errors"
	"github.com/corestoreio/pkg/util/assert"
	"github.com/corestoreio/pkg/util/pseudo"
)

func TestNewDBManagerNonDB_48a8450c0b62e880b2d40acd0bbbd0dc(t *testing.T) {
	ps := pseudo.MustNewService(0, &pseudo.Options{Lang: "de", FloatMaxDecimals: 6})
	_ = ps
	t.Run("CoreConfiguration_Empty", func(t *testing.T) {
		e := new(CoreConfiguration)
		assert.NoError(t, ps.FakeData(e))
		e.Empty()
		assert.Exactly(t, *e, CoreConfiguration{})
	})
	t.Run("CoreConfiguration_Copy", func(t *testing.T) {
		e := new(CoreConfiguration)
		assert.NoError(t, ps.FakeData(e))
		e2 := e.Copy()
		assert.Exactly(t, e, e2)
		assert.NoError(t, ps.FakeData(e))
		assert.NotEqual(t, e, e2)
	})
	t.Run("CoreConfigurations_Validate", func(t *testing.T) {
		c := CoreConfigurations{Data: []*CoreConfiguration{nil}}
		assert.True(t, errors.NotValid.Match(c.Validate()))
	})
	t.Run("SalesOrderStatusState_Empty", func(t *testing.T) {
		e := new(SalesOrderStatusState)
		assert.NoError(t, ps.FakeData(e))
		e.Empty()
		assert.Exactly(t, *e, SalesOrderStatusState{})
	})
	t.Run("SalesOrderStatusState_Copy", func(t *testing.T) {
		e := new(SalesOrderStatusState)
		assert.NoError(t, ps.FakeData(e))
		e2 := e.Copy()
		assert.Exactly(t, e, e2)
		assert.NoError(t, ps.FakeData(e))
		assert.NotEqual(t, e, e2)
	})
	t.Run("SalesOrderStatusStates_Validate", func(t *testing.T) {
		c := SalesOrderStatusStates{Data: []*SalesOrderStatusState{nil}}
		assert.True(t, errors.NotValid.Match(c.Validate()))
	})
}
