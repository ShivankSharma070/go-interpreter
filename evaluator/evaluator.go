package evaluator

import (
	"github.com/ShivankSharma070/go-interpreter/ast"
	"github.com/ShivankSharma070/go-interpreter/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BoolExpression:
		// Creating a new object for every true and fasle is pointless, as there is no difference between two true or false.
		// return &object.Boolean{Value: node.Value}
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}

	return nil
}

func evalPrefixExpression(operator string, value object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(value)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(value object.Object) object.Object {
	switch value {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}
	return result
}

func nativeBoolToBooleanObject(value bool) *object.Boolean {
	if value {
		return TRUE
	}
	return FALSE
}
