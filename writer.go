// SPDX-License-Identifier: MIT

package utils

import (
	"bytes"
	"fmt"
	"strings"
)

// Buffer 提供缓存错误的 bytes.Buffer
type Buffer struct {
	bytes.Buffer
	Err error
}

// StringBuilder 提供缓存错误的 strings.Buffer
type StringBuilder struct {
	strings.Builder
	Err error
}

// NewBuffer 声明 Writer 实例
func NewBuffer() *Buffer {
	return &Buffer{}
}

// WString 写入字符串
func (w *Buffer) WString(str string) *Buffer {
	if w.Err == nil {
		_, w.Err = w.WriteString(str)
	}
	return w
}

// WByte 写入单个字节内容
func (w *Buffer) WByte(b byte) *Buffer {
	return w.WBytes([]byte{b})
}

// WBytes 写入字节内容
func (w *Buffer) WBytes(data []byte) *Buffer {
	if w.Err == nil {
		_, w.Err = w.Write(data)
	}
	return w
}

// WRune 写入单个字节内容
func (w *Buffer) WRune(r rune) *Buffer {
	if w.Err == nil {
		_, w.Err = w.WriteRune(r)
	}
	return w
}

// WRunes 写入字节内容
func (w *Buffer) WRunes(rs []rune) *Buffer {
	for _, r := range rs {
		w.WRune(r)
	}
	return w
}

// Print 相当于 fmt.Fprint(builder, v...)
func (w *Buffer) Print(v ...interface{}) *Buffer {
	if w.Err == nil {
		_, w.Err = fmt.Fprint(w, v...)
	}
	return w
}

// Println 相当于 fmt.Fprintln(builder, v...)
func (w *Buffer) Println(v ...interface{}) *Buffer {
	if w.Err == nil {
		_, w.Err = fmt.Fprintln(w, v...)
	}
	return w
}

// Printf 相当于 fmt.Fprintf(builder, format, v...)
func (w *Buffer) Printf(format string, v ...interface{}) *Buffer {
	if w.Err == nil {
		_, w.Err = fmt.Fprintf(w, format, v...)
	}
	return w
}

// NewStringBuilder 声明 Writer 实例
func NewStringBuilder() *StringBuilder {
	return &StringBuilder{}
}

// WString 写入字符串
func (builder *StringBuilder) WString(str string) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = builder.WriteString(str)
	}
	return builder
}

// WByte 写入单个字节内容
func (builder *StringBuilder) WByte(b byte) *StringBuilder {
	return builder.WBytes([]byte{b})
}

// WBytes 写入字节内容
func (builder *StringBuilder) WBytes(data []byte) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = builder.Write(data)
	}
	return builder
}

// WRune 写入单个字节内容
func (builder *StringBuilder) WRune(r rune) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = builder.WriteRune(r)
	}
	return builder
}

// WRunes 写入字节内容
func (builder *StringBuilder) WRunes(rs []rune) *StringBuilder {
	for _, r := range rs {
		builder.WRune(r)
	}
	return builder
}

// Print 相当于 fmt.Fprint(builder, v...)
func (builder *StringBuilder) Print(v ...interface{}) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = fmt.Fprint(builder, v...)
	}
	return builder
}

// Println 相当于 fmt.Fprintln(builder, v...)
func (builder *StringBuilder) Println(v ...interface{}) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = fmt.Fprintln(builder, v...)
	}
	return builder
}

// Printf 相当于 fmt.Fprintf(builder, format, v...)
func (builder *StringBuilder) Printf(format string, v ...interface{}) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = fmt.Fprintf(builder, format, v...)
	}
	return builder
}
