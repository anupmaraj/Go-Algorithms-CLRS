//CLRS problem 2.1 and 2.1.2 combined
//Implement the insertion sort algorithm to sort an array 
//in both ascending and descending order

package main 

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