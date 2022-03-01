package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args *args
		want int
	}{
		{
			args: &args{35},
			want: 1,
		},
		{
			args: &args{369},
			want: 0,
		},
		{
			args: &args{6227384},
			want: 1,
		},
		{
			args: &args{11},
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
