package link_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, "321", s.String())
	assert.Equal(t, 3, s.Size())
	assert.Equal(t, false, s.IsEmpty())
	pop := s.Pop()
	assert.Equal(t, 3, pop)
	assert.Equal(t, 2, s.Size())
	s.Clear()
	assert.Equal(t, 0, s.size)
}
