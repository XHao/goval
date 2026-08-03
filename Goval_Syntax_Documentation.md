# Goval Syntax Documentation

## Overview

Goval is a lightweight expression language implemented in Go, designed for embedding in Go applications. It targets **rule engine** scenarios: pure-expression evaluation with immutable objects, lambdas, and minimal control flow.

Goval deliberately drops general-purpose scripting features — there are no structs, no `switch`, no `return`, no type annotations, and no mutable container syntax. Custom "types" are expressed as **lambda factories** returning Map literals; all objects and containers are **immutable** once constructed. The only writable targets are local variables (identifiers), used for intermediate results during rule evaluation.

This document describes the syntax rules of the Goval language in detail, including lexical and grammar structures.

## Lexical Rules

### Keywords

Goval defines the following keywords, which cannot be used as identifiers:

```
break, continue, else, for, if, in, null, var
```

plus the boolean literals `true` and `false`.

The following keywords from the legacy grammar have been **removed** and are no longer reserved: `struct`, `this`, `return`, `switch`, `case`, `default`, `Set`, and all primitive type keywords (`boolean`, `byte`, `char`, `short`, `int`, `long`, `float`, `double`, `string`, `List`, `Map`). Types are inferred via `var`; containers use literal syntax only.

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
- `'A'` (Unicode)

#### String Literals

String literals are enclosed in double quotes:

- `"Hello World"`
- `"Line 1\nLine 2"` (with escape characters)

#### Null Literal

- `null`

#### Placeholder Variable

A placeholder variable is written as `#identifier#` and is substituted by the host application before evaluation:

