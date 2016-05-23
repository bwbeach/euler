package integer

import "testing"

func TestFibonacciGenerator(t *testing.T) {
	fib := FibonacciGenerator()
	for i, expected := range []int{1, 1, 2, 3, 5, 8, 13, 21} {
		actual := fib()
		if expected != actual {
			t.Errorf("iteration %v: expected %v but got %v", i, expected, actual)
		}
	}
}
