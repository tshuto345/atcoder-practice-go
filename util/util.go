package util

import (
	"math"
	"strconv"
	"strings"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func absInt(x int) int {
	return int(math.Abs(float64(x)))
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func mapSlice(nn []int, fn func(n int) int) []int {
	_nn := make([]int, len(nn))
	copy(_nn, nn)
	for i := 0; i < len(nn); i++ {
		_nn[i] = fn(_nn[i])
	}

	return _nn
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
