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

package util_test

import (
	"testing"

	"github.com/corestoreio/csfw/util"
	"github.com/stretchr/testify/assert"
)

func TestStringSliceReduceContains(t *testing.T) {
	tests := []struct {
		haveSL util.StringSlice
		haveIN []string
		want   []string
	}{
		{
			util.StringSlice{
				"IFNULL(`scope_table`.`is_visible`, `additional_table`.`is_visible`) AS `is_visible`",
				"IFNULL(`scope_table`.`is_required`, `main_table`.`is_required`) AS `is_required`",
				"IFNULL(`scope_table`.`default_value`, `main_table`.`default_value`) AS `default_value`",
				"IFNULL(`scope_table`.`multiline_count`, `additional_table`.`multiline_count`) AS `multiline_count`",
			},
			[]string{"is_required", "default_value"},
			[]string{
				"IFNULL(`scope_table`.`is_visible`, `additional_table`.`is_visible`) AS `is_visible`",
				"IFNULL(`scope_table`.`multiline_count`, `additional_table`.`multiline_count`) AS `multiline_count`",
			},
		},
		{
			util.StringSlice{"GoLang", "RustLang", "PHP Script", "JScript"},
			[]string{"Script"},
			[]string{"GoLang", "RustLang"},
		},
	}

	for _, test := range tests {
		test.haveSL.ReduceContains(test.haveIN...)
		assert.Equal(t, test.want, test.haveSL.ToString())
		assert.Len(t, test.haveSL, len(test.want))
	}
}

var benchStringSliceReduceContains util.StringSlice
var benchStringSliceReduceContainsData = []string{"is_required", "default_value"}

// BenchmarkStringSliceReduceContains	 1000000	      1841 ns/op	      96 B/op	       2 allocs/op <- Go 1.4.2
// BenchmarkStringSliceReduceContains-4	 1000000	      1250 ns/op	      64 B/op	       1 allocs/op <- Go 1.5
func BenchmarkStringSliceReduceContains(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := util.StringSlice{
			"IFNULL(`scope_table`.`is_visible`, `additional_table`.`is_visible`) AS `is_visible`",
			"IFNULL(`scope_table`.`is_required`, `main_table`.`is_required`) AS `is_required`",
			"IFNULL(`scope_table`.`default_value`, `main_table`.`default_value`) AS `default_value`",
			"IFNULL(`scope_table`.`multiline_count`, `additional_table`.`multiline_count`) AS `multiline_count`",
		}
		l.ReduceContains(benchStringSliceReduceContainsData...)
		benchStringSliceReduceContains = l
	}
}

