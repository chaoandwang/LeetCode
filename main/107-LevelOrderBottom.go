package main
/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	. "LeetCode/main/base"
	"LeetCode/main/container"
	"fmt"
)

type stackElement struct {
	node *TreeNode
	isNil bool
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue1 := container.NewQueue()
	queue2 := container.NewQueue()
	stack := container.NewStack()
	queue1.Push(root)
	for !queue1.IsEmpty() {
		for !queue1.IsEmpty() {
			root = queue1.Pop().(*TreeNode)
			stack.Push(stackElement{node:root})
			if root.Right != nil {
				queue2.Push(root.Right)
			}
			if root.Left != nil {
				queue2.Push(root.Left)
			}
		}
		if !queue2.IsEmpty() {
			stack.Push(stackElement{isNil:true})
			queue1 = queue2
			queue2 = container.NewQueue()
		}
	}
	level := make([]int, 0)
	travel := make([][]int, 0)
	var se stackElement
	for !stack.IsEmpty() {
		se = stack.Pop().(stackElement)
		if !se.isNil {
			level = append(level, se.node.Val)
		} else {
			travel = append(travel, level)
			level = make([]int, 0)
		}
	}
	travel = append(travel, level)
	return travel
}

func main() {
	node20 := &TreeNode{Val:20,Left:&TreeNode{Val:15},Right:&TreeNode{Val:7}}
	node9 := &TreeNode{Val:9,Left:&TreeNode{Val:6},Right:&TreeNode{Val:5}}
	root := &TreeNode{Val:3,Left:node9,Right:node20}
	fmt.Println(levelOrderBottom(root))
}