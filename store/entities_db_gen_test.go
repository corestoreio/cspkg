// Code generated by codegen. DO NOT EDIT.
// Generated by sql/dmlgen. DO NOT EDIT.
// +build csall db

package store

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/corestoreio/pkg/sql/ddl"
	"github.com/corestoreio/pkg/sql/dml"
	"github.com/corestoreio/pkg/sql/dmltest"
	"github.com/corestoreio/pkg/util/assert"
	"github.com/corestoreio/pkg/util/pseudo"
)

func TestNewTablesDB_8fbf75a91e4e6bd670de701be5c9ec85(t *testing.T) {
	db := dmltest.MustConnectDB(t)
	defer dmltest.Close(t, db)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()
	tbls, err := NewTables(ctx, ddl.WithConnPool(db))
	assert.NoError(t, err)
	tblNames := tbls.Tables()
	sort.Strings(tblNames)
	assert.Exactly(t, []string{"store", "store_group", "store_website"}, tblNames)
	err = tbls.Validate(ctx)
	assert.NoError(t, err)
	var ps *pseudo.Service
	ps = pseudo.MustNewService(0, &pseudo.Options{Lang: "de", MaxFloatDecimals: 6},
		pseudo.WithTagFakeFunc("website_id", func(maxLen int) (interface{}, error) {
			return 1, nil
		}),
		pseudo.WithTagFakeFunc("store_id", func(maxLen int) (interface{}, error) {
			return 1, nil
		}),
	)
	t.Run("Store_Entity", func(t *testing.T) {
		tbl := tbls.MustTable(TableNameStore)
		entSELECT := tbl.SelectByPK("*")
		// WithDBR generates the cached SQL string with empty key "".
		entSELECTStmtA := entSELECT.WithDBR().ExpandPlaceHolders()
		entSELECT.WithCacheKey("select_10").Wheres.Reset()
		_, _, err := entSELECT.Where(
			dml.Column("store_id").LessOrEqual().Int(10),
		).ToSQL() // ToSQL generates the new cached SQL string with key select_10
		assert.NoError(t, err)
		entCol := NewStores()
		entINSERT := tbl.Insert().BuildValues()
		entINSERTStmtA := entINSERT.PrepareWithDBR(ctx)
		for i := 0; i < 9; i++ {
			entIn := new(Store)
			if err := ps.FakeData(entIn); err != nil {
				t.Errorf("IDX[%d]: %+v", i, err)
				return
			}
			lID := dmltest.CheckLastInsertID(t, "Error: TestNewTables.Store_Entity")(entINSERTStmtA.Record("", entIn).ExecContext(ctx))
			entINSERTStmtA.Reset()
			entOut := new(Store)
			rowCount, err := entSELECTStmtA.Int64s(lID).Load(ctx, entOut)
			assert.NoError(t, err)
			assert.Exactly(t, uint64(1), rowCount, "IDX%d: RowCount did not match", i)
			assert.Exactly(t, entIn.StoreID, entOut.StoreID, "IDX%d: StoreID should match", lID)
			assert.ExactlyLength(t, 64, &entIn.Code, &entOut.Code, "IDX%d: Code should match", lID)
			assert.Exactly(t, entIn.WebsiteID, entOut.WebsiteID, "IDX%d: WebsiteID should match", lID)
			assert.Exactly(t, entIn.GroupID, entOut.GroupID, "IDX%d: GroupID should match", lID)
			assert.ExactlyLength(t, 255, &entIn.Name, &entOut.Name, "IDX%d: Name should match", lID)
			assert.Exactly(t, entIn.SortOrder, entOut.SortOrder, "IDX%d: SortOrder should match", lID)
			assert.Exactly(t, entIn.IsActive, entOut.IsActive, "IDX%d: IsActive should match", lID)
		}
		dmltest.Close(t, entINSERTStmtA)
		rowCount, err := entSELECTStmtA.WithCacheKey("select_10").Load(ctx, entCol)
		assert.NoError(t, err)
		t.Logf("Collection load rowCount: %d", rowCount)
		entINSERTStmtA = entINSERT.WithCacheKey("row_count_%d", len(entCol.Data)).Replace().SetRowCount(len(entCol.Data)).PrepareWithDBR(ctx)
		lID := dmltest.CheckLastInsertID(t, "Error:  Stores ")(entINSERTStmtA.Record("", entCol).ExecContext(ctx))
		dmltest.Close(t, entINSERTStmtA)
		t.Logf("Last insert ID into: %d", lID)
		t.Logf("INSERT queries: %#v", entINSERT.CachedQueries())
		t.Logf("SELECT queries: %#v", entSELECT.CachedQueries())
	})
	t.Run("StoreGroup_Entity", func(t *testing.T) {
		tbl := tbls.MustTable(TableNameStoreGroup)
		entSELECT := tbl.SelectByPK("*")
		// WithDBR generates the cached SQL string with empty key "".
		entSELECTStmtA := entSELECT.WithDBR().ExpandPlaceHolders()
		entSELECT.WithCacheKey("select_10").Wheres.Reset()
		_, _, err := entSELECT.Where(
			dml.Column("group_id").LessOrEqual().Int(10),
		).ToSQL() // ToSQL generates the new cached SQL string with key select_10
		assert.NoError(t, err)
		entCol := NewStoreGroups()
		entINSERT := tbl.Insert().BuildValues()
		entINSERTStmtA := entINSERT.PrepareWithDBR(ctx)
		for i := 0; i < 9; i++ {
			entIn := new(StoreGroup)
			if err := ps.FakeData(entIn); err != nil {
				t.Errorf("IDX[%d]: %+v", i, err)
				return
			}
			lID := dmltest.CheckLastInsertID(t, "Error: TestNewTables.StoreGroup_Entity")(entINSERTStmtA.Record("", entIn).ExecContext(ctx))
			entINSERTStmtA.Reset()
			entOut := new(StoreGroup)
			rowCount, err := entSELECTStmtA.Int64s(lID).Load(ctx, entOut)
			assert.NoError(t, err)
			assert.Exactly(t, uint64(1), rowCount, "IDX%d: RowCount did not match", i)
			assert.Exactly(t, entIn.GroupID, entOut.GroupID, "IDX%d: GroupID should match", lID)
			assert.Exactly(t, entIn.WebsiteID, entOut.WebsiteID, "IDX%d: WebsiteID should match", lID)
			assert.ExactlyLength(t, 255, &entIn.Name, &entOut.Name, "IDX%d: Name should match", lID)
			assert.Exactly(t, entIn.RootCategoryID, entOut.RootCategoryID, "IDX%d: RootCategoryID should match", lID)
			assert.Exactly(t, entIn.DefaultStoreID, entOut.DefaultStoreID, "IDX%d: DefaultStoreID should match", lID)
			assert.ExactlyLength(t, 64, &entIn.Code, &entOut.Code, "IDX%d: Code should match", lID)
		}
		dmltest.Close(t, entINSERTStmtA)
		rowCount, err := entSELECTStmtA.WithCacheKey("select_10").Load(ctx, entCol)
		assert.NoError(t, err)
		t.Logf("Collection load rowCount: %d", rowCount)
		entINSERTStmtA = entINSERT.WithCacheKey("row_count_%d", len(entCol.Data)).Replace().SetRowCount(len(entCol.Data)).PrepareWithDBR(ctx)
		lID := dmltest.CheckLastInsertID(t, "Error:  StoreGroups ")(entINSERTStmtA.Record("", entCol).ExecContext(ctx))
		dmltest.Close(t, entINSERTStmtA)
		t.Logf("Last insert ID into: %d", lID)
		t.Logf("INSERT queries: %#v", entINSERT.CachedQueries())
		t.Logf("SELECT queries: %#v", entSELECT.CachedQueries())
	})
	t.Run("StoreWebsite_Entity", func(t *testing.T) {
		tbl := tbls.MustTable(TableNameStoreWebsite)
		entSELECT := tbl.SelectByPK("*")
		// WithDBR generates the cached SQL string with empty key "".
		entSELECTStmtA := entSELECT.WithDBR().ExpandPlaceHolders()
		entSELECT.WithCacheKey("select_10").Wheres.Reset()
		_, _, err := entSELECT.Where(
			dml.Column("website_id").LessOrEqual().Int(10),
		).ToSQL() // ToSQL generates the new cached SQL string with key select_10
		assert.NoError(t, err)
		entCol := NewStoreWebsites()
		entINSERT := tbl.Insert().BuildValues()
		entINSERTStmtA := entINSERT.PrepareWithDBR(ctx)
		for i := 0; i < 9; i++ {
			entIn := new(StoreWebsite)
			if err := ps.FakeData(entIn); err != nil {
				t.Errorf("IDX[%d]: %+v", i, err)
				return
			}
			lID := dmltest.CheckLastInsertID(t, "Error: TestNewTables.StoreWebsite_Entity")(entINSERTStmtA.Record("", entIn).ExecContext(ctx))
			entINSERTStmtA.Reset()
			entOut := new(StoreWebsite)
			rowCount, err := entSELECTStmtA.Int64s(lID).Load(ctx, entOut)
			assert.NoError(t, err)
			assert.Exactly(t, uint64(1), rowCount, "IDX%d: RowCount did not match", i)
			assert.Exactly(t, entIn.WebsiteID, entOut.WebsiteID, "IDX%d: WebsiteID should match", lID)
			assert.ExactlyLength(t, 64, &entIn.Code, &entOut.Code, "IDX%d: Code should match", lID)
			assert.ExactlyLength(t, 128, &entIn.Name, &entOut.Name, "IDX%d: Name should match", lID)
			assert.Exactly(t, entIn.SortOrder, entOut.SortOrder, "IDX%d: SortOrder should match", lID)
			assert.Exactly(t, entIn.DefaultGroupID, entOut.DefaultGroupID, "IDX%d: DefaultGroupID should match", lID)
			assert.Exactly(t, entIn.IsDefault, entOut.IsDefault, "IDX%d: IsDefault should match", lID)
		}
		dmltest.Close(t, entINSERTStmtA)
		rowCount, err := entSELECTStmtA.WithCacheKey("select_10").Load(ctx, entCol)
		assert.NoError(t, err)
		t.Logf("Collection load rowCount: %d", rowCount)
		entINSERTStmtA = entINSERT.WithCacheKey("row_count_%d", len(entCol.Data)).Replace().SetRowCount(len(entCol.Data)).PrepareWithDBR(ctx)
		lID := dmltest.CheckLastInsertID(t, "Error:  StoreWebsites ")(entINSERTStmtA.Record("", entCol).ExecContext(ctx))
		dmltest.Close(t, entINSERTStmtA)
		t.Logf("Last insert ID into: %d", lID)
		t.Logf("INSERT queries: %#v", entINSERT.CachedQueries())
		t.Logf("SELECT queries: %#v", entSELECT.CachedQueries())
	})
}
