package main

import (
	"fmt"
	."myGoProject/main/leetcode/ds"
)

func main() {
	node1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node1,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node2,
	}
	node4 := &ListNode{
		Val:  4,
		Next: node2,
	}
	node5 := &ListNode{
		Val:  5,
		Next: node4,
	}
	node := getIntersectionNode2(node5, node3)
	fmt.Println(node.Val)
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodes := make(map[*ListNode]int)
	for headA != nil {
		nodes[headA]++
		headA = headA.Next
	}
	for headB != nil {
		if nodes[headB] > 0 {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := listLen(headA)
	lenB := listLen(headB)
	if lenA > lenB {
		return getIntersectionNode(headB, headA)
	}
	between := lenB - lenA
	for between > 0 {
		headB = headB.Next
		between--
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func listLen(head *ListNode) (len int) {
	for head != nil {
		len++
		head = head.Next
	}
	return
}
