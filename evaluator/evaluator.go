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
	case *ast.InfixExpression:
		right := Eval(node.Right)
		left := Eval(node.Left)
		return evalInfixExpression(node.Operator, right, left)
	}

	return nil
}

func evalPrefixExpression(operator string, value object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(value)
	case "-":
		return evalMinusPrefixOperatorExpression(value)
	default:
		return NULL
	}
}

// Evaluate expresions with minus as prefix operators
func evalMinusPrefixOperatorExpression(value object.Object) object.Object {
	if value.Type() != object.INTEGER_OBJ {
		return NULL
	}

	int_value := value.(*object.Integer).Value
	return &object.Integer{Value: -int_value}
}

// Evaluate bang prefix operations
func evalBangOperatorExpression(value object.Object) object.Object {
	switch value {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	// If its any another value than true, false or null, consider it true, so its bang is false
	default:
		return FALSE
	}
}

func evalInfixExpression(operator string, right, left object.Object) object.Object {
	switch {
	case right.Type() == object.INTEGER_OBJ && left.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, right, left)
	case right.Type() == object.BOOLEAN_OBJ && left.Type() == object.BOOLEAN_OBJ:
		return evalBoolInfixExpression(operator, right, left)
	default:
		return NULL
	}
}

func evalBoolInfixExpression(operator string, right, left object.Object) object.Object {
	switch operator {
	case "==":
		return nativeBoolToBooleanObject(right == left)
	case "!=":
		return nativeBoolToBooleanObject(right != left)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, right, left object.Object) object.Object {
	rightVal := right.(*object.Integer).Value
	leftVal := left.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
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
