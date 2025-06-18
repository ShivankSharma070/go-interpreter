package lexer

import "github.com/ShivankSharma070/go-interpreter/token"

type Lexer struct {
	input        string
	position     int  // Position of current char
	readPosition int  // Position of next char
	ch           byte // Current Character
}

func New(inp string) *Lexer {
	l := &Lexer{
		input: inp,
	}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Eat all the whitespace as it does not matter in the language we are creating
	l.eatWhitespaces()

	switch l.ch {
	case '=':
		tok = token.NewToken(token.EQUAL, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdenOrLiteral(isLetter)
			tok.Type = token.LookUpIden(tok.Literal)
			return tok // Important as positing is already incremented in readIden()
		} else if isDigit(l.ch) {
			tok.Literal = l.readIdenOrLiteral(isDigit)
			tok.Type = token.INT
			return tok // Important as positing is already incremented in readIden()
		} else {
			tok.Type = token.ELLEGAL
			tok.Literal = ""
		}
	}

	l.ReadChar()
	return tok
}

// This function read a identifier or a literal value, it accept a validate func() which return a boolean value
// It reads character continously, till it satisfy validate()
func (l *Lexer) readIdenOrLiteral(validate func(byte) bool) string {
	position := l.position
	for validate(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

// Eat up all the whitespaces, newline, tab characters
func (l *Lexer) eatWhitespaces() {
	for l.ch == '\n' || l.ch == ' ' || l.ch == '\r' || l.ch == '\t' {
		l.ReadChar()
	}
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char == '_')
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
