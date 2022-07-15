package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var res int

	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for i := 0; i < len(boxTypes); i++ {
		for boxTypes[i][0] > 0 && truckSize > 0 {
			res += boxTypes[i][1]
			boxTypes[i][0]--
			truckSize--
		}
	}

	return res
}

func main() {
	arr := [][]int{{1, 3}, {2, 2}, {3, 1}}
	fmt.Println(maximumUnits(arr, 4))
}
