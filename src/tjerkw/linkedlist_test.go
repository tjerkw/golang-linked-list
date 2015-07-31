package tjerkw

import "testing"

func testInput(input []interface{}, t *testing.T) *LinkedList {

	xs := NewLinkedList(input)
	if len(input) == 0 {
		if xs != nil {
			t.Fatalf("Expected nil for empty lists")
		}
		return nil
	} else if xs.Len() != len(input) {
		t.Fatalf("List should be of length %v, but got len %v for %v", len(input), xs.Len(), xs)
	}

	xs, err := xs.Push(9542)
	if err != nil {
		t.Fatal(err)
	}
	if xs.Len() != len(input)+1 {
		t.Fatalf("List should have one item")
	}
	return xs
}

func testFind(xs *LinkedList, searchFor interface{}, expectedIndex int, t *testing.T) {
	index, err := xs.Find(searchFor)
	if err != nil {
		t.Fatal(err)
	}
	if index != expectedIndex {
		t.Fatalf("Got wrong index for %v, got %v, expected %v for %#v", searchFor, index, expectedIndex, xs)
	}
}

func makeGeneric(list []int) []interface{} {
	genericList := make([]interface{}, len(list))
	for i, val := range list {
		genericList[i] = val
	}
	return genericList
}

func TestLinkedList(t *testing.T) {

	testInput(makeGeneric([]int{}), t)
	xs := testInput(makeGeneric([]int{6, 7}), t)
	testFind(xs, 6, 0, t)
	testFind(xs, 7, 1, t)
	xs = testInput(makeGeneric([]int{1, 5, 3}), t)
	testFind(xs, 1, 0, t)
	testFind(xs, 5, 1, t)
	testFind(xs, 3, 2, t)
}
