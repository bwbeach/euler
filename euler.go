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
		if n % 2 == 0 {
			result += n
		}
	}
	return result
}

func getAnswer(problemNumber int) int {
	switch problemNumber {
	case 1: return integer.SumMultiples([]int {3, 5}, 1000)
	case 2: return problem2()
	default: return -1
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
