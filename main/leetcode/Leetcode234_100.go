package main

import . "myGoProject/main/leetcode/ds"

func isPalindrome(head *ListNode) bool {
	mid := midNodeOfList(head)
	head2 := reverseList2(mid)
	for head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head, head2 = head.Next, head2.Next
	}
	return true
}

func midNodeOfList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head
	for p != nil {
		head = head.Next
		p = p.Next
		if p != nil {
			p = p.Next
		}
	}
	return head
}

func reverseList2(head *ListNode) *ListNode {
	newHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	for head != nil {
		p := head
		head = head.Next
		p.Next = newHead.Next
		newHead.Next = p
	}
	return newHead.Next
}

func main() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, node1}
	node3 := &ListNode{3, node2}
	node4 := &ListNode{3, node3}
	node5 := &ListNode{2, node4}
	head := &ListNode{1, node5}
	isPalindrome(head)
}
