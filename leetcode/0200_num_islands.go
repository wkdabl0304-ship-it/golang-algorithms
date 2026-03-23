package main

import "fmt"

// 个人解法：深度优先搜索（DFS），缺点是容易系统栈溢出
//func numIslands(grid [][]byte) int {
//	if len(grid) == 0 {
//		return 0
//	}
//	res := 0
//	rows, columns := len(grid), len(grid[0])
//	var dfs func(i, j int)
//	dfs = func(i, j int) {
//		if i < 0 || i >= rows || j < 0 || j >= columns || grid[i][j] == '0' {
//			return
//		}
//		grid[i][j] = '0'
//		dfs(i-1, j)
//		dfs(i+1, j)
//		dfs(i, j-1)
//		dfs(i, j+1)
//	}
//	for i := 0; i < rows; i++ {
//		for j := 0; j < columns; j++ {
//			if grid[i][j] == '1' {
//				res++
//				dfs(i, j)
//			}
//		}
//	}
//	return res
//}

// 标准解法：广度优先搜索（BFS）
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	rows, columns := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == '1' {
				res++
				grid[i][j] = '0'
				queue := [][]int{{i, j}}
				for len(queue) > 0 {
					current := queue[0]
					queue = queue[1:]
					for _, v := range dirs {
						x, y := current[0]+v[0], current[1]+v[1]
						if x >= 0 && x < rows && y >= 0 && y < columns && grid[x][y] == '1' {
							grid[x][y] = '0'
							queue = append(queue, []int{x, y})
						}
					}
				}
			}
		}
	}
	return res
}

// 思路总结：如果用搜索算法，就是遍历然后去岛屿，无非一个是深度+系统栈，一个是广度+队列

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
