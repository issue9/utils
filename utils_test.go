// SPDX-License-Identifier: MIT

package utils

import (
	"runtime"
	"strings"
	"testing"

	"github.com/issue9/assert"
)

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
