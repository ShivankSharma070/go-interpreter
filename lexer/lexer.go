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
	}

	l.ReadChar()
	return tok
}

