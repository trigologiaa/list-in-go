package list

import "testing"

func TestCircularSinglyLinkedListNewCircularSinglyLinkedList(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	if list == nil {
		t.Fatal("NewCircularSinglyLinkedList returned nil")
	}
	if !list.IsEmpty() {
		t.Error("expected new list to be empty")
	}
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
}

func TestCircularSinglyLinkedListAppendAndPrepend(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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

func TestCircularSinglyLinkedListFindAndContains(t *testing.T) {
	list := NewCircularSinglyLinkedList[string]()
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

func TestCircularSinglyLinkedListRemoveFirst(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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

func TestCircularSinglyLinkedListRemoveLast(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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

func TestCircularSinglyLinkedListRemoveValue(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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

func TestCircularSinglyLinkedListRemoveValueNotFound(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Remove(4)
	if list.Size() != 2 {
		t.Errorf("expected size 2, got %d", list.Size())
	}
}

func TestCircularSinglyLinkedListClear(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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

func TestCircularSinglyLinkedListInsertAt(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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
		if current == nil {
			t.Fatalf("expected node at index %d, got nil", i)
		}
		if current.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, current.Value())
		}
		current = current.Next()
	}
}

func TestCircularSinglyLinkedListInsertAtOutOfBounds(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	err := list.InsertAt(-1, 10)
	if err == nil {
		t.Error("expected error for negative index")
	}
	err = list.InsertAt(1, 10)
	if err == nil {
		t.Error("expected error for index out of range")
	}
}

func TestCircularSinglyLinkedListGetAndSet(t *testing.T) {
	list := NewCircularSinglyLinkedList[string]()
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

func TestCircularSinglyLinkedListGetOutOfBounds(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(1)
	_, err := list.Get(-1)
	if err == nil {
		t.Error("expected error for negative index")
	}
	_, err = list.Get(1)
	if err == nil {
		t.Error("expected error for index out of range")
	}
}

func TestCircularSinglyLinkedListSetOutOfBounds(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	err := list.Set(0, 10)
	if err == nil {
		t.Error("expected error for empty list index")
	}
	list.Append(1)
	err = list.Set(1, 20)
	if err == nil {
		t.Error("expected error for index out of range")
	}
}

func TestCircularSinglyLinkedListReverse(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	for i := 1; i <= 3; i++ {
		list.Append(i)
	}
	list.Reverse()
	expected := []int{3, 2, 1}
	current := list.Head()
	for i, v := range expected {
		if current.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, current.Value())
		}
		current = current.Next()
	}
}

func TestCircularSinglyLinkedListStringEmptyAndNonEmpty(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	got := list.String()
	want := "CircularSinglyLinkedList: []"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	got = list.String()
	want = "CircularSinglyLinkedList: [1] -> [2] -> [3]"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestCircularSinglyLinkedListForEach(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	sum := 0
	list.ForEach(func(val int) { sum += val })
	if sum != 6 {
		t.Errorf("expected sum 6, got %d", sum)
	}
}

func TestCircularSinglyLinkedListFindNotFound(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	node := list.Find(42)
	if node != nil {
		t.Errorf("expected nil, got %v", node)
	}
}

func TestCircularSinglyLinkedListRemoveFirstOnEmptyList(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.RemoveFirst()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after RemoveFirst on empty list")
	}
}

func TestCircularSinglyLinkedListRemoveFirstOnSingleElement(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(42)
	list.RemoveFirst()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after RemoveFirst on single-element list")
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil after removing the only element")
	}
}

func TestCircularSinglyLinkedListRemoveLastOnEmptyList(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.RemoveLast()
	if !list.IsEmpty() {
		t.Error("expected list to remain empty after RemoveLast on empty list")
	}
}

func TestCircularSinglyLinkedListRemoveLastOnSingleElement(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(42)
	list.RemoveLast()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after RemoveLast on single-element list")
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil after removing the only element")
	}
}

func TestCircularSinglyLinkedListRemoveLastWithMultipleElements(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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
	if list.Tail().Next() != list.Head() {
		t.Error("expected tail's next to point to head after RemoveLast")
	}
}

func TestCircularSinglyLinkedListRemoveOnEmptyList(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Remove(42)
	if !list.IsEmpty() {
		t.Error("expected list to remain empty after Remove on empty list")
	}
}

func TestCircularSinglyLinkedListRemoveTailValue(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(3)
	if list.Size() != 2 {
		t.Errorf("expected size 2 after removing tail, got %d", list.Size())
	}
	if list.Tail().Value() != 2 {
		t.Errorf("expected new tail value 2, got %v", list.Tail().Value())
	}
}

func TestCircularSinglyLinkedListRemoveSingleElement(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(42)
	list.Remove(42)
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing the only element")
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil after removing the only element")
	}
}

func TestCircularSinglyLinkedListInsertAtZero(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(1)
	err := list.InsertAt(0, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if list.Head().Value() != 0 {
		t.Errorf("expected head value 0, got %v", list.Head().Value())
	}
}

func TestCircularSinglyLinkedListInsertAtSize(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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

func TestCircularSinglyLinkedListInsertAtMiddle(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
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

func TestCircularSinglyLinkedListReverseEmptyList(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Reverse()
	if !list.IsEmpty() {
		t.Error("expected list to remain empty after Reverse")
	}
}

func TestCircularSinglyLinkedListForEachEmptyList(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	called := false
	list.ForEach(func(val int) { called = true })
	if called {
		t.Error("expected action not to be called on empty list")
	}
}

func TestCircularSinglyLinkedListFindOnEmptyList(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	node := list.Find(42)
	if node != nil {
		t.Error("expected nil when finding in empty list")
	}
}

func TestCircularSinglyLinkedListInsertAtMiddleMultiple(t *testing.T) {
	list := NewCircularSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(4)
	err := list.InsertAt(2, 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{1, 2, 3, 4}
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
