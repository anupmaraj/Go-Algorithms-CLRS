/*Random Search: Pick a random index i into array A. If a[i] is equal to the element x you are searching for, terminate. Else, continue the search by picking a new random index into A.
We continue picking random indices into A until we find an index j such that A[j] =x or until we have checkedpos every element of A. 
Note that we pick from the whole set of indices each time, so that we may examine a given element more than once.

CLRS Problem: 5.2.a*/


package main 

import(
    "fmt"
    "math/rand"
    "time"
)

var checkedpos []int


func RandomSearch(arr []int, n int) (int) {
	var found bool
	// var pos int

	for found != true {
		rand.Seed(time.Now().Unix())
		pos := rand.Intn(len(arr))

		if arr[pos] == n {
			found = true
		}

		for _, value := range checkedpos {
			if pos == value{
				RandomSearch(arr, n)
			} 
		} 
		checkedpos = append(checkedpos,pos)
	}
	fmt.Println(checkedpos)
	return -1
}

// func RandomPosition(max int) int {
// 	rand.Seed(time.Now().Unix())
// 	pos := rand.Intn(max)
// 	return pos
// }

func main(){
	arr := []int{4,1,8,3,10,5}	
	position := RandomSearch(arr, 4)
	fmt.Println("Found at position", position+1)
}