// Copyright 2019 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utils

import (
	"io/ioutil"
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
