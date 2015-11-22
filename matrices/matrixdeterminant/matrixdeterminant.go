/*Determinant of a matrix: https://en.wikipedia.org/wiki/Determinant

This program is for square matrices of size n=2 or n=3 only*/

package main 

import "fmt"

func MatrixDeterminant(m1 [][]int, matrixsize int) int {
	if matrixsize == 2{
		return (m1[0][0]*m1[1][1]) - (m1[0][1]*m1[1][0])
	} 
	return (m1[0][0]*m1[1][1]*m1[2][2]-m1[0][0]*m1[1][2]*m1[2][1]-m1[0][1]*m1[1][0]*m1[2][2]+m1[0][1]*m1[1][2]*m1[2][0]+m1[0][2]*m1[1][0]*m1[2][1]-m1[0][2]*m1[1][1]*m1[2][0])
}

func CheckMatrixSize(m1 [][]int) bool {
    //check for square matrix
    for _, row1 := range m1 {
        if len(row1) != len(m1) {
            return false
        }
    }
    return true
}

func main() {
	m1 := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	matrixsize := 2
	if CheckMatrixSize(m1) == true && (matrixsize ==2 || matrixsize == 3){	
		fmt.Println("Deteminant is: ", MatrixDeterminant(m1,matrixsize))
	} else {
		fmt.Println("Not a square matrix of size n = 2 or n = 3.")
	}
}