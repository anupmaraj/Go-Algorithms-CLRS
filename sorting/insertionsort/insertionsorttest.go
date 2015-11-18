package insertion

import (
	"testing"
	"fmt"
)

func benchmarkInsertionSort(n int, b *testing.B) {
	list := []int{5,6,7,8,2,43,7,8,8,4,2,1,5,6,8,8,4,3,2,5}
	for i := 0; i < b.N; i++ {
		insertionsortdescending(list)
	}
	fmt.Println(list)
}

func BenchmarkInsertionSort1(b *testing.B)    { benchmarkInsertionSort(1, b) }
func BenchmarkInsertionSort5(b *testing.B)   { benchmarkInsertionSort(5, b) }
func BenchmarkInsertionSort10(b *testing.B)  { benchmarkInsertionSort(10, b) }
func BenchmarkInsertionSort15(b *testing.B) { benchmarkInsertionSort(15, b) }