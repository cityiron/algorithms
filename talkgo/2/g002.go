package talkgo

import "errors"

// 题目
// 一个有序数组，从随即一位截断，把前段放在后边，如 4 5 6 7 1 2 3 求中位数

var illegalArray = errors.New("illegal array")

// MiddleNum 算中位数
func MiddleNum(nums []int) (float32, error) {
	l := len(nums)
	if l <= 0 {
		return 0, illegalArray
	}

	if l == 1 {
		return float32(nums[0]), nil
	} else if l == 2 {
		return float32((nums[0] + nums[1]) / 2), nil
	}

	p := Point(nums)

	if isEven(l) {
		target1 := l / 2
		target2 := l/2 + 1
		return (float32(nums[Position(l, p, target1)]) + float32(nums[Position(l, p, target2)])) / 2, nil
	} else {
		return float32(nums[Position(l, p, (l+1)/2)]), nil
	}
}

// Point 得到拐点
func Point(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	n, flag := nums[0], -1

	for i := 1; i < len(nums); i++ {
		cf := -1
		c := nums[i]
		if c < n {
			cf = 1 // 降
		} else if c > n {
			cf = 0 // 升
		}

		if flag >= 0 && cf >= 0 && cf != flag {
			return i
		}

		n = c
		if cf >= 0 {
			flag = cf
		}
	}

	return 0
}

// Position 获取位置
func Position(l, p, t int) int {
	if l-p > t {
		return p + t - 1
	} else {
		return t + p - l - 1
	}
}

func isEven(num int) bool {
	return num&1 == 0
}
