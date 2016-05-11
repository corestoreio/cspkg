// Copyright 2015-2016, Cyrill @ Schumacher.fm and the CoreStore contributors
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

package store_test

import (
	"testing"

	"context"
	"github.com/corestoreio/csfw/storage/dbr"
	"github.com/corestoreio/csfw/store"
	"github.com/corestoreio/csfw/store/scope"
	"github.com/corestoreio/csfw/store/storemock"
	"github.com/corestoreio/csfw/util/errors"
	"github.com/stretchr/testify/assert"
)

func TestContextReaderError(t *testing.T) {
	t.Parallel()
	haveMr, s, err := store.FromContextProvider(context.Background())
	assert.Nil(t, haveMr)
	assert.Nil(t, s)
	assert.True(t, errors.IsNotFound(err))

	ctx := store.WithContextProvider(context.Background(), nil)
	assert.NotNil(t, ctx)
	haveMr, s, err = store.FromContextProvider(ctx)
	assert.Nil(t, haveMr)
	assert.Nil(t, s)
	assert.True(t, errors.IsNotFound(err))

	mr := storemock.NewNullService()
	ctx = store.WithContextProvider(context.Background(), mr)
	assert.NotNil(t, ctx)
	haveMr, s, err = store.FromContextProvider(ctx)
	assert.True(t, errors.IsNotFound(err))
	assert.Nil(t, haveMr)
	assert.Nil(t, s)

}

func TestContextReaderSuccess(t *testing.T) {
	t.Parallel()
	ctx := storemock.WithContextMustService(scope.Option{},
		func(ms *storemock.Storage) {
			ms.MockStore = func() (*store.Store, error) {
				return store.NewStore(
					&store.TableStore{StoreID: 6, Code: dbr.NewNullString("de"), WebsiteID: 1, GroupID: 1, Name: "Germany", SortOrder: 10, IsActive: true},
					&store.TableWebsite{WebsiteID: 1, Code: dbr.NewNullString("euro"), Name: dbr.NewNullString("Europe"), SortOrder: 0, DefaultGroupID: 1, IsDefault: dbr.NewNullBool(true)},
					&store.TableGroup{GroupID: 1, WebsiteID: 1, Name: "DACH Group", RootCategoryID: 2, DefaultStoreID: 6},
				)
			}
		},
	)

	haveMr, s, err := store.FromContextProvider(ctx)
	assert.NoError(t, err)
	assert.Exactly(t, int64(6), s.StoreID())

	s2, err2 := haveMr.Store()
	assert.NoError(t, err2)
	assert.Exactly(t, int64(6), s2.StoreID())

}

func TestWithContextMustService(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r != nil {
			assert.EqualError(t, r.(error), "runtime error: invalid memory address or nil pointer dereference")
		}
	}()
	store.WithContextMustService(scope.Option{}, nil)
}
