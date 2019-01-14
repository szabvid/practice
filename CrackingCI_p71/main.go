// Design an algorithm to print all permutations of a string. For simplicity,
// assume all characters are unique.

package main

import "fmt"

func perm(w []string) []string {
	res := []string{""}
	tmp := []string{}

	for i := 0; i < len(w); i++ { // each letter
		tmp = []string{}
		for j := 0; j < len(res); j++ { // each permutation variant
			for k := 0; k <= len(res[j]); k++ { // each insertion place
				tmp = append(tmp, res[j][0:k]+w[i]+res[j][k:])
			}
		}
		res = tmp
	}

	return res
}

func main() {
	for _, elem := range perm([]string{"a", "b", "c"}) {
		fmt.Println(elem)
	}
}

/*
func perm(w []string) []string {
	if len(w) == 0 {
		return []string{""}
	}
	if len(w) == 1 {
		fmt.Println(w)
		return w
	}

	fLetter := w[0]
	for i := 0; i < len(w); i++ {
		fmt.Println(append(perm(w[1:]), fLetter))
		return perm(w[1:])
	}

	return []string{""}
}

func main() {
	w := [3]string{"a", "b", "c"}
	perm(w[:])
	fmt.Printf("Hello %s", "alma")
}
*/
