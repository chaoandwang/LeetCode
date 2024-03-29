package main

import "fmt"

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	sum := 0
	if len(grid) == 1 {
		for i := 0; i < len(grid[0]); i++ {
			sum += grid[0][i]
		}
		return sum
	}
	if len(grid[0]) == 1 {
		for i := 0; i < len(grid); i++ {
			sum += grid[i][0]
		}
		return sum
	}
	tmp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		tmp[i] = grid[i][:len(grid[i]) - 1]
	}
	return min(minPathSum(grid[:len(grid) - 1]), minPathSum(tmp)) + grid[len(grid) - 1][len(grid[0]) - 1]
}

func main() {
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	fmt.Println(minPathSum(grid))
}

