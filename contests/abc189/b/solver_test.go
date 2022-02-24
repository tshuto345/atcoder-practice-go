package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n, x int
		vp   [][2]int
	}
	tests := []struct {
		args *args
		want int
	}{
		{
			args: &args{2, 15, [][2]int{
				{200, 5},
				{350, 3},
			}},
			want: 2,
		},
		{
			args: &args{2, 10, [][2]int{
				{200, 5},
				{350, 3},
			}},
			want: 2,
		},
		{
			args: &args{3, 1000000, [][2]int{
				{1000, 100},
				{1000, 100},
				{1000, 100},
			}},
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.x, tt.args.vp); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
