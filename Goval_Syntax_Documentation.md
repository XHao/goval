# Goval Syntax Documentation

## Overview

Goval is a lightweight expression language implemented in Go, designed for embedding in Go applications. It targets **rule engine** scenarios: pure-expression evaluation with immutable objects, lambdas, and minimal control flow.

Goval deliberately drops general-purpose scripting features — there are no structs, no `switch`, no `return`, no type annotations, and no mutable container syntax. Custom "types" are expressed as **lambda factories** returning Map literals; all objects and containers are **immutable** once constructed.

This document describes the syntax rules of the Goval language in detail, including lexical structure, grammar, semantics, and usage. All behavior described here is verified by the test suite in `internal/eval`, `internal/syntax`, and `pkg/goval`.

## Lexical Rules

### Keywords

Goval defines the following keywords, which cannot be used as identifiers:

```
break, continue, else, for, if, in, null, var
```

plus the boolean literals `true` and `false`.

The following keywords from the legacy grammar have been **removed** and are no longer reserved: `struct`, `this`, `return`, `switch`, `case`, `default`, `Set`, and all primitive type keywords (`boolean`, `byte`, `char`, `short`, `int`, `long`, `float`, `double`, `string`, `List`, `Map`). These words are now ordinary identifiers. Types are inferred via `var`; containers use literal syntax only.

### Literals

#### Integer Literals

Integer literals in multiple bases are supported. An optional `L`/`l` suffix is accepted and stripped (it does not change the type — Goval integers are always `int64`):

- Decimal: `0`, `123`, `456L`
- Hexadecimal: `0x1A`, `0xFF`, `0X1a`
- Octal: `0755`, `0123`
- Binary: `0b1010`, `0B1111`

Integer literals evaluate to `int64`. Examples: `0x1A` → 26, `0755` → 493, `0b1010` → 10, `123L` → 123.

#### Floating-Point Literals

Floating-point literals in multiple formats are supported. An optional `f`/`F`/`d`/`D` suffix is accepted and stripped (it does not change the type — Goval floats are always `float64`):

- Decimal form: `3.14`, `.5`, `2.`, `1.23f`, `2.5d`
- Scientific notation: `1e10`, `1.5E-5`, `2.5e+3d`

Floating-point literals evaluate to `float64`. Examples: `1e3` → 1000.0, `1.5E-2` → 0.015.

#### Boolean Literals

- `true`
- `false`

Evaluate to `bool`.

#### Character Literals

Character literals are enclosed in single quotes and evaluate to their Unicode code point as an `int64` (rune value):

- `'a'` → 97
- `'\n'` → 10 (newline escape)
- `'\t'` → 9 (tab escape)
- `'\\'` → 92 (escaped backslash)
- `'\''` → 39 (escaped single quote)

Supported escapes: `\n`, `\t`, `\\`, `\'`, `\"`.

#### String Literals

String literals are enclosed in double quotes and evaluate to `string`. Escape sequences are supported:

- `"Hello World"`
- `"Line 1\nLine 2"` (newline)
- `"Tab\there"` (tab)
- `"Back\\slash"` (backslash)
- `"Quote\"inside"` (embedded double quote)

#### Null Literal

- `null`

Evaluates to the null value (distinct from `false` or `0`).

### Identifiers

Identifiers are used to name variables, function parameters, and Map fields. An identifier must begin with a letter, underscore (`_`), or dollar sign (`$`), followed by letters, digits, underscores, or dollar signs.

Supported character ranges include:
- Basic ASCII letters
- Extended ASCII (Latin-1 Supplement)
- Latin Extended-A and Extended-B
- CJK Unified Ideographs
- Hangul Syllables
- Hiragana and Katakana

Examples: `x`, `_tmp`, `$value`, `用户`, `名前`.

### Operators

#### Assignment Operators
```
=
```

