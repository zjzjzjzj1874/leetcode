package day1

import "sort"

var bad int

func firstBadVersion(n int) int {

	sort.Search(n, func(v int) bool {
		return isBadVersion(v)
	})


	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !isBadVersion(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func isBadVersion(tar int) bool {
	return tar == bad
}
