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

// Represents a generic circular doubly linked list.
//
// T must be a comparable type to enable equality-based operations.
type CircularDoublyLinkedList[T comparable] struct {
	tail *DoublyLinkedNode[T]
	size int
}

// Creates and returns a new empty circular doubly linked list.
//
// Returns:
//   - *CircularDoublyLinkedList[T]: Pointer to a new empty list.
//
// Example:
//
//	list := list.NewCircularDoublyLinkedList[string]()
func NewCircularDoublyLinkedList[T comparable]() *CircularDoublyLinkedList[T] {
	return &CircularDoublyLinkedList[T]{}
}

// Returns the first node of the list.
//
// Returns:
//   - *DoublyLinkedNode[T]: Pointer to the head node or nil if the list is empty.
//
// Example:
//
//	head := list.Head()
func (l *CircularDoublyLinkedList[T]) Head() *DoublyLinkedNode[T] {
	if l.Tail() == nil {
		return nil
	}
	return l.Tail().Next()
}

// Returns the last node of the list.
//
// Returns:
//   - *DoublyLinkedNode[T]: Pointer to the tail node or nil if the list is empty.
//
// Example:
//
//	tail := list.Tail()
func (l *CircularDoublyLinkedList[T]) Tail() *DoublyLinkedNode[T] {
	return l.tail
}

// Returns the number of elements in the list.
//
// Returns:
//   - int: Number of elements.
//
// Example:
//
//	fmt.Println(list.Size()) // 3
func (l *CircularDoublyLinkedList[T]) Size() int {
	return l.size
}

// Reports whether the list contains no elements.
//
// Returns:
//   - bool: true if the list is empty, false otherwise.
//
// Example:
//
//	fmt.Println(list.IsEmpty()) // true
func (l *CircularDoublyLinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

// Removes all elements from the list, resetting it to empty.
//
// Example:
//
//	list.Clear()
//	fmt.Println(list.IsEmpty()) // true
func (l *CircularDoublyLinkedList[T]) Clear() {
	l.tail = nil
	l.size = 0
}

// Inserts a new element at the beginning of the list.
//
// Parameters:
//   - value: The value to insert.
//
// Example:
//
//	list.Prepend(5)
func (l *CircularDoublyLinkedList[T]) Prepend(value T) {
	newNode := NewDoublyLinkedNode(value)
	if l.IsEmpty() {
		newNode.next = newNode
		newNode.prev = newNode
		l.tail = newNode
	} else {
		head := l.Head()
		newNode.next = head
		newNode.prev = l.Tail()
		head.prev = newNode
		l.Tail().next = newNode
	}
	l.size++
}

// Inserts a new element at the end of the list.
//
// Parameters:
//   - value: The value to insert.
//
// Example:
//
//	list.Append(10)
func (l *CircularDoublyLinkedList[T]) Append(value T) {
	l.Prepend(value)
	l.tail = l.Tail().Next()
}

// Searches for the first node containing the specified value.
//
// Parameters:
//   - value: The value to search for.
//
// Returns:
//   - *DoublyLinkedNode[T]: Pointer to the node if found, or nil otherwise.
//
// Example:
//
//	node := list.Find(7)
func (l *CircularDoublyLinkedList[T]) Find(value T) *DoublyLinkedNode[T] {
	if l.IsEmpty() {
		return nil
	}
	current := l.Head()
	for range l.Size() {
		if current.Value() == value {
			return current
		}
		current = current.Next()
	}
	return nil
}

// Removes the first element from the list.
//
// If the list is empty, the operation has no effect.
//
// Example:
//
//	list.RemoveFirst()
func (l *CircularDoublyLinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		l.Clear()
		return
	}
	head := l.Head()
	newHead := head.Next()
	l.Tail().next = newHead
	newHead.prev = l.Tail()
	l.size--
}

// Removes the last element from the list.
//
// If the list is empty, the operation has no effect.
//
// Example:
//
//	list.RemoveLast()
func (l *CircularDoublyLinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		l.Clear()
		return
	}
	last := l.Tail()
	prev := last.Prev()
	prev.next = l.Head()
	l.Head().prev = prev
	l.tail = prev
	l.size--
}

