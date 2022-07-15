package main

import "fmt"

func set0s(matrix [][]int) [][]int {
	numrows := len(matrix)
	numcols := len(matrix[0])

	row := make([]int, numrows)
	column := make([]int, numcols)

	for i := 0; i < numrows; i++ {
		for j := 0; j < numcols; j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				column[j] = 1
			}
		}
	}

	for i := 0; i < numrows; i++ {
		for j := 0; j < numcols; j++ {
			if row[i] == 1 || column[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	result := set0s(matrix)
	fmt.Println(result)
}
