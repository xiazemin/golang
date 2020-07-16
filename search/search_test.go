package search

import (
	"testing"
	"fmt"
)

func TestBinarySearch(t *testing.T){
  fmt.Println(binarySearch([]int{1,4,5,7,9,13,45,78},13))
	fmt.Println(insertSearch([]int{1,4,5,7,9,13,45,78},13))
	fmt.Println(fibonacciSearch([]int{1,4,5,7,9,13,45,78},13))


}
