# Goval Syntax Documentation

## Overview

Goval is a lightweight expression language implemented in Go, designed for embedding in Go applications. It is suitable for scenarios such as rule engines and configuration scripts, and supports two type declaration styles: static typing and type inference (`var`).

This document describes the syntax rules of the Goval language in detail, including lexical and grammar structures.

## Lexical Rules

### Keywords

Goval defines the following keywords, which cannot be used as identifiers:

```
boolean, break, byte, case, char, continue, default, double, else, float, 
for, if, in, int, List, long, Map, return, Set, short, string, struct, switch, this, var
```

### Literals

#### Integer Literals

Integer literals in multiple bases are supported:

- Decimal: `0`, `123`, `456L`
- Hexadecimal: `0x1A`, `0xFF`
- Octal: `0755`, `0123`
- Binary: `0b1010`, `0B1111`

#### Floating-Point Literals

Floating-point literals in multiple formats are supported:

- Decimal form: `3.14`, `.5`, `2.`, `1.23f`
- Scientific notation: `1e10`, `1.5E-5`, `2.5e+3d`

#### Boolean Literals

- `true`
- `false`

#### Character Literals

Character literals are enclosed in single quotes:

- `'a'`
- `'\n'` (escape character)
- `'\u0041'` (Unicode)

#### String Literals

String literals are enclosed in double quotes:

- `"Hello World"`
- `"Line 1\nLine 2"` (with escape characters)

#### Null Literal

- `null`

### Identifiers

Identifiers are used to name variables, functions, etc. An identifier must begin with a letter, underscore (`_`), or dollar sign (`$`), followed by letters, digits, underscores, or dollar signs.

Supported character ranges include:
- Basic ASCII letters
- Extended ASCII (Latin-1 Supplement)
- Latin Extended-A and Extended-B
- CJK Unified Ideographs
- Hangul Syllables
- Hiragana and Katakana

### Operators

#### Assignment Operators
```
=, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=
```

#### Arithmetic Operators
```
+, -, *, /, %, ++, --
```

#### Comparison Operators
```
==, !=, <, >, <=, >=, in
```

#### Logical Operators
```
&&, ||, !
```

#### Bitwise Operators
```
&, |, ^, ~, <<, >>
```

#### Other Operators
```
->, ?, :, ., [], ()
```

### Separators
```
(, ), {, }, [, ], ;, ,, .
```

### Comments

Goval supports two comment formats:

- Single-line comment: `// this is a single-line comment`
- Multi-line comment: `/* this is a multi-line comment */`

## Grammar Rules

### Program Structure

A Goval program consists of one or more struct declarations or statements:

```
program : (structDeclaration | statement)+ EOF ;
```

### Type System

#### Primitive Types

Goval supports the following primitive types:

- Integer types: `byte`, `short`, `int`, `long`, `char`
- Floating-point types: `float`, `double`
- Boolean type: `boolean`
- String type: `string`

#### Type Declaration

Types can be declared in the following ways:

1. Primitive types: `int`, `string`, `boolean`, etc.
2. Array types: append `[]` to a primitive type, e.g. `int[]`, `string[][]`
3. Function types: `(T1, T2) -> R` or `(T1 name1, T2 name2) -> R`
4. Container types:
   - List: `List<T>`
   - Map: `Map<K,V>`
   - Set: `Set<T>`

### Variable Declaration

Goval supports two variable declaration styles:

#### Static Typing
```
int a = 1;
string name = "Goval";
boolean flag = true;
```

#### Type Inference with `var`
```
var x = 10;        // inferred as int
var y = "hello";   // inferred as string
var z = 3.14;      // inferred as double
```

Note: Variables declared with `var` must be initialized.

### Struct Declaration

Goval supports struct declarations using the `struct` keyword. A struct can contain fields and methods:

#### Basic Struct
```
struct Point {
    double x,
    double y
}
```

#### Struct with Methods
```
struct Rectangle {
    double width,
    double height,
    
    // calculate area
    area() -> double {
        return this.width * this.height;
    },
    
    // calculate perimeter
    perimeter() -> double {
        return 2 * (this.width + this.height);
    },
    
    // set dimensions (no return value)
    setSize(double w, double h) {
        this.width = w;
        this.height = h;
    }
}
```

