package parser

import (
	"fmt"
	"strconv"

	"github.com/ShivankSharma070/go-interpreter/ast"
	"github.com/ShivankSharma070/go-interpreter/lexer"
	"github.com/ShivankSharma070/go-interpreter/token"
)

// Constants for deciding precedence of operators with parsing them as expression
const (
	_ = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myfunction(x)
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string

	// Maps to associate a token with a parser function
	prefixParserMap map[token.TokenType]prefixParserFunc
	infixParserMap  map[token.TokenType]infixParserFunc
}

type (
	prefixParserFunc func() ast.Expression               // For token found in prefix position
	infixParserFunc  func(ast.Expression) ast.Expression // For token found in infix position
)

func (p *Parser) Errors() []string {
	return p.errors
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixParserMap = map[token.TokenType]prefixParserFunc{}
	p.registerPrefix(token.IDEN, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

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
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// Parsing return statements
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}
	p.nextToken()

	for !p.isCurToken(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
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
func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.isPeekToken(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParserMap[p.currentToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(lit.TokenLiteral(), 0, 64)
	if err != nil {
		err := fmt.Sprintf("Could not parse %q as integer", lit.TokenLiteral())
		p.errors = append(p.errors, err)
	}

	lit.Value = value
	return lit
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

// Helper function to asscoiate parser fucntion for a token in prefix position
func (p *Parser) registerPrefix(tokentype token.TokenType, fn prefixParserFunc) {
	p.prefixParserMap[tokentype] = fn
}

// Helper function to asscoiate parser fucntion for a token in prefix position
func (p *Parser) registerInfix(tokentype token.TokenType, fn prefixParserFunc) {
	p.prefixParserMap[tokentype] = fn
}
