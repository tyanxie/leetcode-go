package main

/*
https://leetcode.cn/problems/delete-node-in-a-linked-list/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {

}