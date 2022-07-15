package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	res := nums[0]
	times := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			times++
		} else {
			times--
			if times < 0 {
				res = nums[i]
				times = 1
			}
		}
	}
	return res
}

func main() {
	arr := []int{3, 2, 3}

	fmt.Println(majorityElement(arr))
}
