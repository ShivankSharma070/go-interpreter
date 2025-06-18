package lexer

import (
	"testing"

	"github.com/ShivankSharma070/go-interpreter/token"
)

func TestLexer(t *testing.T) {
	input := "=+(){},;"

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.EQUAL, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range test {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("Test_%d: Type mismatch Expected:%s Got:%s", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test_%d: Literal mismatch Expected:%s Got:%s", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
