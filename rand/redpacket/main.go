package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i ++{
		redPackage(10, 500)
		fmt.Println("")
	}
}

func redPackage(count,money int) {
	for i:=0;i<count;i++{
		m:=getMoney(count-i,money)
		money-=m
		fmt.Println(m)
	}

}

func getMoney(remainCount,money int)int{
	if remainCount==1{
		return money
	}
	rand.Seed(time.Now().UnixNano())
	return 1+rand.Intn(money/remainCount*2)
}