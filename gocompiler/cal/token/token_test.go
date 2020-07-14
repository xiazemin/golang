package token

import (
	"testing"
	"fmt"
)

func TestNewToken(t *testing.T){
	fmt.Println(NewToken(LPAREN, '('))
}
