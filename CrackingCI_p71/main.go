// Design an algorithm to print all permutations of a string. For simplicity,
// assume all characters are unique.

package main

import "fmt"

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
