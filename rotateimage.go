package main

import "fmt"

func rotate(arr [][]int) [][]int {
	left := 0
	right := len(arr) - 1

	for left < right {
		top, bot := left, right
		for i := 0; i < right-left; i++ {
			toplft := arr[top][left+i]

			arr[top][left+i] = arr[bot-i][left]
			arr[bot-i][left] = arr[bot][right-i]
			arr[bot][right-i] = arr[top+i][right]
			arr[top+i][right] = toplft
		}

		left++
		right--
	}
	return arr
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	res := rotate(arr)
	fmt.Println(res)

}
