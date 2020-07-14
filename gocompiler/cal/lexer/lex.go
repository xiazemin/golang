package lexer

import "github.com/xiazemin/gocompiler/cal/token"

type Lexer struct {
	input        string // 输入
	position     int    // 当前位置
	readPosition int    // 将要读取的位置
	ch           byte   //当前字符
}

func NewLex(input string) *Lexer {
	return &Lexer{input: input,ch:' '}
}

//词法分析器的核心函数是NextToken()用于获取下一个词元。
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok =token.NewToken(token.RPAREN, l.ch)
	case '+':
		tok =token.NewToken(token.PLUS, l.ch)
	case '-':
		tok =token.NewToken(token.MINUS, l.ch)
	case '/':
		tok =token.NewToken(token.SLASH, l.ch)
	case '*':
		tok =token.NewToken(token.ASTERISK, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok =token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok

}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
