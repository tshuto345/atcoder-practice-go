package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	ab := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ab[i][0], &ab[i][1])
	}

	ans := solve(n, k, ab)
	fmt.Println(ans)
}

func solve(n, k int, ab [][2]int) int {
	sort.Slice(ab, func(i, j int) bool {
		return ab[i][0] < ab[j][0]
	})

	at := k

	for i := 0; i < n; i++ {
		if ab[i][0] <= at {
			at += ab[i][1]
		} else {
			break
		}
	}

	return at
}
