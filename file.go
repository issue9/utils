// SPDX-License-Identifier: MIT

package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/issue9/source"
)

// FileExists 判断文件或是文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// DumpGoFile 输出 Go 源代码到 path
//
// 会对源代码作格式化。
//
// Deprecated: 请使用 DumpGoSource 代替
func DumpGoFile(path, content string) error {
	return DumpGoSource(path, []byte(content))
}

// DumpGoSource 输出 Go 源码到 path
//
// 会对源代码作格式化。
//
// Deprecated: 请使用 source.DumpGoSource 代替
func DumpGoSource(path string, content []byte) error {
	return source.DumpGoSource(path, content)
}

// CurrentPath 获取`调用者`所在目录的路径
//
// Deprecated: 请使用 source.CurrentPath 代替
func CurrentPath(path string) string {
	_, fi, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(fi), path)
}

// CurrentDir 获取`调用者`所在的目录
//
// Deprecated: 请使用 source.CurrentDir 代替
func CurrentDir() string {
	_, fi, _, _ := runtime.Caller(1)
	return filepath.Dir(fi)
}

// CurrentFile 获取`调用者`所在的文件
//
// Deprecated: 请使用 source.CurrentFile 代替
func CurrentFile() string {
	_, fi, _, _ := runtime.Caller(1)
	return fi
}

// CurrentLine 获取`调用者`所在的行
//
// Deprecated: 请使用 source.CurrentLine 代替
func CurrentLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

// CurrentFunction 获取`调用者`所在的函数名
//
// Deprecated: 请使用 source.CurrentFunction 代替
func CurrentFunction() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()

	index := strings.LastIndexByte(name, '.')
	if index > 0 {
		name = name[index+1:]
	}

	return name
}
