package main
/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例: 
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

import (
	. "LeetCode/main/base"
	"LeetCode/main/container"
	"fmt"
)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {return false}
	stack := container.NewStack()
	current := 0
	var snode *TreeNodeStackElement
	for root != nil || !stack.IsEmpty() {
		if root != nil {
			current += root.Val
			if current > sum {
				current -= root.Val
				root = nil
			} else if current == sum {
				if root.IsLeaf() {
					return true
				} else {
					current -= root.Val
					root = nil
				}
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
	return false
}

func main() {
	node4 := &TreeNode{Val:4,Left:&TreeNode{Val:11,Left:&TreeNode{Val:7},Right:&TreeNode{Val:2}}}
	node8 := &TreeNode{Val:8,Left:&TreeNode{Val:13},Right:&TreeNode{Val:4,Right:&TreeNode{Val:1}}}
	root := &TreeNode{Val:5,Left:node4,Right:node8}
	fmt.Println(hasPathSum(root, 22))
}