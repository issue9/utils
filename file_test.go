// SPDX-License-Identifier: MIT

package utils

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/issue9/assert"
)

func TestFileExists(t *testing.T) {
	a := assert.New(t)

	// 测试文件
	a.True(FileExists("utils.go"))
	a.False(FileExists("unknown.go"))

	// 测试文件夹
	a.True(FileExists("./"))
	a.False(FileExists("./unknown_dir/"))
}

func TestDumpGoFile(t *testing.T) {
	a := assert.New(t)

	a.NotError(DumpGoFile("./testdata/go.go", "var x=1"))
	content, err := ioutil.ReadFile("./testdata/go.go")
	a.NotError(err)
	a.Equal(string(content), "var x = 1")
}

func TestCurrentPath(t *testing.T) {
	a := assert.New(t)

	dir, err := filepath.Abs("./file.go")
	a.NotError(err).NotEmpty(dir)

	d, err := filepath.Abs(CurrentPath("./file.go"))
	a.NotError(err).NotEmpty(d)

	a.Equal(d, dir)
}

func TestCurrentDir(t *testing.T) {
	a := assert.New(t)

	dir, err := filepath.Abs("./")
	a.NotError(err).NotEmpty(dir)

	a.Equal(CurrentDir(), dir)
}

func TestCurrentFile(t *testing.T) {
	a := assert.New(t)

	filename, err := filepath.Abs("./file_test.go")
	a.NotError(err).NotEmpty(filename)

	a.Equal(CurrentFile(), filepath.FromSlash(filename))
}

func TestCurrentFunction(t *testing.T) {
	a := assert.New(t)

	a.Equal(CurrentFunction(), "TestCurrentFunction")
}

func TestCurrentLine(t *testing.T) {
	a := assert.New(t)

	a.Equal(CurrentLine(), 73)
}
