package list

import "testing"

func TestSinglyLinkedNodeNewSinglyLinkedNode(t *testing.T) {
	node := NewSinglyLinkedNode(10)
	if node == nil {
		t.Fatal("NewSinglyLinkedNode returned nil")
	}
	if node.Value() != 10 {
		t.Errorf("expected value 10, got %v", node.Value())
	}
	if node.Next() != nil {
		t.Errorf("expected next to be nil, got %v", node.Next())
	}
}

func TestSinglyLinkedNodeSetValue(t *testing.T) {
	node := NewSinglyLinkedNode("initial")
	node.SetValue("updated")
	if node.Value() != "updated" {
		t.Errorf("expected value 'updated', got %v", node.Value())
	}
}

func TestSinglyLinkedNodeSetNextAndNext(t *testing.T) {
	node1 := NewSinglyLinkedNode(1)
	node2 := NewSinglyLinkedNode(2)
	node1.SetNext(node2)
	if node1.Next() != node2 {
		t.Errorf("expected next node to be %v, got %v", node2, node1.Next())
	}
}

func TestSinglyLinkedNodeHasNext(t *testing.T) {
	node1 := NewSinglyLinkedNode(1)
	if node1.HasNext() {
		t.Error("expected HasNext to be false for node with nil next")
	}
	node2 := NewSinglyLinkedNode(2)
	node1.SetNext(node2)
	if !node1.HasNext() {
		t.Error("expected HasNext to be true after setting next node")
	}
}
