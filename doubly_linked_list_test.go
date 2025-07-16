package list

import "testing"

func TestDoublyLinkedListNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	if list == nil {
		t.Fatal("NewDoublyLinkedList returned nil")
	}
	if !list.IsEmpty() {
		t.Error("expected new list to be empty")
	}
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
}

func TestDoublyLinkedListAppendAndPrepend(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(2)
	list.Prepend(1)
	list.Append(3)
	expected := []int{1, 2, 3}
	got := list.ToSlice()
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("at index %d, expected %d, got %d", i, v, got[i])
		}
	}
}

func TestDoublyLinkedListInsertAt(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(3)
	err := list.InsertAt(1, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{1, 2, 3}
	got := list.ToSlice()
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("at index %d, expected %d, got %d", i, v, got[i])
		}
	}
	err = list.InsertAt(0, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if list.Head().Value() != 0 {
		t.Errorf("expected head value 0, got %v", list.Head().Value())
	}
	err = list.InsertAt(list.Size(), 4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if list.Tail().Value() != 4 {
		t.Errorf("expected tail value 4, got %v", list.Tail().Value())
	}
	err = list.InsertAt(-1, 99)
	if err == nil {
		t.Error("expected error for negative index")
	}
	err = list.InsertAt(100, 99)
	if err == nil {
		t.Error("expected error for index out of range")
	}
}

func TestDoublyLinkedListGetAndSet(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	node, err := list.Get(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if node.Value() != "b" {
		t.Errorf("expected 'b', got %v", node.Value())
	}
	err = list.Set(1, "z")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	node, _ = list.Get(1)
	if node.Value() != "z" {
		t.Errorf("expected 'z', got %v", node.Value())
	}
	_, err = list.Get(-1)
	if err == nil {
		t.Error("expected error for negative index")
	}
	err = list.Set(3, "x")
	if err == nil {
		t.Error("expected error for out-of-bounds index")
	}
}

func TestDoublyLinkedListRemoveFirstAndLast(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveFirst()
	if list.Head().Value() != 2 {
		t.Errorf("expected head value 2, got %v", list.Head().Value())
	}
	list.RemoveLast()
	if list.Tail().Value() != 2 {
		t.Errorf("expected tail value 2, got %v", list.Tail().Value())
	}
	list.RemoveFirst()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing all elements")
	}
}

func TestDoublyLinkedListRemove(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(2)
	expected := []int{1, 3}
	got := list.ToSlice()
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("at index %d, expected %d, got %d", i, v, got[i])
		}
	}
	list.Remove(1)
	if list.Head().Value() != 3 {
		t.Errorf("expected head value 3, got %v", list.Head().Value())
	}
	list.Remove(3)
	if !list.IsEmpty() {
		t.Error("expected list to be empty")
	}
	list.Remove(99)
}

func TestDoublyLinkedListReverse(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()
	expected := []int{3, 2, 1}
	got := list.ToSlice()
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("at index %d, expected %d, got %d", i, v, got[i])
		}
	}
}

func TestDoublyLinkedListForEach(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	sum := 0
	list.ForEach(func(val int) { sum += val })
	if sum != 6 {
		t.Errorf("expected sum 6, got %d", sum)
	}
}

func TestDoublyLinkedListClear(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDoublyLinkedListToSliceAndString(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	expected := "DoublyLinkedList: [1] ↔ [2] ↔ [3]"
	got := list.String()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestDoublyLinkedListPrependOnEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDoublyLinkedListRemoveFirstOnEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.RemoveFirst()
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil for empty list")
	}
}

func TestDoublyLinkedListRemoveLastOnEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.RemoveLast()
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil for empty list")
	}
}

func TestDoublyLinkedListRemoveLastSingleElement(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(42)
	list.RemoveLast()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing last element")
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Error("expected head and tail to be nil after removing last element")
	}
}

func TestDoublyLinkedListRemoveTailValue(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(3)
	if list.Size() != 2 {
		t.Errorf("expected size 2 after removing tail, got %d", list.Size())
	}
	if list.Tail().Value() != 2 {
		t.Errorf("expected tail value 2 after removing tail, got %v", list.Tail().Value())
	}
	if list.Tail().Next() != nil {
		t.Error("expected tail.Next to be nil after removing tail")
	}
}

func TestDoublyLinkedListStringEmpty(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	got := list.String()
	want := "DoublyLinkedList: []"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
