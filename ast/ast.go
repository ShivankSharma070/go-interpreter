package ast

import (
	"github.com/ShivankSharma070/go-interpreter/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// For a variable binding -> let statements
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// To implement expresssion & node interface
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Here identifier is implementing the expression interface. Why ??
// Although expression are those which produces some values,
// and identifier in let x = 5 does not produces a vlaue
// It is to keep the number of different nodes type minimum
// And a identifier will produce value in statement like let a = b; (b produces value)
type Identifier struct {
	Token token.Token
	Value string
}

// To implement expresssion & node interface
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Token literal will return the literal value of token associated with a node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		// First node's token literal value
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Return Statments
type ReturnStatement struct {
	Token       token.Token
	ReturnValue *Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
