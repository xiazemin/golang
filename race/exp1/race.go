package main

import "time"

func main(){
	var x int
	go func(){
		for{
			x=1
		}
	}()

	go func(){
		for{
			x=2
		}
	}()
	time.Sleep(100*time.Second)
}