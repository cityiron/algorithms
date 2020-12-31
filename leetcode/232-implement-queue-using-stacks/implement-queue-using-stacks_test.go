package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	queue := Constructor()
	assert.Equal(t, true, queue.Empty())
	queue.Push(1)
	assert.Equal(t, false, queue.Empty())
	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, true, queue.Empty())
	queue.Push(3)
	queue.Push(2)
	queue.Push(5)
	assert.Equal(t, false, queue.Empty())
	assert.Equal(t, 3, queue.Peek())
	assert.Equal(t, 3, queue.Pop())
	assert.Equal(t, 2, queue.Pop())
}
