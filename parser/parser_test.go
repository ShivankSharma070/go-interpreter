package parser

import (
	"testing"

	"github.com/ShivankSharma070/go-interpreter/ast"
	"github.com/ShivankSharma070/go-interpreter/lexer"
)

func TestParser(t *testing.T) {
	input := `
	let x 5;
	let  = 10;
	let  100;
	`

	lex := lexer.New(input)
	p := New(lex)

	prog := p.ParseProgram()
	checkForParserErrors(t, p)
	if prog == nil {
		t.Fatal("ParseProgram() return a nil value")
	} else if len(prog.Statements) != 3 {
		t.Fatalf("Returned Program does not container 3 statements got %d", len(prog.Statements))
	}

	test := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range test {
		stmt := prog.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	// Check if it is let statement
	if stmt.TokenLiteral() != "let" {
		t.Errorf("Tokenliteral is not let, got %s\n", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("Statement is not ast.LetStatement, got %T\n", stmt)
		return false
	}

	// Check if the identifier is what we expect
	if letStmt.Name.Value != name {
		t.Errorf("letstmt.name.value is not same as name, got %s", letStmt.Name.Value)
		return false
	}

	return true
}

func checkForParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	for _, err := range errors {
		t.Errorf("Parser error %s", err)

	}
	t.FailNow()
}
