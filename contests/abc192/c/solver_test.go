package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n, k int
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{314, 2},
			want: "693",
		},
		{
			args: &args{1000000000, 100},
			want: "0",
		},
		{
			args: &args{6174, 100000},
			want: "6174",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
