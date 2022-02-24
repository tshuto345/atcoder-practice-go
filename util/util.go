package util

import (
	"math"
	"strconv"
	"strings"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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
