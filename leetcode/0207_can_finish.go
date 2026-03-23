package main

import "fmt"

// 标准解法：入度表 + 邻接表 + 层序遍历（BFS）
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1.初始化入度表和邻接表
	inDegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for _, v := range prerequisites {
		course, preCourse := v[0], v[1]
		inDegree[course]++
		adj[preCourse] = append(adj[preCourse], course)
	}
	// 2.初始化队列
	var queue []int
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	// 3.入队与出队
	count := 0
	for len(queue) > 0 {
		count++
		current := queue[0]
		queue = queue[1:]
		for _, v := range adj[current] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return count == numCourses
}

// 思路总结：涉及图的题目的时候一般用 邻接表 + 层序遍历（BFS）

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites))
}
