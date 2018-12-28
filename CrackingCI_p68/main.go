// Print all positive integer solutions to the equation a^2 + b^2 = c^2 + d^2
// where a, b, c and d are integers between 1 and 1000.

package main

import "fmt"

type FourTuple struct {
	a int
	b int
	c int
	d int
}

func EquationSolver(limit int) map[FourTuple]bool {
	res := make(map[FourTuple]bool)         // using map to check membership easier (for array need to iterate through)
	pairedSquares := make(map[int][][2]int) // for holding a^2 + b^2

	for a := 0; a < limit; a++ {
		for b := a; b < limit; b++ {
			pairedSquares[(a*a)+(b*b)] = append(pairedSquares[(a*a)+(b*b)], [2]int{a, b})
		}
	}

	for c := 0; c < limit; c++ {
		for d := c; d < limit; d++ {
			candidate := (c * c) + (d * d)
			if elements, ok := pairedSquares[candidate]; ok {
				for _, elem := range elements {
					if elem[0] != c && elem[1] != d { // looking for different (a, b) and (c, d) values
						abcd := FourTuple{elem[0], elem[1], c, d}
						bcab := FourTuple{c, d, elem[0], elem[1]}
						if _, ok := res[abcd]; !ok {
							if _, ok := res[bcab]; !ok {
								res[abcd] = true
							}
						}
					}
				}
			}
		}
	}

	return res
}

func main() {
	res := EquationSolver(1000)
	for k, _ := range res {
		fmt.Println(k)
	}
}
