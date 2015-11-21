/*The Maximum Subarray problem is the task of finding the contiguous subarray within a one-dimensional array of 
numbers (containing at least one positive number) which has the largest sum.

CLRS: 4.1.1: FINDING THE MAXIMUM SUM OF SUBARRAY WHEN ALL ELEMENTS ARE NEGATIVE
CLRS: 4.1.2: FINDING THE MAXIMUM SUM OF SUBARRAY THROUGH BRUTE FORCE WITH TIME COMPLEXITY θ(n^2)
CLRS: 4.1.3: FINDING THE MAXIMUM SUM OF SUBARRAY THROUGH BRUTE FORCE

*/

package main 

import "fmt"

func MSS(arr []int) int {
	ans := Min_Element_of_Array(arr)	
	length:= len(arr)														
	for sub_array_size := 1; sub_array_size <= length; sub_array_size++ {
		for start_index := 0; start_index < length; start_index++ {						
			if(start_index + sub_array_size > length){
				break
			}
			sum := 0

			sum += arr[start_index + sub_array_size - 1];

			ans = Maximum(ans,sum)
		}
	}
	return ans
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

func Maximum(a, b int) int{
	if a > b{
		return a
	} else {
		return b
	}
}

func main() {
	arr := []int{-2, -5, 0, -20,-4, -400, -5}
	fmt.Println(MSS(arr))
}