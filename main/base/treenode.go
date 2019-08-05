package base

import "LeetCode/main/container"

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InorderTravel(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	travel := make([]int, 0)
	inorderTravelInternal(root, &travel)
	return travel
}

func inorderTravelInternal(root *TreeNode, travel *[]int) {
	if root == nil {
		return
	}
	inorderTravelInternal(root.Left, travel)
	*travel = append(*travel, root.Val)
	inorderTravelInternal(root.Right, travel)
}

func InorderTravelWithoutRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := container.NewStack()
	travel := make([]int, 0)
	for root != nil || !stack.IsEmpty() {
		if root != nil {
			stack.Push(root)
			root = root.Left
		} else {
			root = stack.Pop().(*TreeNode)
			travel = append(travel, root.Val)
			root = root.Right
		}
	}
	return travel
}

func PreOrderTravel(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	travel := make([]int, 0)
	preOrderTravelInternal(root, &travel)
	return travel
}

func preOrderTravelInternal(root *TreeNode, travel *[]int) {
	if root == nil {
		return
	}
	*travel = append(*travel, root.Val)
	preOrderTravelInternal(root.Left, travel)
	preOrderTravelInternal(root.Right, travel)
}

func PreOrderTravelWithoutRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := container.NewStack()
	travel := make([]int, 0)
	for root != nil || !stack.IsEmpty() {
		if root != nil {
			travel = append(travel, root.Val)
			stack.Push(root)
			root = root.Left
		} else {
			root = stack.Pop().(*TreeNode)
			root = root.Right
		}
	}
	return travel
}

func PostOrderTraver(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	travel := make([]int, 0)
	postOrderTraverInternal(root, &travel)
	return travel
}

func postOrderTraverInternal(root *TreeNode, travel *[]int) {
	if root == nil {
		return
	}
	postOrderTraverInternal(root.Left, travel)
	postOrderTraverInternal(root.Right, travel)
	*travel = append(*travel, root.Val)
}

type treeNodeStackElement struct {
	node *TreeNode
	isLeft bool
}

func PostOrderTraverWithOutRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := container.NewStack()
	travel := make([]int, 0)
	var nodeElment treeNodeStackElement
	for root != nil || !stack.IsEmpty() {
		if root != nil {
			stack.Push(treeNodeStackElement{node:root,isLeft:true})
			root = root.Left
		} else {
			nodeElment = stack.Pop().(treeNodeStackElement)
			if nodeElment.isLeft {
				nodeElment.isLeft = false
				stack.Push(nodeElment)
				root = nodeElment.node.Right
			} else {
				travel = append(travel, nodeElment.node.Val)
				root = nil
			}
		}
	}
	return travel
}

func LevelOrderTravel(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := container.NewQueue()
	queue.Push(root)
	travel := make([]int, 0)
	for !queue.IsEmpty() {
		root = queue.Pop().(*TreeNode)
		travel = append(travel, root.Val)
		if root.Left != nil {
			queue.Push(root.Left)
		}
		if root.Right != nil {
			queue.Push(root.Right)
		}
	}
	return travel
}