package main
/*
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	. "LeetCode/main/base"
	"fmt"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	var i int
	for i = 0; i < len(postorder); i++ {
		if inorder[i] == postorder[len(postorder) - 1] {
			break
		}
	}
	return &TreeNode{Val:inorder[i],Left:buildTree(inorder[:i], postorder[:i]),
	Right:buildTree(inorder[i+1:], postorder[i:len(postorder) - 1])}
}

func main() {
	root := buildTree([]int{9,3,15,20,7}, []int{9,15,7,20,3})
	fmt.Println(LevelOrderTravel(root))
}