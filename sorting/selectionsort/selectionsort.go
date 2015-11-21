/*Selection Sort Algorithm: CLRS Problem 2.2.2

Selection Sort-The algorithm divides the input list into two parts: the sublist of items already sorted, which is built up from left to right at the front (left) of the list, and the sublist of items remaining to be sorted that occupy the rest of the list. Initially, the sorted sublist is empty and the unsorted sublist is the entire input list. The algorithm proceeds by finding the smallest (or largest, depending on sorting order) element in the unsorted sublist, exchanging (swapping) it with the leftmost unsorted element (putting it in sorted order), and moving the sublist boundaries one element to the right.

Worst Case Performance: O(n^2)
Average Case Performance:O(n^2)
Best Case Performance:O(n^2)
Worst Case Space Complexity:o(n) total, O(1) auxiliary

*/

package main

import "fmt"
import "math/rand"

func selectionSort(array []int){
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i] 
	}
}

func main() {
	//array of length n
	n := 12							
	array := make([]int, n) 		
	for i := 0; i <= n - 1; i++ {	
        array[i] = rand.Intn(n)
    }

	fmt.Println("Unsorted array: ", array)
	selectionSort(array)
	fmt.Println("Sorted array: ", array)
}
