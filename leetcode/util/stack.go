package util

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrorStackEmpty = errors.New("stack is empty.")
	ErrorOutOfIndex = errors.New("out of index.")
)

type Stack struct {
	element []interface{}
}

func (s *Stack) Push(value ...interface{}) {
	s.element = append(s.element, value...)
}

func (s *Stack) Top() (value interface{}) {
	if s.Size() > 0 {
		return s.element[s.Size()-1]
	}

	return nil
}

func (s *Stack) Pop() (err error) {
	if s.Size() > 0 {
		s.element = s.element[:s.Size()-1]
		return nil
	}

	return ErrorStackEmpty
}

func (s *Stack) PopAndGet() (value interface{}, err error) {
	if s.Size() > 0 {
		s.element = s.element[:s.Size()-1]
		return s.element[s.Size()-1], nil
	}

	return nil, ErrorStackEmpty
}

func (s *Stack) Swap(o *Stack) {
	if s.Size() == 0 && o.Size() == 0 {
		return
	}

	if s.Size() == 0 {
		s.element = o.element
		o.element = nil
		return
	}

	if o.Size() == 0 {
		o.element = s.element
		s.element = nil
		return
	}

	s.element, o.element = o.element, s.element
}

func (s *Stack) Set(idx int, value interface{}) (err error) {
	if idx >= 0 && s.Size() > 0 && s.Size() > idx {
		s.element[idx] = value
		return nil
	}

	return ErrorOutOfIndex
}

func (s *Stack) Get(idx int) (value interface{}) {
	if idx >= 0 && s.Size() > 0 && s.Size() >= idx {
		return s.element[idx]
	}

	return nil
}

// Size get stack element count.
func (s *Stack) Size() int {
	return len(s.element)
}

// Empty is stack is no element.
func (s *Stack) Empty() bool {
	if s.element == nil || s.Size() == 0 {
		return true
	}

	return false
}

func (s *Stack) Print() {
	for i := len(s.element) - 1; i >= 0; i-- {
		fmt.Println(i, "==>", s.element[i])
	}
}
