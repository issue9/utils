// SPDX-License-Identifier: MIT

package utils

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/issue9/assert"
)

func TestMD5(t *testing.T) {
	a := assert.New(t)
	a.Equal(32, len(MD5("123")))
}

func TestSplitPath(t *testing.T) {
	a := assert.New(t)

	a.Equal([]string{"a", "b"}, SplitPath("/a/b"))
	a.Equal([]string{"a", "b", "c", "d"}, SplitPath("/a/b/c/d"))
	a.Equal([]string{"a", "b", "c", "d"}, SplitPath("/a/b/c/d/"))

	if runtime.GOOS == "windows" {
		a.Equal([]string{"a", "b"}, SplitPath("/a/b"))
		a.Equal([]string{"a", "b", "c", "d"}, SplitPath("/a/b/c/d"))
		a.Equal([]string{"a", "b", "c", "d"}, SplitPath("/a/b/c/d/"))
		a.Equal([]string{"\\\\host\\a", "b"}, SplitPath("\\\\host\\a\\b"))

		a.Equal([]string{"a", "b"}, SplitPath("\\a\\b"))
		a.Equal([]string{"a", "b", "c", "d"}, SplitPath("\\a\\b\\c\\d"))
		a.Equal([]string{"a", "b", "c", "d"}, SplitPath("\\a/b\\c\\d\\"))
		a.Equal([]string{"c:", "a", "b"}, SplitPath("c:\\a\\b"))
	}
}

func TestTraceStack(t *testing.T) {
	a := assert.New(t)

	str, err := TraceStack(1, "message", 12)
	a.NotError(err)
	a.True(strings.HasPrefix(str, "message 12"))
	a.True(strings.Contains(str, "utils_test.go")) // 肯定包含当前文件名
}

func TestHasDuplication(t *testing.T) {
	a := assert.New(t)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	a.Equal(6, HasDuplication(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	stringSlice := []string{"a", "b", "0", "a"}
	a.Equal(3, HasDuplication(stringSlice, func(i, j int) bool {
		return stringSlice[i] == stringSlice[j]
	}))

	type obj struct {
		ID   int
		Name string
		Age  int
	}
	objSlice := []*obj{
		{ID: 1, Name: "5", Age: 1},
		{ID: 2, Name: "4", Age: 2},
		{ID: 3, Name: "3", Age: 3},
		{ID: 4, Name: "2", Age: 4},
		{ID: 5, Name: "5", Age: 5},
		{ID: 1, Name: "1", Age: 6},
	}
	a.Equal(5, HasDuplication(objSlice, func(i, j int) bool {
		return objSlice[i].ID == objSlice[j].ID
	}))
	a.Equal(4, HasDuplication(objSlice, func(i, j int) bool {
		return objSlice[i].Name == objSlice[j].Name
	}))
	a.Equal(-1, HasDuplication(objSlice, func(i, j int) bool {
		return objSlice[i].Age == objSlice[j].Age
	}))

	a.Panic(func() {
		HasDuplication(5, func(i, j int) bool {
			return false
		})
	})
}

func ExampleHasDuplication() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(HasDuplication(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	// Output: 6
}
