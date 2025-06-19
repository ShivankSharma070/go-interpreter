package parser

import (
	"fmt"

	"github.com/ShivankSharma070/go-interpreter/ast"
	"github.com/ShivankSharma070/go-interpreter/lexer"
	"github.com/ShivankSharma070/go-interpreter/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func (p *Parser) Errors() []string {
	return p.errors
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// Read Two tokens so that currentToken and peekToken are set
	p.nextToken()
	p.nextToken()
	return p
}

// Reading next token from our lexer
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// Parsing let Statements
func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDEN) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.isCurToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// expectPeek function reads next token only if the next token is what we expect
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.isPeekToken(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) isCurToken(t token.TokenType) bool {
	return p.currentToken.Type == t

}

func (p *Parser) isPeekToken(t token.TokenType) bool {
	return p.peekToken.Type == t

}

func (p *Parser) peekError(t token.TokenType) {
	err := fmt.Sprintf("Expected next token to be %s , got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, err)
}
