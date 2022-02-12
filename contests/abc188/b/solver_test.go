package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n      int
		aa, bb []int
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{2, []int{-3, 6}, []int{4, 2}},
			want: "Yes",
		},
		{
			args: &args{2, []int{4, 5}, []int{-1, -3}},
			want: "No",
		},
		{
			args: &args{3, []int{1, 3, 5}, []int{3, -6, 3}},
			want: "Yes",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.aa, tt.args.bb); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
