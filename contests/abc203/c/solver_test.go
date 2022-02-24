package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n, k int
		ab   [][2]int
	}
	tests := []struct {
		args *args
		want int
	}{
		{
			args: &args{2, 3, [][2]int{
				{2, 1},
				{5, 10},
			}},
			want: 4,
		},
		{
			args: &args{5, 1000000000, [][2]int{
				{1, 1000000000},
				{2, 1000000000},
				{3, 1000000000},
				{4, 1000000000},
				{5, 1000000000},
			}},
			want: 6000000000,
		},
		{
			args: &args{3, 2, [][2]int{
				{5, 5},
				{2, 1},
				{2, 2},
			}},
			want: 10,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.k, tt.args.ab); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
