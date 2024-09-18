package stacknodes

import "testing"

func TestNewStackNode(t *testing.T) {
	s := NewStack()
	if s.head != nil && s.length != 0 {
		t.Errorf("Expected stack to be empty")
	}
}

func TestStackNodePush(t *testing.T) {
	s := NewStack()
	node := StackNode{}
	s.Push(node)

	if head := s.Peek(); head == nil || s.length == 0 {
		t.Errorf("Expected <head> to be: %v, Got: %v", node, head)
	}

	s.Pop()

}

func TestStackNodePop(t *testing.T) {
	s := NewStack()
	nodes := []StackNode{
		{value: "hello"},
		{value: "random"},
	}

	for _, node := range nodes {
		s.Push(node)
	}
	secondLastNode := s.Peek().prev
	s.Pop()
	if head := s.Peek(); head.value != secondLastNode.value {
		t.Errorf("Expected new head to be: %v, Got: %v", secondLastNode, head)
	}

	for _, node := range nodes {
		s.Push(node)
	}

	for range len(nodes) + 1 {
		s.Pop()
	}
	if s.length != 0 {
		t.Errorf("Expected stack to be empty")
	}
}

func TestStackNodePeek(t *testing.T) {
	s := NewStack()
	if head := s.Peek(); head != nil {
		t.Errorf("Expected to find no head found: %v", head)
	}

	node := StackNode{}
	s.Push(node)
	if head := s.Peek(); head == nil {
		t.Errorf("Expected <head> to be: %v, Got %v", node, head)

	}
}
