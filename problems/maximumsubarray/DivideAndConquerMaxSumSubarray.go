/*The Maximum Subarray problem is the task of finding the contiguous subarray within a one-dimensional array of 
numbers (containing at least one positive number) which has the largest sum.

FINDING THE MAXIMUM SUM OF SUBARRAY THROUGH DIVIDE AND CONQUER WITH TIME COMPLEXITY θ(n log n)

*/

package main 

import "fmt"

func Max_Subarray_Sum(arr []int) int{
	if len(arr) == 1 {
		return arr[0]
	}
	m := len(arr)/2
	left_MSS, right_MSS := arr[:m], arr[m:]

	leftsum := Min_Element_of_Array(arr)
	rightsum := Min_Element_of_Array(arr)
	sum := 0
	for i := m; i < len(arr); i++ {
		sum += arr[i];
		rightsum = Maximum(rightsum, sum);
	}
	sum = 0
	for i := (m-1); i >= 0; i--{
		sum += arr[i]
		leftsum = Maximum(leftsum,sum)
	}
	ans := Maximum(arr[:m],arr[m:])
	return Maximum(ans, leftsum+rightsum)
}

func Maximum(a, b int) int{
	if a > b{
		return a
	} else {
		return b
	}
}

func Min_Element_of_Array(arr []int) int {
	var minimum int
	for _, value := range arr {
		if minimum > value {
			minimum = value
		}
	}
	return minimum
}

func main() {
	arr := []int{3,-2,5,-1, 20}
	fmt.Println(Max_Subarray_Sum(arr))
}