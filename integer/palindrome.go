package integer

import "strconv"

// Returns true iff the string representing the number is a palindrome
func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)

	end := len(s)
	half := (len(s) + 1) / 2
	for i := 0; i <= half; i++ {
		if s[i] != s[end - i - 1] {
			return false
		}
	}
	return true
}
