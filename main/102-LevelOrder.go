package main
/*
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
 */
import (
	. "LeetCode/main/base"
	"LeetCode/main/container"
	"fmt"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue1 := container.NewQueue()
	queue2 := container.NewQueue()
	retList := make([][]int, 0)
	var level []int
	var node *TreeNode
	queue1.Push(root)
	for !queue1.IsEmpty() {
		level = make([]int, 0)
		for !queue1.IsEmpty() {
			node = queue1.Pop().(*TreeNode)
			level = append(level, node.Val)
			if node.Left != nil {
				queue2.Push(node.Left)
			}
			if node.Right != nil {
				queue2.Push(node.Right)
			}
		}
		retList = append(retList, level)
		queue1 = queue2
		queue2 = container.NewQueue()
	}
	return  retList
}

func main() {
	node20 := &TreeNode{Val:20,Left:&TreeNode{Val:15},Right:&TreeNode{Val:7}}
	root := &TreeNode{Val:3,Left:&TreeNode{Val:9},Right:node20}
	fmt.Println(levelOrder(root))
}