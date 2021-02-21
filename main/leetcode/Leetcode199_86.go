package main

import ds "myGoProject/main/leetcode/ds"

func main() {
	rightSideView(nil)
}

func rightSideView(root *ds.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*ds.TreeNode{root}
	res := []int{}
	for len(queue) != 0 {
		size := len(queue)
		for size > 0 {
			p := queue[0]
			queue = queue[1:]
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
			size--
			if size == 0 {
				res = append(res, p.Val)
			}
		}
	}
	return res
}
