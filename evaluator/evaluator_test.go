package evaluator

import (
	"testing"

	"github.com/ShivankSharma070/go-interpreter/lexer"
	"github.com/ShivankSharma070/go-interpreter/object"
	"github.com/ShivankSharma070/go-interpreter/parser"
)

// ========= INTEGER ============
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

// ======== BOOLEAN ==========
func TestEvalBooleanExpresion(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	} {
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		output := testEval(tt.input)
		testBooleanObject(t,output, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Fatalf("obj is not of type object.Boolean, got %T", obj)
		return false
	}

	if result.Value != expected{
		t.Fatalf("result.Value is not %t, got %t",expected, result.Value)
		return false
	}
	return true
}

// ========== PREFIX ==========
func TestBangOperator (t *testing.T) {
	tests := []struct{
		input string
		expected bool
	} {
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		output := testEval(tt.input)
		testBooleanObject(t, output, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	return Eval(program)
}
