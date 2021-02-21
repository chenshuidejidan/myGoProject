package ds

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree() *TreeNode {
	node0 := &TreeNode{0, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node5 := &TreeNode{5, nil, nil}
	node4 := &TreeNode{4, node3, node5}
	node2 := &TreeNode{2, node0, node4}
	node7 := &TreeNode{7, nil, nil}
	node9 := &TreeNode{9, nil, nil}
	node8 := &TreeNode{8, node7, node9}
	node6 := &TreeNode{6, node2, node8}
	return node6
}
