package main

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/

// 最小堆实现
type KthLargest struct {
	heap *IntHeap
	size int
}

func Constructor(k int, nums []int) KthLargest {
	ih := new(IntHeap)
	heap.Init(ih)
	for _, v := range nums {
		if ih.Len() < k {
			heap.Push(ih, v)
		} else {
			_, min := ih.Top()
			if v > min {
				heap.Pop(ih)
				heap.Push(ih, v)
			}
		}
	}

	return KthLargest{heap: ih, size: k}
}

func (this *KthLargest) Add(val int) int {
	b, min := this.heap.Top()
	if !b {
		heap.Push(this.heap, val)
		return val
	}

	if this.heap.Len() < this.size {
		heap.Push(this.heap, val)
		_, min = this.heap.Top()
		return min
	}

	if val > min {
		heap.Pop(this.heap)
		heap.Push(this.heap, val)
		_, min = this.heap.Top()
	}

	return min
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Top() (bool, int) {
	if h.Len() == 0 {
		return false, 0
	}

	old := *h
	return true, old[0]
}

type KthLargest_SortPro struct {
	// k 个元素的数组
	array []int
	// 当前元素个数
	size int
}

func Constructor_SortPro(k int, nums []int) KthLargest_SortPro {
	re := KthLargest_SortPro{array: []int{}, size: k}
	if len(nums) == 0 {
		return re
	}

	for _, v := range nums {
		re.Add(v)
	}

	return re
}

func (this *KthLargest_SortPro) Add(val int) int {
	// 未满
	if len(this.array) < this.size {
		// 找到位置，第一个位置最小值不特意取出来比较了，因为循环第一个就匹配到了
		p := -1
		for i := 0; i < len(this.array); i++ {
			if this.array[i] >= val {
				p = i
				break
			}
		}

		// val 是最大值
		if p == -1 {
			this.array = append(this.array, val)
		} else {
			max := this.array[len(this.array)-1]
			for i := len(this.array) - 2; i >= p; i-- {
				this.array[i+1] = this.array[i]
			}
			this.array = append(this.array, max)
			this.array[p] = val
		}
	} else {
		p := -1
		for i := 0; i < len(this.array); i++ {
			if this.array[i] >= val {
				p = i
				break
			}
		}

		// val 是最大值，这个地方可以直接拿最大值比较
		if p == -1 {
			this.array = this.array[1:]
			this.array = append(this.array, val)
		} else if p == 0 {

		} else {
			for i := 0; i < p-1; i++ {
				this.array[i] = this.array[i+1]
			}
			this.array[p-1] = val
		}
	}

	return this.array[0]
}

type KthLargest_Sort struct {
	// k 个元素的数组
	array []int
	size  int
}

func Constructor_Sort(k int, nums []int) KthLargest_Sort {
	re := KthLargest_Sort{array: []int{}, size: k}
	if len(nums) == 0 {
		return re
	}
	// 排序
	sort.Ints(nums)
	// 符合要求的元素设置到 array
	s := len(nums) - k
	if s < 0 {
		s = 0
	}
	for i := s; i < len(nums); i++ {
		re.array = append(re.array, nums[i])
	}

	return re
}

func (this *KthLargest_Sort) Add(val int) int {
	// 先加上，然后取前K个
	this.array = append(this.array, val)
	sort.Ints(this.array)
	if len(this.array) > this.size {
		this.array = this.array[1:]
	}

	return this.array[0]
}

type KthLargest_Sort2 struct {
	// k 个元素的数组
	array []int
	size  int
	full  bool
}

func Constructor_Sort2(k int, nums []int) KthLargest_Sort2 {
	re := KthLargest_Sort2{array: []int{}, size: k}
	if len(nums) == 0 {
		return re
	}
	// 排序
	sort.Ints(nums)
	// 符合要求的元素设置到 array
	s := len(nums) - k
	if s < 0 {
		s = 0
	}
	for i := s; i < len(nums); i++ {
		re.array = append(re.array, nums[i])
	}
	re.setFull()

	return re
}

func (this *KthLargest_Sort2) Add(val int) int {
	// 没满
	if !this.full {
		this.array = append(this.array, val)
		sort.Ints(this.array)
		this.setFull()
		return this.array[0]
	}

	// 满了，就考虑是否替换

	// 如果最小值都比新值大或等于新值，那么不用处理
	min := this.array[0]
	if min >= val {
		return this.array[0]
	}

	// 如果是最大值，就直接移动，省去遍历
	max := this.array[len(this.array)-1]
	if val > max {
		this.array = this.array[1:]
		this.array = append(this.array, val)
		return this.array[0]
	}

	big := -1
	for i, v := range this.array {
		if v >= val {
			big = i
			break
		}
	}

	for i := 0; i < big-1; i++ {
		this.array[i], this.array[i+1] = this.array[i+1], this.array[i+2]
	}

	this.array[big-1] = val
	return this.array[0]
}

func (this *KthLargest_Sort2) setFull() {
	if len(this.array) >= this.size {
		this.full = true
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
