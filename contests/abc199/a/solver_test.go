package main

import (
	"strconv"
	"testing"
)

func TestA(t *testing.T) {
	type args struct {
		a, b, c int
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{2, 2, 4},
			want: "Yes",
		},
		{
			args: &args{10, 10, 10},
			want: "No",
		},
		{
			args: &args{3, 4, 5},
			want: "No",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
