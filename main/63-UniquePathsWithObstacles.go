package main

import "fmt"

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？



网格中的障碍物和空位置分别用 1 和 0 来表示。

说明：m 和 n 的值均不超过 100。

示例 1:

输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) < 1 || len(obstacleGrid[0]) < 1 {
		return 0
	}
	if obstacleGrid[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1] == 1 {
		return 0
	}
	if len(obstacleGrid) == 1 {
		for i := 0; i < len(obstacleGrid[0]); i++ {
			if obstacleGrid[0][i] == 1 {
				return 0
			}
		}
		return 1
	}
	if len(obstacleGrid[0]) == 1 {
		for i := 0; i < len(obstacleGrid); i++ {
			if obstacleGrid[i][0] == 1 {
				return 0
			}
		}
		return 1
	}
	tmp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		tmp[i] = obstacleGrid[i][:len(obstacleGrid[i]) - 1]
	}
	return uniquePathsWithObstacles(obstacleGrid[:len(obstacleGrid) - 1]) +
		uniquePathsWithObstacles(tmp)
}

func main() {
	grid := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	fmt.Println(uniquePathsWithObstacles(grid))
}