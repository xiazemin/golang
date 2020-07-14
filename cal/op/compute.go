package op

func Compute(source string)string{
	var op Operator
	var sta Stack
	for i:=0;i<len(source);i++{
		if source[i]>='0' && source[i]<='9'{
			sta.Push(string(source[i]))
		}else{
			opr:=sta.Pop()
			opl:=sta.Pop()
			r:=op.Compute(opl,opr,string(source[i]))
			sta.Push(r)
		}
	}
	return sta.Pop()
}
