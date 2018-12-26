// Given two sorted arrays, find the number of elements in common.
// The arrays are the same length and each has all distinct elements.

package main

func CommonFinder(A []int, B []int) []int {
	var commons []int
	a, b := 0, 0

	for a < len(A) && b < len(B) {
		if A[a] == B[b] {
			commons = append(commons, A[a])
			a++
			b++
		} else if A[a] > B[b] {
			b++
		} else if A[a] < B[b] {
			a++
		}
	}

	return commons
}

func main() {

}
