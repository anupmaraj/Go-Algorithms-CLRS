/* SQUARE MATRIX MULTIPLICATION: The Theory behind multiplication of matrices can be found here: https://en.wikipedia.org/wiki/Matrix_multiplication

CLRS 4.2*/

package main 

import "fmt"

func MatrixMultiplication(m1, m2 [][]int) [][]int{
	res := [][]int{{0,0},{0,0}}
	squarematrix := 2
	for i := 0; i < squarematrix; i++ {
		for j := 0; j < squarematrix; j++ {
			res[i][j] = 0
			for k := 0; k < squarematrix; k++ {
				res[i][j] = res[i][j] + (m1[i][k] * m2[k][j])
			}
		}	
	}
	return res
}

func CheckMatrixSize(m1, m2 [][]int) bool {
    //check if row lengths match
    if len(m1) != len(m2) {
        return false
    }
    //check if column lengths match
    for i, row1 := range m1 {
        row2 := m2[i]
        //checks for square matrix
        if len(row1) != len(m2) || len(row1) != len(row2) {
            return false
        }
    }
    return true
}

func main() {
	m1 := [][]int{{1,2},{4,6}}
	m2 := [][]int{{7,9}, {11,12}}
	if CheckMatrixSize(m1, m2) == true {	
		fmt.Println(MatrixMultiplication(m1,m2))
	} else {
		fmt.Println("Not a square matrix.")
	}
}