package main
/*
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
 */
import (
	. "LeetCode/main/base"
	"LeetCode/main/container"
	"fmt"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack1 := container.NewStack()
	stack2 := container.NewStack()
	retList := make([][]int, 0)
	leftToRight := true
	var level []int
	var node *TreeNode
	stack1.Push(root)
	for !stack1.IsEmpty() {
		level = make([]int, 0)
		for !stack1.IsEmpty() {
			node = stack1.Pop().(*TreeNode)
			level = append(level, node.Val)
			if leftToRight {
				if node.Left != nil {
					stack2.Push(node.Left)
				}
				if node.Right != nil {
					stack2.Push(node.Right)
				}
			} else {
				if node.Right != nil {
					stack2.Push(node.Right)
				}
				if node.Left != nil {
					stack2.Push(node.Left)
				}
			}
		}
		stack1 = stack2
		stack2 = container.NewStack()
		retList = append(retList, level)
		leftToRight = !leftToRight
	}
	return retList
}

func main() {
	node20 := &TreeNode{Val:20,Left:&TreeNode{Val:15},Right:&TreeNode{Val:7}}
	node9 := &TreeNode{Val:9,Left:&TreeNode{Val:6},Right:&TreeNode{Val:5}}
	root := &TreeNode{Val:3,Left:node9,Right:node20}
	fmt.Println(zigzagLevelOrder(root))
}