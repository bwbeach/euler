package integer

// True if n is a multiple of any of the numbers in
// factors
func isMultipleOfAny(n int, factors []int) bool {
	for _, f := range factors {
		if n % f == 0 {
			return true
		}
	}
	return false
}

// Returns the sum of all integers in the range [1, max) that
// are multiples of any of the given factors.
func SumMultiples(factors []int, max int) int {
	result := 0
	for i := 1; i < max; i++ {
		if isMultipleOfAny(i, factors) {
			result += i;
		}
	}
	return result
}
