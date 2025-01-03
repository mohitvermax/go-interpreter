/*It will take source code as input and output the tokens that represent the source
code. It will go through its input and output the next token it recognizes. It doesn’t need to
buffer or save tokens, since there will only be one method called NextToken(), which will output
the next token.
*/

package lexer

import (
	"testing"

	"github.com/mohitvermax/go-interpreter/token"
)

func TestNextToken(t *testing.T){
	input := `=+(){},;`

	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral	string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	
	l:= New(input)

	for i, tt := range tests{
		tok:= l.NextToken()

		if tok.Type != tt.expectedType{
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
			i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			i, tt.expectedLiteral, tok.Literal)
			}
	}
}