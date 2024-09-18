package stackarr

type Item string

type Stack struct {
	Items []Item
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item Item) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() {
	if !s.isValid() {
		return
	}

	s.Items = s.Items[:len(s.Items)-2]
}

func (s *Stack) Peek() Item {
	if !s.isValid() {
		return Item("")
	}

	return s.Items[0]
}

func (s *Stack) isValid() bool {
	if len(s.Items) > 0 {
		return true
	}
	return false
}
