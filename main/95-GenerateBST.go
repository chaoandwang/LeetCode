package main

import "fmt"

/*
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例:

输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func generate(interval, n int) []*TreeNode {
	if n < 1 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{Val:n+interval}}
	}
	roots := make([]*TreeNode, 0)
	for _,root := range generate(interval + 1, n - 1) {
		if root == nil {
			continue
		}
		roots = append(roots, &TreeNode{Val:1+interval, Right:root})
	}
	for i := 2; i < n; i++ {
		for _,lroot := range generate(interval, i - 1) {
			for _,rroot := range generate(interval + i, n - i) {
				if lroot == nil || rroot == nil {
					continue
				}
				roots = append(roots, &TreeNode{Val:i+interval, Left:lroot, Right:rroot})
			}
		}
	}
	for _,root := range generate(interval, n -1) {
		if root == nil {
			continue
		}
		roots = append(roots, &TreeNode{Val:n + interval, Left:root})
	}
	return roots
}


func generateTrees(n int) []*TreeNode {
	return generate(0, n)
}

type queue struct {
	 que []interface{}
	 index int
}

func (q *queue) push(i interface{}) {
	q.que = append(q.que, i)
}

func (q *queue) pop() interface{} {
	q.index += 1
	ret := q.que[q.index]
	return ret
}

func (q *queue) isEmpty() bool {
	if len(q.que) == q.index + 1 {
		return true
	}
	return false
}

func newQueue() *queue {
	return &queue{que:make([]interface{}, 0), index:-1}
}

func printTree(root *TreeNode) []string {
	if root == nil {
		return []string{"nil"}
	}
	desc := make([]string, 0)
	que := newQueue()
	que.push(root)
	for !que.isEmpty() {
		tmp := que.pop().(*TreeNode)
		if tmp != nil {
			desc = append(desc, fmt.Sprintf("%d", tmp.Val))
			if tmp.Left == nil && tmp.Right == nil {
				continue
			}
			que.push(tmp.Left)
			que.push(tmp.Right)
		} else {
			desc = append(desc, fmt.Sprint("nil"))
		}
	}
	return desc
}

func printTrees(roots []*TreeNode) [][]string {
	if len(roots) < 1 {
		return [][]string{{"nil"}}
	}
	descs := make([][]string, 0)
	for _,root := range roots {
		descs = append(descs, printTree(root))
	}
	return descs
}

func main() {
	roots := generateTrees(3)
	fmt.Println(printTrees(roots))
}