package main

import "fmt"

const (
	RED          = 0
	BLACK        = 1
	DOUBLE_BLACK = 2
)

type node struct {
	val         int
	left, right *node
	color       int //红色0，黑色1，双重黑色2
}

func newNode(val int) *node {
	return &node{
		val:   val,
		left:  NilNode,
		right: NilNode,
		color: 0,
	}
}

var NilNode *node = &node{
	val:   0,
	left:  nil,
	right: nil,
	color: 1,
}

func Clear(p **node) {
	*p = NilNode
}

func (root *node) Insert(value int) *node {
	root = root.__insert(value)
	root.color = BLACK //强制把根节点染色成黑色
	return root
}

func (root *node) __insert(value int) *node {
	if root == NilNode {
		return newNode(value)
	}
	if root.val == value {
		return root
	}
	if root.val < value {
		root.right = root.right.__insert(value)
	} else {
		root.left = root.left.__insert(value)
	}
	return root.insertMaintain() //递归修改
}

func (root *node) insertMaintain() *node {
	if root.left.color == RED && root.right.color == RED { //第一种情况，只需要调整帽子

	} else if left := root.left; left.color == RED && left.hasRedChild() { //第二种情况：左子树中发生冲突
		if left.right.color == RED {
			root.left = left.leftRotate()
		}
		root = root.rightRotate()
	} else if right := root.right; right.color == RED && right.hasRedChild() { //第二种情况：右子树中发生冲突
		if right.left.color == RED {
			root.right = right.rightRotate()
		}
		root = root.leftRotate()
	} else { //不需要调整
		return root
	}
	root.color, root.left.color, root.right.color = RED, BLACK, BLACK //两种情况最后都需要调整root帽子的颜色
	return root
}

func (root *node) Delete(value int) *node {
	root = root.__delete(value)
	root.color = BLACK
	return root
}

func (root *node) __delete(value int) *node {
	if root == NilNode {
		return NilNode
	}
	if value < root.val {
		root.left = root.left.__delete(value)
	} else if value > root.val {
		root.right = root.right.__delete(value)
	} else {

		if root.left == NilNode || root.right == NilNode { //度为0或者度为1的节点的删除
			var temp *node
			if temp = root.left; temp == NilNode {
				temp = root.right
			}
			temp.color += root.color
			root = nil //删除root
			return temp
		} else {
			temp := root.predecessor()
			root.val = temp.val
			root.left = root.left.__delete(temp.val)
		}
	}
	return root.deleteMaintain()
}

func (root *node) deleteMaintain() *node {
	if root.left.color != DOUBLE_BLACK && root.right.color != DOUBLE_BLACK { //没有双重黑色节点，无需调整
		return root
	}
	if root.hasRedChild() { //4。有双重黑节点,双重黑节点的兄弟是红色节点，旋转至没有红色子节点的情况
		root.color = RED            //原黑色根节点变为红色
		if root.left.color == RED { //左子树是红色
			root = root.rightRotate()
			root.color = BLACK //新的红根节点变为黑色
			root.right = root.right.deleteMaintain()
		} else { //右子树是红色
			root = root.leftRotate()
			root.color = BLACK //新的红根节点变为黑色
			root.left = root.left.deleteMaintain()
		}
		return root
	}

	if root.left.color == DOUBLE_BLACK && !root.right.hasRedChild() ||
		root.right.color == DOUBLE_BLACK && !root.left.hasRedChild() { //1. 双重黑的兄弟黑，且兄弟的两个孩子都是黑色
		root.left.color -= 1
		root.right.color -= 1
		root.color += 1
		return root
	}

	if root.left.color == DOUBLE_BLACK {
		if right, rright := root.right, root.right.right; right.right.color != RED { //2。情况二 RL
			right = right.rightRotate()
			right.color, rright.color = rright.color, right.color
		}

		root.leftRotate()            //RR
		root.color = root.left.color //新根节点颜色变为原来根节点的颜色
	} else {
		if left, lleft := root.left, root.left.left; lleft.color != RED { //2。情况二 LR
			left = left.leftRotate()
			left.color, lleft.color = lleft.color, left.color
		}

		root.rightRotate()            //LL
		root.color = root.right.color //新根节点颜色变为原来根节点的颜色
	}
	root.left.color, root.right.color = BLACK, BLACK
	return root
}

func (root *node) predecessor() *node { //寻找中序前驱节点:左子树的最右孩子
	temp := root.left
	for temp != NilNode && temp.right != NilNode {
		temp = temp.right
	}
	return temp
}

func (root *node) hasRedChild() bool {
	return root.left.color == RED || root.right.color == RED
}

func (root *node) leftRotate() *node {
	temp := root.right
	root.right = temp.left
	temp.left = root
	return temp
}

func (root *node) rightRotate() *node {
	temp := root.left
	root.left = temp.right
	temp.right = root
	return temp
}

func (root *node) ToString() {
	if root == NilNode {
		return
	}
	root.print()
	root.left.ToString()
	root.right.ToString()
}
func (root *node) print() {
	fmt.Printf("%d | %d, l : %d, r: %d\n", root.color, root.val, root.left.val, root.right.val)
}

func Test() {
	var op, val int
	root := NilNode
	fmt.Println("=============RED_BLACK_TREE_TEST===============")

	for {
		fmt.Print("\n->")
		fmt.Scanf("%d %d", &op, &val)
		switch op {
		case 1: //插入
			root = root.Insert(val)
		case 0: //删除
			root = root.Delete(val)
		case -1:
			break
		case 2:
			Clear(&root)
		default:
			continue
		}
		root.ToString()
	}
}
