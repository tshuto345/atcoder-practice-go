package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	ans := solve(n, k)
	fmt.Println(ans)
}

func solve(n, k int) string {

	m := n
	for i := 0; i < k; i++ {
		m = f(m)
	}
	return strconv.Itoa(m)
}

func reverse(ss []string) {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
}

func g1(n int) int {
	ss := strings.Split(fmt.Sprintf("%d", n), "")
	sort.Strings(ss)
	i, _ := strconv.Atoi(strings.Join(ss, ""))
	return i
}

func g2(n int) int {
	ss := strings.Split(fmt.Sprintf("%d", n), "")
	sort.Strings(ss)
	reverse(ss)
	i, _ := strconv.Atoi(strings.Join(ss, ""))
	return i
}

func f(n int) int {
	return g2(n) - g1(n)
}
