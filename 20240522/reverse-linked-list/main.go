package main

/*
https://leetcode.cn/problems/reverse-linked-list
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代法
func reverseList(head *ListNode) *ListNode {
	var current = head
	var previous *ListNode
	var next *ListNode
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

func main() {

}
