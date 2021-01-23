package main

import (
	"fmt"
	. "myGoProject/main/leetcode/ds"
)

func main() {
	list := (&ListNode{}).NewTestList()
	reorderList(list)
	fmt.Println(list.String())
}

func reorderList(head *ListNode)  {
	mid := midNode(head)
	p := reverseList(mid.Next)
	mid.Next = nil
	for p!=nil && head!=nil {
		temp := p
		if p != nil {
			p = p.Next
		}
		temp.Next = head.Next
		head.Next = temp
	}
}

func midNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	fast, low := head, head
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		low = low.Next
	}
	return low
}

func reverseList(head *ListNode) *ListNode {
	p := head
	newHead := &ListNode{0, nil}
	for p != nil {
		temp := p
		p = p.Next
		temp.Next = newHead.Next
		newHead.Next = temp
	}
	return newHead.Next
}
