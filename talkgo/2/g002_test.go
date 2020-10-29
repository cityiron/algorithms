package talkgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	assert.Equal(t, Point([]int{}), 0)
	assert.Equal(t, Point([]int{8, 8, 9, 10, 1, 2, 3, 4, 5, 6}), 4)
	assert.Equal(t, Point([]int{6, 5, 3, 2, 1, 13, 10, 9, 8, 6}), 5)
	assert.Equal(t, Point([]int{4, 5, 6, 7, 8, 10, -11, -9, -8, -5, -4, -1, 0}), 6)
}

var testCases = []struct {
	array  []int
	result float32
	err    error
	desc   string
}{
	{
		array:  []int{},
		result: 0,
		err:    illegalArray,
		desc:   "empty result",
	},
	{
		array:  []int{8, 11, 12, 13, 14, 16, 19, -1, 2, 4, 5, 6},
		result: 9.5,
		desc:   "even",
	},
	{
		array:  []int{8, 8, 9, 10, 11, 12, 15, 17, 1, 2, 3, 4, 5, 6, 7},
		result: 8,
		desc:   "odd",
	},
}

func BenchmarkMiddleNum(b *testing.B) {
	for _, test := range testCases {
		test := test
		b.Run(test.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := MiddleNum(test.array)
				assert.Equal(b, err, test.err)
				assert.Equal(b, result, test.result)
			}
		})
	}
}

func Test(t *testing.T) {
	result, err := MiddleNum(testCases[1].array)
	assert.Equal(t, err, testCases[1].err)
	assert.Equal(t, result, testCases[1].result)
}
