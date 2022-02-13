package main

import (
	"fmt"
)

func main() {
	var n, s, d int
	fmt.Scan(&n)
	fmt.Scan(&s)
	fmt.Scan(&d)

	xy := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&xy[i][0])
		fmt.Scan(&xy[i][1])
	}

	ans := solve(n, s, d, xy)
	fmt.Println(ans)
}

func solve(n, s, d int, xy [][2]int) string {
	for i := 0; i < n; i++ {
		if xy[i][0] < s && xy[i][1] > d {
			return "Yes"
		}
	}
	return "No"
}
