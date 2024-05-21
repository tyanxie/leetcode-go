package main

/*
https://leetcode.cn/problems/merge-sorted-array
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	i := len(nums1) - 1
	for m >= 0 && n >= 0 {
		num1 := nums1[m]
		num2 := nums2[n]
		if num1 >= num2 {
			m--
			nums1[i] = num1
		} else {
			n--
			nums1[i] = num2
		}
		i--
	}

	for n >= 0 {
		nums1[n] = nums2[n]
		n--
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
