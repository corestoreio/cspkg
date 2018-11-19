// Copyright 2015-2017, Cyrill @ Schumacher.fm and the CoreStore contributors
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

package dmlgen_test

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/corestoreio/errors"
	"github.com/corestoreio/pkg/sql/ddl"
	"github.com/corestoreio/pkg/sql/dmlgen"
	"github.com/corestoreio/pkg/sql/dmltest"
	"github.com/corestoreio/pkg/storage/null"
	"github.com/corestoreio/pkg/util/assert"
)

var _ = null.JSONMarshalFn

/*
SELECT
  concat('col_',
         replace(
             replace(
                 replace(
                     replace(COLUMN_TYPE, '(', '_')
                     , ')', '')
                 , ' ', '_')
             , ',', '_')
  )
    AS ColName,
  COLUMN_TYPE,
  IF(IS_NULLABLE = 'NO', 'NOT NULL', ''),
  ' DEFAULT',
  COLUMN_DEFAULT,
  ','
FROM information_schema.COLUMNS
WHERE
  table_schema = 'magento22' AND
  column_type IN (SELECT column_type
                  FROM information_schema.`COLUMNS`
                  GROUP BY column_type)
GROUP BY COLUMN_TYPE
ORDER BY COLUMN_TYPE
*/

func writeFile(t *testing.T, outFile string, w func(io.Writer) error) {
	f, err := os.Create(outFile)
	assert.NoError(t, err)
	defer dmltest.Close(t, f)
	err = w(f)
	assert.NoError(t, err, "%+v", err)
}

// TestNewTables_Generated writes a Go and Proto file to the testdata directory for manual
// review for different tables. This test also analyzes the foreign keys
// pointing to customer_entity. No tests of the generated source code are
// getting executed because API gets developed, still.
func TestNewTables_Generated(t *testing.T) {
	db := dmltest.MustConnectDB(t)
	defer dmltest.Close(t, db)

	// defer dmltest.SQLDumpLoad(t, "testdata/test_*.sql", nil)()
	dmltest.SQLDumpLoad(t, "testdata/test_*.sql", nil)

	ctx := context.Background()
	ts, err := dmlgen.NewTables("testdata",

		dmlgen.WithLoadColumns(ctx, db.DB, "dmlgen_types", "core_config_data", "customer_entity"),
		dmlgen.WithTableOption(
			"customer_entity", &dmlgen.TableOption{
				Encoders: []string{"json", "protobuf"},
			}),

		// dmlgen.WithLoadColumns(ctx, db.DB, "dmlgen_types"),

		dmlgen.WithTableOption(
			"core_config_data", &dmlgen.TableOption{
				Encoders: []string{"json", "protobuf"},
				CustomStructTags: []string{
					"path", `json:"x_path" xml:"y_path"`,
					"scope_id", `json:"scope_id" xml:"scope_id"`,
				},
				StructTags: []string{"json"},
				ColumnAliases: map[string][]string{
					"path": {"storage_location", "config_directory"},
				},
				UniquifiedColumns: []string{"path"},
			}),

		dmlgen.WithTable("core_config_data", ddl.Columns{
			&ddl.Column{Field: "path", Pos: 5, Default: null.MakeString("'general'"), Null: "NO", DataType: "varchar", CharMaxLength: null.MakeInt64(255), ColumnType: "varchar(255)", Comment: "Config Path overwritten"},
		}, "overwrite"),

		dmlgen.WithTableOption(
			"dmlgen_types", &dmlgen.TableOption{
				Encoders:          []string{"json", "binary", "protobuf"},
				StructTags:        []string{"json", "protobuf"},
				UniquifiedColumns: []string{"price_12_4a", "col_longtext_2", "col_int_1", "col_int_2", "has_smallint_5", "col_date_2", "col_blob"},
				Comment:           "Just another comment.\n//easyjson:json",
			}),
		dmlgen.WithTableOption(
			"dmlgen_types", &dmlgen.TableOption{
				Encoders: []string{"json", "protobuf"},
			}),

		dmlgen.WithColumnAliasesFromForeignKeys(ctx, db.DB),
	)
	assert.NoError(t, err)

	writeFile(t, "testdata/output_gen.go", ts.WriteGo)
	writeFile(t, "testdata/output_gen.proto", ts.WriteProto)
	// Generates for all proto files the Go source code.
	err = dmlgen.GenerateProto("./testdata")
	assert.NoError(t, err, "%+v", err)
}

func TestInfoSchemaForeignKeys(t *testing.T) {

	t.Skip("One time test. Use when needed to regenerate the code")

	db := dmltest.MustConnectDB(t)
	defer dmltest.Close(t, db)

	ts, err := dmlgen.NewTables("testdata",
		dmlgen.WithTableOption("KEY_COLUMN_USAGE", &dmlgen.TableOption{
			Encoders:          []string{"json", "binary"},
			UniquifiedColumns: []string{"TABLE_NAME", "COLUMN_NAME"},
		}),
		dmlgen.WithLoadColumns(context.Background(), db.DB, "KEY_COLUMN_USAGE"),
	)
	assert.NoError(t, err)

	writeFile(t, "testdata/KEY_COLUMN_USAGE_gen.go", ts.WriteGo)
}

