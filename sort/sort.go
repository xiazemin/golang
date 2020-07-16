package sort

import "github.com/golang/go/src/fmt"

func BubbleSort(a[]int){
	for i:=0;i<len(a);i++{
		for j:=i+1;j<len(a);j++{
			if a[i]>a[j]{
				a[i],a[j]=a[j],a[i]
			}
		}
	}
}

func BubbleSort1(a[]int ,min,max int){
	if max==0{
		return
	}
	for i:=min;i<max;i++{
		if a[i]>a[max]{
			a[i],a[max]=a[max],a[i]
		}
	}
	 BubbleSort1(a,min,max-1)
}

func BucketSort(a[]int)[]int{
	size:=5
	min:=a[0]
	max:=a[0]
	for i:=0;i<len(a);i++{
		if a[i]<min{
			min=a[i]
		}
		if a[i]>max{
			max=a[i]
		}
	}

	fmt.Println(min,max)
	max-=min
	max/=size
	fmt.Println(min,max)
	bucket:=make([][]int,max+1)
	for i:=0;i<len(a);i++{
		index:=(a[i]-min)/size
		bucket[index]=append(bucket[index],a[i])
	}
	for i:=0;i<=max;i++{
		if len(bucket[i])>0{
			BubbleSort(bucket[i])
		}
	}
fmt.Println(bucket)
	a=[]int{}
	for _,v:=range bucket{
		a=append(a,v...)
		fmt.Println(a)
	}
return a
}


func CountSort(a[]int)[]int{
	min:=a[0]
	max:=a[0]
	for i:=0;i<len(a);i++{
		if a[i]<min{
			min=a[i]
		}
		if a[i]>max{
			max=a[i]
		}
	}

	nums:=make([]int,max-min+1)
	for _,v:=range a{
		nums[v-min]++
	}
	a=[]int{}
	for i:=0;i<max-min+1;i++{
		for j:=0;j<nums[i];j++{
			a=append(a,i+min)
		}
	}
	return a
}

func HeapSort(a []int){
	l:=len(a)-1
	for{
		if l<1{
			break
		}
		for p:=l/2;p>=0;p--{
			swap(a,p,l)

		}
		a[0],a[l]=a[l],a[0]
		l--
	}
}

func swap(a []int,p,l int){
	lc:=2*p+1
	rc:=2*p+2
	if  lc<=l && a[p]<a[lc]{
		a[p],a[lc]=a[lc],a[p]
	}
	if rc<=l && a[p]<a[rc]{
		a[p],a[rc]=a[rc],a[p]
	}
}


func InsertSort(a[]int){
	for i:=1;i<len(a);i++{
		for j:=i;j>0;j--{
			if a[j]<a[j-1]{
				a[j],a[j-1]=a[j-1],a[j]
			}else{
				break
			}
		}
	}
}

func MergeSort(a[]int)[]int{
	if len(a)<2{
		return a
	}
	m:=len(a)/2
	return merge(MergeSort(a[:m]),MergeSort(a[m:]))
}

func merge(a,b []int)(r []int){
	ia:=0
	ib:=0

	for ia<len(a)&&ib<len(b){
		if a[ia]<b[ib]{
			r=append(r,a[ia])
			ia++
		}else{
			r=append(r,b[ib])
			ib++
		}
	}

	if ia<len(a){
		r=append(r,a[ia:]...)
	}
	if ib<len(b){
		r=append(r,b[ib:]...)
	}
	fmt.Println(a,b,r,ia,ib)
	return r
}

func QuickSort(a[]int,l,r int){
	if r<=l{
		return
	}
	l0:=l
	r0:=r
	p:=l
	for l<r {
		for p < r && a[p] <= a[r] {
			r--
		}
		if p < r {
			a[p], a[r] = a[r], a[p]
			p = r
		}
		for l < p && a[l] <= a[p] {
			l++
		}
		if l < p {
			a[l], a[p] = a[p], a[l]
			p = l
		}
	}

	QuickSort(a,l0,p-1)
	QuickSort(a,p+1,r0)
	return
}

func diglen(v int)int{
	if v<0{
		v=-v
	}
	i:=0
	for v>0{
		i++
		v/=10
	}
	return i
}

func getDig(v,i int)int{
    l:=diglen(v)
	if i>l-1{
		return 0
	}
	for i>0 {
		v/=10
		i--
	}
	return v%10
}
func RadixSort(a[]int) {
	baseList := make([][]int, 10)
	max:=diglen(a[0])
	for i:=0;i<len(a);i++{
		if diglen(a[i])>max{
			max=diglen(a[i])
		}
	}
	for i:=0;i<max;i++{
		for _,v:=range a{
			baseList[getDig(v,i)]=append(baseList[getDig(v,i)],v)
		}
		j:=0
		for i,v:=range baseList{
			if len(v)>0{
				for _,v1:=range v{
					a[j]=v1
					j++
				}
			}
			baseList[i]=[]int{}
		}
	}
}

func SelectSort(a []int,l,r int)  {
	if l==r{
		return
	}
	min:=l
	for i:=l;i<=r;i++{
		if a[i]<a[min]{
			min=i
		}
	}
	a[l],a[min]=a[min],a[l]
	SelectSort(a,l+1,r)
}

func shellsort(a[]int,l,r int)  {
  inc:=len(a)/3+1
	for{
		if inc<1{
			break
		}
		for i:=l;i<=r;i+=inc{
			for j:=i+inc;j<=r;j+=inc{
				if a[j-inc]>a[j] {
					a[j - inc], a[j] = a[j], a[j - inc]
					fmt.Println(i,j,j-inc)
				}
			}

		}
		fmt.Println(inc)
		inc--
	}
}