package main
/*
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	. "LeetCode/main/base"
	"fmt"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	var i int
	for i = 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	if i >= len(inorder) {
		return nil
	}
	return &TreeNode{Val:preorder[0],Left:buildTree(preorder[1:i+1], inorder[:i]),
	Right:buildTree(preorder[i+1:], inorder[i+1:])}
}

func main() {
	root := buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
	fmt.Println(LevelOrderTravel(root))
}