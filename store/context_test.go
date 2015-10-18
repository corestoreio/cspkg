// Copyright 2015, Cyrill @ Schumacher.fm and the CoreStore contributors
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

	"github.com/corestoreio/csfw/store"
	storemock "github.com/corestoreio/csfw/store/mock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestContextManagerReader(t *testing.T) {
	mr := storemock.NewNullManager()
	ctx := store.NewContextManagerReader(context.Background(), mr)
	haveMr, ok := store.FromContextManagerReader(ctx)
	assert.True(t, ok)
	assert.Exactly(t, mr, haveMr)

	ctx = store.NewContextManagerReader(context.Background(), nil)
	store.FromContextManagerReader(ctx)
}
