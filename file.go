// Copyright 2019 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utils

import (
	"go/format"
	"io/ioutil"
	"os"
)

// FileExists 判断文件或是文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// DumpGoFile 输出 Go 源代码到 path
//
// 会对源代码作格式化。
func DumpGoFile(path, content string) error {
	return DumpGoSource(path, []byte(content))
}

// DumpGoSource 输出 Go 源码到 path
func DumpGoSource(path string, content []byte) error {
	src, err := format.Source(content)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, src, os.ModePerm)
}
