package main

/*
https://leetcode.cn/problems/remove-duplicates-from-sorted-array
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	slow, fast := 0, 1
	for fast < len(nums) {
		for fast < len(nums) && nums[fast] == nums[slow] {
			fast++
		}
		if fast >= len(nums) {
			break
		}
		nums[slow+1] = nums[fast]
		slow++
	}
	return slow + 1
}

func main() {

}
