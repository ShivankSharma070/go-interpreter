package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(t TokenType, char byte) Token {
	return Token{t, string(char)}
}

const (
	EOF     = "EOF"
	ELLEGAL = "ELLEGAL"

	// Identifier and Literals
	IDEN = "IDEN" // Variable names
	INT  = "INT"

	// Operators
	EQUAL = "="
	PLUS  = "+"

	//Delimeters
	SEMICOLON = ";"
	COMMA     = ","

	// Braces
	LPAREN = "("
	LBRACE = "{"
	RPAREN = ")"
	RBRACE = "}"

	// Keywords
	FUNC = "FUNCTION"
	LET  = "LET"
)
