package main

import "fmt"

type Solve struct {
	nums []int
	k    int
}

func maxwindow(arr []int, k int) []int {
	result := make([]int, 0)
	q := make([]int, 0)

	for i, num := range arr {
		for len(q) > 0 && arr[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		if i >= k-1 {
			result = append(result, arr[q[0]])
		}

		if q[0] == i-k+1 {
			q = q[1:]
		}
	}

	return result
}

func main() {
	arr := []int{4, 0, -1, 3, 5, 3, 6, 8}
	fmt.Println(maxwindow(arr, 3))
}
