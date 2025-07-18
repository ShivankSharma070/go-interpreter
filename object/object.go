package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ShivankSharma070/go-interpreter/ast"
)

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	NULL_OBJ         = "NULL"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ = "STRING"
	BUILTIN_OBJ = "BUILTIN"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Null struct{}

func (n *Null) Inspect() string  { return "NULL" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

type String struct {
	Value string
}

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType {return STRING_OBJ}

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Inspect() string { return r.Value.Inspect() }
func (r *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

type Error struct {
	Message string
}

func (e *Error) Inspect() string  { return "Error: " + e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ }

type FunctionLiteral struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (fe *FunctionLiteral) Inspect() string {
	var result bytes.Buffer
	params := []string{}
	for _, f := range fe.Parameters {
		params = append(params, f.String())
	}

	result.WriteString("fn (")
	result.WriteString(strings.Join(params, ", "))
	result.WriteString(") {\n")
	result.WriteString(fe.Body.String())
	result.WriteString("\n}")
	return result.String()
}

func (fe *FunctionLiteral) Type() ObjectType {
	return FUNCTION_OBJ
}

// ============ ENVIRONMENT ==============
type Environment struct {
	Store map[string]Object
	Outer *Environment
}

func NewEnclosingEnvironment(enclosingEnv *Environment) *Environment{
	env := NewEnvironment()
	env.Outer = enclosingEnv
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{Store: s, Outer: nil }
}

func (e *Environment) Get(name string) (Object, bool) {
	value, ok := e.Store[name]
	if !ok && e.Outer != nil {
		value, ok = e.Outer.Get(name)
	}
	return value, ok
}

func (e *Environment) Set(name string, value Object) Object {
	e.Store[name] = value
	return value
}

// ================== BUILT-IN FUNCTION ===================

type BuiltInFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltInFunction
}

func (bu *Builtin) Inspect() string{return "builtin function"}
func (bu *Builtin) Type() ObjectType { return BUILTIN_OBJ}
