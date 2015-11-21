/*Merge Sort - Divide the unsorted list into n sublists, each containing 1 element (a list of 1 element is considered sorted). Then repeatedly merge sublists to produce new sorted sublists until there is only 1 sublist remaining. This will be the sorted list.

Worst Case Performance: O(n log n)
Average Case Performance:O(n log n) typical, O(n) natural variant
Best Case Performance:O(n log n)
Worst Case Space Complexity:O(n) total, O(n) auxiliary

*/

package main 

import "fmt"

func Merge(left, right []int) []int {
    result := make([]int, 0, len(left) + len(right))
    
    for len(left) > 0 || len(right) > 0 {			
        if len(left) == 0 {
            return append(result, right...)
        }
        if len(right) == 0 {
            return append(result, left...)
        }
        if left[0] <= right[0] {
            result = append(result, left[0])
            left = left[1:]
        } else {
            result = append(result, right[0])
            right = right[1:]
        }
    }
    return result
}

func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    middle := len(arr) / 2
    left := MergeSort(arr[:middle])
    right := MergeSort(arr[middle:])
    return Merge(left, right)
}

func main() {
    arr := []int{10,4,3,2,7,9,5,33,6,39,453454,88,12,16}
    fmt.Println("Initial array is:", arr)
    fmt.Println("Sorted array is: ", MergeSort(arr))
}
