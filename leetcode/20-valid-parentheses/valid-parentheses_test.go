package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	index  int
	input  string
	result bool
}{
	{
		index:  1,
		input:  "()",
		result: true,
	},
	{
		index:  2,
		input:  "()[]{}",
		result: true,
	},
	{
		index:  3,
		input:  "(]",
		result: false,
	},
	{
		index:  4,
		input:  "([)]",
		result: false,
	},
	{
		index:  5,
		input:  "{[]}",
		result: true,
	},
	{
		index:  6,
		input:  "({[]})",
		result: true,
	},
	{
		index:  7,
		input:  "({[[][]{}]})",
		result: true,
	},
}

func Test_isValid(t *testing.T) {
	for _, tc := range tests {
		i := tc.input
		r := tc.result
		if !assert.Equal(t, r, isValid(i)) {
			t.Logf("current is %d", tc.index)
		}
	}
}

func Test_isValid2(t *testing.T) {
	for _, tc := range tests {
		i := tc.input
		r := tc.result
		if !assert.Equal(t, r, isValid2(i)) {
			t.Logf("current is %d", tc.index)
		}
	}
}

func Test_isValid3(t *testing.T) {
	for _, tc := range tests {
		i := tc.input
		r := tc.result
		if !assert.Equal(t, r, isValid3(i)) {
			t.Logf("current is %d", tc.index)
		}
	}
}
