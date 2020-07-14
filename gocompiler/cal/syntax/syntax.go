package syntax

import (
	"bytes"
	"github.com/xiazemin/gocompiler/cal/token"

)

/*
计算表达式的语法
两种情况
1,++i。也就是某个操作符作为前缀与后面数字发生反应。同样还包括我们的-1
2,还有一种更加常见的情况1 + 2。操作符在中间
3,填写一个数字类似于12。这也是一个计算表达式。

 */
type Expression interface {
	String() string
}

type IntegerLiteralExpression struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteralExpression) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" ")
	out.WriteString(ie.Operator)
	out.WriteString(" ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func Eval(exp Expression) int64 {
	switch node := exp.(type) {
	case *IntegerLiteralExpression:
		return node.Value
	case *PrefixExpression:
		rightV := Eval(node.Right)
		return evalPrefixExpression(node.Operator, rightV)
	case *InfixExpression:
		leftV := Eval(node.Left)
		rightV := Eval(node.Right)
		return evalInfixExpression(leftV, node.Operator, rightV)
	}

	return 0
}

func evalPrefixExpression(operator string, right int64) int64{
	if operator != "-" {
		return 0
	}
	return -right
}


func evalInfixExpression(left int64, operator string, right int64) int64 {

	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right != 0 {
			return left / right
		} else {
			return 0
		}
	default:
		return 0
	}

}



