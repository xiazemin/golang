package op

import (
	"strconv"
	"fmt"
)

type Operator struct {

}

func (o*Operator)Priority(c byte)int{
	switch c {
	case '+','-':
		return 1;
	case '*','/':
		return 2;
	case '(',')':
		return 3;
	default:
	}
	return 0;
}

func (o*Operator)IsOperator(c byte)bool{
	arr:=[]byte{'+','-','*','/','(',')'}
	for i:=0;i<len(arr);i++{
		if arr[i]==c{
			return true
		}
	}
	return false
}

func (o*Operator)Compute(opl,opr,op string) string {
    l,_:=strconv.ParseInt(opl,10,64)
	r,_:=strconv.ParseInt(opr,10,64)
	switch op {
	case "+":
		return fmt.Sprint(l+r)
	case "-":
		return fmt.Sprint(l-r)
	case "*":
		return fmt.Sprint(l*r)
	case "/":
		return fmt.Sprint(l/r)
	}
	return "0"
}