func TestWithCustomStructTags(t *testing.T) {
	t.Parallel()

	t.Run("unbalanced should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					assert.True(t, errors.Fatal.Match(err), "%s", err)
				} else {
					t.Errorf("Panic should contain an error but got:\n%+v", r)
				}
			} else {
				t.Error("Expecting a panic but got nothing")
			}
		}()

		tbl, err := dmlgen.NewTables("testdata",
			dmlgen.WithTable("table", ddl.Columns{&ddl.Column{Field: "config_id"}}),
			dmlgen.WithTableOption("table", &dmlgen.TableOption{
				CustomStructTags: []string{"unbalanced"},
			}),
		)
		assert.Nil(t, tbl)
		assert.NoError(t, err)
	})

	t.Run("table not found", func(t *testing.T) {
		tbls, err := dmlgen.NewTables("test",
			dmlgen.WithTableOption("tableNOTFOUND", &dmlgen.TableOption{
				CustomStructTags: []string{"column", "db:..."},
			}),
		)
		assert.Nil(t, tbls)
		assert.True(t, errors.NotFound.Match(err), "%+v", err)
	})

	t.Run("column not found", func(t *testing.T) {
		tbls, err := dmlgen.NewTables("test",
			dmlgen.WithTableOption("core_config_data", &dmlgen.TableOption{
				CustomStructTags: []string{"scope_id", "toml:..."},
			}),
			dmlgen.WithTable("core_config_data", ddl.Columns{
				&ddl.Column{Field: "config_id"},
			}),
		)
		assert.Nil(t, tbls)
		assert.True(t, errors.NotFound.Match(err), "%+v", err)
	})
}

func TestWithStructTags(t *testing.T) {
	t.Parallel()

	t.Run("table not found", func(t *testing.T) {
		tbls, err := dmlgen.NewTables("test",
			dmlgen.WithTableOption("tableNOTFOUND", &dmlgen.TableOption{
				StructTags: []string{"unbalanced"},
			}),
		)
		assert.Nil(t, tbls)
		assert.True(t, errors.NotFound.Match(err), "%+v", err)
	})

	t.Run("struct tag not supported", func(t *testing.T) {
		tbls, err := dmlgen.NewTables("test",
			dmlgen.WithTableOption("core_config_data", &dmlgen.TableOption{
				StructTags: []string{"hjson"},
			}),
			dmlgen.WithTable("core_config_data", ddl.Columns{
				&ddl.Column{Field: "config_id"},
			}),
		)
		assert.Nil(t, tbls)
		assert.True(t, errors.NotSupported.Match(err), "%+v", err)
	})

	t.Run("al available struct tags", func(t *testing.T) {
		tbls, err := dmlgen.NewTables("test",
			dmlgen.WithTableOption("core_config_data", &dmlgen.TableOption{
				StructTags: []string{"bson", "db", "env", "json", "toml", "yaml", "xml"},
			}),
			dmlgen.WithTable("core_config_data", ddl.Columns{
				&ddl.Column{Field: "config_id"},
			}),
		)
		assert.NoError(t, err)
		have := tbls.Tables["core_config_data"].Columns.ByField("config_id").GoString()
		assert.Exactly(t, "&ddl.Column{Field: \"config_id\", StructTag: \"bson:\\\"config_id,omitempty\\\" db:\\\"config_id\\\" env:\\\"config_id\\\" json:\\\"config_id,omitempty\\\" toml:\\\"config_id\\\" yaml:\\\"config_id,omitempty\\\" xml:\\\"config_id,omitempty\\\"\", }", have)
	})
}

func TestWithColumnAliases(t *testing.T) {
	t.Parallel()

	t.Run("table not found", func(t *testing.T) {
		tbls, err := dmlgen.NewTables("test",
			dmlgen.WithTableOption("tableNOTFOUND", &dmlgen.TableOption{
				ColumnAliases: map[string][]string{"column": {"alias"}},
			}),
		)
		assert.Nil(t, tbls)
		assert.True(t, errors.NotFound.Match(err), "%+v", err)
	})

	t.Run("column not found", func(t *testing.T) {
		tbls, err := dmlgen.NewTables("test",
			dmlgen.WithTableOption("tableNOTFOUND", &dmlgen.TableOption{
				ColumnAliases: map[string][]string{"scope_id": {"scopeID"}},
			}),
			dmlgen.WithTable("core_config_data", ddl.Columns{
				&ddl.Column{Field: "config_id"},
			}),
		)
		assert.Nil(t, tbls)
		assert.True(t, errors.NotFound.Match(err), "%+v", err)
	})
}

func TestWithUniquifiedColumns(t *testing.T) {
	t.Parallel()

	t.Run("column not found", func(t *testing.T) {
		tbls, err := dmlgen.NewTables("test",
			dmlgen.WithTableOption("core_config_data", &dmlgen.TableOption{
				UniquifiedColumns: []string{"scope_id", "scopeID"},
			}),

			dmlgen.WithTable("core_config_data", ddl.Columns{
				&ddl.Column{Field: "config_id"},
			}),
		)
		assert.Nil(t, tbls)
		assert.True(t, errors.NotFound.Match(err), "%+v", err)
	})
}
