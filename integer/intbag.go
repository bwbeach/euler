package integer

import "sort"

type IntBag map[int]int

func NewIntBag() *IntBag {
	m := map[int]int{}
	return (*IntBag)(& m)
}

func (bag *IntBag) asMap() *map[int]int {
	return (*map[int]int)(bag)
}

func (bag *IntBag) Add(n int) {
	m := bag.asMap()
	(*m)[n] = (*m)[n] + 1
}

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

func union(a *IntBag, b *IntBag) *IntBag {
	result := map[int]int{}
	for k, v := range *(a.asMap()) {
		result[k] = v
	}
	for k, v := range *(b.asMap()) {
		result[k] = max(v, result[k])
	}
	return (*IntBag)(&result)
}
