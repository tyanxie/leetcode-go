package main

import (
	"fmt"
)

/*
https://leetcode.cn/problems/longest-substring-without-repeating-characters/
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	l, r := 0, 1
	res := 0
	m := make(map[byte]struct{})
	for l < len(s) && r < len(s) {
		m[s[l]] = struct{}{}
		if _, ok := m[s[r]]; ok {
			length := r - l
			if length > res {
				res = length
			}
			delete(m, s[l])
			l++
			if l == r {
				r++
			}
		} else {
			m[s[r]] = struct{}{}
			r++
		}
	}
	length := r - l
	if length > res {
		return length
	}
	return res
}

func main() {
	length := lengthOfLongestSubstring("pwwkew")
	fmt.Println(length)
}