func TestStringSliceUpdate(t *testing.T) {
	tests := []struct {
		haveSL util.StringSlice
		haveD  string
		haveI  int
		err    error
		want   []string
	}{
		{
			haveSL: util.StringSlice{
				"IFNULL(`scope_table`.`is_visible`, `additional_table`.`is_visible`) AS `is_visible`",
				"IFNULL(`scope_table`.`is_required`, `main_table`.`is_required`) AS `is_required`",
				"IFNULL(`scope_table`.`default_value`, `main_table`.`default_value`) AS `default_value`",
				"IFNULL(`scope_table`.`multiline_count`, `additional_table`.`multiline_count`) AS `multiline_count`",
			},
			haveD: "default_value",
			haveI: 1,
			err:   nil,
			want: []string{
				"IFNULL(`scope_table`.`is_visible`, `additional_table`.`is_visible`) AS `is_visible`",
				"default_value",
				"IFNULL(`scope_table`.`default_value`, `main_table`.`default_value`) AS `default_value`",
				"IFNULL(`scope_table`.`multiline_count`, `additional_table`.`multiline_count`) AS `multiline_count`",
			},
		},
		{
			haveSL: util.StringSlice{
				"IFNULL(`scope_table`.`is_visible`, `additional_table`.`is_visible`) AS `is_visible`",
				"IFNULL(`scope_table`.`is_required`, `main_table`.`is_required`) AS `is_required`",
				"IFNULL(`scope_table`.`default_value`, `main_table`.`default_value`) AS `default_value`",
				"IFNULL(`scope_table`.`multiline_count`, `additional_table`.`multiline_count`) AS `multiline_count`",
			},
			haveD: "default_value",
			haveI: 6,
			err:   util.ErrOutOfRange,
			want: []string{
				"IFNULL(`scope_table`.`is_visible`, `additional_table`.`is_visible`) AS `is_visible`",
				"IFNULL(`scope_table`.`is_required`, `main_table`.`is_required`) AS `is_required`",
				"IFNULL(`scope_table`.`default_value`, `main_table`.`default_value`) AS `default_value`",
				"IFNULL(`scope_table`.`multiline_count`, `additional_table`.`multiline_count`) AS `multiline_count`",
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.err, test.haveSL.Update(test.haveI, test.haveD))
		assert.Equal(t, test.want, test.haveSL.ToString())
	}
}

func TestStringSlice(t *testing.T) {
	l := util.StringSlice{"Maybe", "GoLang", "should", "have", "generics", "but", "who", "needs", "them", "?", ";-)"}
	assert.Len(t, l, l.Len())
	assert.Equal(t, 1, l.Index("GoLang"))
	assert.Equal(t, -1, l.Index("Golang"))
	assert.True(t, l.Contains("GoLang"))
	assert.False(t, l.Contains("Golang"))

	l2 := util.StringSlice{"Maybe", "GoLang"}
	l2.Map(func(s string) string {
		return s + "2"
	})
	assert.Equal(t, []string{"Maybe2", "GoLang2"}, l2.ToString())
	l2.Append("will", "be")
	assert.Equal(t, []string{"Maybe2", "GoLang2", "will", "be"}, l2.ToString())

}

func TestStringSliceDelete(t *testing.T) {
	l := util.StringSlice{"Maybe", "GoLang", "should"}
	assert.NoError(t, l.Delete(1))
	assert.Equal(t, []string{"Maybe", "should"}, l.ToString())
	assert.NoError(t, l.Delete(1))
	assert.Equal(t, []string{"Maybe"}, l.ToString())
	assert.EqualError(t, l.Delete(1), util.ErrOutOfRange.Error())
}

func TestStringSliceReduce(t *testing.T) {
	l := util.StringSlice{"Maybe", "GoLang", "should"}
	assert.EqualValues(t, []string{"GoLang"}, l.Reduce(func(s string) bool {
		return s == "GoLang"
	}).ToString())
	assert.Len(t, l, 1)
}

func TestStringSliceFilter(t *testing.T) {
	l := util.StringSlice{"All", "Go", "Code", "is"}
	rl := l.Filter(func(s string) bool {
		return s == "Go"
	}).ToString()
	assert.EqualValues(t, []string{"Go"}, rl)
	assert.Len(t, l, 4)

	l.Append("incredible", "easy", ",", "sometimes")
	assert.Len(t, l, 8)
	assert.EqualValues(t, []string{"Go"}, rl)
}

func TestStringSliceUnique(t *testing.T) {
	l := util.StringSlice{"Maybe", "GoLang", "GoLang", "GoLang", "or", "or", "RostLang", "RostLang"}
	assert.Equal(t, []string{"Maybe", "GoLang", "or", "RostLang"}, l.Unique().ToString())
}

var benchStringSliceUnique util.StringSlice

// BenchmarkStringSliceUnique	 2000000	       612 ns/op	     160 B/op	       2 allocs/op <- Go 1.4.2
// BenchmarkStringSliceUnique  	10000000	       179 ns/op	     128 B/op	       1 allocs/op <- Go 1.5
func BenchmarkStringSliceUnique(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := util.StringSlice{"Maybe", "GoLang", "GoLang", "GoLang", "or", "or", "RostLang", "RostLang"}
		l.Unique()
		benchStringSliceUnique = l
	}
}

func TestStringSliceSplit(t *testing.T) {
	l := util.StringSlice{"a", "b"}
	assert.Equal(t, []string{"a", "b", "c", "d"}, l.Split("c,d", ",").ToString())
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", ""}, l.Split("e,f,", ",").ToString())
}

