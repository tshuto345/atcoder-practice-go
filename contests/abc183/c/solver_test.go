package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n, k int
		ttt  [][]int
	}
	tests := []struct {
		args *args
		want int
	}{
		{
			args: &args{4, 330, [][]int{
				{0, 1, 10, 100},
				{1, 0, 20, 200},
				{10, 20, 0, 300},
				{100, 200, 300, 0},
			}},
			want: 2,
		},
		{
			args: &args{5, 5, [][]int{
				{0, 1, 1, 1, 1},
				{1, 0, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 1, 1, 0},
			}},
			want: 24,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.k, tt.args.ttt); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
