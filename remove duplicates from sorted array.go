package main

import (
	"fmt"
)

func removeDuplicates(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	j := 1
	for i := 1; i < len(arr); i++ {
		if arr[i-1] != arr[i] {
			arr[j] = arr[i]
			j++
		}
	}

	return arr[0:j]
}

func main() {
	var arr = []int{1, 1, 2, 2, 2, 3, 3}
	fmt.Println(removeDuplicates(arr))
	arr = []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4}
	fmt.Println(removeDuplicates(arr))
	arr = []int{1}
	fmt.Println(removeDuplicates(arr))
}
