package main

/*
https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii
*/

// 0 0 1 1 1 2 2 2

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	slow, fast := 2, 2
	for fast < length {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {

}
