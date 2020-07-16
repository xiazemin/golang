package sort

import (
	"testing"
	"github.com/golang/go/src/fmt"
)

func TestSort(t *testing.T){
	a:=[]int{1,3,2,7,5,9,6,4,4,1,3}
	fmt.Println(a)
	//BubbleSort(a)
	//BubbleSort1(a,0,len(a)-1)
	//a=BucketSort(a)
	//a=CountSort(a)
	//HeapSort(a)
	//InsertSort(a)
	//a=MergeSort(a)
	//QuickSort(a,0,len(a)-1)
	//RadixSort(a)
	//SelectSort(a,0,len(a)-1)
	shellsort(a,0,len(a)-1)
	fmt.Println(a)


}


