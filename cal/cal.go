package main

import (
	"fmt"
	"github.com/xiazemin/cal/op"
)

func main() {
	//fmt.Println(op.Translate("3+3*4"),"\t",op.Compute(op.Translate("3+3*4")))
	//fmt.Println(op.Translate("(2+1)*3"),"\t",op.Compute(op.Translate("(2+1)*3")))
	//fmt.Println(op.Translate("2*(3+5)-4+7/1"),"\t",op.Compute(op.Translate("2*(3+5)-4+7/1")))
	//35+2*4-71/+
	//fmt.Println(op.Translate("3*4/5"),"\t",op.Compute(op.Translate("3*4/5")))
	fmt.Println(op.Translate("3+3*4/5+7*(6+7*8/2-2*(6+9*3))"),op.Compute(op.Translate("3+3*4/5+7*(6+7*8/2-2*(6+9*3))")))
}

