package integer

import "testing"

func TestIsPalindrome(t *testing.T) {
	expected := true
	actual := IsPalindrome(323)
	if expected != actual {
		t.Errorf("For %d, expected %v but got %v", 323, expected, actual)
	}

	expected = false
	actual = IsPalindrome(324)
	if expected != actual {
		t.Errorf("For %d, expected %v but got %v", 324, expected, actual)
	}

	// TODO: more cases
}
