package main
/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
 
说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	. "LeetCode/main/base"
	"LeetCode/main/container"
	"fmt"
)

func FoundPath(root, p *TreeNode) *container.Stack {
	if root == nil || p == nil {
		return nil
	}
	var enode *TreeNodeStackElement
	stack := container.NewStack()
	for root != nil || !stack.IsEmpty() {
		if root != nil {
			stack.Push(&TreeNodeStackElement{root, true})
			if root.Val == p.Val {
				return stack
			}
			root = root.Left
		} else {
			enode = stack.Pop().(*TreeNodeStackElement)
			if enode.IsLeft && enode.Node.Right != nil {
				stack.Push(&TreeNodeStackElement{enode.Node, false})
				root = enode.Node.Right
			} else {
				root = nil
			}
		}
	}
	return stack
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q== nil {
		return nil
	}
	stack := FoundPath(root, p)
	var path *container.Stack
	for !stack.IsEmpty() {
		root = stack.Pop().(*TreeNodeStackElement).Node
		path = FoundPath(root, q)
		if !path.IsEmpty() {
			return root
		}
	}
	return nil
}

func main()  {
	node4 := &TreeNode{Val:4}
	node5 := &TreeNode{Val:5,Left:&TreeNode{Val:6},Right:&TreeNode{Val:2,Left:&TreeNode{Val:7},Right:node4}}
	node1 := &TreeNode{Val:1,Left:&TreeNode{Val:0},Right:&TreeNode{Val:8}}
	root := &TreeNode{Val:3, Left:node5, Right:node1}
	pNode := lowestCommonAncestor(root, node5, node1)
	if pNode == nil {
		fmt.Print("No Found")
	} else {
		fmt.Println(pNode.Val)
	}
	pNode = lowestCommonAncestor(root, node5, node4)
	if pNode == nil {
		fmt.Print("No Found")
	} else  {
		fmt.Println(pNode.Val)
	}
}
