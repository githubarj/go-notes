package main

import "fmt"

func createMatrix (rows , cols int) [][]int {
	matrix := [][]int{}
	for i := 0 ; i < rows ; i++ {
			row := []int{}
			for j := 0 ; j < cols ; j++ {
				row = append(row, i*j)
			}
			matrix = append(matrix, row)
	}
	return matrix
}

func printMatrix (m [][]int) {
	for i := 0 ; i < len(m) ; i ++ {
		fmt.Printf("%v\n", m[i])
	}
}

func main () {
	matrix := createMatrix(10,10)
	fmt.Printf("%v \n", matrix)
	printMatrix(matrix)
}