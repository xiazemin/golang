package main

import (
	"github.com/xiazemin/gocompiler/cal/lexer"
	"github.com/xiazemin/gocompiler/cal/syntax"
	"fmt"
)

func Calc(input string) int64 {
	lexer := lexer.NewLex(input)
	fmt.Printf("%#v\n",lexer)
	parser := syntax.NewParser(lexer)
        fmt.Printf("%#v\n",parser)
	exp := parser.ParseExpression(syntax.LOWEST)
	fmt.Printf("%#v\n",exp)
	return syntax.Eval(exp)
}

func main () {
	r:=Calc("5 + 5/2*(3-7);")
	fmt.Println(r)
}