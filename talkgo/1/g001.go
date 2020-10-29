package talkgo

// 题目：
// 输入一个递增排序的数组和一个数字 S，在数组中查找两个数，使得他们的和正好是 S，如果有多对数字的和等于 S，输出两个数的乘积最小的。

// TwoSum2For 暴力法
func TwoSum2For(sum int, array []int) []int {
	b, arr := checkArray(array)

	if !b {
		return arr
	}

	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j > i; j-- {
			if sum == array[i]+array[j] {
				return []int{array[i], array[j]}
			}
		}
	}

	return []int{}
}

// TwoSumMap map 缓存法
func TwoSumMap(sum int, array []int) []int {
	b, arr := checkArray(array)

	if !b {
		return arr
	}

	// 通过 map 缓存数据
	cache := make(map[int]int, len(array))
	for i := range array {
		cache[array[i]] = i
	}

	// 返回值在数组中的位置
	for i := range array {
		// 通过减法找到另一个值，直接去 map 查找
		if _, ok := cache[sum-array[i]]; ok {
			return []int{array[i], sum - array[i]}
		}
	}

	return []int{}
}

// TwoSum2Point 双指针法
func TwoSum2Point(sum int, array []int) []int {
	b, arr := checkArray(array)

	if !b {
		return arr
	}

	left, right := 0, len(array)-1
	for left < right {
		s := array[left] + array[right]
		if sum == s {
			return []int{array[left], array[right]}
		} else if sum > s {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

func checkArray(array []int) (bool, []int) {
	if len(array) > 0 {
		return true, nil
	}

	return false, []int{}
}

//func main() {
//	fmt.Println("暴力破解法")
//	ints := FindTwoSum1(10, []int{})
//	fmt.Printf("result %v, expect %v", ints, []int{})
//	fmt.Println()
//	ints = FindTwoSum1(10, []int{1, 2, 4, 7, 11, 16})
//	fmt.Printf("result %v, expect %v", ints, []int{})
//	fmt.Println()
//	ints = FindTwoSum1(15, []int{1, 2, 4, 7, 8, 11, 15})
//	fmt.Printf("result %v expect %v", ints, []int{4, 11})
//	fmt.Println()
//	ints = FindTwoSum1(4, []int{-11, -5, -3, 0, 4, 7, 9, 11, 15})
//	fmt.Printf("result %v expect %v", ints, []int{-11, 15})
//	fmt.Println()
//	fmt.Println("Map缓存法")
//	ints = FindTwoSum2(10, []int{})
//	fmt.Printf("result %v, expect %v", ints, []int{})
//	fmt.Println()
//	ints = FindTwoSum2(10, []int{1, 2, 4, 7, 11, 16})
//	fmt.Printf("result %v, expect %v", ints, []int{})
//	fmt.Println()
//	ints = FindTwoSum2(15, []int{1, 2, 4, 7, 8, 11, 15})
//	fmt.Printf("result %v expect %v", ints, []int{4, 11})
//	fmt.Println()
//	ints = FindTwoSum2(4, []int{-11, -5, -3, 0, 4, 7, 9, 11, 15})
//	fmt.Printf("result %v expect %v", ints, []int{-11, 15})
//}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
