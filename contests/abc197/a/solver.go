package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	ans := solve(s)
	fmt.Println(ans)
}

func solve(s string) string {
	rs := []rune(s)
	x, rs := rs[0], rs[1:]
	rs = append(rs, x)
	return string(rs)
}
