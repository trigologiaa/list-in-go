// Package list provides generic linked list data structures and nodes in Go.
//
// It includes implementations for singly linked lists, doubly linked lists, and
// their circular variants, as well as the corresponding node types. All lists are
// generic and work with any comparable type T.
//
// The package offers a rich set of operations such as insertion, deletion, search,
// traversal, reversal, and random access.
//
// Both linear and circular lists support iteration that respects their structural
// properties.
//
// ## Provided Types:
//
//   - SinglyLinkedList[T]:
//     A linear singly linked list where each node points to the next node.
//   - CircularSinglyLinkedList[T]:
//     A circular singly linked list where the last node points back to the first
//     node.
//   - DoublyLinkedList[T]:
//     A linear doubly linked list where each node points to both the next and
//     previous nodes.
//   - CircularDoublyLinkedList[T]:
//     A circular doubly linked list where the last node points to the first node
//     and vice versa.
//   - SinglyLinkedNode[T]:
//     A node for singly linked lists, storing a value and a pointer to the next
//     node.
//   - DoublyLinkedNode[T]:
//     A node for doubly linked lists, storing a value and pointers to both the
//     next and previous nodes.
//
// ## Features:
//
//   - Generic (works with any comparable type T)
//   - Insertion at head, tail, or arbitrary index
//   - Removal by value, head, or tail
//   - Search and containment checks
//   - Traversal and ForEach iteration
//   - Reversal of list order
//   - Conversion to slices for interoperability
//
// ## Examples:
//
// SinglyLinkedList:
//
//	list := list.NewSinglyLinkedList[int]()
//	list.Append(1)
//	list.Prepend(0)
//	fmt.Println(list) // SinglyLinkedList: [0] -> [1]
//	list.Reverse()
//	fmt.Println(list) // SinglyLinkedList: [1] -> [0]
//
// CircularSinglyLinkedList:
//
//	clist := list.NewCircularSinglyLinkedList[int]()
//	clist.Append(1)
//	clist.Append(2)
//	clist.Append(3)
//	fmt.Println(clist) // CircularSinglyLinkedList: [1] -> [2] -> [3]
//	clist.Remove(2)
//	fmt.Println(clist) // CircularSinglyLinkedList: [1] -> [3]
//
// DoublyLinkedList:
//
//	dlist := list.NewDoublyLinkedList[int]()
//	dlist.Append(10)
//	dlist.Prepend(5)
//	dlist.InsertAt(1, 7)
//	fmt.Println(dlist) // DoublyLinkedList: [5] ↔ [7] ↔ [10]
//	dlist.Reverse()
//	fmt.Println(dlist) // DoublyLinkedList: [10] ↔ [7] ↔ [5]
//
// CircularDoublyLinkedList:
//
//	cdlist := list.NewCircularDoublyLinkedList[int]()
//	cdlist.Append(10)
//	cdlist.Append(20)
//	cdlist.Prepend(5)
//	fmt.Println(cdlist) // CircularDoublyLinkedList: [5] <-> [10] <-> [20]
//	cdlist.Reverse()
//	fmt.Println(cdlist) // CircularDoublyLinkedList: [20] <-> [10] <-> [5]
//
// ## Notes:
//
// All lists are dynamic in size and support O(1) insertion and removal at the ends
// (head/tail).
//
// Random access operations (Get, Set) have O(n) complexity due to linear traversal.
package list

import "fmt"

// A generic singly linked list storing elements of type T.
//
// T must be comparable to allow element equality checks.
type SinglyLinkedList[T comparable] struct {
	head *SinglyLinkedNode[T]
	tail *SinglyLinkedNode[T]
	size int
}

// Creates and returns a new empty singly linked list.
//
// Returns:
//   - *SinglyLinkedList[T]: A pointer to an empty list.
//
// Example:
//
//	list := list.NewSinglyLinkedList[string]()
func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Returns the first node of the list.
//
// Returns:
//   - *SinglyLinkedNode[T]: Pointer to the head node or nil if empty.
//
// Example:
//
//	head := list.Head()
func (l *SinglyLinkedList[T]) Head() *SinglyLinkedNode[T] {
	return l.head
}

// Returns the last node of the list.
//
// Returns:
//   - *SinglyLinkedNode[T]: Pointer to the tail node or nil if empty.
//
// Example:
//
//	tail := list.Tail()
func (l *SinglyLinkedList[T]) Tail() *SinglyLinkedNode[T] {
	return l.tail
}

// Returns the number of elements in the list.
//
// Returns:
//   - int: Count of nodes.
//
// Example:
//
//	fmt.Println(list.Size())
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// Returns true if the list contains no elements.
//
// Returns:
//   - bool: true if list is empty; false otherwise.
//
// Example:
//
//	if list.IsEmpty() {
//	    fmt.Println("List is empty")
//	}
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

