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
			args: &args{6},
			want: 6,
		},
		{
			args: &args{20},
			want: 17,
		},
		{
			args: &args{100000},
			want: 30555,
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

func TestBase(t *testing.T) {
	type args struct {
		n int
		b int
	}
	tests := []struct {
		args *args
		want int
	}{
		{
			args: &args{0, 7},
			want: 0,
		},
		{
			args: &args{1, 7},
			want: 1,
		},
		{
			args: &args{6, 7},
			want: 6,
		},
		{
			args: &args{7, 7},
			want: 10,
		},
		{
			args: &args{8, 7},
			want: 11,
		},
		{
			args: &args{0, 3},
			want: 0,
		},
		{
			args: &args{1, 3},
			want: 1,
		},
		{
			args: &args{6, 3},
			want: 20,
		},
		{
			args: &args{7, 3},
			want: 21,
		},
		{
			args: &args{8, 3},
			want: 22,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := base(tt.args.n, tt.args.b)
			if err != nil {
				t.Errorf("\nexpected: err == nil: %v\n", err)
			}
			if got != tt.want {
				t.Errorf("\nexpected: %v\nresult  : %v", tt.want, got)
			}
		})
	}
}
