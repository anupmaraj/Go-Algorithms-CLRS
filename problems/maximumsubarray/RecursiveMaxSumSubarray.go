/*The Maximum Subarray problem is the task of finding the contiguous subarray within a one-dimensional array of 
numbers (containing at least one positive number) which has the largest sum.

FINDING THE MAXIMUM SUM OF SUBARRAY THROUGH KADANE'S ALGORITHM WITH TIME COMPLEXITY θ(n)

CLRS: 4.1.1: FINDING THE MAXIMUM SUM OF SUBARRAY WHEN ALL ELEMENTS ARE NEGATIVE
*/

package main 

import "fmt"

func MSS(arr []int) int {
 	var ans, sum int 
	for i := 0; i < len(arr); i++{
		ans = max(ans, arr[i]) 	
	}
	if ans < 0 { 	
		return ans
	}
	for i := 0 ; i < len(arr); i++{
		if sum + arr[i] > 0 {
			sum += arr[i]
		} else {
			sum  = 0
		}
		ans = max(ans,sum)
	}
	return ans
}

func max(a, b int) int{
	if a > b{
		return a
	} else {
		return b
	}
}

func main() {
	arr := []int{3,-2,5,-1, 20, 8}
	fmt.Println(MSS(arr))
}