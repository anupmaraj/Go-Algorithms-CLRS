/*CLRS Prblem 2.1.3: Linear search scans through the sequence, 
looking for value 'v'. Using a loop invariant, prove that 
your algorithm is correct.

Linear Search-linear search or sequential search is a method for finding a particular value in a list that checks each element in sequence until the desired element is found or the list is exhausted.

*/

//CLRS asks to return nil value for array index variable if 
//value not found but in golang ->cannot assign nil value to int

package main 

import "fmt"
import "math/rand"

func linearsearch(a []int, v int){
	//to keep track of more than one occurance: counter
	var counter int											
	for i:=0; i<len(a); i++{
		if a[i]==v{
			fmt.Println("Value", v, "at position", i+1)
			counter++
		}
	}
	if counter==0{ 
		fmt.Println("Value", v, "not found in array")
	}
}

func main(){
	// array length stored in n
	n := 10												
	A := make([]int, n) 							
	for i := 0; i <= n - 1; i++ {	
        A[i] = rand.Intn(n)
    }
    fmt.Println("Array: ",A)
    //in case you want to define the array A yourself--> A:= []int{2,3,4,5,}			
	//value being searched V
	V := 0							
	linearsearch(A, V)
}
