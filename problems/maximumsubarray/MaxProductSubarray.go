package main

import "fmt"

/*
https://oj.leetcode.com/problems/maximum-product-subarray/
Find the contiguous subarray within an array (containing at least one number) which has the largest product.
For example, given the array [2,3,-2,4],
the contiguous subarray [2,3] has the largest product = 6.
*/

func main() {
	array := []int{4, 5, -400, 1, -2, -10000}
	fmt.Println(maximumProductSubarray(array))
}
func maximumProductSubarray(array []int) int {
	var maxEndingHere, minEndingHere, maxSoFar int = 1, 1, 1
	for key, _ := range array {
		switch {
		case array[key] > 0:
			{
				maxEndingHere = maxEndingHere * array[key]
				minEndingHere = min(minEndingHere*array[key], 1)
			}
		case array[key] == 0:
			{
				maxEndingHere = 1
				minEndingHere = 1
			}
		default:
			{
				temp := maxEndingHere
				maxEndingHere = max(minEndingHere*array[key], 1)
				minEndingHere = temp * array[key]
			}
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}