package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	ans := solve(n)
	fmt.Println(ans)
}

func solve(n int) int {
	ans := 0

	for i := 1; i <= n; i++ {
		i8, _ := base(i, 8)
		if !leftContainsRight(i, 7) && !leftContainsRight(i8, 7) {
			ans++
		}
	}

	return ans
}

func base(n, b int) (int, error) {
	if n == 0 {
		return 0, nil
	}

	x := ""
	r := 0

	for n > 0 {
		r = n % b
		n = (n - r) / b
		x = strconv.Itoa(r) + x
	}

	return strconv.Atoi(x)
}

func leftContainsRight(l, r int) bool {
	return strings.Contains(strconv.Itoa(l), strconv.Itoa(r))
}
