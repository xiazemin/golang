package op

func Translate(source string) string {
	if len(source) <= 0 {
		return source
	}
	var ops Stack
	var lastOrder string
	var o Operator
	for i := 0; i < len(source); i++ {
		if source[i] >= '0' && source[i] <= '9' {
			lastOrder += string(source[i])
		} else if o.IsOperator(source[i]) {
			if o.Priority(source[i]) == 2 {
				for !ops.Empty() && (o.Priority([]byte(ops.Top())[0]) == 2){
					lastOrder += ops.Pop()
				}
				ops.Push(string(source[i]))
			} else if o.Priority(source[i]) == 1 {
				for !ops.Empty() && (o.Priority([]byte(ops.Top())[0]) == 2 || o.Priority([]byte(ops.Top())[0]) == 1) {
					lastOrder += ops.Pop()
				}
				ops.Push(string(source[i]))
			} else if o.Priority(source[i]) == 3 {
				if source[i] == '(' {
					ops.Push(string(source[i]))
				} else if source[i] == ')' {
					for !ops.Empty() && ops.Top() != "(" {
						lastOrder += string(ops.Pop())
					}
					ops.Pop()
				}
			}
		}
	}
	for !ops.Empty() {
		lastOrder += string(ops.Pop())
	}
	return lastOrder
}
