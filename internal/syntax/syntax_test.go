package syntax

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

// TestSyntaxChecker_NoError tests that a valid program parses without errors.
func TestSyntaxChecker_NoError(t *testing.T) {
	checker := NewSyntaxChecker()
	input := antlr.NewInputStream("int a = 1 + 2; return a;")

	tree, err := checker.Check(input)

	assert.NoError(t, err)
	assert.NotNil(t, tree)
}

// TestSyntaxChecker_WithError tests that a program with a syntax error is correctly identified.
func TestSyntaxChecker_WithError(t *testing.T) {
	checker := NewSyntaxChecker()
	// An invalid expression (missing operand)
	input := antlr.NewInputStream("int a = 1 + ; return a")

	tree, err := checker.Check(input)

	assert.Error(t, err)
	assert.Nil(t, tree)
	assert.Contains(t, err.Error(), "syntax error at line 1")
}

// TestSyntaxChecker_MultipleErrors tests that the checker can capture multiple syntax errors.
func TestSyntaxChecker_MultipleErrors(t *testing.T) {
	checker := NewSyntaxChecker()
	// Invalid expression with multiple issues (missing operands)
	input := antlr.NewInputStream("int a = 1 + ; int b = 2 +")

	_, errors := checker.CheckWithAllErrors(input)

	// Tree might not be nil even with errors in ANTLR, but there should be errors captured
	assert.Greater(t, len(errors), 0, "Should capture syntax errors")

	// Verify the checker also detects errors via Check method
	_, err := checker.Check(antlr.NewInputStream("int a = 1 + ; int b = 2 +"))
	assert.Error(t, err, "Check method should return error for invalid syntax")
}

// TestSyntaxErrorListener_CollectsMultipleErrors tests that the error listener correctly collects errors.
func TestSyntaxErrorListener_CollectsMultipleErrors(t *testing.T) {
	listener := &SyntaxErrorListener{}

	// Simulate multiple syntax errors
	listener.SyntaxError(nil, nil, 1, 10, "error 1", nil)
	listener.SyntaxError(nil, nil, 2, 15, "error 2", nil)

	assert.True(t, listener.HasErrors())
	assert.Equal(t, 2, len(listener.GetAllErrors()))

	firstError := listener.GetFirstError()
	assert.NotNil(t, firstError)
	assert.Contains(t, firstError.Error(), "error 1")
}

// TestSyntaxChecker_IsValid tests the IsValid helper method.
func TestSyntaxChecker_IsValid(t *testing.T) {
	checker := NewSyntaxChecker()

	// Valid input
	validInput := antlr.NewInputStream("int a = 1 + 2; return a;")
	assert.True(t, checker.IsValid(validInput))

	// Invalid input
	invalidInput := antlr.NewInputStream("int a = 1 + ;")
	assert.False(t, checker.IsValid(invalidInput))
}

// TestSyntaxChecker_Reset tests the Reset method clears previous errors.
func TestSyntaxChecker_Reset(t *testing.T) {
	checker := NewSyntaxChecker()

	// First check with invalid input
	invalidInput := antlr.NewInputStream("int a = 1 + ;")
	_, err := checker.Check(invalidInput)
	assert.Error(t, err)
	assert.Greater(t, checker.GetErrorCount(), 0)

	// Reset and verify
	checker.Reset()
	assert.Equal(t, 0, checker.GetErrorCount())
}

// TestSyntaxChecker_EmptyInput tests that an empty input results in an error.
func TestSyntaxChecker_EmptyInput(t *testing.T) {
	checker := NewSyntaxChecker()
	emptyInput := antlr.NewInputStream("")

	_, err := checker.Check(emptyInput)
	assert.Error(t, err, "Empty input should result in an error")
}

