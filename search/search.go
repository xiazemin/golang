package search

import "github.com/golang/go/src/fmt"

func binarySearch(a []int, key int) int {
	low := 0
	high := len(a) - 1
	for low < high {
		/**
		利用位与（&）提取出两个数相同的部分，利用异或（^）拿出两个数不同的部分的和，相同的部分加上不同部分的和除2即得到两个数的平均值
		异或： 相同得零，不同得1 == 男男等零，女女得零，男女得子
		avg = (a&b)  + (a^b)>>1;
		或者
		avg = (a&b)  + (a^b)/2;
		*/
		mid := low&high + (low^high)>>1
		fmt.Println(low, mid, high)
		if key > a[mid] {
			low = mid + 1
		} else if key < a[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func insertSearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	time := 0
	for low < high {
		time += 1
		// 计算mid值是插值算法的核心代码 实际上就是
		mid := low + int((high-low)*(key-arr[low])/(arr[high]-arr[low]))
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

/**
基本思想：利用黄金分割 0.168 ：1 来确定中间值；也是二分查找一种改进版
用文字来说，就是费波那契数列由0和1开始，之后的费波那契系数就是由之前的两数相加而得出。
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233…… 特别指出：0不是第一项，而是第零项
数列的值为: F(0)=0，F(1)=1, F(n)=F(n-1)+F(n-2)（n>=2，n∈N*） n为数组下标
	|--------------- F(K)-1 ---------------|
	low					  mid        high
	|______________________|_______________|
	|------- F(K-1)-1 -----|--- F(K-2)-1 --|
他要求开始表中记录的个数为某个斐波那契数小1，即n = F(k)-1；开始将key值（要查找的数据）与第F(k-1)位置的记录进行比较(即mid = low + F(k-1) - 1)，比较结果也分为三种：
  （1）相等，mid位置的元素即为所求；
  （2）大于，low=mid+1，k-=2。说明：low=mid+1 :说明待查找的元素在[mid+1,high]范围内，k-=2 :说明范围[mid+1,high]内的元素个数为n-(F(k-1))= Fk-1-F(k-1)=Fk-F(k-1)-1=F(k-2)-1个，所以可以递归的应用斐波那契查找。
  （3）小于，high=mid-1，k-=1。说明：low=mid+1说明待查找的元素在[low,mid-1]范围内，k-=1 说明范围[low,mid-1]内的元素个数为F(k-1)-1个，所以可以递归 的应用斐波那契查找
*/
func fibonacciSearch(arr []int, key int) int {
	// 生成裴波那契数列，因为我们要满足 len(arr) = F(k) - 1
	fibArr := make([]int, 0)
	// 因为 斐波那契数列的性质我们知道数据递增的特别快，所以我们这里随机选择 生成的数列长度 36 够用了
	// [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368 75025 121393 196418 317811 514229 832040 1346269 2178309 3524578 5702887 9227465 14930352]
	for i := 0; i <= 36; i++ {
		fibArr = append(fibArr, fibonacci(i))
	}
	//fmt.Println(fibArr)

	// 确定待查找数组在裴波那契数列的位置
	k := 0
	n := len(arr)

	// 此处 n > fib[k]-1 也是别有深意的
	// 若n恰好是裴波那契数列上某一项，且要查找的元素正好在最后一位，此时必须将数组长度填充到数列下一项的数字
	for n > fibArr[k]-1 {
		k = k + 1
	}
	//fmt.Println(k, fibArr[k])
	// 将待查找数组填充到指定的长度
	for i := n; i < fibArr[k]; i++ {
		arr = append(arr, 0)
	}
	low, high := 0, n-1
	for low <= high {
		// 获取黄金分割位置元素下标
		mid := low + fibArr[k-1] - 1
		if key < arr[mid] {
			// 若key比这个元素小, 则key值应该在low至mid - 1之间，剩下的范围个数为F(k-1) - 1
			high = mid - 1
			k -= 1
		} else if key > arr[mid] {
			// 若key比这个元素大, 则key至应该在mid + 1至high之间，剩下的元素个数为F(k) - F(k-1) - 1 = F(k-2) - 1
			low = mid + 1
			k -= 2
		} else {
			if mid < n {
				return mid
			} else {
				return n - 1
			}
		}
	}
	return -1
}

/**
生成 斐波那契数列
*/

// 最屌写法
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	var fibarry = [3]int{0, 1, 0}
	for i := 2; i <= n; i++ {
		fibarry[2] = fibarry[0] + fibarry[1]
		fibarry[0] = fibarry[1]
		fibarry[1] = fibarry[2]
	}
	return fibarry[2]
}

//递归实现
func Fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return Fibo1(n-1) + Fibo1(n-2)
	} else {
		return -1
	}
}

//迭代实现
func Fibo2(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		a, b := 1, 1
		result := 0
		for i := 3; i <= n; i++ {
			result = a + b
			a, b = b, result
		}
		return result
	}
}

//利用闭包
func Fibo3(n int) int {
	if n < 0 {
		return -1
	} else {
		f := Fibonacci()
		result := 0
		for i := 0; i < n; i++ {
			result = f()
		}
		return result
	}
}
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
