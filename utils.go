// SPDX-License-Identifier: MIT

// Package utils 一些常用功能的集合
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/issue9/localeutil"
	"github.com/issue9/sliceutil"
	"github.com/issue9/source"

	"golang.org/x/text/language"
)

// GetSystemLanguageTag 返回当前系统的本地化信息
//
// *nix 系统会使用 LANG 环境变量中的值，windows 在 LANG
// 环境变量不存在的情况下，调用 GetUserDefaultLocaleName 函数获取。
//
// Deprecated: 请使用 localeutil.SystemLanguageTag 代替
func GetSystemLanguageTag() (language.Tag, error) {
	return localeutil.SystemLanguageTag()
}

// MD5 将一段字符串转换成 md5 编码
func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// TraceStack 返回调用者的堆栈信息
//
// Deprecated: 请使用 source.TraceStack 代替
func TraceStack(level int, msg ...interface{}) (string, error) {
	return source.TraceStack(level+1, msg...)
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
	}

	if len(path) > index {
		ret = append(ret, path[index:])
	}

	return ret
}

// HasDuplication 检测数组中是否包含重复的值
//
// slice 需要检测的数组或是切片，其它类型会 panic；
// eq 对比数组中两个值是否相等，相等需要返回 true；
// 返回值表示存在相等值时，第二个值在数组中的下标值；
func HasDuplication(slice interface{}, eq func(i, j int) bool) int {
	return sliceutil.Dup(slice, eq)
}
