package evaluator

import (
	"fmt"
	"github.com/ShivankSharma070/go-interpreter/ast"
	"github.com/ShivankSharma070/go-interpreter/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BoolExpression:
		// Creating a new object for every true and fasle is pointless, as there is no difference between two true or false.
		// return &object.Boolean{Value: node.Value}
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		return evalInfixExpression(node.Operator, right, left)
	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements, env)
	case *ast.IfElseExpression:
		return evalIfExpression(node, env)
	case *ast.ReturnStatement:
		value := Eval(node.ReturnValue, env)
		if isError(value) {
			return value
		}
		return &object.ReturnValue{Value: value}
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.Identifier:
		return evalIdentifier(node, env)

	}
	return nil
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
	return	newError("identifier not found: %s" , node.Value)
	}

	return val
}

func evalIfExpression(node *ast.IfElseExpression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(node.Consequence, env)
	} else if node.Alternative != nil {
		return Eval(node.Alternative, env)
	} else {
		return NULL
	}
}

func evalPrefixExpression(operator string, value object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(value)
	case "-":
		return evalMinusPrefixOperatorExpression(value)
	default:
		return newError("unknown operator: %s%s", operator, value.Type())
	}
}

// Evaluate expresions with minus as prefix operators
func evalMinusPrefixOperatorExpression(value object.Object) object.Object {
	if value.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", value.Type())
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
	case right.Type() != left.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unkown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalBoolInfixExpression(operator string, right, left object.Object) object.Object {
	switch operator {
	case "==":
		return nativeBoolToBooleanObject(right == left)
	case "!=":
		return nativeBoolToBooleanObject(right != left)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
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
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}

}

func evalBlockStatement(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt, env)
		if result != nil {
			if result.Type() == object.RETURN_VALUE_OBJ || result.Type() == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func evalProgram(node *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, stmt := range node.Statements {
		result = Eval(stmt, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func nativeBoolToBooleanObject(value bool) *object.Boolean {
	if value {
		return TRUE
	}
	return FALSE
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case TRUE:
		return true
	case FALSE:
		return false
	case NULL:
		return false
	default:
		return true
	}
}

func newError(format string, a ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
