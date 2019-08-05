package main
/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	. "LeetCode/main/base"
	"fmt"
)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	leftChildTreeIsValid := true
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		leftChildTreeIsValid = isValidBST(root.Left)
	}
	rightChildTreeIsValid := true
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		rightChildTreeIsValid = isValidBST(root.Right)
	}
	return leftChildTreeIsValid && rightChildTreeIsValid
}

func main()  {
	node1 := &TreeNode{Val:1}
	node2 := &TreeNode{Val:2}
	node3 := &TreeNode{Val:3}
	node4 := &TreeNode{Val:4}
	node5 := &TreeNode{Val:5}
	node6 := &TreeNode{Val:6}

	node2.Left = node1
	node2.Right = node3
	fmt.Println(isValidBST(node2))

	node5.Left = node1
	node5.Right = node4
	node4.Left = node3
	node4.Right = node6
	fmt.Println(isValidBST(node5))
}

