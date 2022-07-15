package main

import (
	"fmt"
)

func sort(arr []int) {

	if len(arr) < 1 {
		return
	}
	i := 0
	k := 0
	j := len(arr) - 1

	for k <= j {
		//fmt.Println(arr)
		if arr[k] == 0 {
			arr[i], arr[k] = arr[k], arr[i]
			i++
			k++
		} else if arr[k] == 1 {
			k++
		} else {
			arr[j], arr[k] = arr[k], arr[j]
			j--
		}
	}
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}

	sort(arr)
	fmt.Println(arr)
}
