/* The Theory behind addition of matrices can be found here: https://en.wikipedia.org/wiki/Matrix_addition
*/

package main 

import "fmt"

func MatrixAddition(m1 [][]int, m2 [][]int, rows int, columns int) [][]int{
	res := [][]int{{0,0,0},{0,0,0}}
	for i := 0; i < rows; i++ {
		for j:= 0; j < columns; j++ {
			res[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return res
}

func main() {
	m1 := [][]int{{1,2,3},{4,5,6}}
	m2 := [][]int{{7,8,9}, {10,11,12}}
	//check if lengths of both matrices are equal
	if len(m1) == len(m2){	
		rows := len(m1)
		columns := 3
		fmt.Println(MatrixAddition(m1,m2,rows,columns))
	} else {
		fmt.Println("Matrices are not of the same size. Addition not possible.")
	}
}