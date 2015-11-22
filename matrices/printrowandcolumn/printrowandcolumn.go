/* Functions to print the rows and columns.

In Golang, arrays are one dimensional, and 2D matrices are represented as arrays of arrays.*/

package main 

import "fmt"

func PrintRows(x [][]int, n int) {
	fmt.Println("Rows are:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ",x[i][j])
		} 
		fmt.Printf("\n \n")
	}
}

func PrintColumns(x [][]int, n int) {
	fmt.Println("Columns are:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(x[j][i:i+1])
		} 
		fmt.Printf("\n")
	}
}

func main() {
	//square matrix of size n
	n := 3
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	PrintRows(matrix, n)
	PrintColumns(matrix, n)
}
