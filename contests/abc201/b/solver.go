package main

import (
	"fmt"
	"sort"
)

type mountain struct {
	s string
	t int
}

func main() {
	var n int
	fmt.Scan(&n)

	st := make([]mountain, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&st[i].s)
		fmt.Scan(&st[i].t)
	}

	ans := solve(n, st)
	fmt.Println(ans)
}

func solve(n int, st []mountain) string {
	sort.Slice(st, func(i int, j int) bool {
		return st[i].t > st[j].t
	})
	return st[1].s
}
