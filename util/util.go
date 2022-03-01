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
