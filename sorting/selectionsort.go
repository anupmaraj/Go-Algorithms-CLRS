//Selection Sort Algorithm: CLRS Problem 2.2.2

package main

import "fmt"
import "math/rand"

func selectionSort(array []int){
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i] 	//swap values
	}
}

func main() {
	n := 10							//array length
	array := make([]int, n) 		//generate array
	for i := 0; i <= n - 1; i++ {	//fill araay with random int values
        array[i] = rand.Intn(n)
    }

	fmt.Println("Unsorted array: ", array)
	selectionSort(array)
	fmt.Println("Sorted array: ", array)
}