package main
/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	. "LeetCode/main/base"
	"LeetCode/main/container"
	"fmt"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var node *TreeNode
	depth := 0
	queue1 := container.NewQueue()
	queue2 := container.NewQueue()
	queue1.Push(root)
	for !queue1.IsEmpty() {
		depth += 1
		for !queue1.IsEmpty() {
			node = queue1.Pop().(*TreeNode)
			if node.Left != nil {
				queue2.Push(node.Left)
			}
			if node.Right != nil {
				queue2.Push(node.Right)
			}
		}
		queue1 = queue2
		queue2 = container.NewQueue()
	}
	return depth
}

func main() {
	node20 := &TreeNode{Val:20,Left:&TreeNode{Val:15},Right:&TreeNode{Val:7}}
	root := &TreeNode{Val:3,Left:&TreeNode{Val:9},Right:node20}
	fmt.Println(maxDepth(root))
}