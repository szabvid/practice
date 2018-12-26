package main

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		A   []int
		e   int
		exp int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 4, 4},
		{[]int{1, 2, 3, 4, 5, 6}, 0, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4, 4},
		{[]int{0}, 0, 0},
		{[]int{1}, 4, -1},
	}

	for _, c := range cases {
		result := BinarySearch(c.A, c.e)
		if result != c.exp {
			t.Errorf("Incorrect: got %d, want: %d.\n", result, c.exp)
		}
	}

}
