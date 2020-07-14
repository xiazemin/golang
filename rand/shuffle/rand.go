package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	intArr := []int{1,2,3,4,5,6,7,8,9}
	for i := 0; i < 10; i++{
		shuffle(intArr)
		fmt.Println(intArr)
	}
}

func shuffle(arr []int){
	rand.Seed(time.Now().UnixNano())

	for i:= len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

}