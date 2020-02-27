package link_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedQueue(t *testing.T) {
	q := NewLinkQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	assert.Equal(t, "1234", q.String())
	pop := q.Pop()
	assert.Equal(t, 1, pop)
	assert.Equal(t, 4, q.Size())
	deq := q.Dequeue()
	assert.Equal(t, 1, deq)
	assert.Equal(t, 3, q.Size())
	assert.Equal(t, "234", q.String())
}