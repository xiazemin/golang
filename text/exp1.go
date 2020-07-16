package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multi("123456", "87891"), 123456*87891)
}

func multi(num1, num2 string) string {
	n1, e1 := str2int(num1)
	if e1 != nil {
		return ""
	}

	n2, e2 := str2int(num2)
	if e2 != nil {
		return ""
	}
	fmt.Println(n1, n2)
	r := make([]int64, len(n1)+len(n2))
	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			r[i+j+1] += (r[i+j] + n1[i]*n2[j]) / 10
			r[i+j] = (r[i+j] + n1[i]*n2[j]) % 10
		}
	}
	j := len(r) - 1

	for j > 0 && r[j] == 0 {
		j--
	}

	val := ""
	for i := 0; i <= j; i++ {
		val = fmt.Sprintf("%d", r[i]) + val
	}

	return val
}

type Num2IntError struct {
}

func (this Num2IntError) Error() string {
	return "num is empty"
}
func str2int(num string) ([]int64, error) {
	if len(num) == 0 {
		var e Num2IntError
		return nil, e
	}
	var r []int64
	for i := 0; i < len(num); i++ {
		v, err := strconv.ParseInt(string([]byte{num[i]}), 10, 10)
		if err != nil {
			return nil, err
		}
		r = append([]int64{v}, r...)
	}
	return r, nil
}
