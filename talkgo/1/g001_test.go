package talkgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	sum    int
	array  []int
	result []int
	desc   string
}{
	{
		sum:    10,
		array:  []int{},
		result: []int{},
		desc:   "empty result",
	},
	{
		sum:    10,
		array:  []int{-1, 2, 4, 5, 6, 8, 11, 12, 13, 14, 16, 19},
		result: []int{-1, 11},
		desc:   "left side",
	},
	{
		sum:    5,
		array:  []int{-11, -9, -5, 2, 4, 5, 6, 8, 11, 12, 13, 14, 16, 19},
		result: []int{-11, 16},
		desc:   "two side",
	},
	{
		sum:    24,
		array:  []int{-11, -9, -5, 2, 4, 5, 6, 8, 11, 12, 13, 14, 16, 19},
		result: []int{5, 19},
		desc:   "right side",
	},
	{
		sum:    0,
		array:  []int{-11, -9, -5, 2, 4, 5, 6, 8, 11, 12, 13, 14, 16, 19},
		result: []int{-11, 11},
		desc:   "zero",
	},
	{
		sum:    -3,
		array:  []int{-11, -9, -5, 2, 4, 5, 6, 8, 11, 12, 13, 14, 16, 19},
		result: []int{-11, 8},
		desc:   "negative",
	},
	{
		sum: 98,
		array: []int{-11, -9, -5, 1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 16, 19, 22, 23, 24, 25, 29, 31, 33, 34, 35, 36,
			37, 39, 40, 41, 43, 45, 46, 47, 48, 50, 51, 52, 53, 54, 56, 57, 59, 60, 61, 62, 63, 64, 66, 70, 78, 89, 90},
		result: []int{34, 64},
		desc:   "more",
	},
}

//pkg: algorithms/talkgo/1
//BenchmarkFindTwoSumMap
//BenchmarkFindTwoSumMap/empty_result
//BenchmarkFindTwoSumMap/empty_result-4         	  438326	      2477 ns/op
//BenchmarkFindTwoSumMap/left_side
//BenchmarkFindTwoSumMap/left_side-4            	  360903	      4123 ns/op
//BenchmarkFindTwoSumMap/two_side
//BenchmarkFindTwoSumMap/two_side-4             	  354012	      3779 ns/op
//BenchmarkFindTwoSumMap/right_side
//BenchmarkFindTwoSumMap/right_side-4           	  325746	      4469 ns/op
//BenchmarkFindTwoSumMap/zero
//BenchmarkFindTwoSumMap/zero-4                 	  350618	      3610 ns/op
//BenchmarkFindTwoSumMap/negative
//BenchmarkFindTwoSumMap/negative-4             	  350402	      3508 ns/op
//PASS
func BenchmarkFindTwoSumMap(b *testing.B) {
	for _, test := range testCases {
		test := test
		b.Run(test.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result := TwoSumMap(test.sum, test.array)
				assert.Equal(b, result, test.result)
			}
		})
	}
}

//pkg: algorithms/talkgo/1
//BenchmarkFindTwoSum2For
//BenchmarkFindTwoSum2For/empty_result
//BenchmarkFindTwoSum2For/empty_result-4         	  941133	      1571 ns/op
//BenchmarkFindTwoSum2For/left_side
//BenchmarkFindTwoSum2For/left_side-4            	  829220	      1470 ns/op
//BenchmarkFindTwoSum2For/two_side
//BenchmarkFindTwoSum2For/two_side-4             	  821710	      1581 ns/op
//BenchmarkFindTwoSum2For/right_side
//BenchmarkFindTwoSum2For/right_side-4           	  706478	      1783 ns/op
//BenchmarkFindTwoSum2For/zero
//BenchmarkFindTwoSum2For/zero-4                 	  680068	      1714 ns/op
//BenchmarkFindTwoSum2For/negative
//BenchmarkFindTwoSum2For/negative-4             	  826993	      1490 ns/op
//PASS
func BenchmarkFindTwoSum2For(b *testing.B) {
	for _, test := range testCases {
		test := test
		b.Run(test.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result := TwoSum2For(test.sum, test.array)
				assert.Equal(b, result, test.result)
			}
		})
	}
}

//pkg: algorithms/talkgo/1
//BenchmarkTwoSum2Point
//BenchmarkTwoSum2Point/empty_result
//BenchmarkTwoSum2Point/empty_result-4         	  632310	      1825 ns/op
//BenchmarkTwoSum2Point/left_side
//BenchmarkTwoSum2Point/left_side-4            	  799206	      2043 ns/op
//BenchmarkTwoSum2Point/two_side
//BenchmarkTwoSum2Point/two_side-4             	  708943	      2888 ns/op
//BenchmarkTwoSum2Point/right_side
//BenchmarkTwoSum2Point/right_side-4           	  588092	      2186 ns/op
//BenchmarkTwoSum2Point/zero
//BenchmarkTwoSum2Point/zero-4                 	  653221	      1885 ns/op
//BenchmarkTwoSum2Point/negative
//BenchmarkTwoSum2Point/negative-4             	  765879	      2345 ns/op
//BenchmarkTwoSum2Point/more
//BenchmarkTwoSum2Point/more-4                 	  568328	      2492 ns/op
//PASS
func BenchmarkTwoSum2Point(b *testing.B) {
	for _, test := range testCases {
		test := test
		b.Run(test.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result := TwoSum2Point(test.sum, test.array)
				assert.Equal(b, result, test.result)
			}
		})
	}
}
