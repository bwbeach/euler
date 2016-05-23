package integer

import "sort"

// IntBag is a bag of integers
type IntBag map[int]int

// Creates a new, empty bag
func NewIntBag() *IntBag {
	m := map[int]int{}
	return (*IntBag)(& m)
}

// internal function that returns the map that represents a bag
func (bag *IntBag) asMap() *map[int]int {
	return (*map[int]int)(bag)
}

// Adds a number to a bag
func (bag *IntBag) Add(n int) {
	m := bag.asMap()
	(*m)[n] = (*m)[n] + 1
}

// Converts a bag of numbers to a slice, with the numbers in order,
// and repeated the number of times they are in the bag.
func (bag *IntBag) ToSlice() []int {
	m := bag.asMap()
	values := []int{}
	for k, _ := range *m {
		values = append(values, k)
	}
	sort.Ints(values)

	result := []int{}
	for _, value := range values {
		count := (*m)[value]
		for i := 0; i < count; i++ {
			result = append(result, value)
		}
	}
	return result
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// Returns the "union" of two bags, where the number of times each
// number occurs is the maximum of the how many times it occurs in each
// input bag.
func Union(a *IntBag, b *IntBag) *IntBag {
	result := map[int]int{}
	for k, v := range *(a.asMap()) {
		result[k] = v
	}
	for k, v := range *(b.asMap()) {
		result[k] = max(v, result[k])
	}
	return (*IntBag)(&result)
}
