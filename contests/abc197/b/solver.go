package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var h, w, x, y int
	fmt.Scan(&h)
	fmt.Scan(&w)
	fmt.Scan(&x)
	fmt.Scan(&y)

	ss := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&ss[i])
	}

	ans := solve(h, w, x, y, ss)
	fmt.Println(ans)
}

func solve(h, w, x, y int, ss []string) string {
	x, y = x-1, y-1

	sss := make([][]string, h)
	for i := 0; i < h; i++ {
		sss[i] = strings.Split(ss[i], "")
	}

	ans := 1
	for i := x - 1; i >= 0; i-- {
		if sss[i][y] == "#" {
			break
		}
		ans++
	}
	for i := x + 1; i < h; i++ {
		if sss[i][y] == "#" {
			break
		}
		ans++
	}
	for j := y - 1; j >= 0; j-- {
		if sss[x][j] == "#" {
			break
		}
		ans++
	}
	for j := y + 1; j < w; j++ {
		if sss[x][j] == "#" {
			break
		}
		ans++
	}

	return strconv.Itoa(ans)
}
