package stacknodes

type StackNode struct {
	value string
	prev  *StackNode
	next  *StackNode
}

type Stack struct {
	length int
	head   *StackNode
}

func NewStack() *Stack {
	return &Stack{
		length: 0,
		head:   nil,
	}
}

func (s *Stack) Push(node StackNode) {
	s.length++

	if s.head == nil {
		s.head = &node
		return
	}

	node.prev = s.head
	s.head = &node
}

func (s *Stack) Pop() {
	s.length--

	if s.length == 0 {
		s.head = nil
		return
	}

	head := s.head
	s.head = head.prev
}

func (s *Stack) Peek() *StackNode {
	return s.head
}
