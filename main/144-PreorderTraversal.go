package main
/*
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	. "LeetCode/main/base"
	"fmt"
)
func preorderTraversal(root *TreeNode) []int {
	return PreOrderTravel(root)
}

func preorderTraversal2(root *TreeNode) []int {
	return PreOrderTravelWithoutRecursion(root)
}

func main() {
	node3 := &TreeNode{Val:3}
	node2 := &TreeNode{Val:2,Left:node3}
	root := &TreeNode{Val:1,Right:node2}
	fmt.Println(preorderTraversal(root))
	fmt.Println(preorderTraversal2(root))
}