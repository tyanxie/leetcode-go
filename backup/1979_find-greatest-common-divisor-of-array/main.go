package main

import "fmt"

/*
https://leetcode.cn/problems/find-greatest-common-divisor-of-array
*/

func findGCD(nums []int) int {
	switch len(nums) {
	case 0:
		return 1
	case 1:
		return nums[0]
	}

	minNum, maxNum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}

	if minNum == maxNum {
		return minNum
	}

	for i := maxNum; i >= 2; i-- {
		if maxNum%i == 0 && minNum%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	fmt.Println(findGCD([]int{6, 9}))
}
