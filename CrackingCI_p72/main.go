// Find all paris with sum k within an array (assuming all distinct elements).

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

func main() {

}
