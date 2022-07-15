package main

import "fmt"

func main() {
	var row int
	var temp int = 1
	fmt.Print("Enter number of rows : ")
	fmt.Scan(&row)

	for i := 0; i < row; i++ {

		for j := 1; j <= row-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
			fmt.Printf(" %d", temp)
		}
		fmt.Println("")
	}
}
