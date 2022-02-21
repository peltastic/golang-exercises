package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		}
		return &ListNode{carry, nil}
	}

	val1 := 0
	var l1Next *ListNode
	if l1 != nil {
		val1 = l1.Val
		l1Next = l1.Next
	}

	val2 := 0
	var l2Next *ListNode
	if l2 != nil {
		val2 = l2.Val
		l2Next = l2.Next
	}

	sum := val1 + val2 + carry
	carry_next := 0
	if sum >= 10 {
		carry_next = 1
		sum -= 10
	}

	return &ListNode{sum, addTwoNumbersHelper(l1Next, l2Next, carry_next)}
}
