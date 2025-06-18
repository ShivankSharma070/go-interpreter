package token

type TokenType string

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
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keyword = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIden(iden string) TokenType {
	if tokType, ok := keyword[iden]; ok {
		return tokType
	}
	return IDEN
}

func NewToken(t TokenType, char byte) Token {
	return Token{t, string(char)}
}
