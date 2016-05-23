package integer

// Writes primes forever to the channel.
func generatePrimes(c chan int) {
	// Special case for the only even prime
	c <- 2

	// List of primes generated so far, in order.
	found := make([]int, 0)

	// One to prime the pump
	c <- 3
	found = append(found, 3)

	// Generate forever, adding primes to 'found' as
	// they are found.
	for n := 5; true; n += 2 {
		for _, p := range found {
			if n < p*p {
				c <- n
				found = append(found, n)
				break
			}
			if n%p == 0 {
				break
			}
		}
	}
}

// Returns a channel from which prime ints can be read.
func PrimeChannel() chan int {
	c := make(chan int)
	go generatePrimes(c)
	return c
}

// Returns the largest prime factor of a number
func LargestPrimeFactor(n int) int {
	result := 1
	primes := PrimeChannel()
	for 1 < n {
		p := <-primes
		for n%p == 0 {
			result = p
			n = n / p
		}
	}
	return result
}

// Returns the prime factors for n as a bag of ints
func PrimeFactors(n int) *IntBag {
	result := NewIntBag()
	primes := PrimeChannel()
	for 1 < n {
		p := <-primes
		for n%p == 0 {
			result.Add(p)
			n = n / p
		}
	}
	return result
}
