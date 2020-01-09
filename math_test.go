// SPDX-License-Identifier: MIT

package utils

import (
	"testing"

	"github.com/issue9/assert"
)

func TestRound(t *testing.T) {
	a := assert.New(t)

	a.Equal(Round(1.1), 1).
		Equal(Round(1.5), 2).
		Equal(Round(0.9), 1).
		Equal(Round(0.5), 1).
		Equal(Round(0.4), 0)
}
