package ds

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) NewTestList() *ListNode {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	return node1
}

func (l *ListNode) String() (s string) {
	b := new(bytes.Buffer)
	b.WriteByte('[')
	head := l
	for l != nil {
		if l != head {
			b.Write([]byte(", "))
		}
		b.Write([]byte(strconv.Itoa(l.Val)))
		l = l.Next
	}
	b.WriteByte(']')
	return b.String()
}