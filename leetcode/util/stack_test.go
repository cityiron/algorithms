package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	stack := new(Stack)
	assert.Equal(t, 0, stack.Size())
	stack.Print()
	assert.Nil(t, stack.Top())
	err := stack.Pop()
	assert.Error(t, ErrorStackEmpty, err)
	stack.Push("{", "[", "]")
	stack.Print()
	assert.Equal(t, 3, stack.Size())
	assert.Equal(t, "]", stack.Top().(string))
	assert.Equal(t, "[", stack.Get(1))
	err = stack.Set(5, "}")
	assert.Error(t, ErrorOutOfIndex, err)
	stack.Push(")")
	err = stack.Set(3, "}")
	assert.NoError(t, err)
	assert.Equal(t, 4, stack.Size())
	err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, stack.Size())
	stack.Print()
	v, err := stack.PopAndGet()
	assert.NoError(t, err)
	assert.Equal(t, "[", v)
}
