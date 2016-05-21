package integer

import "testing"

func TestPrimeChannel(t *testing.T) {
	c := PrimeChannel()
	expected := []int { 2, 3, 5, 7, 11, 13, 17, 19 }
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
