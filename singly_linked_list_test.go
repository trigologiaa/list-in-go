package list

import "testing"

func TestSinglyLinkedListNewSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	if list == nil {
		t.Fatal("NewSinglyLinkedList returned nil")
	}
	if !list.IsEmpty() {
		t.Error("expected new list to be empty")
	}
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
}

func TestSinglyLinkedListAppendAndPrepend(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	if list.Size() != 3 {
		t.Errorf("expected size 3, got %d", list.Size())
	}
	expected := []int{0, 1, 2}
	current := list.Head()
	for i, val := range expected {
		if current == nil || current.Value() != val {
			t.Errorf("at index %d, expected %d, got %v", i, val, current)
		}
		current = current.Next()
	}
}

func TestSinglyLinkedListFindAndContains(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	if !list.Contains("b") {
		t.Error("expected list to contain 'b'")
	}
	if list.Contains("z") {
		t.Error("did not expect list to contain 'z'")
	}
	node := list.Find("c")
	if node == nil || node.Value() != "c" {
		t.Error("expected to find node with value 'c'")
	}
}

func TestSinglyLinkedListRemoveFirst(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.RemoveFirst()
	if list.Size() != 1 {
		t.Errorf("expected size 1, got %d", list.Size())
	}
	if list.Head().Value() != 2 {
		t.Errorf("expected head value 2, got %v", list.Head().Value())
	}
}

func TestSinglyLinkedListRemoveLast(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.RemoveLast()
	if list.Size() != 1 {
		t.Errorf("expected size 1, got %d", list.Size())
	}
	if list.Tail().Value() != 1 {
		t.Errorf("expected tail value 1, got %v", list.Tail().Value())
	}
}

func TestSinglyLinkedListRemove(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(2)
	if list.Size() != 2 {
		t.Errorf("expected size 2, got %d", list.Size())
	}
	if list.Contains(2) {
		t.Error("expected value 2 to be removed")
	}
}

func TestSinglyLinkedListInsertAt(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(3)
	err := list.InsertAt(1, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if list.Size() != 3 {
		t.Errorf("expected size 3, got %d", list.Size())
	}
	values := []int{1, 2, 3}
	current := list.Head()
	for i, v := range values {
		if current.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, current.Value())
		}
		current = current.Next()
	}
}

func TestSinglyLinkedListGetAndSet(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	list.Append("a")
	list.Append("b")
	node, err := list.Get(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if node.Value() != "b" {
		t.Errorf("expected value 'b', got %v", node.Value())
	}
	err = list.Set(1, "z")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	node, _ = list.Get(1)
	if node.Value() != "z" {
		t.Errorf("expected updated value 'z', got %v", node.Value())
	}
}

func TestSinglyLinkedListReverse(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	for i := 1; i <= 3; i++ {
		list.Append(i)
	}
	list.Reverse()
	values := []int{3, 2, 1}
	current := list.Head()
	for i, v := range values {
		if current.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, current.Value())
		}
		current = current.Next()
	}
}

func TestSinglyLinkedListClear(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Clear()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after Clear()")
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil after Clear()")
	}
}

func TestSinglyLinkedListForEach(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	sum := 0
	list.ForEach(func(val int) { sum += val })
	if sum != 6 {
		t.Errorf("expected sum 6, got %d", sum)
	}
}

func TestSinglyLinkedListPrependOnEmptyList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Prepend(42)
	if list.Head() == nil || list.Head().Value() != 42 {
		t.Errorf("expected head value 42, got %v", list.Head())
	}
	if list.Tail() == nil || list.Tail().Value() != 42 {
		t.Errorf("expected tail value 42, got %v", list.Tail())
	}
	if list.Size() != 1 {
		t.Errorf("expected size 1, got %d", list.Size())
	}
}

func TestSinglyLinkedListRemoveFirstOnEmptyList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.RemoveFirst()
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil for empty list")
	}
}

func TestSinglyLinkedListRemoveFirstSingleElement(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(42)
	list.RemoveFirst()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing the only element")
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil after removing the only element")
	}
}

