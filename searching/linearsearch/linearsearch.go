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
	var counter int		//to keep track of more than one occurance
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
	n := 10							//array length
	A := make([]int, n) 			//generate array
	for i := 0; i <= n - 1; i++ {	//fill araay with random int values
        A[i] = rand.Intn(n)
    }
    fmt.Println("Array: ",A)
    //A:= []int{2,3,4,5,}		//in case you want to define the array yourself
	V := 0		//value being searched
	linearsearch(A, V)
}
