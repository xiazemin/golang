package syntax

import (
	"testing"
	"github.com/xiazemin/gocompiler/cal/lexer"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := "250"
	var expectValue int64 = 250

	l := lexer.NewLex(input)
	p := NewParser(l)


	//checkParseErrors(t, p)
	expression := p.ParseExpression(LOWEST)
	testInterLiteral(t, expression, expectValue)
}


func testInterLiteral(t *testing.T, il Expression, value int64) bool {
	integ, ok := il.(*IntegerLiteralExpression)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	return true
}

func TestParsingPrefixExpression(t *testing.T) {
	input := "-15"
	expectedOp := "-"
	var expectedValue int64 =  15


	l := lexer.NewLex(input)
	p := NewParser(l)
	//checkParseErrors(t, p)

	expression := p.ParseExpression(LOWEST)
	exp, ok := expression.(*PrefixExpression)

	if !ok {
		t.Fatalf("stmt is not PrefixExpression, got=%T", exp)
	}

	if exp.Operator != expectedOp {
		t.Fatalf("exp.Operator is not %s, go=%s", expectedOp, exp.Operator)
	}

	testInterLiteral(t, exp.Right, expectedValue)
}

func TestParsingInfixExpression(t *testing.T) {
	infixTests := []struct{
		input string
		leftValue int64
		operator string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
	}

	for _, tt := range infixTests {
		l := lexer.NewLex(tt.input)
		p := NewParser(l)
		//checkParseErrors(t, p)

		expression := p.ParseExpression(LOWEST)
		exp, ok := expression.(*InfixExpression)

		if !ok {
			t.Fatalf("exp is not InfixExpression, got=%T", exp)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s, go=%s", tt.operator, exp.Operator)
		}

		testInterLiteral(t, exp.Left, tt.leftValue)
		testInterLiteral(t, exp.Right, tt.rightValue)
	}
}


