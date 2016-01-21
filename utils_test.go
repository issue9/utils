// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utils

import (
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
