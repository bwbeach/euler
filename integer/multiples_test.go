package integer

import "testing"

func TestSumMultiples(t *testing.T) {
	actual := SumMultiples([]int {3, 5}, 10)
	if 23 != actual {
		t.Errorf("expected 23, got %v", actual)
	}
}
