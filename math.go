// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utils

import "math"

// Round 对一个浮点数进行四舍五入取整
func Round(v float64) int64 {
	return int64(math.Round(v))
}