#### Method Syntax

The basic syntax format of a method:
```
methodName(parameter1 type1, parameter2 type2) -> returnType {
    // method body
    return value;
}
```

- The method name is followed by a parameter list.
- If there is a return value, use `-> returnType` to specify the return type.
- If there is no return value, `-> returnType` can be omitted.
- Use the `this` keyword inside a method to reference the current instance.

#### The `this` Keyword

Inside struct methods, the `this` keyword can be used to access the current instance's fields and methods:

```
struct Counter {
    int count,
    
    increment() {
        this.count++;          // access field
    },
    
    getValue() -> int {
        return this.count;     // return field value
    },
    
    reset() {
        this.count = 0;        // set field value
    },
    
    doubleValue() -> int {
        return this.getValue() * 2;  // call another method
    }
}
```

### Struct Literals

Structs can be instantiated using literal syntax:

```
Point{x: 1.0, y: 2.0}
Rectangle{width: 10.0, height: 5.0}
```

### Container Literals

#### List Literals
```
List{1, 2, 3}    // using the List keyword
[1, 2, 3]        // syntactic sugar
```

#### Map Literals
```
Map{"key1": value1, "key2": value2}
```

#### Set Literals
```
Set{1, 2, 3}
```

### Control Flow Statements

#### If Statement
```
if (condition) {
    // statement block
} else {
    // statement block
}
```

#### Switch Statement
```
switch (expression) {
    case value1:
        // statement
        break;
    case value2:
        // statement
        break;
    default:
        // statement
}
```

#### Loop Statements

##### Basic For Loop
```
for (init; condition; update) {
    // statement block
}
```

##### Enhanced For Loop
```
for (type variable : collection) {
    // statement block
}
```

#### Control Statements

- `break`: exits a loop
- `continue`: skips the current loop iteration
- `return`: returns from a function

### Lambda Expressions

Goval supports lambda expressions:

```
// simple form
x -> x * 2

// with parameter types
(int x, int y) -> x + y

// with a statement block
(x, y) -> {
    int sum = x + y;
    return sum * 2;
}
```

### Expression Blocks

An expression block is a special statement block where the value of the last expression becomes the block's return value:

```
{
    int x = 10;
    int y = 20;
    x + y  // return value of the block
}
```

## Parser Generation

Goval uses ANTLR4 as the parser generator. Lexical rules and grammar rules are defined separately in `RuleExprLexer.g4` and `RuleExprParser.g4`.

Run the `generate.sh` script to generate the Go code:

```bash
./generate.sh
```

This will generate the parser and visitor code in the `internal/ast` directory.

## Usage Examples

The following is a simple example of using Goval:

```
// variable declaration
var x = 10;
string message = "hello";

// struct declaration
struct Calculator {
    int value,
    
    // set value
    setValue(int v) {
        this.value = v;
    },
    
    // get value
    getValue() -> int {
        return this.value;
    },
    
    // addition
    add(int n) -> int {
        this.value += n;
        return this.value;
    },
    
    // reset
    reset() {
        this.value = 0;
    }
}

// struct instantiation
Calculator calc = Calculator{value: 10};

// control flow
if (x > 5) {
    return "greater than 5";
} else {
    return "less than or equal to 5";
}

// lambda expression
var multiply = (a, b) -> a * b;

// container usage
List<int> numbers = [1, 2, 3, 4, 5];
Map<string, int> scores = Map{"Alice": 95, "Bob": 87};
```

## Summary

Goval provides a concise yet powerful expression language, especially suited for embedding in Go applications. It supports both static typing and `var` type inference, rich container types, structs with methods, control flow statements, and lambda expressions. Through the `this` keyword, struct methods can conveniently access the instance's fields and other methods, enabling developers to write object-oriented style code while maintaining type safety and high performance.

### Key Features

- **Type System**: Supports static type declarations and `var` type inference
- **Structs**: Supports fields and methods for object-oriented programming
- **`this` Keyword**: Access the current instance inside methods
- **Container Types**: Built-in List, Map, and Set support
- **Lambda Expressions**: Functional programming support
- **Control Flow**: Full if/else, for, and switch support
- **Expression Blocks**: Supports complex expression composition
