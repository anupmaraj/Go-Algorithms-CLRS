/*CLRS Prblem 2.1.3: Linear search scans through the sequence, 
looking for value 'v'. Using a loop invariant, prove that 
your algorithm is correct.*/

//CLRS asks to return nil value for array index variable if 
//value not found but in golang ->cannot assign nil value to int

package main 

import "fmt"

func linearsearch(a []int, v int){
	var counter int		//to keep track of more than one occurance
	for i:=0; i<len(a); i++{
		if a[i]==v{
			fmt.Println("Value at position", i+1)
			counter++
		}
	}
	if counter==0{ 
		fmt.Println("Value not found in array")
	}
}

func main(){
	A := []int{4,7,2,12,6,8,4,7,9,21}
	V := 4		//value being searched
	linearsearch(A, V)
}