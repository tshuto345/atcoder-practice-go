package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)

	ans := solve(n)
	for _, i := range ans {
		fmt.Println(i)
	}
}

func solve(n int) []int {
	left := []int{}
	right := []int{}
	rn := int(math.Sqrt(float64(n)))

	for i := 1; i <= rn; i++ {
		if n%i == 0 {
			a, b := i, n/i
			left = append(left, a)
			if a != b {
				right = append([]int{b}, right...)
			}
		}
	}

	ans := make([]int, 0, len(left)+len(right))
	ans = append(ans, left...)
	ans = append(ans, right...)

	return ans
}
