package integer

import (
	"fmt"
	"testing"
)

func TestPrimeChannel(t *testing.T) {
	c := PrimeChannel()
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19}
	for _, p := range expected {
		actual := <-c
		if actual != p {
			t.Errorf("Expected %d but got %d", p, actual)
		}
	}
}

func TestLargestPrimeFactor(t *testing.T) {
	expected := 29
	actual := LargestPrimeFactor(13195)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestFactorNumber(t *testing.T) {
	checkFactoring("[2]", 2, t)
	checkFactoring("[2 2 3]", 12, t)
	checkFactoring("[2 2 5 5]", 100, t)
}

func checkFactoring(expected string, n int, t *testing.T) {
	factors := PrimeFactors(n)
	actual := fmt.Sprintf("%v", factors.ToSlice())
	if expected != actual {
		t.Errorf("Expected %v but got %v for %v", expected, actual, n)
	}
}

