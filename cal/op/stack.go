package op

type Stack struct {
	data []string
}

func (s*Stack)Push(val string){
	s.data=append(s.data,val)
}

func (s*Stack)Pop()string{
	val:=s.data[len(s.data)-1]
	s.data=s.data[:len(s.data)-1]
	return val
}

func (s*Stack)Empty()bool{
	return len(s.data)<=0
}

func (s*Stack)Top() string {
	return s.data[len(s.data)-1]
}