package main

import (
	"fmt"
)

/*
https://leetcode.cn/problems/kth-largest-element-in-an-array
*/

/*
桶排序
*/

func findKthLargest(nums []int, k int) int {
	buckets := make([]int, 20001)
	for _, num := range nums {
		buckets[num+10000]++
	}

	for i := 20000; i >= 0; i-- {
		k = k - buckets[i]
		if k <= 0 {
			return i - 10000
		}
	}
	return 0
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
