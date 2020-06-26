// SPDX-License-Identifier: MIT

package utils

import (
	"reflect"
	"testing"

	"github.com/issue9/assert"
)

type s1 struct {
	ID    int64
	Value string
	hide  int // 不可导出
	Slice []int
}

type s2 struct {
	s1
	S2 string
	S3 *s1
	S4 s1
}

func TestMerge(t *testing.T) {
	a := assert.New(t)

	v1 := &s1{ID: 1, Value: ""}
	v2 := &s1{ID: 2, Value: "2"}
	v3 := &s1{ID: 3, Value: ""}

	a.NotError(Merge(true, v1, v2, v3))
	a.Equal(v1.ID, 3).Equal(v1.Value, "2")

	// 以数组扩展的方式传递参数
	v1 = &s1{ID: 1, Value: ""}
	v2 = &s1{ID: 2, Value: "2"}
	v3 = &s1{ID: 3, Value: ""}
	elems := []interface{}{v1, v2, v3}
	a.NotError(Merge(true, elems...))
	a.Equal(elems[0].(*s1).ID, 3).Equal(elems[0].(*s1).Value, "2")

	// 参数太少
	a.Panic(func() {
		Merge(true, v1)
	})

	// 类型不同
	a.Panic(func() {
		Merge(true, v1, 5)
	})

	v4 := &s2{}
	v5 := &s2{S4: s1{ID: 5}}
	a.NotError(Merge(true, v4, v5))
	a.Empty(v4.Value).Equal(v4.S4.ID, 5)
}

func TestMergeBase(t *testing.T) {
	a := assert.New(t)

	v1 := &s1{ID: 1, Value: "", hide: 1}
	v2 := &s1{ID: 2, Value: "2", hide: 2}
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.ID, 2).Equal(v1.Value, "2").Equal(v1.hide, 1)

	// 零值，不会覆盖
	v1.ID = 1
	v2.ID = 0
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.ID, 1).Equal(v1.Value, "2").Equal(v1.hide, 1)
}

func TestMergeAnonymous(t *testing.T) {
	a := assert.New(t)

	v1 := &s2{s1: s1{ID: 1, Value: "", hide: 1}, S2: "1"}
	v2 := &s2{s1: s1{ID: 2, Value: "2", hide: 2}, S2: "2"}
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.ID, 2).Equal(v1.Value, "2").Equal(v1.hide, 1)
}

func TestMergeNest(t *testing.T) {
	a := assert.New(t)

	v1 := &s2{S2: "1", S3: &s1{}}
	v2 := &s2{S2: "2", S3: &s1{ID: 2, hide: 2}}

	// deep 为 true 时，会依次赋值子元素，hide 不可导出，所以被忽略
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.ID, 0).Equal(v1.S3.ID, 2).Equal(v1.S3.hide, 0)

	// deep为 false 时，会将整个 v2.S3 赋给给 v1.S3
	a.NotError(merge(false, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.ID, 0).Equal(v1.S3.ID, 2).Equal(v1.S3.hide, 2)

	// 自动初始化 v1.S3
	v1.S3 = nil
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.ID, 0).Equal(v1.S3.ID, 2).Equal(v1.S3.hide, 0)

	// 不会初始化 v1.S3
	v1.S3 = nil
	v2.S3 = nil
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.ID, 0).Nil(v1.S3)
}

func TestMergeSlice(t *testing.T) {
	a := assert.New(t)

	v1 := &s1{Slice: []int{1, 1}}
	v2 := &s1{Slice: []int{2, 2}}
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.Slice, []int{2, 2})

	// 空值
	v1.Slice = []int{1, 1}
	v2.Slice = nil
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.Slice, []int{1, 1})

	// 长度为 0
	v2.Slice = []int{}
	a.NotError(merge(true, reflect.ValueOf(v1).Elem(), reflect.ValueOf(v2).Elem()))
	a.Equal(v1.Slice, []int{1, 1})
}
