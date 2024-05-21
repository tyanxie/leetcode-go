package main

/*
https://leetcode.cn/problems/add-two-numbers/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	res := new(ListNode)
	l := res
	val1 := l1.Val
	val2 := l2.Val
	carry := false

	for {

		if l1 != nil {
			val1 = l1.Val
		} else {
			val1 = 0
		}

		if l2 != nil {
			val2 = l2.Val
		} else {
			val2 = 0
		}

		l.Val = val1 + val2
		if carry {
			l.Val += 1
		}

		if l.Val >= 10 {
			l.Val = l.Val % 10
			carry = true
		} else {
			carry = false
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			break
		}

		l.Next = new(ListNode)
		l = l.Next
	}

	if carry {
		l.Next = &ListNode{Val: 1}
	}

	return res
}

func main() {

}
