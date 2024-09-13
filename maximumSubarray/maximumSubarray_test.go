package maximumSubarray

import (
	"testing"
)

func TestMaxSubArraryDpKandane(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArraryDpKandane(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArraryDpKandane() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubArraryGreedy(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArraryGreedy(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArraryGreedy() = %v, want %v", got, tt.want)
			}
		})
	}
}
