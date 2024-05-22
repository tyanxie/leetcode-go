package main

import (
	"fmt"
	"strconv"
)

/*
https://leetcode.cn/problems/summary-ranges
*/

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return make([]string, 0)
	} else if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	res := make([]string, 0)
	l, r := 0, 1
	length := len(nums)

	for r < length {
		for r < length && (nums[r] == nums[r-1] || nums[r] == nums[r-1]+1) {
			r++
		}
		res = append(res, toScope(nums[l], nums[r-1]))
		l = r
		r++
	}

	if l < length {
		res = append(res, toScope(nums[l], nums[r-1]))
	}
	return res
}

func toScope(l, r int) string {
	if l == r {
		return strconv.Itoa(l)
	}
	return strconv.Itoa(l) + "->" + strconv.Itoa(r)
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
