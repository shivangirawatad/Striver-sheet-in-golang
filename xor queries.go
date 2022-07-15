package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}

	q := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}

	fmt.Println(xor(arr, q))
}

func xor(arr []int, q [][]int) []int {

	prefix := make([]int, len(arr)+1)
	prefix[0] = 0

	for i := 1; i < len(arr)+1; i++ {

		prefix[i] = prefix[i-1] ^ arr[i-1]
	}

	res := make([]int, len(q))
	for i, q := range q {

		res[i] = prefix[q[0]] ^ prefix[q[1]+1]
	}

	return res
}
