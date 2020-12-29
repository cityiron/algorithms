package leetcode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]
	array := [][]int{[]int{3, 0, 8, 4}, []int{2, 4, 5, 7}, []int{9, 2, 6, 3}, []int{0, 3, 1, 0}}
	maxIncreaseKeepingSkyline(array)
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	result := 0
	rmap := make(map[int]int, 4)
	cmap := make(map[int]int, 4)
	for i := range grid {
		for i2 := range grid[i] {
			setMax(rmap, i, grid[i][i2])
			setMax(cmap, i2, grid[i][i2])
			fmt.Printf("row:%d, value:%d \r\n", i, grid[i][i2])
		}
	}

	for i := range grid {
		for i2 := range grid[i] {
			result += min(rmap[i], cmap[i2]) - grid[i][i2]
		}
	}

	return result
}

func setMax(rmap map[int]int, p, i int) {
	if _, ok := rmap[p]; ok {
		rmap[p] = max(rmap[p], i)
		return
	}

	rmap[p] = i
}

func max(x, y int) int {
	if x <= y {
		return y
	}

	return x
}

func min(x, y int) int {
	if x <= y {
		return x
	}

	return y
}
