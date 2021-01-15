package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type Step struct {
	value  int
	result int
}

var tests = []struct {
	key   int
	input []int
	step  []Step
}{
	{
		key: 5,
		input: []int{
			9, 2, 4, 6, 7, 2, 5, 3,
		},
		step: []Step{
			{
				value:  7,
				result: 5,
			},
			{
				value:  4,
				result: 5,
			},
			{
				value:  11,
				result: 6,
			},
			{
				value:  7,
				result: 7,
			},
			{
				value:  8,
				result: 7,
			},
		},
	},
	{
		key:   1,
		input: []int{},
		step: []Step{
			{
				value:  -3,
				result: -3,
			},
			{
				value:  -2,
				result: -2,
			},
			{
				value:  -4,
				result: -2,
			},
			{
				value:  0,
				result: 0,
			},
			{
				value:  4,
				result: 4,
			},
		},
	},
	{
		key:   2,
		input: []int{0},
		step: []Step{
			{
				value:  -1,
				result: -1,
			},
			{
				value:  1,
				result: 0,
			},
			{
				value:  -2,
				result: 0,
			},
			{
				value:  -4,
				result: 0,
			},
			{
				value:  3,
				result: 1,
			},
		},
	},
	{
		key:   3,
		input: []int{5, -1},
		step: []Step{
			{
				value:  2,
				result: -1,
			},
			{
				value:  1,
				result: 1,
			},
			{
				value:  -1,
				result: 1,
			},
			{
				value:  3,
				result: 2,
			},
			{
				value:  4,
				result: 3,
			},
		},
	},
	{
		key:   7,
		input: []int{-10, 1, 3, 1, 4, 10, 3, 9, 4, 5, 1},
		step: []Step{
			{
				value:  3,
				result: 3,
			},
			{
				value:  2,
				result: 3,
			},
			{
				value:  3,
				result: 3,
			},
			{
				value:  1,
				result: 3,
			},
			{
				value:  2,
				result: 3,
			},
			{
				value:  4,
				result: 3,
			},
			{
				value:  5,
				result: 4,
			},
			{
				value:  4,
				result: 4,
			},
		},
	},
}

func TestSort(t *testing.T) {
	i := []int{
		9, 2, 4, 6, 7, 2, 5, 3,
	}
	sort.Ints(i)

	t.Log(i) // [2 2 3 4 5 6 7 9]
}

func TestKthLargest_SortPro(t *testing.T) {
	for _, v := range tests {
		cs := Constructor_SortPro(v.key, v.input)
		for _, v := range v.step {
			assert.Equal(t, v.result, cs.Add(v.value))
		}
	}
}

func TestKthLargest_Sort(t *testing.T) {
	for _, v := range tests {
		cs := Constructor_Sort(v.key, v.input)
		for _, v := range v.step {
			assert.Equal(t, v.result, cs.Add(v.value))
		}
	}
}

func TestKthLargest_Sort2(t *testing.T) {
	for _, v := range tests {
		cs := Constructor_Sort2(v.key, v.input)
		for _, v := range v.step {
			assert.Equal(t, v.result, cs.Add(v.value))
		}
	}
}

func TestArray(t *testing.T) {
	array := []int{
		9, 2, 4, 6, 7, 2, 5, 3,
	}

	sort.Ints(array)

	t.Log(array[1:1])
	t.Log(array[1:])
}

func TestIntHeap(t *testing.T) {
	for _, v := range tests {
		t.Logf(" ====== start ===== top : %d, init array : %v", v.key, v.input)
		cs := Constructor(v.key, v.input)
		for _, v := range v.step {
			b := assert.Equal(t, v.result, cs.Add(v.value))
			if !b {
				t.Logf(" ====== fail ===== value : %d, array : %v", v.value, cs.heap)
			}
		}
	}
}
