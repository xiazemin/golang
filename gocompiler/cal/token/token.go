package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"

	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LPAREN = "("
	RPAREN = ")"
)

type Token struct {
	Type    string //对应我们上面的词元类型
	Literal string // 实际的string字符
}

func NewToken(t string, ch byte) Token {
	return Token{Type: t, Literal:string([]byte{ch})}
}
