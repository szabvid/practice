package main

import (
	"reflect"
	"testing"
)

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

func TestSumPairFinder(t *testing.T) {
	cases := []struct {
		A   []int
		sum int
		exp []Pair
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, []Pair{{1, 4}, {2, 3}}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, []Pair{}},
		{[]int{1, 2, 10}, 11, []Pair{{1, 10}}},
	}

	for _, c := range cases {
		result := SumPairFinder(c.A, c.sum)

		for i, v := range result {
			if !reflect.DeepEqual(v, c.exp[i]) {
				t.Errorf("Incorrect: got %v\n want: %v\n", result, c.exp)
			}
		}
	}
}
