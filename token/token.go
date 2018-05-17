package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOE     = "EOF"

	// Identifiers & Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN  = "("
	RPARENT = ")"
	LBRACE  = "{"
	RBRACE  = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
