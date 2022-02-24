package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n  int
		xy [][2]int
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{4, [][2]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 1},
			}},
			want: "Yes",
		},
		{
			args: &args{14, [][2]int{
				{5, 5},
				{0, 1},
				{2, 5},
				{8, 0},
				{2, 1},
				{0, 0},
				{3, 6},
				{8, 6},
				{5, 9},
				{7, 9},
				{3, 4},
				{9, 2},
				{9, 8},
				{7, 2},
			}},
			want: "No",
		},
		{
			args: &args{9, [][2]int{
				{8, 2},
				{2, 3},
				{1, 3},
				{3, 7},
				{1, 0},
				{8, 8},
				{5, 6},
				{9, 7},
				{0, 1},
			}},
			want: "Yes",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.xy); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