func TestStringSliceJoin(t *testing.T) {
	l := util.StringSlice{"a", "b"}
	assert.Equal(t, "a,b", l.Join(","))
}

func TestStringSliceSort(t *testing.T) {
	l := util.StringSlice{"c", "a", "z", "b"}
	assert.Equal(t, "a,b,c,z", l.Sort().Join(","))
}

func TestStringSliceAny(t *testing.T) {
	l := util.StringSlice{"c", "a", "z", "b"}
	assert.True(t, l.Any(func(s string) bool {
		return s == "z"
	}))
	assert.False(t, l.Any(func(s string) bool {
		return s == "zx"
	}))
}

func TestStringSliceAll(t *testing.T) {
	l := util.StringSlice{"c", "a", "z", "b"}
	assert.True(t, l.All(func(s string) bool {
		return len(s) == 1
	}))
	l.Append("xx")
	assert.False(t, l.All(func(s string) bool {
		return len(s) == 1
	}))
}

func TestStringSliceSplitStringer8(t *testing.T) {
	tests := []struct {
		haveName  string
		haveIndex []uint8
		want      util.StringSlice
	}{
		{
			"ScopeAbsentScopeDefaultScopeWebsiteScopeGroupScopeStore",
			[]uint8{0, 11, 23, 35, 45, 55},
			util.StringSlice{"ScopeAbsent", "ScopeDefault", "ScopeWebsite", "ScopeGroup", "ScopeStore"},
		},
		{
			"TypeCustomTypeHiddenTypeObscureTypeMultiselectTypeSelectTypeTextTypeTime",
			[]uint8{10, 20, 31, 46, 56, 64, 72},
			util.StringSlice{"TypeHidden", "TypeObscure", "TypeMultiselect", "TypeSelect", "TypeText", "TypeTime"},
		},
	}
	for _, test := range tests {
		var a util.StringSlice
		have := a.SplitStringer8(test.haveName, test.haveIndex...)
		assert.Exactly(t, test.want, have)
	}
}

var benchStringSliceSplitStringer8 util.StringSlice

// BenchmarkStringSliceSplitStringer8	 1000000	      1041 ns/op	     240 B/op	       4 allocs/op <- Go 1.4.2
// BenchmarkStringSliceSplitStringer8-4	 2000000	       673 ns/op	     240 B/op	       4 allocs/op <- Go 1.5
func BenchmarkStringSliceSplitStringer8(b *testing.B) {
	const _ScopeGroup_name = "ScopeAbsentScopeDefaultScopeWebsiteScopeGroupScopeStore"
	var _ScopeGroup_index = [...]uint8{0, 11, 23, 35, 45, 55}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchStringSliceSplitStringer8.SplitStringer8(_ScopeGroup_name, _ScopeGroup_index[:]...)
		benchStringSliceSplitStringer8 = nil
	}
}

func TestStringSliceContainsReverse(t *testing.T) {
	tests := []struct {
		have string
		in   util.StringSlice
		want bool
	}{
		{"I live in the black forest", util.StringSlice{"black"}, true},
		{"I live in the black forest", util.StringSlice{"blagg", "forest"}, true},
		{"I live in the black forest", util.StringSlice{"blagg", "wald"}, false},
		{"We don't have any Internet connection", nil, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, test.in.ContainsReverse(test.have), "Test: %#v", test)
	}
}

func TestStringSliceStartsWithReverse(t *testing.T) {
	tests := []struct {
		have string
		in   util.StringSlice
		want bool
	}{
		{"grand_total", util.StringSlice{"grand_"}, true},
		{"base_discount_amount", util.StringSlice{"amount"}, false},
		{"base_grand_total", util.StringSlice{"grand_", "base_"}, true},
		{"base_grand_total", util.StringSlice{"xgrand_", "zbase_"}, false},
		{"base_grand_total", nil, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, test.in.StartsWithReverse(test.have), "Test: %#v", test)
	}
}
