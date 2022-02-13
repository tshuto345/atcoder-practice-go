package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	fmt.Scan(&x)

	aa := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&aa[i])
	}

	ans := solve(n, x, aa)
	fmt.Println(ans)
}

func solve(n, x int, aa []int) string {
	f := []string{}
	for i := 0; i < n; i++ {
		if x != aa[i] {
			f = append(f, strconv.Itoa(aa[i]))
		}
	}
	return strings.Join(f, " ")
}
