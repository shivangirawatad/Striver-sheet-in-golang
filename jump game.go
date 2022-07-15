package main

import "fmt"

func jumpGame(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}

func main() {
	arr := []int{2, 3, 1, 1, 4}

	output := jumpGame(arr)
	fmt.Println(output)
}
