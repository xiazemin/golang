package main

import "fmt"

type A int

func (a A) String() string {
	return fmt.Sprintf("%v", a)
	//return fmt.Sprintf("%d", int(a))
}
func (a A) Error() string {
	return fmt.Sprintf("failed to login(%x)", a)
	//死循环的原因在于：fmt.Sprintf会通过接口查询知道a是一个接口类型，所以就会调用a的Error函数，但这个fmt.Sprintf本身就是在Error函数里调用的，所以就构成循环调用了。

}
func main() {
	var a A
	fmt.Println("this will never print", a)
}

//runtime: goroutine stack exceeds 1000000000-byte limit
//fatal error: stack overflow