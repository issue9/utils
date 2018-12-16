// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package utils 一些常用的函数集合。
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// MD5 将一段字符串转换成 md5 编码
func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// FileExists 判断文件或是文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// TraceStack 返回调用者的堆栈信息
func TraceStack(level int, msg ...interface{}) (string, error) {
	var w strings.Builder
	var err error

	if len(msg) > 0 {
		if _, err = fmt.Fprintln(&w, msg...); err != nil {
			return "", err
		}
	}

	ws := func(str string) {
		if err == nil {
			_, err = w.WriteString(str)
		}
	}

	for i := level; true; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		ws(file)
		ws(":")
		ws(strconv.Itoa(line))
		ws("\n")
	}

	if err != nil {
		return "", err
	}

	return w.String(), nil
}

// SplitPath 将路径按分隔符分隔成字符串数组。比如：
//  /a/b/c  ==>  []string{"a", "b", "c"}
func SplitPath(path string) []string {
	vol := filepath.VolumeName(path)
	ret := make([]string, 0, 10)

	index := 0
	if len(vol) > 0 {
		ret = append(ret, vol)
		path = path[len(vol)+1:]
	}
	for i := 0; i < len(path); i++ {
		if os.IsPathSeparator(path[i]) {
			if i > index {
				ret = append(ret, path[index:i])
			}
			index = i + 1 // 过滤掉此符号
		}
	} // end for

	if len(path) > index {
		ret = append(ret, path[index:])
	}

	return ret
}
