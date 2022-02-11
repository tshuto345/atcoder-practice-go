package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{"abc"},
			want: "bca",
		},
		{
			args: &args{"aab"},
			want: "aba",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.s); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
