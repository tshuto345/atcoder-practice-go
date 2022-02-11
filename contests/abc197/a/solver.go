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
	rs = append(rs[1:], rs[0])
	return string(rs)
}
