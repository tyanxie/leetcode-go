package main

import (
	"sort"
)

/*
https://leetcode.cn/problems/majority-element
*/

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
}
