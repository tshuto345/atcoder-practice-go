package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	ans := solve(n)
	fmt.Println(ans)
}

func solve(n int) int {
	digits := getDigits(n)
	digitsMod3 := mapSlice(digits, func(n int) int { return n % 3 })
	count := count(digitsMod3)
	totalMod3 := sum(digitsMod3) % 3

	switch totalMod3 {
	case 0:
		return 0
	case 1:
		switch {
		case count[1] > 0 && len(digits) > 1:
			return 1
		case len(digits) > 2:
			return 2
		default:
			return -1
		}
	case 2:
		switch {
		case count[2] > 0 && len(digits) > 1:
			return 1
		case len(digits) > 2:
			return 2
		default:
			return -1
		}
	}

	panic("unreachable")
}

func getDigits(n int) []int {
	digits := []int{}
	for n > 0 {
		r := n % 10
		n = n / 10
		digits = append([]int{r}, digits...)
	}

	return digits
}

func mapSlice(nn []int, fn func(n int) int) []int {
	_nn := make([]int, len(nn))
	copy(_nn, nn)
	for i := 0; i < len(nn); i++ {
		_nn[i] = fn(_nn[i])
	}

	return _nn
}

func count(nn []int) map[int]int {
	c := make(map[int]int)
	for _, n := range nn {
		c[n]++
	}

	return c
}

func sum(nn []int) int {
	t := 0
	for _, n := range nn {
		t += n
	}

	return t
}
