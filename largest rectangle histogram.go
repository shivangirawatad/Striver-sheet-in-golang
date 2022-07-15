package main

import "fmt"

func largestarea(heights []int) int {
	max := 0
	hLen := len(heights)
	for i, iv := range heights {
		sum := iv

		for l := i - 1; l >= 0; l-- {
			if heights[l] < iv {
				break
			}
			sum += iv
		}
		for r := i + 1; r < hLen; r++ {
			if heights[r] < iv {
				break
			}
			sum += iv
		}

		if max < sum {
			max = sum
		}
	}
	return max
}

func main() {
	arr := []int{2, 4}
	fmt.Println(largestarea(arr))
}
