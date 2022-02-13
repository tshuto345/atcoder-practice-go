package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n  int
		st []mountain
	}
	tests := []struct {
		args *args
		want string
	}{
		{
			args: &args{3, []mountain{
				{"Everest", 8849},
				{"K2", 8611},
				{"Kangchenjunga", 8586},
			}},
			want: "K2",
		},
		{
			args: &args{4, []mountain{
				{"Kita", 3193},
				{"Aino", 3189},
				{"Fuji", 3776},
				{"Okuhotaka", 3190},
			}},
			want: "Kita",
		},
		{
			args: &args{4, []mountain{
				{"QCFium", 2846},
				{"chokudai", 2992},
				{"kyoprofriends", 2432},
				{"penguinman", 2390},
			}},
			want: "QCFium",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.st); got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