- `#user#`
- `#order.amount#`

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
=
```

Only simple assignment (`=`) is supported. Compound assignment operators (`+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `|=`, `^=`, `<<=`, `>>=`) and increment/decrement operators (`++`, `--`) have been removed — they conflict with the immutability rules (see [Immutability Rules](#immutability-rules)).

#### Arithmetic Operators
```
+, -, *, /, %
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

A Goval program consists of zero or more statements:

```
program : statement* EOF ;
```

### Statements

A statement is one of:

```
statement
    : block
    | ifStatement
    | forStatement
    | breakStatement
    | continueStatement
    | expressionStatement
    | localVariableDeclarationStatement
    | SEMI
    ;
```

### Variable Declaration

Variables are declared with `var` and **must be initialized** (the type is inferred from the initializer):

```
var x = 10;           // int
var name = "Goval";   // string
var flag = true;      // boolean
var pi = 3.14;        // float/double
var multiply = (a, b) -> a * b;   // lambda
```

Multiple declarators are allowed in one statement:

```
var x = 1, y = 2, z = 3;
```

Static type declarations (`int a = 1;`, `string name = "Goval";`) and type annotations on parameters are **not supported** — use `var` for all declarations.

### Assignment

Assignment rebinds an existing variable. The left-hand side **must be a bare identifier** — this is the immutability rule (see [Immutability Rules](#immutability-rules)):

```
x = 20;
result = x + 1;
```

Field assignment (`p.name = ...`) and element assignment (`lst[i] = ...`, `m["k"] = ...`) are **syntactically rejected**.

### Control Flow

#### If / Else If / Else
```
if (condition) {
    // statement block
} else if (otherCondition) {
    // statement block
} else {
    // statement block
}
```

The `else` clause is optional. `switch`/`case`/`default` are **not supported** — use `if`/`else if` chains instead.

#### For-In Loop

Goval supports only the for-in form of the `for` loop. The three-part C-style `for (init; cond; update)` form is **not supported**.

Iterate over a List (value binding):
```
for x in [1, 2, 3] {
    // use x
}
```

Iterate over a Map (key, value binding):
```
for k, v in userMap {
    // use k and v
}
```

Iterate over a string (character binding):
```
for ch in "hello" {
    // use ch
}
```

#### Break and Continue

- `break;` — exits the enclosing `for` loop.
- `continue;` — skips to the next iteration of the enclosing `for` loop.

`break` and `continue` are only legal inside a `for` body; using them elsewhere is a semantic error. `return` is **not supported** — lambdas and expression blocks return their trailing expression (see [Expression Blocks](#expression-blocks)).

### Immutability Rules

Goval enforces a **fully immutable** model for objects and containers:

1. **Object fields are read-only.** Once a Map-based object is constructed, its fields cannot be reassigned. `p.name = "x"` is a syntax error.
2. **Container elements are read-only.** `lst[i] = ...` and `m["k"] = ...` are syntax errors.
3. **The only writable target is a local variable identifier.** `x = ...` is allowed at any scope; `.field =` and `[i] =` are rejected by the parser.
4. **Container modification goes through built-in functions** that return new containers, leaving the original unchanged:

```
lst2 = append(lst, x)        // returns a new List; lst is unchanged
m2 = put(m, "k", v)          // returns a new Map
lst2 = removeAt(lst, 0)      // returns a new List without the element at index 0
```

5. **No `this`, no `return`.** Methods are lambdas that capture constructor parameters via closure; they never reference the enclosing object itself. Lambdas return the value of their body's trailing expression.

### Objects: Lambda Factory + Map Literal

Custom "types" need no dedicated syntax. A type is a **factory function** that returns a Map literal: fields are Map keys, methods are closure values.

```
Person = (name, age) -> {
    name: name,
    age: age,
    greet: () -> "hi " + name      // closure captures constructor param; no `this`
}

p = Person("alice", 30)
p.name                              // "alice"  — Map lookup
p.greet()                           // "hi alice" — fetch lambda, then call
```

- Accessing `p.name` is a Map lookup; `p.greet()` is fetching the lambda value and invoking it. Both are the standard `.identifier` and `.identifier()` postfix forms — no special method-call rule.
- Each call to the factory produces an independent object (its own closure, its own Map).
- Methods capture **constructor parameters**, not the Map's current fields. This is the accepted trade-off of the immutable model: construct once, evaluate read-only, produce a result.

A method with no parameters:
```
Counter = (start) -> {
    value: start,
    twice: () -> start * 2
}
```

A method with parameters:
```
Calculator = (base) -> {
    base: base,
    add: (n) -> base + n,
    scale: (factor, offset) -> base * factor + offset
}
```

### Container Literals

#### List Literals
```
[1, 2, 3]
["a", "b", "c"]
[]
```

#### Map Literals
```
{"key1": value1, "key2": value2}
{}
{"name": "alice", "age": 30}
```

Map literals double as the object body inside a lambda factory (see above).

`Set` has no literal syntax and no keyword. Set operations are provided by built-in functions (e.g. `setOf`, `unique`) rather than grammar-level syntax.

### Lambda Expressions

A lambda has the form `parameters -> body`:

```
// single parameter, no parentheses needed
x -> x * 2

// zero parameters
() -> "hello"

// multiple parameters
(a, b) -> a + b
```

Parameters are **bare identifiers** — no type annotations. The body is either a single expression or an expression block:

```
// expression body
(x, y) -> x + y

// block body: statements followed by a trailing expression
(x, y) -> {
    var sum = x + y;
    sum * 2          // trailing expression — the lambda's return value
}
```

A lambda captures variables from its enclosing scope by closure.

### Expression Blocks

An expression block is a `{ statements... expression }` form where the **trailing expression** becomes the block's value. Expression blocks are used as lambda block bodies and can also appear as a primary expression:

```
{
    var x = 10;
    var y = 20;
    x + y            // the value of the block
}
```

The block's last element **must be an expression** (not a statement); there is no `return` keyword. This is how lambdas and blocks produce values.

### Postfix Access

All access forms are **read-only**:

```
p.name              // field access (Map lookup)
lst[i]              // subscript access (List/Map/string)
f(args)             // function call
obj.method(args)    // method call (fetch lambda, then call)
```

## Parser Generation

Goval uses ANTLR4 as the parser generator. Lexical rules and grammar rules are defined separately in `grammar/RuleExprLexer.g4` and `grammar/RuleExprParser.g4`.

Run the `generate.sh` script to generate the Go code:

```bash
./generate.sh
```

This regenerates the parser and visitor code in the `internal/ast` directory.

## Static Checks

The `SyntaxChecker` (in `internal/syntax`) validates a parse tree beyond what the grammar enforces:

1. **Lvalue check** — the left side of an assignment must be an identifier. `.field =` and `[i] =` are rejected.
2. **For-in target check** — the expression after `in` must be a traversable type (List, Map, or string).
3. **break/continue scope** — `break` and `continue` are only legal inside a `for` body.
4. **Lambda block tail** — a lambda block body must end with an expression (the return value).

## Usage Example

```
// Rule: discount VIP users' large orders

Order = (amount, userId) -> {
    amount: amount,
    userId: userId,
    discounted: (rate) -> amount * rate
}
User = (id, level) -> { id: id, level: level }

users = [
    User("u1", "gold"),
    User("u2", "vip")
]
userMap = {}
for u in users {
    userMap = put(userMap, u.id, u)
}

order = Order(500, "u2")
user = userMap[order.userId]
final = user.level == "vip" && order.amount > 100
        ? order.discounted(0.8)
        : order.amount
```

## Summary

Goval is a concise expression language for rule engines, built on three ideas:

- **Objects as lambda factories + Map literals** — no `struct`, no `this`, no method-declaration syntax. A "type" is a function returning a Map; methods are closures that capture constructor parameters.
- **Full immutability** — object fields and container elements are read-only. The only writable targets are local variables (identifiers). Container updates go through built-in functions that return new containers.
- **Expression-oriented control flow** — `if`/`else`, `for`-`in`, `break`/`continue`. No `switch`, no `return`, no three-part `for`. Lambdas and expression blocks return their trailing expression.

### Key Features

- **Type Inference**: `var` declarations with mandatory initialization; no type annotations.
- **Lambda Factories**: custom types as functions returning Map literals.
- **Immutability**: identifier-only lvalues; containers modified via built-in functions.
- **Container Literals**: List `[...]` and Map `{...}`; Set via built-in functions only.
- **Control Flow**: `if`/`else`, `for`-`in` over List/Map/string, `break`/`continue`.
- **Expression Blocks**: `{ stmts; expr }` — trailing expression is the block's value.
- **Lambda Closures**: `(params) -> expr` or `(params) -> { stmts; expr }`.
