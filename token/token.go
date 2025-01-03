package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

// These are the TokenTypes that we will be using in our lexer
// ILLEGAL signifies a token/character we don’t know about and EOF stands for “end of file”

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" 
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

