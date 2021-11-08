package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{1, 2}, 3},
		{[]int{3, 2}, 5},
		{[]int{3, 3}, 6},
		{[]int{1, 2}, 3},
	}

	for _, v := range tests {
		x := MySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
