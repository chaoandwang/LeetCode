package main
/*
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
 */

import (
	. "LeetCode/main/base"
	"LeetCode/main/container"
	"fmt"
)

func path(stack *container.Stack) []int {
	tmp := container.NewStack()
	pathList := make([]int, 0)
	for !stack.IsEmpty() {
		tmp.Push(stack.Pop())
	}
	for !tmp.IsEmpty() {
		node := tmp.Pop().(*TreeNodeStackElement)
		stack.Push(node)
		pathList = append(pathList, node.Node.Val)
	}
	return pathList
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {return [][]int{}}
	stack := container.NewStack()
	current := 0
	paths := make([][]int, 0)
	var snode *TreeNodeStackElement
	for root != nil || !stack.IsEmpty() {
		if root != nil {
			current += root.Val
			if current > sum {
				current -= root.Val
				root = nil
			} else if current == sum {
				if root.IsLeaf() {
					stack.Push(&TreeNodeStackElement{root, true})
					paths = append(paths, path(stack))
					stack.Pop()
				}
				current -= root.Val
				root = nil
			} else {
				stack.Push(&TreeNodeStackElement{root, true})
				root = root.Left
			}
		} else {
			snode = stack.Pop().(*TreeNodeStackElement)
			current -= snode.Node.Val
			if snode.IsLeft && snode.Node.Right != nil {
				stack.Push(&TreeNodeStackElement{snode.Node, false})
				current += snode.Node.Val
				root = snode.Node.Right
			} else {
				root = nil
			}
		}
	}
	return paths
}

func main() {
	node4 := &TreeNode{Val:4,Left:&TreeNode{Val:11,Left:&TreeNode{Val:7},Right:&TreeNode{Val:2}}}
	node8 := &TreeNode{Val:8,Left:&TreeNode{Val:13},Right:&TreeNode{Val:4,Left:&TreeNode{Val:5},Right:&TreeNode{Val:1}}}
	root := &TreeNode{Val:5,Left:node4,Right:node8}
	fmt.Println(pathSum(root, 22))
}