// Deletes the first occurrence of the specified value from the list.
//
// If the value is not found, the list remains unchanged.
//
// Parameters:
//   - value: The value to remove.
//
// Example:
//
//	list.Remove(10)
func (l *CircularDoublyLinkedList[T]) Remove(value T) {
	if l.IsEmpty() {
		return
	}
	current := l.Head()
	for range l.Size() {
		if current.Value() == value {
			if l.Size() == 1 {
				l.Clear()
				return
			}
			prev := current.Prev()
			next := current.Next()
			prev.next = next
			next.prev = prev
			if current == l.Tail() {
				l.tail = prev
			}
			l.size--
			return
		}
		current = current.Next()
	}
}

// Returns a string representation of the list.
//
// Returns:
//   - string: A human-readable string representation.
//
// Example:
//
//	fmt.Println(list.String()) // CircularSinglyLinkedList: [1] -> [2] -> [3]
func (l *CircularDoublyLinkedList[T]) String() string {
	if l.IsEmpty() {
		return "CircularDoublyLinkedList: []"
	}
	result := "CircularDoublyLinkedList: "
	current := l.Head()
	for i := range l.Size() {
		result += fmt.Sprintf("[%v]", current.Value())
		if i < l.Size()-1 {
			result += " <-> "
		}
		current = current.Next()
	}
	return result
}

// Inserts a new element at the specified index.
//
// Parameters:
//   - index: Position at which to insert (0-based).
//   - value: The value to insert.
//
// Returns:
//   - error: If index is out of bounds.
//
// Example:
//
//	err := list.InsertAt(2, 99)
func (l *CircularDoublyLinkedList[T]) InsertAt(index int, value T) error {
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
	current := l.Head()
	for range index {
		current = current.Next()
	}
	newNode := NewDoublyLinkedNode(value)
	prev := current.Prev()
	prev.next = newNode
	newNode.prev = prev
	newNode.next = current
	current.prev = newNode
	l.size++
	return nil
}

// Retrieves the node at the specified index.
//
// Parameters:
//   - index: Position of the node (0-based).
//
// Returns:
//   - *SinglyLinkedNode[T]: Pointer to the node.
//   - error: If index is out of bounds.
//
// Example:
//
//	node, err := list.Get(1)
func (l *CircularDoublyLinkedList[T]) Get(index int) (*DoublyLinkedNode[T], error) {
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
//   - index: Position of the node (0-based).
//   - value: New value to set.
//
// Returns:
//   - error: If index is out of bounds.
//
// Example:
//
//	err := list.Set(0, 42)
func (l *CircularDoublyLinkedList[T]) Set(index int, value T) error {
	node, err := l.Get(index)
	if err != nil {
		return err
	}
	node.value = value
	return nil
}

// Reverses the order of elements in the list.
//
// Example:
//
//	list.Reverse()
func (l *CircularDoublyLinkedList[T]) Reverse() {
	if l.IsEmpty() || l.Size() == 1 {
		return
	}
	current := l.Head()
	originalHead := current
	for range l.size {
		current.next, current.prev = current.Prev(), current.Next()
		current = current.Prev()
	}
	l.tail = originalHead
}

// Reports whether the list contains the specified value.
//
// Parameters:
//   - value: The value to search for.
//
// Returns:
//   - bool: true if found, false otherwise.
//
// Example:
//
//	fmt.Println(list.Contains(5)) // true
func (l *CircularDoublyLinkedList[T]) Contains(value T) bool {
	return l.Find(value) != nil
}

// Applies a provided function to each element in the list.
//
// Parameters:
//   - action: A function to apply to each element.
//
// Example:
//
//	list.ForEach(func(v int) { fmt.Println(v) })
func (l *CircularDoublyLinkedList[T]) ForEach(action func(T)) {
	if l.IsEmpty() {
		return
	}
	current := l.Head()
	for range l.Size() {
		action(current.Value())
		current = current.Next()
	}
}
