package list

import "testing"

func TestNewCircularDoublyLinkedList(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	if list == nil {
		t.Fatal("NewCircularDoublyLinkedList returned nil")
	}
	if !list.IsEmpty() {
		t.Error("expected new list to be empty")
	}
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
	if list.Head() != nil {
		t.Error("expected Head to be nil on empty list")
	}
	if list.Tail() != nil {
		t.Error("expected Tail to be nil on empty list")
	}
}

func TestAppendAndPrepend(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	if list.Size() != 3 {
		t.Errorf("expected size 3, got %d", list.Size())
	}
	expected := []int{0, 1, 2}
	curr := list.Head()
	for i, v := range expected {
		if curr == nil {
			t.Fatalf("expected node at index %d, got nil", i)
		}
		if curr.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, curr.Value())
		}
		curr = curr.Next()
	}
}

func TestFindAndContains(t *testing.T) {
	list := NewCircularDoublyLinkedList[string]()
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

func TestRemoveFirst(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.RemoveFirst()
	if list.Size() != 1 {
		t.Errorf("expected size 1, got %d", list.Size())
	}
	if list.Head().Value() != 2 {
		t.Errorf("expected head value 2, got %v", list.Head().Value())
	}
	list.RemoveFirst()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing all elements")
	}
}

func TestRemoveLast(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.RemoveLast()
	if list.Size() != 1 {
		t.Errorf("expected size 1, got %d", list.Size())
	}
	if list.Tail().Value() != 1 {
		t.Errorf("expected tail value 1, got %v", list.Tail().Value())
	}
	list.RemoveLast()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing all elements")
	}
}

func TestRemoveValue(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
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
	list.Remove(42)
	if list.Size() != 2 {
		t.Errorf("expected size 2 after removing non-existing value, got %d", list.Size())
	}
}

func TestClear(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
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

func TestInsertAt(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Append(1)
	list.Append(3)
	err := list.InsertAt(1, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if list.Size() != 3 {
		t.Errorf("expected size 3, got %d", list.Size())
	}
	expected := []int{1, 2, 3}
	curr := list.Head()
	for i, v := range expected {
		if curr == nil {
			t.Fatalf("expected node at index %d, got nil", i)
		}
		if curr.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, curr.Value())
		}
		curr = curr.Next()
	}
}

func TestInsertAtOutOfBounds(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	err := list.InsertAt(-1, 10)
	if err == nil {
		t.Error("expected error for negative index")
	}
	err = list.InsertAt(1, 10)
	if err == nil {
		t.Error("expected error for index out of range")
	}
}

func TestGetSet(t *testing.T) {
	list := NewCircularDoublyLinkedList[string]()
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
	_, err = list.Get(-1)
	if err == nil {
		t.Error("expected error for negative index")
	}
	err = list.Set(5, "x")
	if err == nil {
		t.Error("expected error for out of bounds index")
	}
}

func TestReverse(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	for i := 1; i <= 3; i++ {
		list.Append(i)
	}
	list.Reverse()
	expected := []int{3, 2, 1}
	curr := list.Head()
	for i, v := range expected {
		if curr.Value() != v {
			t.Errorf("at index %d, expected %d, got %v", i, v, curr.Value())
		}
		curr = curr.Next()
	}
}

func TestForEach(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	sum := 0
	list.ForEach(func(val int) { sum += val })
	if sum != 6 {
		t.Errorf("expected sum 6, got %d", sum)
	}
}

func TestStringEmptyAndNonEmpty(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	got := list.String()
	want := "CircularDoublyLinkedList: []"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	got = list.String()
	want = "CircularDoublyLinkedList: [1] <-> [2] <-> [3]"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestFindOnEmptyList(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	node := list.Find(42)
	if node != nil {
		t.Error("expected nil when finding in empty list")
	}
}

func TestRemoveFirstOnEmptyList(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.RemoveFirst()
	if !list.IsEmpty() {
		t.Error("expected list to remain empty after RemoveFirst on empty list")
	}
}

func TestRemoveLastOnEmptyList(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.RemoveLast()
	if !list.IsEmpty() {
		t.Error("expected list to remain empty after RemoveLast on empty list")
	}
}

func TestRemoveOnEmptyList(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Remove(42)
	if !list.IsEmpty() {
		t.Error("expected list to remain empty after Remove on empty list")
	}
}

func TestRemoveSingleElement(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Append(42)
	list.Remove(42)
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing the only element")
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil after removing the only element")
	}
}

func TestRemoveTailValue(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
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

func TestInsertAtZero(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	err := list.InsertAt(0, 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if list.Head().Value() != 100 {
		t.Errorf("expected head value 100, got %v", list.Head().Value())
	}
}

func TestInsertAtSize(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	err := list.InsertAt(list.Size(), 300)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if list.Tail().Value() != 300 {
		t.Errorf("expected tail value 300, got %v", list.Tail().Value())
	}
}

func TestReverseEmptyList(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	list.Reverse()
	if !list.IsEmpty() {
		t.Error("expected list to remain empty after Reverse on empty list")
	}
}

func TestForEachEmptyList(t *testing.T) {
	list := NewCircularDoublyLinkedList[int]()
	called := false
	list.ForEach(func(val int) { called = true })
	if called {
		t.Error("expected action not to be called on empty list")
	}
}