Only simple assignment (`=`) is supported. Compound assignment operators (`+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `|=`, `^=`, `<<=`, `>>=`) and increment/decrement operators (`++`, `--`) are **not supported** — they are rejected by the parser.

#### Arithmetic Operators
```
+, -, *, /, %
```

- `+` works on `int`/`float` (numeric addition) and `string` (concatenation). Mixing string and numeric operands is a runtime error.
- `/` on integers is integer division; on floats is float division. Division by zero is a runtime error.
- `%` is integer modulo (int operands only). Modulo by zero is a runtime error.

#### Comparison Operators
```
==, !=, <, >, <=, >=, in
```

- `==`/`!=` compare same-typed values.
- `<`, `>`, `<=`, `>=` work on `int`/`float`/`string`. Comparing mismatched types is a runtime error.
- `in` tests containment:
  - `x in list` — true if `x` equals an element of the List.
  - `k in map` — true if `k` (string) is a key of the Map.
  - `sub in str` — true if `sub` is a substring of `str`.

#### Logical Operators
```
&&, ||, !
```

Operands must be `bool`; otherwise a runtime error. `&&` and `||` short-circuit: the right operand is not evaluated if the left determines the result.

#### Bitwise Operators
```
&, |, ^, ~, <<, >>
```

Operands must be `int` (`int64`); otherwise a runtime error. `~` is unary bitwise NOT. `<<`/`>>` are left/right shift.

#### Other Operators
```
->, ?, :, ., [], ()
```

- `->` lambda arrow.
- `? :` ternary conditional.
- `.` field access (Map lookup).
- `[]` subscript access.
- `()` function/method call / grouping.

### Operator Precedence

Precedence is defined by the grammar rule hierarchy (lowest to highest). Within the same level, binary operators are left-associative; assignment and ternary are right-associative.

| Level | Operators | Associativity |
|-------|-----------|---------------|
| 1 (lowest) | `=` (assignment) | right |
| 2 | `? :` (ternary) | right |
| 3 | `\|\|` | left |
| 4 | `&&` | left |
| 5 | `\|` | left |
| 6 | `^` | left |
| 7 | `&` | left |
| 8 | `==` `!=` | left |
| 9 | `<` `>` `<=` `>=` `in` | left |
| 10 | `<<` `>>` | left |
| 11 | `+` `-` | left |
| 12 | `*` `/` `%` | left |
| 13 | `+` `-` `!` `~` (unary, prefix) | — |
| 14 (highest) | `[]` `.` `()` (postfix) | left |

Examples:
- `2 + 3 * 4` → 14 (multiplication binds tighter)
- `1 << 2 + 1` → 8 (i.e. `1 << (2+1)`)
- `0xF0 & 0x0F | 0x10` → 16 (i.e. `(0xF0 & 0x0F) | 0x10`; `&` binds tighter than `|`)
- `false ? 1 : true ? 2 : 3` → 2 (ternary right-associative: `false ? 1 : (true ? 2 : 3)`)

### Separators
```
(, ), {, }, [, ], ;, ,, .
```

Semicolons (`;`) are **optional** — statements may be separated by newlines or semicolons.

### Comments

Goval supports two comment formats, both skipped by the lexer:

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
var pi = 3.14;        // float
var multiply = (a, b) -> a * b;   // lambda
```

Multiple declarators are allowed in one statement:

```
var x = 1, y = 2, z = 3;
```

Static type declarations (`int a = 1;`, `string name = "Goval";`) and type annotations on parameters are **not supported** — use `var` for all declarations.

### Assignment and Single-Assignment Semantics

Goval enforces **single assignment**: a variable can be bound at most once. Once a name is bound (via `var` or `=`), any subsequent assignment to the same name is a compile-time error.

The left-hand side of an assignment **must be a bare identifier**. Field assignment (`p.name = ...`) and element assignment (`lst[i] = ...`, `m["k"] = ...`) are **syntactically rejected** by the parser.

```
x = 20;              // OK — first binding of x (equivalent to var x = 20)
result = x + 1;      // OK — first binding of result

var y = 1;
y = 2;               // ERROR — cannot rebind variable 'y' (single assignment)
```

**Important consequence:** because reassignment is forbidden, a loop body **cannot accumulate results into an outer variable**. For example, this is a compile error:

```
var s = 0;
for x in [1, 2, 3] {
    s = s + x        // ERROR: cannot rebind 's'
}
```

To aggregate values across a loop, use the built-in functions `reduce`, `map`, `filter`, which handle accumulation internally and return a new value.

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

The `else` clause is optional. Each branch introduces its own scope. `if` is a statement and does not produce a value. `switch`/`case`/`default` are **not supported** — use `if`/`else if` chains instead.

#### For-In Loop

Goval supports only the for-in form of the `for` loop. The three-part C-style `for (init; cond; update)` form is **not supported**.

Iterate over a List (value binding):
```
for x in [1, 2, 3] {
    // x is each element (int)
}
```

Iterate over a Map (key, value binding):
```
for k, v in userMap {
    // k is the key (string), v is the value
}
```

Iterate over a string (character binding):
```
for ch in "hello" {
    // ch is each character as a single-character string
}
```

> **Note:** string iteration yields single-character **strings**, not rune integers. (This is consistent with string subscripting — see [Postfix Access](#postfix-access).)

A single-identifier form `for k in map` binds `k` to each key (string).

#### Break and Continue

- `break;` — exits the enclosing `for` loop.
- `continue;` — skips to the next iteration of the enclosing `for` loop.

`break` and `continue` are only legal inside a `for` body; using them elsewhere is a semantic error. `return` is **not supported** — lambdas and expression blocks return their trailing expression (see [Expression Blocks](#expression-blocks)).

### Immutability Rules

Goval enforces a **fully immutable** model for objects and containers:

1. **Object fields are read-only.** Once a Map-based object is constructed, its fields cannot be reassigned. `p.name = "x"` is a syntax error.
2. **Container elements are read-only.** `lst[i] = ...` and `m["k"] = ...` are syntax errors.
3. **Single assignment.** A local variable can be bound at most once. Reassignment (`x = ...` after `x` is already bound) is a compile-time error. The only writable target is a bare identifier being bound for the first time.
4. **Container modification goes through built-in functions** that return new containers, leaving the original unchanged:

```
var lst2 = append(lst, x)        // returns a new List; lst is unchanged
var m2 = put(m, "k", v)          // returns a new Map
var lst3 = removeAt(lst, 0)      // returns a new List without the element at index 0
```

5. **No `this`, no `return`.** Methods are lambdas that capture constructor parameters via closure; they never reference the enclosing object itself. Lambdas return the value of their body's trailing expression.

### Objects: Lambda Factory + Map Literal

Custom "types" need no dedicated syntax. A type is a **factory function** that returns a Map literal: fields are Map keys, methods are closure values.

```
var Person = (name, age) -> {
    name: name,
    age: age,
    greet: () -> "hi " + name      // closure captures constructor param; no `this`
}

var p = Person("alice", 30)
p.name                              // "alice"  — Map lookup
p.greet()                           // "hi alice" — fetch lambda, then call
```

- Accessing `p.name` is a Map lookup; `p.greet()` is fetching the lambda value and invoking it. Both are the standard `.identifier` and `.identifier()` postfix forms — no special method-call rule.
- Each call to the factory produces an independent object (its own closure, its own Map).
- Methods capture **constructor parameters**, not the Map's current fields. This is the accepted trade-off of the immutable model: construct once, evaluate read-only, produce a result.

A method with parameters:
```
var Calculator = (base) -> {
    base: base,
    add: (n) -> base + n,
    scale: (factor, offset) -> base * factor + offset
}
```

> **Map bare-identifier key shorthand:** inside a Map literal, `{name: expr}` uses the identifier `name` as the string key (equivalent to `{"name": expr}`). This is the common form for object fields, as in `{ name: name, greet: () -> ... }`.

### Container Literals

#### List Literals
```
[1, 2, 3]
["a", "b", "c"]
[]
```

A List evaluates to `[]Value` (a list of Goval values). Elements may be of mixed types.

#### Map Literals
```
{"key1": value1, "key2": value2}
{}
{"name": "alice", "age": 30}
```

Map keys **must be strings** (either string literals or bare identifiers via the shorthand above). Using a non-string key is a runtime error. Map literals double as the object body inside a lambda factory (see above).

`Set` has no literal syntax and no keyword.

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
    sum          // trailing expression — the lambda's return value
}
```

A lambda captures variables from its enclosing scope by closure (capturing the value at the point of definition).

### Expression Blocks

An expression block is a `{ statements... expression }` form where the **trailing expression** becomes the block's value. Expression blocks are used as lambda block bodies and can also appear as a primary expression:

```
{
    var x = 10;
    var y = 20;
    x + y            // the value of the block (30)
}
```

The block's last element **must be an expression** (not a statement); there is no `return` keyword. This is how lambdas and blocks produce values. Expression blocks may be nested.

### Postfix Access

All access forms are **read-only**:

```
p.name              // field access (Map lookup)
lst[i]              // subscript access
f(args)             // function call
obj.method(args)    // method call (fetch lambda, then call)
```

Subscript access `base[i]` behavior depends on the base type:
- **List:** `i` must be an int index in `[0, len)`. Out-of-range or negative indices are a runtime error. Returns the element.
- **Map:** `i` must be a string key. Returns the value, or `null` if the key is absent.
- **string:** `i` must be an int index in `[0, rune-length)`. Returns a **single-character string** (not a rune integer). Out-of-range is a runtime error.

Access chains left-associatively, so `a.b.c` and `m["a"]["b"]` work as expected.

## Built-in Functions

Goval provides these built-in functions, available in every scope:

| Function | Signature | Description |
|----------|-----------|-------------|
| `reduce(list, init, fn)` | `(List, Value, (acc, x) -> Value) -> Value` | Left-fold: applies `fn` to accumulate elements of `list` starting from `init`. |
| `map(list, fn)` | `(List, (x) -> Value) -> List` | Returns a new List with `fn` applied to each element. |
| `filter(list, fn)` | `(List, (x) -> bool) -> List` | Returns a new List containing only elements for which `fn` returns true. |
| `find(list, fn)` | `(List, (x) -> bool) -> Value` | Returns the first element for which `fn` returns true, or `null` if none match. |
| `append(list, item)` | `(List, Value) -> List` | Returns a new List with `item` added at the end. |
| `put(map, key, value)` | `(Map, string, Value) -> Map` | Returns a new Map with `key` set to `value`. |
| `removeAt(list, index)` | `(List, int) -> List` | Returns a new List with the element at `index` removed. |
| `len(v)` | `(List\|string\|Map) -> int` | Returns the length of a List, string, or Map. |
| `range(start, end)` | `(int, int) -> List` | Returns a new List of integers `[start, start+1, ..., end-1]`. Empty if `start >= end`. |

All built-ins return **new** values; they never mutate their inputs (consistent with the immutability rules).

## Parser Generation

Goval uses ANTLR4 as the parser generator. Lexical rules and grammar rules are defined separately in `grammar/RuleExprLexer.g4` and `grammar/RuleExprParser.g4`.

Run the `generate.sh` script to generate the Go code:

```bash
./generate.sh
```

This regenerates the parser and visitor code in the `internal/ast` directory.

## Static Checks

The public `goval.Evaluate` runs the full pipeline — parsing, semantic checks, compilation, evaluation — and errors from any stage are returned as `error`:

1. **Grammar-level** — the left side of an assignment must be a bare identifier; `.field =` and `[i] =` are parse errors. Unsupported syntax (compound assignment, `++`/`--`, C-style three-part `for`, `switch`, `return`) is likewise rejected at parse time.
2. **Semantic checks** — the `SyntaxChecker` (in `internal/syntax`) validates the parse tree beyond what the grammar enforces: `break` and `continue` are only legal inside a `for` body.
3. **Compile-time checks** — single assignment is enforced in the `eval` compiler: rebinding an already-declared variable is a compile-time error.

## Embedding API

Goval is embedded via the public `goval.Evaluate` function:

```go
v, err := goval.Evaluate(source string, context map[string]interface{}) (interface{}, error)
```

- `source` is a Goval program string.
- `context` injects Go values as global variables. Supported Go types: `int`, `int64`, `float64`, `float32`, `bool`, `string`, `nil`, `[]interface{}`, `map[string]interface{}`.
- The result is a Go native value: `int64`, `float64`, `bool`, `string`, `nil`, `[]interface{}`, or `map[string]interface{}`.
- Errors from every stage are returned as `error` — never as a panic to the caller: syntax errors (rejected input, unsupported operators, non-identifier lvalues), semantic errors (`break`/`continue` outside a loop), compile errors (single-assignment violations), and runtime panics (e.g. division by zero, type mismatches) are all recovered and returned.

```go
v, err := goval.Evaluate("1 + 2 * 3", nil)
// v == int64(7)

ctx := map[string]interface{}{"x": float64(10)}
v, err := goval.Evaluate("x > 5", ctx)
// v == true
```

## Usage Example

```
// Rule: discount VIP users' large orders.
// Objects are lambda factories returning Map literals; everything is immutable.

var Order = (amount, userId) -> {
    amount: amount,
    userId: userId,
    discounted: (rate) -> amount * rate
}
var User = (id, level) -> { id: id, level: level }

var users = [
    User("u1", "gold"),
    User("u2", "vip")
]

// Build a lookup map id -> user via reduce (single assignment forbids
// reassigning an outer variable inside a for loop, so use reduce).
var userMap = reduce(users, {}, (acc, u) -> put(acc, u.id, u))

var order = Order(500, "u2")
var user = userMap[order.userId]
var final = user.level == "vip" && order.amount > 100
        ? order.discounted(0.8)
        : order.amount
```

> Note how `reduce` is used to build `userMap` instead of a mutating `for` loop — this is a direct consequence of the single-assignment rule.

## Summary

Goval is a concise expression language for rule engines, built on three ideas:

- **Objects as lambda factories + Map literals** — no `struct`, no `this`, no method-declaration syntax. A "type" is a function returning a Map; methods are closures that capture constructor parameters.
- **Full immutability + single assignment** — object fields and container elements are read-only; a local variable can be bound at most once. Container updates go through built-in functions that return new containers.
- **Expression-oriented control flow** — `if`/`else`, `for`-`in`, `break`/`continue`. No `switch`, no `return`, no three-part `for`. Lambdas and expression blocks return their trailing expression.

### Key Features

- **Type Inference**: `var` declarations with mandatory initialization; no type annotations.
- **Lambda Factories**: custom types as functions returning Map literals.
- **Immutability & Single Assignment**: identifier-only lvalues; each variable bound once; containers modified via built-in functions.
- **Container Literals**: List `[...]` and Map `{...}` (string keys, with bare-identifier shorthand).
- **Control Flow**: `if`/`else`, `for`-`in` over List/Map/string, `break`/`continue`.
- **Expression Blocks**: `{ stmts; expr }` — trailing expression is the block's value.
- **Lambda Closures**: `(params) -> expr` or `(params) -> { stmts; expr }`.
- **Rich Operators**: arithmetic, comparison, logical (short-circuit), bitwise, shift, `in`, ternary — with C-like precedence.
- **Built-in Functions**: `reduce`, `map`, `filter`, `find`, `append`, `put`, `removeAt`, `len`, `range`.
