package lexer
import (
	"testing"
	"github.com/xiazemin/gocompiler/cal/token"
	"github.com/golang/go/src/fmt"
)

func TestTokenizer(t *testing.T) {
	input := `(5 + -10 * 2 + 15 / 3) * 2`
	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.INT, "10"},
		{token.ASTERISK, "*"},
		{token.INT, "2"},
		{token.PLUS, "+"},
		{token.INT, "15"},
		{token.SLASH, "/"},
		{token.INT, "3"},
		{token.RPAREN, ")"},
		{token.ASTERISK, "*"},
		{token.INT, "2"},
	}

	l := NewLex(input)
fmt.Println(l,l.ch==' ')
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(l,tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
