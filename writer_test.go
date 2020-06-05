// SPDX-License-Identifier: MIT

package utils

import (
	"testing"

	"github.com/issue9/assert"
)

func TestBuffer(t *testing.T) {
	a := assert.New(t)

	buf := NewBuffer()
	buf.WString("123").WRunes([]rune("456")).WByte('7').
		Printf("%s", "89").Print("10").Println()

	a.Equal(buf.String(), "12345678910\n").
		NotError(buf.Err)
}

func TestStringBuilder(t *testing.T) {
	a := assert.New(t)

	buf := NewStringBuilder()
	buf.WString("123").WRunes([]rune("456")).WByte('7').
		Printf("%s", "89").Print("10").Println()

	a.Equal(buf.String(), "12345678910\n").
		NotError(buf.Err)
}
