# List - Generic Linked List Implementations in Go

A fully featured generic **Linked List** package in Go supporting **singly linked**, **doubly linked**, and their **circular** variants.

This package provides dynamic data structures for efficient insertion, deletion, traversal, and manipulation of elements in linear and circular forms.

---

## Table of Contents

- [List - Generic Linked List Implementations in Go](#list---generic-linked-list-implementations-in-go)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Usage](#usage)
  - [Running Tests](#running-tests)
  - [Design Notes](#design-notes)
  - [Examples](#examples)
    - [CircularSinglyLinkedList](#circularsinglylinkedlist)
    - [CircularDoublyLinkedList](#circulardoublylinkedlist)
  - [Author](#author)
  - [License](#license)
  - [Contact](#contact)

---

## Features

- **Generic**: works with any comparable type (`List[T comparable]` in Go 1.18+).

- Supports multiple list types:

  - `SinglyLinkedList[T]`
  - `DoublyLinkedList[T]`
  - `CircularSinglyLinkedList[T]`
  - `CircularDoublyLinkedList[T]`

- Core list operations:

  - `Append(T)` — add an element at the end
  - `Prepend(T)` — add an element at the start
  - `InsertAt(index, T)` — insert at a specific index
  - `Remove(T)` — remove a specific value
  - `RemoveFirst()`, `RemoveLast()`
  - `Get(index)`, `Set(index, T)`
  - `Find(T)`, `Contains(T)`
  - `Clear()` — empties the list
  - `Reverse()` — reverses the order of elements in-place
  - `ForEach(func(T))` — iterate over all elements
  - `ToSlice() []T` — returns a slice copy of list elements
  - `String() string` — human-readable representation

- Handles edge cases gracefully (empty list operations are safe).

- Fully documented using GoDoc comments for easy browsing on `pkg.go.dev`.

---

## Usage

```go
package main

import (
	"fmt"
	"your/module/path/list" // replace with your actual import path
)

func main() {
	// Example with DoublyLinkedList
	dlist := list.NewDoublyLinkedList[int]()
	dlist.Append(10)
	dlist.Prepend(5)
	dlist.InsertAt(1, 7)
	fmt.Println(dlist) // Output: DoublyLinkedList: [5] ↔ [7] ↔ [10]

	// Remove and reverse
	dlist.Remove(7)
	dlist.Reverse()
	fmt.Println(dlist) // Output: DoublyLinkedList: [10] ↔ [5]
}
```

---

## Running Tests

This package includes **comprehensive unit tests** using Go’s `testing` package.

To run all tests:

```bash
go test ./list -v
```

This will verify:

- Core operations (append, prepend, insert, remove).
- Edge cases (empty lists, single-element lists).
- Circular and doubly linked behavior.
- Utility methods like `Reverse`, `ToSlice`, `Contains`.

For test coverage:

```bash
go test ./list -cover
```

---

## Design Notes

- **Internals**:

  - Singly linked lists use `SinglyLinkedNode[T]` nodes.
  - Doubly linked lists use `DoublyLinkedNode[T]` nodes.
  - Circular variants link tail nodes back to head nodes for continuous iteration.
- **Generics**: all list types use Go 1.18+ type parameters (`T comparable`).
- **Error Handling**: `Get` and `Set` return idiomatic Go errors on out-of-bounds indexes.
- **String Representations**:

  - `SinglyLinkedList`: `[A] -> [B] -> [C]`
  - `DoublyLinkedList`: `[A] ↔ [B] ↔ [C]`
  - Circular variants indicate cycles naturally.

---

## Examples

### CircularSinglyLinkedList

```go
clist := list.NewCircularSinglyLinkedList[int]()
clist.Append(1)
clist.Append(2)
clist.Append(3)
fmt.Println(clist) // CircularSinglyLinkedList: [1] -> [2] -> [3]

clist.Reverse()
fmt.Println(clist) // CircularSinglyLinkedList: [3] -> [2] -> [1]
```

### CircularDoublyLinkedList

```go
cdlist := list.NewCircularDoublyLinkedList[string]()
cdlist.Append("A")
cdlist.Append("B")
cdlist.Prepend("Start")
fmt.Println(cdlist) // CircularDoublyLinkedList: [Start] <-> [A] <-> [B]

cdlist.RemoveLast()
fmt.Println(cdlist) // CircularDoublyLinkedList: [Start] <-> [A]
```

---

## Author

trigologiaa

---

## License

This project is released under the MIT License. You’re free to use, modify, and distribute.

---

## Contact

For questions, suggestions, or contributions: open an issue or contact the author.