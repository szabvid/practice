// Find all pairs with sum k within an array (assuming all distinct elements).

package main

func BinarySearch(A []int, e int) int {
	l, r := 0, len(A)-1

	for r >= l {
		m := (l + r) / 2

		if e == A[m] {
			return e
		} else if e < A[m] {
			r = m - 1
		} else if e > A[m] {
			l = m + 1
		}
	}

	return -1
}

type Pair struct {
	first  int
	second int
}

func SumPairFinder(A []int, sum int) []Pair {
	var res []Pair

	for i := 0; i < len(A)-1; i++ {
		pairCandidate := sum - A[i]

		// search in the rest (among the higher values)
		if BinarySearch(A[i+1:], pairCandidate) != -1 {
			res = append(res, Pair{A[i], pairCandidate})
		}
	}

	return res
}

func main() {

}
