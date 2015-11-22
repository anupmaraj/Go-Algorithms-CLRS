/*Bubble Sort-repeatedly steps through the list to be sorted, 
compares each pair of adjacent items and swaps them if they 
are in the wrong order. 

Worst Case Performance: O(n^2)
Average Case Performance:O(n^2)
Best Case Performance:O(n^2)
Worst Case Space Complexity:O(1) auxiliary

*/

package main 

import "fmt"

func BubbleSort(arr []int){
	//index value for array
	length := len(arr)-1			
	for i:=0; i<length-1; i++{
		for j:=length; j>i; j--{
			if arr[j] < arr[j-1]{ //Check if previous value is greater than current value
				arr[j], arr[j-1] = arr[j-1], arr[j] //Swap previous value with current value
			}
		}
	}
}

func main(){
    arr := []int{10,4,3,2,7,9,5,88,12,16}
    fmt.Println("Initial array is:", arr)
    BubbleSort(arr)
    fmt.Println("Sorted array is: ", arr)
}
