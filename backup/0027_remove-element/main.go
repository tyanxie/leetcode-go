package main

/*
https://leetcode.cn/problems/remove-element
*/

func removeElement(nums []int, val int) int {
	r := len(nums) - 1
	for r >= 0 && nums[r] == val {
		r--
	}

	for l := 0; l < r; l++ {
		if nums[l] == val {
			nums[l] = nums[r]
			r--
			for r >= 0 && nums[r] == val {
				r--
			}
		}
	}

	return r + 1
}

func main() {

}
