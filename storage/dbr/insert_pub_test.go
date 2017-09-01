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

package dbr_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/corestoreio/csfw/storage/dbr"
	"github.com/corestoreio/csfw/util/cstesting"
	"github.com/corestoreio/errors"
	"github.com/corestoreio/log/logw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync/atomic"
	"testing"
	"time"
)

var _ dbr.ArgumentsAppender = (*someRecord)(nil)

type someRecord struct {
	SomethingID int
	UserID      int64
	Other       bool
}

func (sr someRecord) appendArgs(args dbr.Arguments, column string) (_ dbr.Arguments, err error) {
	switch column {
	case "something_id":
		args = args.Int(sr.SomethingID)
	case "user_id":
		args = args.Int64(sr.UserID)
	case "other":
		args = args.Bool(sr.Other)
	default:
		err = errors.NewNotFoundf("[dbr_test] Column %q not found", column)
	}
	return args, err
}

func (sr someRecord) AppendArgs(args dbr.Arguments, columns []string) (_ dbr.Arguments, err error) {
	l := len(columns)
	if l == 1 {
		return sr.appendArgs(args, columns[0])
	}
	if l == 0 {
		return args.Int(sr.SomethingID).Int64(sr.UserID).Bool(sr.Other), nil // except auto inc column ;-)
	}
	for _, col := range columns {
		if args, err = sr.appendArgs(args, col); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return args, err
}

func TestInsert_Bind(t *testing.T) {
	t.Parallel()
	objs := []someRecord{{1, 88, false}, {2, 99, true}, {3, 101, true}}
	wantArgs := []interface{}{int64(1), int64(88), false, int64(2), int64(99), true, int64(3), int64(101), true, int64(99)}

	t.Run("valid with multiple records", func(t *testing.T) {
		compareToSQL(t,
			dbr.NewInsert("a").
				AddColumns("something_id", "user_id", "other").
				BindRecord(objs[0]).BindRecord(objs[1], objs[2]).
				AddOnDuplicateKey(
					dbr.Column("something_id").Int64(99),
					dbr.Column("user_id").Values(),
				),
			nil,
			"INSERT INTO `a` (`something_id`,`user_id`,`other`) VALUES (?,?,?),(?,?,?),(?,?,?) ON DUPLICATE KEY UPDATE `something_id`=?, `user_id`=VALUES(`user_id`)",
			"INSERT INTO `a` (`something_id`,`user_id`,`other`) VALUES (1,88,0),(2,99,1),(3,101,1) ON DUPLICATE KEY UPDATE `something_id`=99, `user_id`=VALUES(`user_id`)",
			wantArgs...,
		)
	})
	t.Run("without columns, all columns requested", func(t *testing.T) {
		compareToSQL(t,
			dbr.NewInsert("a").
				SetRecordValueCount(3).
				BindRecord(objs[0]).BindRecord(objs[1], objs[2]).
				AddOnDuplicateKey(
					dbr.Column("something_id").Int64(99),
					dbr.Column("user_id").Values(),
				),
			nil,
			"INSERT INTO `a` VALUES (?,?,?),(?,?,?),(?,?,?) ON DUPLICATE KEY UPDATE `something_id`=?, `user_id`=VALUES(`user_id`)",
			"INSERT INTO `a` VALUES (1,88,0),(2,99,1),(3,101,1) ON DUPLICATE KEY UPDATE `something_id`=99, `user_id`=VALUES(`user_id`)",
			wantArgs...,
		)
	})
	t.Run("column not found", func(t *testing.T) {
		objs := []someRecord{{1, 88, false}, {2, 99, true}}
		compareToSQL(t,
			dbr.NewInsert("a").AddColumns("something_it", "user_id", "other").BindRecord(objs[0]).BindRecord(objs[1]),
			errors.IsNotFound,
			"",
			"",
		)
	})
}

func TestInsert_Prepare(t *testing.T) {
	t.Parallel()

	t.Run("ToSQL Error", func(t *testing.T) {
		in := &dbr.Insert{}
		in.AddColumns("a", "b")
		stmt, err := in.Prepare(context.TODO())
		assert.Nil(t, stmt)
		assert.True(t, errors.IsEmpty(err))
	})

	t.Run("DB Error", func(t *testing.T) {
		in := &dbr.Insert{
			Into: "table",
		}
		in.DB = dbMock{
			error: errors.NewAlreadyClosedf("Who closed myself?"),
		}
		in.AddColumns("a", "b").AddValues(1, true)

		stmt, err := in.Prepare(context.TODO())
		assert.Nil(t, stmt)
		assert.True(t, errors.IsAlreadyClosed(err), "%+v", err)
	})

	t.Run("ExecArgs One Row", func(t *testing.T) {
		dbc, dbMock := cstesting.MockDB(t)
		defer cstesting.MockClose(t, dbc, dbMock)

		prep := dbMock.ExpectPrepare(cstesting.SQLMockQuoteMeta("INSERT INTO `customer_entity` (`email`,`group_id`,`created_at`) VALUES (?,?,?)"))
		prep.ExpectExec().WithArgs("a@b.c", 33, now()).WillReturnResult(sqlmock.NewResult(4, 0))
		prep.ExpectExec().WithArgs("x@y.z", 44, now().Add(time.Minute)).WillReturnResult(sqlmock.NewResult(5, 0))

		stmt, err := dbr.NewInsert("customer_entity").
			AddColumns("email", "group_id", "created_at").
			WithDB(dbc.DB).
			Prepare(context.TODO())
		require.NoError(t, err, "failed creating a prepared statement")
		defer func() {
			require.NoError(t, stmt.Close(), "Close on a prepared statement")
		}()

		tests := []struct {
			email      string
			groupID    int
			created_at time.Time
			insertID   int64
		}{
			{"a@b.c", 33, now(), 4},
			{"x@y.z", 44, now().Add(time.Minute), 5},
		}

		args := dbr.MakeArgs(3)
		for i, test := range tests {
			args = args[:0]

			res, err := stmt.
				WithArguments(args.Str(test.email).Int(test.groupID).Time(test.created_at)).
				ExecContext(context.TODO())
			if err != nil {
				t.Fatalf("Index %d => %+v", i, err)
			}
			lid, err := res.LastInsertId()
			if err != nil {
				t.Fatalf("Result index %d with error: %s", i, err)
			}
			assert.Exactly(t, test.insertID, lid, "Index %d has different LastInsertIDs", i)
		}
	})

	t.Run("ExecArgs Multi Row", func(t *testing.T) {
		dbc, dbMock := cstesting.MockDB(t)
		defer cstesting.MockClose(t, dbc, dbMock)

		prep := dbMock.ExpectPrepare(cstesting.SQLMockQuoteMeta("INSERT INTO `customer_entity` (`email`,`group_id`) VALUES (?,?),(?,?)"))
		prep.ExpectExec().WithArgs("a@b.c", 33, "d@e.f", 33).WillReturnResult(sqlmock.NewResult(6, 0))
		prep.ExpectExec().WithArgs("x@y.z", 44, "u@v.w", 44).WillReturnResult(sqlmock.NewResult(7, 0))

		stmt, err := dbr.NewInsert("customer_entity").
			AddColumns("email", "group_id").
			SetRowCount(2).
			WithDB(dbc.DB).
			Prepare(context.TODO())
		require.NoError(t, err)
		defer func() {
			require.NoError(t, stmt.Close(), "Close on a prepared statement")
		}()

		tests := []struct {
			email1   string
			groupID1 int
			email2   string
			groupID2 int
			insertID int64
		}{
			{"a@b.c", 33, "d@e.f", 33, 6},
			{"x@y.z", 44, "u@v.w", 44, 7},
		}

		args := dbr.MakeArgs(4)
		for i, test := range tests {
			args = args[:0]

			res, err := stmt.
				WithArguments(args.Str(test.email1).Int(test.groupID1).Str(test.email2).Int(test.groupID2)).
				ExecContext(context.TODO())
			if err != nil {
				t.Fatalf("Index %d => %+v", i, err)
			}
			lid, err := res.LastInsertId()
			if err != nil {
				t.Fatalf("Result index %d with error: %s", i, err)
			}
			assert.Exactly(t, test.insertID, lid, "Index %d has different LastInsertIDs", i)
		}
	})

	t.Run("ExecRecord One Row", func(t *testing.T) {
		dbc, dbMock := cstesting.MockDB(t)
		defer cstesting.MockClose(t, dbc, dbMock)

		prep := dbMock.ExpectPrepare(cstesting.SQLMockQuoteMeta("INSERT INTO `dbr_person` (`name`,`email`) VALUES (?,?)"))
		prep.ExpectExec().WithArgs("Peter Gopher", "peter@gopher.go").WillReturnResult(sqlmock.NewResult(4, 0))
		prep.ExpectExec().WithArgs("John Doe", "john@doe.go").WillReturnResult(sqlmock.NewResult(5, 0))

		stmt, err := dbr.NewInsert("dbr_person").
			AddColumns("name", "email").
			WithDB(dbc.DB).
			Prepare(context.TODO())
		require.NoError(t, err, "failed creating a prepared statement")
		defer func() {
			require.NoError(t, stmt.Close(), "Close on a prepared statement")
		}()

		tests := []struct {
			name     string
			email    string
			insertID int64
		}{
			{"Peter Gopher", "peter@gopher.go", 4},
			{"John Doe", "john@doe.go", 5},
		}

		for i, test := range tests {

			p := &dbrPerson{
				Name:  test.name,
				Email: dbr.MakeNullString(test.email),
			}

			res, err := stmt.WithRecords(p).ExecContext(context.TODO())
			if err != nil {
				t.Fatalf("Index %d => %+v", i, err)
			}
			lid, err := res.LastInsertId()
			if err != nil {
				t.Fatalf("Result index %d with error: %s", i, err)
			}
			assert.Exactly(t, test.insertID, lid, "Index %d has different LastInsertIDs", i)
			assert.Exactly(t, test.insertID, p.ID, "Index %d and model p has different LastInsertIDs", i)
		}
	})

	t.Run("ExecContext", func(t *testing.T) {
		dbc, dbMock := cstesting.MockDB(t)
		defer cstesting.MockClose(t, dbc, dbMock)

		prep := dbMock.ExpectPrepare(cstesting.SQLMockQuoteMeta("INSERT INTO `dbr_person` (`name`,`email`) VALUES (?,?)"))
		prep.ExpectExec().WithArgs("Peter Gopher", "peter@gopher.go").WillReturnResult(sqlmock.NewResult(4, 0))

		stmt, err := dbr.NewInsert("dbr_person").
			AddColumns("name", "email").
			WithDB(dbc.DB).
			Prepare(context.TODO())
		require.NoError(t, err, "failed creating a prepared statement")
		defer func() {
			require.NoError(t, stmt.Close(), "Close on a prepared statement")
		}()

		res, err := stmt.ExecContext(context.TODO(), "Peter Gopher", "peter@gopher.go")
		require.NoError(t, err, "failed to execute ExecContext")

		lid, err := res.LastInsertId()
		if err != nil {
			t.Fatal(err)
		}
		assert.Exactly(t, int64(4), lid, "Different LastInsertIDs")
	})
}

func TestInsert_WithLogger(t *testing.T) {
	uniID := new(int32)
	rConn := createRealSession(t)
	defer cstesting.Close(t, rConn)

	var uniqueIDFunc = func() string {
		return fmt.Sprintf("UNIQ%02d", atomic.AddInt32(uniID, 4))
	}

	buf := new(bytes.Buffer)
	lg := logw.NewLog(
		logw.WithLevel(logw.LevelDebug),
		logw.WithWriter(buf),
		logw.WithFlag(0), // no flags at all
	)
	require.NoError(t, rConn.Options(dbr.WithLogger(lg, uniqueIDFunc)))

	t.Run("ConnPool", func(t *testing.T) {
		d := rConn.InsertInto("dbr_people").Replace().AddColumns("email", "name").AddValues("a@b.c", "John")

		t.Run("Exec", func(t *testing.T) {
			defer func() {
				buf.Reset()
				d.IsInterpolate = false
			}()
			_, err := d.Interpolate().Exec(context.TODO())
			require.NoError(t, err)

			assert.Exactly(t, "DEBUG Exec conn_pool_id: \"UNIQ04\" insert_id: \"UNIQ08\" table: \"dbr_people\" duration: 0 sql: \"REPLACE /*ID:UNIQ08*/ INTO `dbr_people` (`email`,`name`) VALUES ('a@b.c','John')\"\n",
				buf.String())
		})

		t.Run("Prepare", func(t *testing.T) {
			defer buf.Reset()
			stmt, err := d.Prepare(context.TODO())
			require.NoError(t, err)
			defer stmt.Close()

			assert.Exactly(t, "DEBUG Prepare conn_pool_id: \"UNIQ04\" insert_id: \"UNIQ08\" table: \"dbr_people\" duration: 0 sql: \"REPLACE /*ID:UNIQ08*/ INTO `dbr_people` (`email`,`name`) VALUES (?,?)\"\n",
				buf.String())
		})

		t.Run("Tx Commit", func(t *testing.T) {
			defer buf.Reset()
			tx, err := rConn.BeginTx(context.TODO(), nil)
			require.NoError(t, err)
			require.NoError(t, tx.Wrap(func() error {
				_, err := tx.InsertInto("dbr_people").Replace().AddColumns("email", "name").AddValues("a@b.c", "John").Interpolate().Exec(context.TODO())
				return err
			}))
			assert.Exactly(t, "DEBUG BeginTx conn_pool_id: \"UNIQ04\" tx_id: \"UNIQ12\"\nDEBUG Exec conn_pool_id: \"UNIQ04\" tx_id: \"UNIQ12\" insert_id: \"UNIQ16\" table: \"dbr_people\" duration: 0 sql: \"REPLACE /*ID:UNIQ16*/ INTO `dbr_people` (`email`,`name`) VALUES ('a@b.c','John')\"\nDEBUG Commit conn_pool_id: \"UNIQ04\" tx_id: \"UNIQ12\" duration: 0\n",
				buf.String())
		})
	})

	t.Run("Conn", func(t *testing.T) {
		conn, err := rConn.Conn(context.TODO())
		require.NoError(t, err)

		d := conn.InsertInto("dbr_people").Replace().AddColumns("email", "name").AddValues("a@b.zeh", "J0hn")

		t.Run("Exec", func(t *testing.T) {
			defer func() {
				buf.Reset()
				d.IsInterpolate = false
			}()

			_, err := d.Interpolate().Exec(context.TODO())
			require.NoError(t, err)

			assert.Exactly(t, "DEBUG Exec conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" insert_id: \"UNIQ24\" table: \"dbr_people\" duration: 0 sql: \"REPLACE /*ID:UNIQ24*/ INTO `dbr_people` (`email`,`name`) VALUES ('a@b.zeh','J0hn')\"\n",
				buf.String())
		})

		t.Run("Prepare", func(t *testing.T) {
			defer buf.Reset()

			stmt, err := d.Prepare(context.TODO())
			require.NoError(t, err)
			defer stmt.Close()

			assert.Exactly(t, "DEBUG Prepare conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" insert_id: \"UNIQ24\" table: \"dbr_people\" duration: 0 sql: \"REPLACE /*ID:UNIQ24*/ INTO `dbr_people` (`email`,`name`) VALUES (?,?)\"\n",
				buf.String())
		})

		t.Run("Prepare Exec", func(t *testing.T) {
			defer buf.Reset()

			stmt, err := d.Prepare(context.TODO())
			require.NoError(t, err)
			defer stmt.Close()

			_, err = stmt.Exec(context.TODO(), "mail@e.de", "Hans")
			require.NoError(t, err)

			assert.Exactly(t, "DEBUG Prepare conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" insert_id: \"UNIQ24\" table: \"dbr_people\" duration: 0 sql: \"REPLACE /*ID:UNIQ24*/ INTO `dbr_people` (`email`,`name`) VALUES (?,?)\"\nDEBUG Exec conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" insert_id: \"UNIQ24\" table: \"dbr_people\" duration: 0 arg_len: 2\n",
				buf.String())
		})

		t.Run("Tx Commit", func(t *testing.T) {
			defer buf.Reset()
			tx, err := conn.BeginTx(context.TODO(), nil)
			require.NoError(t, err)
			require.NoError(t, tx.Wrap(func() error {
				_, err := tx.InsertInto("dbr_people").Replace().AddColumns("email", "name").AddValues("a@b.c", "John").Interpolate().Exec(context.TODO())
				return err
			}))

			assert.Exactly(t, "DEBUG BeginTx conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" tx_id: \"UNIQ28\"\nDEBUG Exec conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" tx_id: \"UNIQ28\" insert_id: \"UNIQ32\" table: \"dbr_people\" duration: 0 sql: \"REPLACE /*ID:UNIQ32*/ INTO `dbr_people` (`email`,`name`) VALUES ('a@b.c','John')\"\nDEBUG Commit conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" tx_id: \"UNIQ28\" duration: 0\n",
				buf.String())
		})

		t.Run("Tx Rollback", func(t *testing.T) {
			defer buf.Reset()
			tx, err := conn.BeginTx(context.TODO(), nil)
			require.NoError(t, err)
			require.Error(t, tx.Wrap(func() error {
				_, err := tx.InsertInto("dbr_people").Replace().AddColumns("email", "name").
					AddArguments(dbr.MakeArgs(2).PlaceHolder().PlaceHolder()).Interpolate().Exec(context.TODO())
				return err
			}))

			assert.Exactly(t, "DEBUG BeginTx conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" tx_id: \"UNIQ36\"\nDEBUG Exec conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" tx_id: \"UNIQ36\" insert_id: \"UNIQ40\" table: \"dbr_people\" duration: 0 sql: \"REPLACE /*ID:UNIQ40*/ INTO `dbr_people` (`email`,`name`) VALUES (?,?)\"\nDEBUG Rollback conn_pool_id: \"UNIQ04\" conn_id: \"UNIQ20\" tx_id: \"UNIQ36\" duration: 0\n",
				buf.String())
		})
	})
}