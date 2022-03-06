package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	ttt := make([][]int, n)
	for i := 0; i < n; i++ {
		ttt[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &ttt[i][j])
		}
	}

	ans := solve(n, k, ttt)
	fmt.Println(ans)
}

func solve(n, k int, ttt [][]int) int {
	ans := 0

	perm := intPerm(n)
	for _, p := range perm {
		if p[0] != 0 {
			continue
		}

		l := 0
		for i := 0; i < len(p)-1; i++ {
			l += ttt[p[i]][p[i+1]]
		}
		l += ttt[p[len(p)-1]][0]

		if l == k {
			ans++
		}
	}

	return ans
}

func remove(ar []int, i int) []int {
	tmp := make([]int, len(ar))
	copy(tmp, ar)

	return append(tmp[0:i], tmp[i+1:]...)
}

func permutation(ar []int) [][]int {
	var result [][]int

	if len(ar) == 1 {
		return append(result, ar)
	}

	for i, a := range ar {
		for _, b := range permutation(remove(ar, i)) {
			result = append(result, append([]int{a}, b...))
		}
	}

	return result
}

func intRange(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}

	return s
}

func intPerm(n int) [][]int {
	return permutation(intRange(n))
}
