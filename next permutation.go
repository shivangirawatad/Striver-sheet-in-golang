package main

import (
	"fmt"
)

func nextPermutation(nums []int) []int {
	l := len(nums)
	idx := -1
	for i := l - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			idx = i
			break
		}
	}
	low, high := idx+1, l-1

	if idx >= 0 {
		for low <= high {
			mid := (low + high) / 2
			if nums[mid] > nums[idx] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		nums[idx], nums[high] = nums[high], nums[idx]
	}

	for i, j := idx+1, l-1; i < j && nums[i] > nums[j]; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func main() {
	arr := []int{1, 2, 3}
	output := nextPermutation(arr)
	fmt.Println(output)
}
