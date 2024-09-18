package queuearr

import (
	"fmt"
	"log/slog"
)

type QItem string

type Queue struct {
	cap   int
	Items []QItem
}

func NewQ(cap int) *Queue {
	return &Queue{
		cap: cap,
	}
}

func (q *Queue) Enq(item QItem) {
	if q.isFull() {
		slog.Error(fmt.Sprintf("Queue is full, Ignoring <%v>", item))
		return
	}

	q.Items = append(q.Items, item)
}

func (q *Queue) DQ() {
	if q.isEmpty() {
		slog.Error("Queue is empty")
		return
	}
	q.Items = q.Items[len(q.Items)-2:]

}

func (q *Queue) Peek() QItem {
	if q.isEmpty() {
		return QItem("")
	}
	return q.Items[0]
}

func (q *Queue) isFull() bool {
	if len(q.Items) >= q.cap {
		return true
	}
	return false

}

func (q *Queue) isEmpty() bool {
	if len(q.Items) == 0 {
		return true
	}
	return false
}