// Test various literal types
func TestSyntaxChecker_Literals(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Integer literals
		{"Decimal Integer", "int a = 42;"},
		{"Hexadecimal Integer", "int b = 0xFF;"},
		{"Octal Integer", "int c = 0755;"},
		{"Binary Integer", "int d = 0b1010;"},
		{"Long Integer", "long e = 123L;"},

		// Floating point literals
		{"Double Literal", "double f = 3.14;"},
		{"Float Literal", "float g = 2.5f;"},
		{"Scientific Notation Double", "double h = 1.23e-4;"},

		// Boolean literals
		{"Boolean True", "boolean i = true;"},
		{"Boolean False", "boolean j = false;"},

		// Character literals
		{"Character Literal", "char k = 'A';"},
		{"Character Escape Sequence", "char l = '\\n';"},

		// String literals
		{"String Literal", "String m = \"Hello\";"},
		{"String with Escaped Quotes", "String n = \"She said \\\"Hello\\\"\";"},

		// Null literal
		{"Null Literal", "String o = null;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test variable declarations of different types
func TestSyntaxChecker_VariableDeclarations(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Static type declarations
		{"Static Type Single", "int a = 10;"},
		{"Static Type Multiple", "int b, c = 20;"},
		{"Static Type String", "string name = \"Alice\";"},

		// Var declarations (type inferred)
		{"Var Declaration Int", "var x = 42;"},
		{"Var Declaration Double", "var y = 3.14;"},
		{"Var Declaration Boolean", "var z = true;"},
		{"Var Declaration Multiple", "var a = 1, b = 2;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test array declarations and initializations
func TestSyntaxChecker_Arrays(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Array declarations
		{"Single Dimension Array Declaration", "int[] numbers;"},
		{"String Array Declaration", "String[] names;"},
		{"Two Dimensional Array Declaration", "int[][] matrix;"},

		// Array initializations
		{"Single Dimension Array Initialization", "int[] arr = {1, 2, 3};"},
		{"String Array Initialization", "String[] strs = {\"a\", \"b\"};"},
		{"Two Dimensional Array Initialization", "int[][] grid = {{1, 2}, {3, 4}};"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test control structures
func TestSyntaxChecker_ControlStructures(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// If statements
		{"If Statement Block", "if (true) { int a = 1; }"},
		{"If-Else Statement Block", "if (true) { int a = 1; } else { int b = 2; }"},
		{"If-Else Statement Single", "if (x > 0) return x; else return -x;"},

		// For loops
		{"Basic For Loop", "for (int i = 0; i < 10; i++) { sum += i; }"},
		{"Enhanced For Loop", "for (String s : strings) { process(s); }"},

		// Switch statements
		{"Switch Statement", "switch (value) { case 1: return 1; default: return 0; }"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test various expression types
func TestSyntaxChecker_Expressions(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Arithmetic expressions
		{"Arithmetic Operator Precedence", "int a = 1 + 2 * 3;"},
		{"Arithmetic Parentheses", "int b = (1 + 2) * 3;"},
		{"Unary Minus", "int c = -5;"},
		{"Pre Increment", "int d = ++i;"},

		// Logical expressions
		{"Logical AND", "boolean a = true && false;"},
		{"Logical NOT", "boolean b = !condition;"},
		{"Logical OR", "boolean c = x > 0 || y < 0;"},

		// Comparison expressions
		{"Equality Comparison", "boolean a = x == y;"},
		{"Inequality Comparison", "boolean b = x != y;"},
		{"Greater Than Or Equal", "boolean c = x >= y;"},

		// Ternary expressions
		{"Ternary Expression", "int a = condition ? 1 : 0;"},
		{"Ternary Expression String", "String b = score >= 60 ? \"Pass\" : \"Fail\";"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test function calls and method invocations
func TestSyntaxChecker_FunctionCalls(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Simple function calls
		{"Function Call With Args", "int result = calculate(1, 2);"},
		{"Function Call No Args", "processData();"},

		// Method invocations
		{"Method Invocation", "String upper = text.toUpperCase();"},
		{"Field Access", "int len = array.length;"},

		// Chained method calls
		{"Chained Method Calls", "String result = text.trim().toLowerCase();"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker.Reset()
			input := antlr.NewInputStream(tt.input)
			assert.True(t, checker.IsValid(input), "Input: %s", tt.input)
		})
	}
}

// Test lambda expressions
func TestSyntaxChecker_Lambdas(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Simple lambda (expression form)
		{"Lambda Simple", "var square = (x) -> x * x;"},

		// Lambda with explicit types (expression form)
		{"Lambda With Types", "var add = (int a, int b) -> a + b;"},

		// Lambda with block body (unified syntax)
		{"Lambda Block Body", "var process = (int x) -> int { return x * 2; };"},

		// Lambda with no parameters
		{"Lambda No Parameters", "var supplier = () -> 42;"},

		// Lambda with type inference and block
		{"Lambda Type Inference Block", "var calculator = (x, y) -> x + y;"},

		// Expression block (no explicit return needed)
		{"Expression Block Simple", "var result = { 1 + 2 };"},

		// Expression block with statements (requires explicit expression at end)
		{"Expression Block With Statements", "var compute = { var x = 10; var y = 20; x + y };"},

		// Expression block in lambda
		{"Lambda Expression Block", "var process = (int x) -> int { var doubled = x * 2; doubled + 1 };"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test struct declarations and usage
func TestSyntaxChecker_Structs(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Struct declaration with fields only
		{"Struct Declaration", "struct Point { double x, double y }"},

		// Struct with various field types
		{"Struct With Mixed Types", "struct Person { string name, int age, boolean active }"},

		// Empty struct
		{"Empty Struct", "struct Empty { }"},

		// Struct instantiation
		{"Struct Instantiation", "Point p = Point{x: 1.0, y: 2.0};"},

		// Struct with methods
		{"Struct With Method", "struct Calculator { int value, add(int x) -> int { return this.value + x; } }"},

		// Struct with multiple methods
		{"Struct With Multiple Methods", `struct Rectangle { 
			double width, 
			double height,
			area() -> double { return this.width * this.height; },
			perimeter() -> double { return 2 * (this.width + this.height); }
		}`},

		// Struct with method without return type
		{"Struct Method No Return", "struct Logger { string prefix, log(string msg) { print(this.prefix + msg); } }"},

		// Struct with method taking multiple parameters
		{"Struct Method Multiple Params", "struct Math { multiply(int a, int b) -> int { return a * b; } }"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test container literals (List, Map, Set)
func TestSyntaxChecker_ContainerLiterals(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// List literals
		{"List Literal Typed", "List<int> numbers = List{1, 2, 3};"},
		{"List Literal Inferred", "var strings = [\"a\", \"b\", \"c\"];"},

		// Map literals
		{"Map Literal Typed", "Map<string, int> scores = Map{\"Alice\": 95, \"Bob\": 87};"},

		// Set literals
		{"Set Literal", "Set<int> primes = Set{2, 3, 5, 7};"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test assignment expressions
func TestSyntaxChecker_Assignments(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Simple assignment
		{"Simple Assignment", "x = 10;"},

		// Compound assignments
		{"Add Assignment", "x += 5;"},
		{"Multiply Assignment", "y *= 2;"},
		{"Subtract Assignment", "z -= 3;"},
		{"Divide Assignment", "count /= 4;"},
		{"Modulo Assignment", "value %= 7;"},
		{"Bitwise OR Assignment", "flags |= mask;"},
		{"Left Shift Assignment", "data <<= 2;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test return statements
func TestSyntaxChecker_ReturnStatements(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Return without value
		{"Return No Value", "return;"},

		// Return with value
		{"Return Integer", "return 42;"},
		{"Return String", "return \"Hello\";"},
		{"Return Expression", "return x > 0;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test break and continue statements
func TestSyntaxChecker_BreakContinue(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Break statement
		{"Break Statement", "break;"},

		// Continue statement
		{"Continue Statement", "continue;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test placeholder variables
func TestSyntaxChecker_PlaceholderVariables(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Placeholder variables
		{"Placeholder Integer", "int score = #user_score#;"},
		{"Placeholder Boolean", "boolean active = #is_active#;"},
		{"Placeholder String", "string name = #user_name#;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test complex, integrated programs
func TestSyntaxChecker_ComplexPrograms(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		{
			"Complex Struct Program",
			`
		struct Point { double x, double y }
		struct Rectangle { Point topLeft, Point bottomRight }
		
		Rectangle rect = Rectangle{
			topLeft: Point{x: 0.0, y: 0.0},
			bottomRight: Point{x: 10.0, y: 5.0}
		};
		
		double area = (rect.bottomRight.x - rect.topLeft.x) * (rect.bottomRight.y - rect.topLeft.y);
		return area;
	`,
		},
		{
			"Program With Control Flow",
			`
		int sum = 0;
		for (int i = 1; i <= 10; i++) {
			if (i % 2 == 0) {
				sum += i;
			}
		}
		return sum;
	`,
		},
		{
			"Program With Functions and Lambdas",
			`
		var square = (x) -> x * x;
		var numbers = [1, 2, 3, 4, 5];
		var squares = [];
		
		for (int n : numbers) {
			squares.add(square(n));
		}
		
		return squares;
	`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Using raw string literal backticks to avoid escaping newlines/quotes
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Failed input:\n%s", tt.input)
		})
	}
}

// Test struct methods specifically
func TestSyntaxChecker_StructMethods(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Simple method with return type
		{"Method With Return Type", "struct Calculator { getValue() -> int { return 42; } }"},

		// Method without return type (void method)
		{"Method Without Return Type", "struct Logger { log(string msg) { print(msg); } }"},

		// Method with parameters
		{"Method With Parameters", "struct Math { add(int a, int b) -> int { return a + b; } }"},

		// Method using 'this' keyword
		{"Method Using This", "struct Counter { int count, increment() { this.count++; } }"},

		// Method with complex logic
		{"Method With Complex Logic", `struct Validator { 
			isValid(string input) -> boolean { 
				if (input == null) { return false; }
				return input.length() > 0;
			}
		}`},

		// Struct with mixed fields and methods
		{"Mixed Fields And Methods", `struct Person {
			string name,
			int age,
			getName() -> string { return this.name; },
			setAge(int newAge) { this.age = newAge; },
			isAdult() -> boolean { return this.age >= 18; }
		}`},

		// Method calling other methods
		{"Method Calling Other Methods", `struct Calculator {
			int value,
			setValue(int v) { this.value = v; },
			getValue() -> int { return this.value; },
			doubleValue() -> int { return this.getValue() * 2; }
		}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}

// Test 'this' keyword usage
func TestSyntaxChecker_ThisKeyword(t *testing.T) {
	checker := NewSyntaxChecker()

	tests := []struct {
		name  string
		input string
	}{
		// Using 'this' to access fields
		{"This Field Access", "struct Test { int value, getValue() -> int { return this.value; } }"},

		// Using 'this' to call methods
		{"This Method Call", "struct Test { getValue() -> int { return 42; }, getDouble() -> int { return this.getValue() * 2; } }"},

		// Using 'this' in assignments
		{"This Assignment", "struct Test { int value, setValue(int v) { this.value = v; } }"},

		// Using 'this' in complex expressions
		{"This Complex Expression", "struct Test { int a, int b, sum() -> int { return this.a + this.b; } }"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, checker.IsValid(antlr.NewInputStream(tt.input)), "Input: %s", tt.input)
		})
	}
}
