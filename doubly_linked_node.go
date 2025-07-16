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

// Represents a node in a doubly linked list.
//
// Each node holds a value of type T and pointers to the next and previous nodes in
// the list. T must be a comparable type to allow equality checks.
type DoublyLinkedNode[T comparable] struct {
	value T
	next  *DoublyLinkedNode[T]
	prev  *DoublyLinkedNode[T]
}

// Creates and returns a new doubly linked node with the given value
//
// Parameters:
//   - value: The value to store in the node.
//
// Returns:
//   - *DoublyLinkedNode[T]: Pointer to the newly created node.
//
// Example:
//
//	node := list.NewDoublyLinkedNode[string]("hello")
func NewDoublyLinkedNode[T comparable](value T) *DoublyLinkedNode[T] {
	return &DoublyLinkedNode[T]{value: value}
}

// Updates the value stored in the node.
//
// Parameters:
//   - value: The new value to set.
//
// Example:
//
//	node.SetValue(100)
func (n *DoublyLinkedNode[T]) SetValue(value T) {
	n.value = value
}

// Returns the value stored in the node.
//
// Returns:
//   - T: The node’s value.
//
// Example:
//
//	v := node.Value()
func (n *DoublyLinkedNode[T]) Value() T {
	return n.value
}

// Updates the next pointer of the node.
//
// Parameters:
//   - next: Pointer to the next node.
//
// Example:
//
//	node.SetNext(nextNode)
func (n *DoublyLinkedNode[T]) SetNext(next *DoublyLinkedNode[T]) {
	n.next = next
}

// Returns the next node in the list.
//
// Returns:
//   - *DoublyLinkedNode[T]: Pointer to the next node or nil if none.
//
// Example:
//
//	next := node.Next()
func (n *DoublyLinkedNode[T]) Next() *DoublyLinkedNode[T] {
	return n.next
}

// Reports whether the node has a next node.
//
// Returns:
//   - bool: true if there is a next node; false otherwise.
//
// Example:
//
//	if node.HasNext() {
//	    fmt.Println("Next node exists")
//	}
func (n *DoublyLinkedNode[T]) HasNext() bool {
	return n.next != nil
}

// Updates the previous pointer of the node.
//
// Parameters:
//   - prev: Pointer to the previous node.
//
// Example:
//
//	node.SetPrev(prevNode)
func (n *DoublyLinkedNode[T]) SetPrev(prev *DoublyLinkedNode[T]) {
	n.prev = prev
}

// Returns the previous node in the list.
//
// Returns:
//   - *DoublyLinkedNode[T]: Pointer to the previous node or nil if none.
//
// Example:
//
//	prev := node.Prev()
func (n *DoublyLinkedNode[T]) Prev() *DoublyLinkedNode[T] {
	return n.prev
}

// Reports whether the node has a previous node.
//
// Returns:
//   - bool: true if there is a previous node; false otherwise.
//
// Example:
//
//	if node.HasPrev() {
//	    fmt.Println("Previous node exists")
//	}
func (n *DoublyLinkedNode[T]) HasPrev() bool {
	return n.prev != nil
}
