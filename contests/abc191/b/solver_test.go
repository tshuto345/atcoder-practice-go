package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n, x int
		aa   []int
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{5, 5, []int{3, 5, 6, 5, 4}},
			want: "3 6 4",
		},
		{
			args: &args{3, 3, []int{3, 3, 3}},
			want: "",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.x, tt.args.aa); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
