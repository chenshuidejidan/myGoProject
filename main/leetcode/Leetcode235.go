package main

import (
	"fmt"
	. "myGoProject/main/leetcode/ds"
)

func main() {
	root := NewTree()
	ancestor := lowestCommonAncestor(root, root.Left, root.Right)
	fmt.Printf("%+v\n", ancestor)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if p == q {
		return p
	}
	stack1, stack2 := postOrderFindNode(root, p), postOrderFindNode(root, q)
	i := 0
	for i = 0; i < len(stack1) && i < len(stack2) && stack1[i] == stack2[i]; i++ {
	}
	if i > 0 {
		return stack1[i-1]
	} else {
		return nil
	}
}

func postOrderFindNode(root, target *TreeNode) []*TreeNode {
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	pre, p := root, root
	for len(stack) > 0 {
		p = stack[len(stack)-1]
		if p.Left != nil && p.Left != pre && p.Right != pre {
			stack = append(stack, p.Left)
		} else if p.Right != nil && p.Right != pre {
			stack = append(stack, p.Right)
		} else {
			pre = stack[len(stack)-1]
			if pre == target {
				return stack
			}
			stack = stack[:len(stack)-1]
		}
	}
	return nil
}
