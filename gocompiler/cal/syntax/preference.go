package syntax

import "github.com/xiazemin/gocompiler/cal/token"

//定义出以下几个级别的优先级，与各符号对应的优先级
const (
	_ int = iota
	LOWEST
	SUM     // +, -
	PRODUCT // *, /
	PREFIX  // -X
	CALL    // (X)
)

var precedences = map[string]int{
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}
