package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int

	if len(candidates) == 0 {
		return ans
	}

	bt(&ans, make([]int, 0), candidates, 0, target)

	return ans
}

func bt(ans *[][]int, temp, nums []int, idx, target int) {
	if target < 0 || idx > len(nums) {
		return
	}

	if target == 0 {
		cpyTmp := make([]int, len(temp))
		copy(cpyTmp, temp)
		*ans = append(*ans, cpyTmp)
	}

	for i := idx; i < len(nums); i++ {
		temp = append(temp, nums[i])
		bt(ans, temp, nums, i, target-nums[i])
		temp = temp[:len(temp)-1]
	}
}

func main() {
	arr := []int{2, 3, 6, 7}
	target := 7

	fmt.Println(combinationSum(arr, target))
}
