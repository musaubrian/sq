package stackarr

import "testing"

func TestStackArrNewStack(t *testing.T) {
	s := NewStack()
	if len(s.Items) > 0 {
		t.Errorf("Expected an empty stack, got %d", len(s.Items))
	}
}

func TestStackArrPush(t *testing.T) {
	s := NewStack()
	items := []Item{"hello", "random", "okay"}

	for _, item := range items {
		s.Push(item)
	}

	if len(s.Items) == 0 {
		t.Errorf("Expected a non-empty stack got %d", len(s.Items))
	}

	if len(s.Items) != len(items) {
		t.Errorf("Expected a stack of size: %d, Got %d", len(items), len(s.Items))
	}

}

func TestStackArrPop(t *testing.T) {
	s := NewStack()

	prevCount := len(s.Items)
	s.Pop()
	if len(s.Items) != 0 {
		t.Errorf("Expected stack to be empty")
	}

	items := []Item{"hello", "random", "okay"}
	for _, item := range items {
		s.Push(item)
	}

	prevCount = len(s.Items)
	s.Pop()

	if len(s.Items) >= prevCount {
		t.Errorf("Expected stack to reduce in size")
	}
}
