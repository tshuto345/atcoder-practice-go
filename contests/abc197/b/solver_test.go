package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		h, w, x, y int
		ss         []string
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{4, 4, 2, 2, []string{
				"##..",
				"...#",
				"#.#.",
				".#.#",
			}},
			want: "4",
		},
		{
			args: &args{3, 5, 1, 4, []string{
				"#....",
				"#####",
				"....#",
			}},
			want: "4",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.h, tt.args.w, tt.args.x, tt.args.y, tt.args.ss); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
