// Copyright 2019 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utils

import (
	"go/format"
	"os"
)

// FileExists 判断文件或是文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// DumpFile 输出 content 到 path
func DumpFile(path string, content []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	_, err = file.Write(content)
	return err
}

// DumpGoFile 输出 Go 源代码到 path
func DumpGoFile(path, content string) error {
	src, err := format.Source([]byte(content))
	if err != nil {
		return err
	}

	return DumpFile(path, src)
}
