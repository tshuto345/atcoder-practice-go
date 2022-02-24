package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		a, b int
	}
	tests := []struct {
		args *args
		want int
	}{
		{
			args: &args{2, 3},
			want: 6,
		},
		{
			args: &args{123, 456},
			want: 18696,
		},
		{
			args: &args{100000, 99999},
			want: 9999900000,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
