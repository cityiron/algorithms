package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoArray(t *testing.T) {
	stack := Constructor()
	assert.Equal(t, true, stack.Empty())
	stack.Push(1)
	assert.Equal(t, false, stack.Empty())
	assert.Equal(t, 1, stack.Top())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, true, stack.Empty())
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Top())
	stack.Push(3)
	assert.Equal(t, 3, stack.Top())
	assert.Equal(t, false, stack.Empty())
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, true, stack.Empty())
}

func TestOneQueue(t *testing.T) {
	stack := Constructor1()
	assert.Equal(t, true, stack.Empty())
	stack.Push(1)
	assert.Equal(t, false, stack.Empty())
	assert.Equal(t, 1, stack.Top())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, true, stack.Empty())
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Top())
	stack.Push(3)
	assert.Equal(t, 3, stack.Top())
	assert.Equal(t, false, stack.Empty())
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, true, stack.Empty())
}

func TestOther(t *testing.T) {
	stack := Constructor2()
	assert.Equal(t, true, stack.Empty())
	stack.Push(1)
	assert.Equal(t, false, stack.Empty())
	assert.Equal(t, 1, stack.Top())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, true, stack.Empty())
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Top())
	stack.Push(3)
	assert.Equal(t, 3, stack.Top())
	assert.Equal(t, false, stack.Empty())
}
