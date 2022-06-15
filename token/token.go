package token

// TokenType using a string might not lead to the same performance
// as using an int or a byte would, but for us a string is perfect.
type TokenType string

type Token struct {
	Type    TokenType
	Literal string // hold the literal value of the token
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, x, y, ...
	INT   = "INT"   // 123456

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