func TestSinglyLinkedListRemoveLastOnEmptyList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.RemoveLast()
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
}

func TestSinglyLinkedListRemoveLastSingleElement(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(42)
	list.RemoveLast()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing last element")
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil after removing last element")
	}
}

func TestSinglyLinkedListRemoveLastMultipleElements(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveLast()
	if list.Size() != 2 {
		t.Errorf("expected size 2 after RemoveLast, got %d", list.Size())
	}
	if list.Tail().Value() != 2 {
		t.Errorf("expected tail value 2 after RemoveLast, got %v", list.Tail().Value())
	}
	if list.Tail().Next() != nil {
		t.Error("expected tail's next to be nil after RemoveLast")
	}
}

func TestSinglyLinkedListRemoveValueNotFound(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Remove(3)
	if list.Size() != 2 {
		t.Errorf("expected size 2, got %d", list.Size())
	}
}

func TestSinglyLinkedListRemoveHeadValue(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Remove(1)
	if list.Size() != 1 {
		t.Errorf("expected size 1, got %d", list.Size())
	}
	if list.Head().Value() != 2 {
		t.Errorf("expected head value 2, got %v", list.Head().Value())
	}
}

func TestSinglyLinkedListRemoveMiddleValue(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(2)
	if list.Size() != 2 {
		t.Errorf("expected size 2, got %d", list.Size())
	}
	if list.Contains(2) {
		t.Error("value 2 should be removed")
	}
}

func TestSinglyLinkedListRemoveTailValue(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(3)
	if list.Size() != 2 {
		t.Errorf("expected size 2, got %d", list.Size())
	}
	if list.Tail().Value() != 2 {
		t.Errorf("expected tail value 2, got %v", list.Tail().Value())
	}
}

func TestSinglyLinkedListStringEmpty(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	got := list.String()
	want := "SinglyLinkedList: []"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestSinglyLinkedListStringWithElements(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	got := list.String()
	want := "SinglyLinkedList: [1] -> [2] -> [3]"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestSinglyLinkedListInsertAtIndexOutOfBounds(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	err := list.InsertAt(-1, 10)
	if err == nil {
		t.Error("expected error for negative index")
	}
	err = list.InsertAt(1, 20)
	if err == nil {
		t.Error("expected error for index out of range")
	}
}

func TestSinglyLinkedListInsertAtIndexZero(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	err := list.InsertAt(0, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if list.Head().Value() != 0 {
		t.Errorf("expected head value 0, got %v", list.Head().Value())
	}
}

func TestSinglyLinkedListInsertAtIndexSize(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	err := list.InsertAt(list.Size(), 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if list.Tail().Value() != 3 {
		t.Errorf("expected tail value 3, got %v", list.Tail().Value())
	}
}

func TestSinglyLinkedListInsertAtMiddleIndex(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(3)
	err := list.InsertAt(1, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{1, 2, 3}
	current := list.Head()
	for i, v := range expected {
		if current == nil {
			t.Fatalf("expected node at index %d, got nil", i)
		}
		if current.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, current.Value())
		}
		current = current.Next()
	}
}

func TestSinglyLinkedListGetIndexOutOfBounds(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	_, err := list.Get(-1)
	if err == nil {
		t.Error("expected error for negative index")
	}
	_, err = list.Get(2)
	if err == nil {
		t.Error("expected error for index >= size")
	}
}

func TestSinglyLinkedListSetIndexOutOfBounds(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	err := list.Set(2, 10)
	if err == nil || err.Error() != "index 2 out of bounds" {
		t.Errorf("expected error 'index 2 out of bounds', got %v", err)
	}
}

func TestSinglyLinkedListInsertAtIndexGreaterThan2(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(3)
	list.Append(4)
	err := list.InsertAt(2, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{1, 3, 2, 4}
	current := list.Head()
	for i, v := range expected {
		if current == nil {
			t.Fatalf("expected node at index %d, got nil", i)
		}
		if current.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, current.Value())
		}
		current = current.Next()
	}
}
