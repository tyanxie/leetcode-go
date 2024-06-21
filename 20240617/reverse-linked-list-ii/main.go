package main

import "fmt"

/*
https://leetcode.cn/problems/reverse-linked-list-ii
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	var prev *ListNode
	l := head
	for i := left; i > 1; i-- {
		prev = l
		l = l.Next
	}

	r := l
	for i := right - left; i > 0; i-- {
		r = r.Next
	}
	rn := r.Next
	r.Next = nil

	l = reverse(l)

	if prev == nil {
		if rn != nil {
			for r.Next != nil {
				r = r.Next
			}
			r.Next = rn
		}
		return l
	}

	prev.Next = l
	if rn != nil {
		for l.Next != nil {
			l = l.Next
		}
		l.Next = rn
	}

	return head
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	curr := head
	next := curr.Next
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	list := []int{1, 2, 3, 4, 5}

	head := &ListNode{
		Val: list[0],
	}
	prev := head
	for i := 1; i < len(list); i++ {
		prev.Next = &ListNode{
			Val: list[i],
		}
		prev = prev.Next
	}

	head = reverseBetween(head, 1, 4)

	curr := head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
