package integer

import (
	"fmt"
	"testing"
)

func TestUnion(t *testing.T) {
	bag1 := NewIntBag()
	checkMap("[]", bag1, "empty", t)
	bag1.Add(2)
	bag1.Add(2)
	bag1.Add(1)
	checkMap("[1 2 2]", bag1, "bag1", t)

	bag2 := NewIntBag()
	bag2.Add(1)
	bag2.Add(1)
	bag2.Add(2)
	bag2.Add(3)
	checkMap("[1 1 2 3]", bag2, "bag2", t)

	bag3 := union(bag1, bag2)
	checkMap("[1 1 2 2 3]", bag3, "bag3", t)
}

func checkMap(expected_str string, actual *IntBag, testName string, t *testing.T) {
	actual_array := actual.ToSlice()
	actual_str := fmt.Sprintf("%v", actual_array)
	if expected_str != actual_str {
		t.Errorf("%v: expected %v but got %v", testName, expected_str, actual_str)
	}
}
