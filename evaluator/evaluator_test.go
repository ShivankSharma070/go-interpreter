package evaluator

import (
	"testing"

	"github.com/ShivankSharma070/go-interpreter/lexer"
	"github.com/ShivankSharma070/go-interpreter/object"
	"github.com/ShivankSharma070/go-interpreter/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}
	for _, tt := range tests {
		output := testEval(tt.input)
		testIntegerObject(t, output, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	integerObj, ok := obj.(*object.Integer)
	if !ok {
		t.Fatalf("Object not of type object.Integer, got %T", obj)
		return false
	}

	if integerObj.Value != expected {
		t.Fatalf("integerObj.value is not %d, got %d", expected, integerObj.Value)
		return false
	}

	return true
}
