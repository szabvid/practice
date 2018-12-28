package main

import "testing"

func TestEquationSolver(t *testing.T) {
	abcdList := EquationSolver(1000)

	for abcd, _ := range abcdList {
		a, b, c, d := abcd.a, abcd.b, abcd.c, abcd.d
		if (a*a)+(b*b) != (c*c)+(d*d) {
			t.Errorf("Incorrect: {%d, %d, %d, %d}.\n", a, b, c, d)
		}
	}
}
