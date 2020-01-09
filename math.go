// SPDX-License-Identifier: MIT

package utils

import "math"

// Round 对一个浮点数进行四舍五入取整
func Round(v float64) int64 {
	return int64(math.Round(v))
}
