/*Fibonacci Sequence: each number is the sum of the previous two numbers
CLRS 3.22
In this program, i is the ith Fibonacci number*/

package main 

import "fmt"

func Fibonacci(i int){
	var a, b, fib int = 0, 1, 0
	for ; i>2;i--{
		fib = a + b
		a, b = b, fib
	}
	fmt.Println(fib)
}

func main(){
	var i int
	fmt.Println("Enter the Fibaonnaci Number you want: ")
	fmt.Scanf("%d", &i)
	Fibonacci(i)
}