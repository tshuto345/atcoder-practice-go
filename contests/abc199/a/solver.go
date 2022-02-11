package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	ans := solve(a, b, c)
	fmt.Println(ans)
}

func solve(a, b, c int) string {
	return boolStr(powInt(a, 2)+powInt(b, 2) < powInt(c, 2), "Yes", "No")
}

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
