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

// Represents a node in a singly linked list, storing a value
// of type T and a pointer to the next node.
//
// T must be a comparable type to allow equality checks when needed.
type SinglyLinkedNode[T comparable] struct {
	value T
	next  *SinglyLinkedNode[T]
}

// Creates a new singly linked list node containing the given
// value and with no next node.
//
// Parameters:
//   - value: The value to store in the node.
//
// Returns:
//   - *SinglyLinkedNode[T]: A pointer to the newly created node.
//
// Example:
//
//	node := NewSinglyLinkedNode[string]("hello")
func NewSinglyLinkedNode[T comparable](value T) *SinglyLinkedNode[T] {
	return &SinglyLinkedNode[T]{value: value}
}

// Updates the value stored in the node.
//
// Parameters:
//   - value: The new value to set.
//
// Example:
//
//	node.SetValue(42)
func (n *SinglyLinkedNode[T]) SetValue(value T) {
	n.value = value
}

// Returns the value stored in the node.
//
// Returns:
//   - T: The value stored.
//
// Example:
//
//	v := node.Value()
func (n *SinglyLinkedNode[T]) Value() T {
	return n.value
}

// Sets the pointer to the next node in the list.
//
// Parameters:
//   - newNext: Pointer to the node that should follow this node.
//
// Example:
//
//	node.SetNext(nextNode)
func (n *SinglyLinkedNode[T]) SetNext(newNext *SinglyLinkedNode[T]) {
	n.next = newNext
}

// Returns the pointer to the next node in the list.
//
// Returns:
//   - *SinglyLinkedNode[T]: Pointer to the next node, or nil if none.
//
// Example:
//
//	next := node.Next()
func (n *SinglyLinkedNode[T]) Next() *SinglyLinkedNode[T] {
	return n.next
}

// Reports whether this node points to another one.
//
// Returns:
//   - bool: true if the node has a next node; false if next is nil.
//
// Example:
//
//	if node.HasNext() {
//	   fmt.Println("There is a next node")
//	}
func (n *SinglyLinkedNode[T]) HasNext() bool {
	return n.next != nil
}
