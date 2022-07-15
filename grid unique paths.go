package main

import "fmt"

func uniquePaths(m int, n int) int {
	cur := make([]int, m)
	prev := make([]int, m)
	for i := 0; i < m; i++ {
		cur[i] = 1
		prev[i] = 1
	}

	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			cur[i] = cur[i-1] + prev[i]
		}
		cur, prev = prev, cur
	}
	return prev[m-1]
}

func main() {
	output1 := uniquePaths(2, 2)
	fmt.Println(output1)

	output2 := uniquePaths(2, 3)
	fmt.Println(output2)
}
