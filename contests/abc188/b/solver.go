package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	aa := make([]int, n)
	bb := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&aa[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&bb[i])
	}

	ans := solve(n, aa, bb)
	fmt.Println(ans)
}

func solve(n int, aa, bb []int) string {
	var prod int
	for i := 0; i < n; i++ {
		prod += aa[i] * bb[i]
	}
	if prod == 0 {
		return "Yes"
	} else {
		return "No"
	}
}
