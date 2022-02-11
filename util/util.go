package util

import (
	"math"
)

func boolStr(b bool, ts, fs string) string {
	if b {
		return ts
	} else {
		return fs
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
