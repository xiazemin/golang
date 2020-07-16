package search

import "fmt"

func HeapSearchK (arr []int, topk int) {

	// 初始化原始最小堆
	smallHeapArr := buildSmallHeap(arr, topk)
	for i := topk; i < len(arr); i ++ {
		// 如果当前原始比最小堆的根元素大，那么替换根，且重新调整最小堆
		if arr[i] > smallHeapArr[0] {
			swapRoot(smallHeapArr, arr[i])
		}
	}
	// 最大的K个数
	fmt.Println(smallHeapArr)
}

//建立小顶堆 
func buildSmallHeap(arr []int,topk int) []int{
	smallHeapArr := arr[:topk]
	for i := topk>>1 - 1; i >= 0; i-- {
		adjustSmallHeap(smallHeapArr, i, topk)
	}
	return smallHeapArr
}
// 调整最小堆
func adjustSmallHeap(arr []int, parent, length int) {
	i := parent
	for  {
		lchild := 2*parent + 1
		rchild := 2 *parent + 2
		if lchild < length && arr[lchild] < arr[i] {
			i = lchild
		}
		// 右节点和根
		if rchild < length && arr[rchild] < arr[i]{
			i = rchild
		}
		// 互换位置
		if parent != i {
			arr[i], arr[parent] = arr[parent], arr[i]
			parent = i
		}else {
			break
		}
	}
}

// 替换根部，且重新构造最小堆
func swapRoot(arr []int, root int) {
	arr[0] = root // 新的根
	// 重新调整堆
	adjustSmallHeap(arr, 0, len(arr))
}