package list

import "testing"

func TestDoublyLinkedNodeNewDoublyLinkedNode(t *testing.T) {
	node := NewDoublyLinkedNode(42)
	if node == nil {
		t.Fatal("expected NewDoublyLinkedNode to return non-nil node")
	}
	if node.Value() != 42 {
		t.Errorf("expected value 42, got %v", node.Value())
	}
	if node.Next() != nil {
		t.Error("expected Next to be nil for new node")
	}
	if node.Prev() != nil {
		t.Error("expected Prev to be nil for new node")
	}
}

func TestDoublyLinkedNodeSetAndGetValue(t *testing.T) {
	node := NewDoublyLinkedNode("hello")
	node.SetValue("world")
	if node.Value() != "world" {
		t.Errorf("expected value 'world', got %v", node.Value())
	}
}

func TestDoublyLinkedNodeSetAndGetNext(t *testing.T) {
	node1 := NewDoublyLinkedNode(1)
	node2 := NewDoublyLinkedNode(2)
	node1.SetNext(node2)
	if node1.Next() != node2 {
		t.Error("expected Next to return node2")
	}
	if !node1.HasNext() {
		t.Error("expected HasNext to return true")
	}
}

func TestDoublyLinkedNodeSetAndGetPrev(t *testing.T) {
	node1 := NewDoublyLinkedNode(1)
	node0 := NewDoublyLinkedNode(0)
	node1.SetPrev(node0)
	if node1.Prev() != node0 {
		t.Error("expected Prev to return node0")
	}
	if !node1.HasPrev() {
		t.Error("expected HasPrev to return true")
	}
}

func TestDoublyLinkedNodeHasNextAndHasPrevNil(t *testing.T) {
	node := NewDoublyLinkedNode(5)
	if node.HasNext() {
		t.Error("expected HasNext to return false for node with nil next")
	}
	if node.HasPrev() {
		t.Error("expected HasPrev to return false for node with nil prev")
	}
}
