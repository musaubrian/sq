package queuenodes

// I visualize queues inserting to the right
// tail[p| Val] <- [p| Val] <- [p| Val]head

type QNode struct {
	Value string
	Prev  *QNode
}

type Queue struct {
	len  int
	head *QNode
	tail *QNode
}

func NewQ() *Queue {
	return &Queue{
		len: 0,
	}
}

func (q *Queue) Enq(node QNode) {
	q.len++
	if q.head == nil {
		q.head = &node
		q.tail = &node
	}
	q.tail.Prev = &node
	q.tail = &node
}

func (q *Queue) Dq() {
	q.len--
	q.head = q.head.Prev

}
func (q *Queue) Peek() *QNode {
	return q.head
}
