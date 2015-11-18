/*CLRS problem 2.1 and 2.1.2 combined
Implement the Insertion sort algorithm to sort an array in both ascending and descending order

Insertion Sort-simple sorting algorithm that builds the final sorted array (or list) one item at a time. 

Worst Case Performance: O(n^2)
Average Case Performance:Ω(n) comparisons, O(1) swaps
Best Case Performance:O(n^2) comparisons, swaps
Worst Case Space Complexity:O(n) total, O(1) auxiliary

*/

package insertionsort 

import "fmt"

func insertionsortdescending(list []int){
	for j:=1; j<len(list);j++{
		key:=list[j]
		i:=j-1
		for i>=0 && list[i] < key{
			list[i+1] = list[i]
			i= i-1
		}
		list[i+1]=key
	}
}

func insertionsortascending(list []int){
    for j:=1; j<len(list); j++{
        key:= list[j]
        i:=j-1
        for i>=0 && list[i] > key{
            list[i+1] = list[i]
            i = i-1
        } 
        list[i+1] = key
    }
}


func main(){
	list:= []int{12,4,56,22,10,7,1,0, 99, 654, 88}
	fmt.Println("\nUnsorted List=", list)

	insertionsortdescending(list)
	fmt.Println("\nDescending Sorted List=", list)	

	insertionsortascending(list)
    fmt.Println("\nAscending Sorted List", list)

}
