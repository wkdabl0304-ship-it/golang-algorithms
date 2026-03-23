package main

import "fmt"

// 标准解法：广度优先搜索（BFS）+ 层序遍历，题目难点在于大量边界条件的控制
func orangesRotting(grid [][]int) int {
	var queue [][]int
	total, count := 0, 0
	rows, columns := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 1 || grid[i][j] == 2 {
				total++
			}
			if grid[i][j] == 2 {
				count++
				queue = append(queue, []int{i, j})
			}
		}
	}
	if total == 0 {
		return 0
	}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	res := -1
	for len(queue) > 0 {
		levelNum := len(queue)
		for i := 0; i < levelNum; i++ {
			current := queue[0]
			queue = queue[1:]
			// 单次污染
			for _, v := range dirs {
				x, y := current[0]+v[0], current[1]+v[1]
				if x >= 0 && x < rows && y >= 0 && y < columns && grid[x][y] == 1 {
					count++
					grid[x][y] = 2
					queue = append(queue, []int{x, y})
				}
			}
		}
		res++
	}
	if count != total {
		res = -1
	}
	return res
}

func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}
