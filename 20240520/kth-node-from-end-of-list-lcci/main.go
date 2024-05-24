package main

/*
https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	fast := head
	slow := head
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

func main() {

}
