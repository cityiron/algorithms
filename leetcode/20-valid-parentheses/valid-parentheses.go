package leetcode

import "strings"

// https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	if len(s)&1 == 1 {
		return false
	}
	var stack []byte
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := range s {
		if pairs[s[i]] > 0 { // 这里如果没匹配到，是 byte 默认值 0
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] { // 栈顶 Top() 匹配期望值
				return false
			}
			stack = stack[:len(stack)-1] // 移除栈顶 Pop()
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

func isValid2(s string) bool {
	var stack []byte
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := range s {
		if pairs[s[i]] > 0 { // 这里如果没匹配到，是 byte 默认值 0
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] { // 栈顶 Top() 匹配期望值
				return false
			}
			stack = stack[:len(stack)-1] // 移除栈顶 Pop()
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

// 连连看写法，时间复杂度会比前面的高
func isValid3(s string) bool {
	var l int
	for ok := true; ok; ok = l != len(s) {
		l = len(s)
		s = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "()", ""), "{}", ""), "[]", "")
	}
	return len(s) == 0
}


