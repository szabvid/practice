package main

import "testing"

func TestCommonFinder(t *testing.T) {
	cases := []struct {
		A   []int
		B   []int
		res int
	}{
		{[]int{13, 27, 35, 40, 49, 55, 59}, []int{17, 35, 39, 40, 55, 58, 60}, 3},
		{[]int{13, 27, 35, 40, 49, 55, 59}, []int{14, 28, 36, 41, 50, 56, 60}, 0},
		{[]int{13, 27, 35, 40, 49, 55, 59}, []int{13, 27, 35, 40, 49, 55, 59}, 7},
	}

	for _, c := range cases {
		commons := commonFinder(c.A, c.B)
		if len(commons) != c.res {
			t.Errorf("Incorrect: got %d, want: %d.\n", len(commons), c.res)
		}
	}

}
