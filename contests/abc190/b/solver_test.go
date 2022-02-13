package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n, s, d int
		xy      [][2]int
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{4, 9, 9, [][2]int{{5, 5}, {15, 5}, {5, 15}, {15, 15}}},
			want: "Yes",
		},
		{
			args: &args{3, 691, 273, [][2]int{{691, 997}, {593, 273}, {691, 273}}},
			want: "No",
		},
		{
			args: &args{7, 100, 100, [][2]int{{10, 11}, {12, 67}, {192, 79}, {154, 197}, {142, 158}, {20, 25}, {17, 108}}},
			want: "Yes",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.s, tt.args.d, tt.args.xy); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
