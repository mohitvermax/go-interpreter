package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// These are the TokenTypes that we will be using in our lexer
// ILLEGAL signifies a token/character we don’t know about and EOF stands for “end of file”

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

/*
LookupIdent checks the keywords table to see whether the given identifier is in fact a keyword.
If it is, it returns the keyword’s TokenType constant. If it isn’t, we just get back token.IDENT,
which is the TokenType for all user-defined identifiers.
*/
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
