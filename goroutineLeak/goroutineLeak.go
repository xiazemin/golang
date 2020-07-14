package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {

	fmt.Println("main before")
	go func (){
	   fmt.Println("go func before")
		time.Sleep(1000)
		fmt.Println("go func after")
	}()
	fmt.Println("main after",runtime.NumGoroutine())
}
