package lexer

import (
	"testing"

	"github.com/ShivankSharma070/go-interpreter/token"
)

func TestLexer(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let sum = fn (a, b) {
		a + b;
	};

	let result = sum(five,ten);
	`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDEN, "five"},
		{token.EQUAL, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDEN, "ten"},
		{token.EQUAL, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDEN, "sum"},
		{token.EQUAL, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDEN, "a"},
		{token.COMMA, ","},
		{token.IDEN, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDEN, "a"},
		{token.PLUS, "+"},
		{token.IDEN, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDEN, "result"},
		{token.EQUAL, "="},
		{token.IDEN, "sum"},
		{token.LPAREN, "("},
		{token.IDEN, "five"},
		{token.COMMA, ","},
		{token.IDEN, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range test {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("Test_%d: Type mismatch Expected:%q Got:%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test_%d: Literal mismatch Expected:%q Got:%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
