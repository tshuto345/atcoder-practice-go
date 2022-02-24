package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)

	ans := solve(a, b)
	fmt.Println(ans)
}

func solve(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
