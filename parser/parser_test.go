package parser

import (
	"fmt"
	"testing"

	"github.com/ShivankSharma070/go-interpreter/ast"
	"github.com/ShivankSharma070/go-interpreter/lexer"
)

func TestReturnParser(t *testing.T) {
	input := `
	return 1234;
	return myvar;
	return 5;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Error("ParseProgram returned nil")
		return
	}

	if len(program.Statements) != 3 {
		t.Errorf("Program does not container 3 statements. It contains %d", len(program.Statements))
		return
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt is not *ast.ReturnStatement, got %T", stmt)
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("TokenLiteral() is not 'return', got %s", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	checkForParserErrors(t, p)

	if len(prog.Statements) != 1 {
		t.Fatalf("Now enough statement in program, got %d", len(prog.Statements))
	}

	for _, stmt := range prog.Statements {

		expStmt, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Statement is not of *ast.ExpressionStatement, got %T", stmt)
		}

		idenStmt, ok := expStmt.Expression.(*ast.Identifier)
		if !ok {
			t.Fatalf("expStmt.Expression is not of *ast.Identifier, got %T", stmt)
		}

		if idenStmt.Value != "foobar" {
			t.Fatalf("idenStmt.value is not equal to %s, got %s", "foobar", idenStmt.Value)
		}

		if idenStmt.TokenLiteral() != "foobar" {
			t.Fatalf("idenStmt.tokenLiteral() is not equal to %s, got %s", "foobar", idenStmt.TokenLiteral())
		}
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := `5;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkForParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not have enough statements, got %d", len(program.Statements))
	}

	expStmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not ast.ExpressionStatement, got %T", program.Statements[0])
	}

	if !testIntegerLiteral(t, expStmt.Expression, 5) {
		return
	}

	// literal, ok := expStmt.Expression.(*ast.IntegerLiteral)
	// if !ok {
	// 	t.Fatalf("expStmt.Expression is not ast.IntegerLiteral, got %T", expStmt.Expression)
	// }
	//
	// if literal.TokenLiteral() != "5" {
	// 	t.Fatalf("literal.Tokenliteral is not equal to %s, got %s", "5", literal.TokenLiteral())
	// }
	//
	// if literal.Value != 5 {
	// 	t.Fatalf("literal.value is not equal to %d, got %d", 5, literal.Value)
	// }
}
func TestPrefixExpressionParsing(t *testing.T) {
	tests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		parser := New(l)

		prog := parser.ParseProgram()

		checkForParserErrors(t, parser)
		if len(prog.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. Got %d", 1, len(prog.Statements))
		}

		expStmt, ok := prog.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("prog.Statement is not of type *ast.ExpressionStatement, got %T", prog.Statements[0])
		}

		prefixStmt, ok := expStmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expStmt.Expression is not of type *ast.PrefixExpression, got %T", expStmt.Expression)
		}

		if prefixStmt.Operator != tt.operator {
			t.Fatalf("prefixStmt.Operator is not %s, got %s", tt.operator, prefixStmt.Operator)
		}

		if !testIntegerLiteral(t, prefixStmt.Right, tt.integerValue) {
			return
		}
	}

}

func TestLetParser(t *testing.T) {
	input := `
	let x = 5;
	let y  = 10;
	let foo =  100;
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

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
	intStmt, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Exp is not of type ast.IntegerLiteral, got %T", exp)
		return false
	}

	if intStmt.Value != value {
		t.Fatalf("intStmt.Value is not %d, got %d", value, intStmt.Value)
		return false
	}

	if intStmt.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Fatalf("intStmt.TokenLiteral() is not equal to %s, got %s", fmt.Sprintf("%d", value), intStmt.TokenLiteral())
		return false
	}

	return true
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
