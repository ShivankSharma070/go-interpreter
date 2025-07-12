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
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
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
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
	}

	for _, tt := range tests {
		output := testEval(tt.input)
		testBooleanObject(t, output, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Fatalf("obj is not of type object.Boolean, got %T", obj)
		return false
	}

	if result.Value != expected {
		t.Fatalf("result.Value is not %t, got %t", expected, result.Value)
		return false
	}
	return true
}

// ========== PREFIX ==========
func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
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

// ================ Conditional Expressions ===========
func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Fatalf("object is not object.NULL, got %T (%+v)", obj, obj)
		return false
	}
	return true
}

// ================== RETURN ================
func TestReturnExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
			if (10 > 1) {
				if (10 > 1) {
					return 10
				}
				return 1
			}
			`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

// =============== Errors ====================
func TestErrors(t *testing.T) {
	tests := []struct {
		input            string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
if (10 > 1) {
if (10 > 1) {
return true + false;
}
return 1;
}
`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errorObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("No errror object retruned, got %T (%+v)", evaluated, evaluated)
			continue
		}

		if errorObj.Message != tt.expectedMessage{
			t.Errorf("ErrorObj.Message is not %s, got %s", tt.expectedMessage,errorObj.Message)
		}
	}

}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	return Eval(program)
}
