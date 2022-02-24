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

	xy := make([][2]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &xy[i][0], &xy[i][1])
	}

	ans := solve(n, xy)
	fmt.Println(ans)
}

func solve(n int, xy [][2]int) string {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if onLine(xy[i], xy[j], xy[k]) {
					return "Yes"
				}
			}
		}
	}

	return "No"
}

func onLine(xy1, xy2, xy3 [2]int) bool {
	vs := [2][2]int{
		{xy2[0] - xy1[0], xy2[1] - xy1[1]},
		{xy3[0] - xy1[0], xy3[1] - xy1[1]},
	}

	return vs[0][0]*vs[1][1]-vs[0][1]*vs[1][0] == 0
}
