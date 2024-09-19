package queuenodes

import "testing"

func TestNewNodeQ(t *testing.T) {
	q := NewQ()

	if q.head != nil || q.tail != nil {
		t.Errorf("Expected both tail and head to be nil, Got: [head]: %v\n[tail]: %v", q.head, q.tail)
	}
}
func TestNodeQEnq(t *testing.T) {
	q := NewQ()
	node := QNode{
		Value: "hello",
	}
	q.Enq(node)

	if q.Peek().Value != node.Value {
		t.Errorf("Expected head to be: %v, Got: %v", q.Peek().Value, node)
	}

	if q.head.Value != q.tail.Value {
		t.Errorf("Expected head and tail to be the same <%v>, Got: [head]: %v\n[tail]: %v", node, q.head, q.tail)
	}

	nodes := []QNode{
		{Value: "world"},
		{Value: "last"},
	}

	for _, node := range nodes {
		q.Enq(node)
	}
	expected := nodes[len(nodes)-1]
	if q.tail.Value != expected.Value {
		t.Errorf("Expected tail to be: %v, Got: %v", q.tail.Value, expected.Value)
	}

}
func TestNodeQDq(t *testing.T) {
	q := NewQ()

	nodes := []QNode{
		{Value: "hello"},
		{Value: "world"},
		{Value: "last"},
	}

	for _, node := range nodes {
		q.Enq(node)
	}

	for range 2 {
		q.Dq()
	}
	expected := nodes[len(nodes)-1]
	if q.Peek().Value != expected.Value {
		t.Errorf("Expected head to be: %v, Got: %v", expected, q.Peek())
	}

}
func TestNodeQPeek(t *testing.T) {
	q := NewQ()

	nodes := []QNode{
		{Value: "hello"},
		{Value: "world"},
		{Value: "last"},
	}

	for _, node := range nodes {
		q.Enq(node)
	}

	if head := q.Peek(); head == nil {
		t.Errorf("Expected head to be %v, Got: %v", nodes[0], q.head)
	}
}
