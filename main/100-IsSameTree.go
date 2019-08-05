package main
/*
New
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true
示例 2:

输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false
示例 3:

输入:       1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

输出: false
 */
import (
	. "LeetCode/main/base"
	"fmt"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	node2 := &TreeNode{Val:2}
	node3 := &TreeNode{Val:3}

	root1 := &TreeNode{Val:1,Left:node2,Right:node3}
	root2 := &TreeNode{Val:1,Left:node2,Right:node3}
	fmt.Println(isSameTree(root1,root2))

	root1.Right = nil
	root2.Left = nil
	root2.Right = node2
	fmt.Println(isSameTree(root1, root2))

	root1.Right = &TreeNode{Val:1}
	root2.Left = &TreeNode{Val:1}
	fmt.Println(isSameTree(root1, root2))
}