package queuearr

import "testing"

func TestNewQ(t *testing.T) {
	val := 10
	q := NewQ(val)

	if q.cap != val {
		t.Errorf("Expected Queue Cap to be <%d>, Got: <%d>", val, q.cap)
	}
	if len(q.Items) != 0 {
		t.Errorf("Expected an empty Queue, Got <%d> items", len(q.Items))
	}
}
func TestEnQ(t *testing.T) {
	q := NewQ(5)
	items := []QItem{"hello", "random", "jambo"}

	for _, item := range items {
		q.Enq(item)
	}

	if len(q.Items) != len(items) {
		t.Errorf("Expected queue of <%d> items, Got <%d> items", len(items), len(q.Items))
	}

}
func TestDQ(t *testing.T) {
	q := NewQ(5)
	items := []QItem{"hello", "random", "jambo"}

	for _, item := range items {
		q.Enq(item)
	}
	q.DQ()
	if len(q.Items) != len(items)-1 {
		t.Errorf("Expected <%d> items, Got <%d> items", len(items)-1, len(q.Items))
	}
}

func TestPeek(t *testing.T) {
	q := NewQ(2)

	items := []QItem{"hello", "random", "jambo"}

	for _, item := range items {
		q.Enq(item)
	}

	if res := q.Peek(); res != items[0] {
		t.Errorf("Expected <%v>, Got: <%v>", items[0], res)

	}

}
