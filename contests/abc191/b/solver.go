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
	ansstr := make([]string, len(ans))
	for _, v := range ans {
		ansstr = append(ansstr, strconv.Itoa(v))
	}
	fmt.Println(strings.Join(ansstr, " "))
}

func solve(n, x int, aa []int) []int {
	f := []int{}
	for i := 0; i < n; i++ {
		if x != aa[i] {
			f = append(f, aa[i])
		}
	}
	return f
}
