package main
/*
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	. "LeetCode/main/base"
	"fmt"
)

func postorderTraversal(root *TreeNode) []int {
	return PostOrderTraver(root)
}

func postorderTraversal2(root *TreeNode) []int {
	return PostOrderTraverWithOutRecursion(root)
}

func main() {
	node3 := &TreeNode{Val:3}
	node2 := &TreeNode{Val:2,Left:node3}
	root := &TreeNode{Val:1,Right:node2}
	fmt.Println(postorderTraversal(root))
	fmt.Println(postorderTraversal2(root))
}