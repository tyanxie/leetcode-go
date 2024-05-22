package main

/*
https://leetcode.cn/problems/house-robber
*/

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}

	res := make([]int, length)
	res[0] = nums[0]
	if nums[0] >= nums[1] {
		res[1] = nums[0]
	} else {
		res[1] = nums[1]
	}
	if length == 2 {
		return res[1]
	}

	if nums[0]+nums[2] >= nums[1] {
		res[2] = nums[0] + nums[2]
	} else {
		res[2] = nums[1]
	}
	if length == 3 {
		return res[2]
	}

	for i := 3; i < length; i++ {
		l := res[i-3] + nums[i]
		r := res[i-2] + nums[i]

		if l >= r {
			res[i] = l
		} else {
			res[i] = r
		}

		if res[i-1] > res[i] {
			res[i] = res[i-1]
		}
	}
	return res[length-1]
}

func main() {

}
