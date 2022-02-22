package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args *args
		want []int
	}{
		{
			args: &args{6},
			want: []int{1, 2, 3, 6},
		},
		{
			args: &args{720},
			want: []int{
				1,
				2,
				3,
				4,
				5,
				6,
				8,
				9,
				10,
				12,
				15,
				16,
				18,
				20,
				24,
				30,
				36,
				40,
				45,
				48,
				60,
				72,
				80,
				90,
				120,
				144,
				180,
				240,
				360,
				720,
			},
		},
		{
			args: &args{1000000007},
			want: []int{1, 1000000007},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