// Removes all elements from the list, resetting it to empty.
//
// Example:
//
//	list.Clear()
func (l *SinglyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Inserts a new element at the start of the list.
//
// Parameters:
//   - value: Element to insert.
//
// Example:
//
//	list.Prepend(5)
func (l *SinglyLinkedList[T]) Prepend(value T) {
	newNode := NewSinglyLinkedNode(value)
	newNode.next = l.Head()
	l.head = newNode
	if l.Tail() == nil {
		l.tail = newNode
	}
	l.size++
}

// Adds a new element at the end of the list.
//
// Parameters:
//   - value: Element to insert.
//
// Example:
//
//	list.Append(10)
func (l *SinglyLinkedList[T]) Append(value T) {
	newNode := NewSinglyLinkedNode(value)
	if l.Head() == nil {
		l.head = newNode
	}
	if l.Tail() != nil {
		l.Tail().next = newNode
	}
	l.tail = newNode
	l.size++
}

// Searches for the first node containing the specified value.
//
// Parameters:
//   - value: Value to search for.
//
// Returns:
//   - *SinglyLinkedNode[T]: Pointer to the node if found; nil otherwise.
//
// Example:
//
//	node := list.Find(5)
func (l *SinglyLinkedList[T]) Find(value T) *SinglyLinkedNode[T] {
	for current := l.Head(); current != nil; current = current.Next() {
		if current.Value() == value {
			return current
		}
	}
	return nil
}

// Removes the first element from the list.
//
// Does nothing if the list is empty.
//
// Example:
//
//	list.RemoveFirst()
func (l *SinglyLinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}
	l.head = l.Head().Next()
	if l.Head() == nil {
		l.tail = nil
	}
	l.size--
}

// Removes the last element from the list.
//
// Does nothing if the list is empty.
//
// Example:
//
//	list.RemoveLast()
func (l *SinglyLinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		current := l.Head()
		for current.Next() != l.Tail() {
			current = current.Next()
		}
		current.SetNext(nil)
		l.tail = current
	}
	l.size--
}

// Deletes the first node found with the specified value.
//
// Does nothing if the value is not found.
//
// Parameters:
//   - value: Element to remove.
//
// Example:
//
//	list.Remove(3)
func (l *SinglyLinkedList[T]) Remove(value T) {
	node := l.Find(value)
	if node == nil {
		return
	}
	if node == l.Head() {
		l.RemoveFirst()
		return
	}
	prev := l.Head()
	for prev.Next() != node {
		prev = prev.Next()
	}
	prev.next = node.Next()
	if node == l.Tail() {
		l.tail = prev
	}
	l.size--
}

// Returns a string representation of the list.
//
// Returns:
//   - string: Formatted string of elements.
//
// Example:
//
//	fmt.Println(list.String())
func (l *SinglyLinkedList[T]) String() string {
	if l.IsEmpty() {
		return "SinglyLinkedList: []"
	}
	result := "SinglyLinkedList: "
	current := l.Head()
	for {
		result += fmt.Sprintf("[%v]", current.Value())
		if !current.HasNext() {
			break
		}
		result += " -> "
		current = current.Next()
	}
	return result
}

// Inserts a new element at the specified index.
//
// Parameters:
//   - index: Position to insert at (0-based).
//   - value: Element to insert.
//
// Returns:
//   - error: If index is out of bounds.
//
// Example:
//
//	list.InsertAt(1, 42)
func (l *SinglyLinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > l.Size() {
		return fmt.Errorf("index %d out of bounds", index)
	}
	if index == 0 {
		l.Prepend(value)
		return nil
	}
	if index == l.Size() {
		l.Append(value)
		return nil
	}
	newNode := NewSinglyLinkedNode(value)
	current := l.Head()
	for range index - 1 {
		current = current.Next()
	}
	newNode.next = current.Next()
	current.next = newNode
	l.size++
	return nil
}

// Returns the node at the specified index.
//
// Parameters:
//   - index: Zero-based index.
//
// Returns:
//   - *SinglyLinkedNode[T]: Pointer to the node.
//   - error: If index is out of bounds.
//
// Example:
//
//	node, err := list.Get(0)
func (l *SinglyLinkedList[T]) Get(index int) (*SinglyLinkedNode[T], error) {
	if index < 0 || index >= l.Size() {
		return nil, fmt.Errorf("index %d out of bounds", index)
	}
	current := l.Head()
	for range index {
		current = current.Next()
	}
	return current, nil
}

// Updates the value of the node at the specified index.
//
// Parameters:
//   - index: Zero-based index.
//   - value: New value to set.
//
// Returns:
//   - error: If index is out of bounds.
//
// Example:
//
//	list.Set(1, 99)
func (l *SinglyLinkedList[T]) Set(index int, value T) error {
	node, err := l.Get(index)
	if err != nil {
		return fmt.Errorf("index %d out of bounds", index)
	}
	node.SetValue(value)
	return nil
}

// Reverses the order of nodes in the list.
//
// Example:
//
//	list.Reverse()
func (l *SinglyLinkedList[T]) Reverse() {
	var prev *SinglyLinkedNode[T]
	current := l.Head()
	l.tail = l.Head()
	for current != nil {
		next := current.Next()
		current.SetNext(prev)
		prev = current
		current = next
	}
	l.head = prev
}

// Returns true if the list contains the specified value.
//
// Parameters:
//   - value: Value to check.
//
// Returns:
//   - bool: true if value exists in the list; false otherwise.
//
// Example:
//
//	if list.Contains(10) {
//	   fmt.Println("Found 10")
//	}
func (l *SinglyLinkedList[T]) Contains(value T) bool {
	return l.Find(value) != nil
}

// Executes the given action for each element in the list.
//
// Parameters:
//   - action: Function to execute on each value.
//
// Example:
//
//	list.ForEach(func(v int) {
//	    fmt.Println(v)
//	})
func (l *SinglyLinkedList[T]) ForEach(action func(T)) {
	for current := l.Head(); current != nil; current = current.Next() {
		action(current.Value())
	}
}
