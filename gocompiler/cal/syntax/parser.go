package syntax

import (
	"fmt"
	"github.com/xiazemin/gocompiler/cal/lexer"
	"github.com/xiazemin/gocompiler/cal/token"
	"strconv"
)

// parser
type (
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
)


//用一个结构parser来把我们的字符串变成expression。parser里面包含我们上一步的lexer。以及存储error的数组。当前的词元和下一个词元。另外针对于上面提到的两种不同的expression。利用不同的处理方法。
type Parser struct {
	l              *lexer.Lexer
	errors         []string
	curToken       token.Token
	peekToken      token.Token
	prefixParseFns map[string]prefixParseFn
	infixParseFns  map[string]infixParseFn
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[string]prefixParseFn)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)

	p.infixParseFns = make(map[string]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()
	return p
}


//我们的核心函数是将lexer要变成ast，这个核心函数是ParseExpression
//用到的算法叫做pratt parser
func (p *Parser) ParseExpression(precedence int) Expression {
	/*
	递归主函数ParseExpression
我们通过当前优先级和下一个token的优先级进行对比，如果这个优先级比下一个优先级别低，那就变成infix。用parseInfixExpression处理。如果这个优先级等于或者比下一个优先级高，那就变成了prefix。用parsePrefixExpression处理
	 */
	prefix := p.prefixParseFns[p.curToken.Type] //registerPrefix
	if prefix == nil {
		fmt.Println(p)
	}
	returnExp := prefix()

	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return returnExp
		}

		p.nextToken()
		returnExp = infix(returnExp)
	}

	return returnExp
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t string) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t string) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instend",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
func (p *Parser) expectPeek(t string) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// 往结构体里面筛处理方法
func (p *Parser) registerPrefix(tokenType string, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType string, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseIntegerLiteral() Expression {

	lit := &IntegerLiteralExpression{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parsePrefixExpression() Expression {

	expression := &PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.ParseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left Expression) Expression {

	expression := &InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()

	expression.Right = p.ParseExpression(precedence)

	return expression
}

func (p *Parser) parseGroupedExpression() Expression {
	p.nextToken()
	exp := p.ParseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}

