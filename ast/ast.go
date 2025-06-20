package ast

import (
	"bytes"

	"github.com/ShivankSharma070/go-interpreter/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

// Token literal will return the literal value of token associated with a node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		// First node's token literal value
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var buf bytes.Buffer
	for _, stmt := range p.Statements {
		buf.WriteString(stmt.String())
	}

	return buf.String()
}

// Let Statements
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// To implement statement & node interface
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var buf bytes.Buffer

	buf.WriteString(ls.TokenLiteral() + " ")
	buf.WriteString(ls.Name.String())
	buf.WriteString(" = ")

	if ls.Value != nil {
		buf.WriteString(ls.Value.String())
	}

	buf.WriteString(";")
	return buf.String()
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

// To implement expression & node interface
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// Return Statments
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var buf bytes.Buffer
	buf.WriteString(rs.TokenLiteral() + " ")
	buf.WriteString(rs.ReturnValue.String())

	buf.WriteString(";")
	return buf.String()
}

// Expression statement
// 5+10, (5*10)+5, foo(a,b) etc
// We are treating expression as statements because, we want to allow one line containing only expression as a statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
