package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)

	vp := make([][2]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &vp[i][0], &vp[i][1])
	}

	ans := solve(n, x, vp)
	fmt.Println(ans)
}

func solve(n, x int, vp [][2]int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += vp[i][0] * vp[i][1]
		if x*100 < total {
			return i + 1
		}
	}

	return -1
}
