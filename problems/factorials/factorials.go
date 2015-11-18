/* Program to calculate the Factorial for any number 'n'

CLRS Section 3.18
*/

package factorial

import "fmt"

func factorial(n int){
	var res int = n
	for  ; n>1; n--{
	res = res*(n-1)
	}
	fmt. Println(res)
}

func main(){
	var n int = 10
	// fmt.Print("Enter number whose factorial you want to calculate: ")
 //    fmt.Scanf("%d", &n)
	factorial(n)
}