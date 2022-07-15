package main

import "fmt"

func longestseq(nums []int) int {
	set := make(map[int]bool)

	for idx := range nums {
		set[nums[idx]] = true
	}

	maxLen := 0
	for num := range set {
		if set[num-1] {
			continue
		}

		curLen := 0

		for set[num] {
			curLen++
			num++
		}

		maxLen = max(maxLen, curLen)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	arr := []int{100, 200, 1, 3, 2, 4}
	fmt.Println(longestseq(arr))
}
