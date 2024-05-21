package main

import "fmt"

/*
https://leetcode.cn/problems/minimum-size-subarray-sum
*/

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		if nums[0] >= target {
			return 1
		}
		return 0
	}

	slow, fast := 0, 0
	res := length + 1
	val := 0
	for ; fast < length; fast++ {
		val += nums[fast]
		for val >= target {
			curr := fast - slow + 1
			if curr < res {
				res = curr
			}
			val -= nums[slow]
			slow++
		}
	}
	if res > length {
		return 0
	}
	return res
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
