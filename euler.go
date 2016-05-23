package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/bwbeach/euler/integer"
)

func problem2() int {
	fib := integer.FibonacciGenerator()
	result := 0
	for n := fib(); n < 4000000; n = fib() {
		if n%2 == 0 {
			result += n
		}
	}
	return result
}

func problem4() int {
	result := 1
	for a := 100; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			product := a * b
			if integer.IsPalindrome(product) && result < product {
				result = product
			}
		}
	}
	return result
}

func problem5() int {
	factors := integer.NewIntBag()
	for i := 2; i < 20; i++ {
		factors = integer.Union(factors, integer.PrimeFactors(i))
	}

	result := 1
	for _, f := range factors.ToSlice() {
		result = result * f
	}
	return result
}

func getAnswer(problemNumber int) int {
	switch problemNumber {
	case 1:
		return integer.SumMultiples([]int{3, 5}, 1000)
	case 2:
		return problem2()
	case 3:
		return integer.LargestPrimeFactor(600851475143)
	case 4:
		return problem4()
	case 5:
		return problem5()
	default:
		return -1
	}
}

func main() {
	flag.Parse()
	problemNumber, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Printf("Bad problem number: %v\n", flag.Arg(0))
		return
	}

	fmt.Printf("%v\n", getAnswer(problemNumber))
}
