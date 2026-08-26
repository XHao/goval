# Goval Syntax Documentation

## Overview

Goval is a lightweight expression language implemented in Go, designed for embedding in Go applications. It targets **rule engine** scenarios: expression evaluation with lambda-based objects and minimal control flow.

Goval deliberately drops general-purpose scripting features — there are no structs, no `switch`, no `return`, no type annotations. Custom "types" are expressed as **lambda factories** returning Map literals. Variables can be reassigned and object fields / container elements can be written in place (v2 semantics), while the language stays sandboxed: no I/O, and host data is deep-copied on the way in and out.

This document describes the syntax rules of the Goval language in detail, including lexical structure, grammar, semantics, and usage. All behavior described here is verified by the test suite in `internal/eval`, `internal/syntax`, and `pkg/goval`.

## Lexical Rules

### Keywords

Goval defines the following keywords, which cannot be used as identifiers:

```
break, continue, else, for, if, in, null, var, this, capture
```

plus the boolean literals `true` and `false`.

`this` is bound at method-call sites (see [Objects](#objects-lambda-factory--map-literal)); `capture` is a compile-time pragma that disables implicit-`this` rewriting (see [Lambda Expressions](#lambda-expressions)).

The following keywords from the legacy grammar have been **removed** and are no longer reserved: `struct`, `return`, `switch`, `case`, `default`, `Set`, and all primitive type keywords (`boolean`, `byte`, `char`, `short`, `int`, `long`, `float`, `double`, `string`, `List`, `Map`). These words are now ordinary identifiers. Types are inferred via `var`; containers use literal syntax only.

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

Supported escapes (shared by string literals): `\n`, `\t`, `\r`, `\0`, `\\`, `\'`, `\"`.

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

Only simple assignment (`=`) is supported. The left-hand side may be a bare identifier, a field access (`p.name = ...`), or a subscript (`lst[i] = ...`, `m["k"] = ...`) — see [Assignment Semantics](#assignment-semantics-v2). Compound assignment operators (`+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `|=`, `^=`, `<<=`, `>>=`) and increment/decrement operators (`++`, `--`) are **not supported** — they are rejected by the parser.

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

- `==`/`!=` compare **by value**. Numeric operands compare numerically across `int`/`float` (`1 == 1.0` → `true`; Go JSON numbers inject as `float64`, so `amount == 100` matches an int literal). Lists and Maps compare deeply — equal length / key set plus element-wise value equality (`[1, 2] == [1, 2]` → `true`). Values of different kinds are simply unequal (`1 == "1"`, `null == 0`, `true == 1`); lambdas compare by identity.
- `<`, `>`, `<=`, `>=` work on `int`/`float`/`string`. Comparing mismatched types is a runtime error.
- `in` tests containment:
  - `x in list` — true if `x` equals an element of the List (value equality; `1 in [1.0]` → `true`).
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
- `? :` ternary conditional. Both branches are full expressions — assignments may appear bare in either branch (`cond ? x = 1 : y = 2`). The ternary is right-associative. The condition must be `bool` — anything else is a runtime error (consistent with `if`).
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

Semicolons (`;`) are **optional** — statements may be separated by newlines or semicolons. Note that newlines are plain whitespace, not automatic statement terminators: if a line **starts with `[` or `(`**, it continues the previous expression (e.g. `var a = [1]` followed by `[2]` parses as `a = [1][2]`). Write a semicolon at the end of such lines.

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

### Assignment Semantics (v2)

Goval allows **reassignment**: a variable declared with `var` can be updated afterwards. An assignment writes to the nearest enclosing binding of the name (walking the scope chain), so assignments inside blocks, `if` branches, and loop bodies update outer variables:

```
var s = 0;
for x in [1, 2, 3] {
    s = s + x        // OK — updates the outer s (s == 6 after the loop)
}
```

**Strict assignment:** the target of a bare `x = ...` must already be bound (via `var`, an enclosing scope, the embedding context, or a builtin). Assigning to an undeclared name is a **compile-time error** — this catches typos that would otherwise silently create a new variable:

```
var userId = 1;
usreId = 2;           // ERROR — undefined variable 'usreId'
```

**Field and element assignment** are supported in place:

```
p.name = "bob";       // writes the Map field
lst[0] = 9;           // writes the List element (bounds-checked)
m["k"] = 1;           // writes the Map entry (key must be a string)
a.b.c = 3;            // nested chains work
```

- Strings are immutable: `s[0] = 'x'` is a runtime error.
- Assignments are **expressions** — they evaluate to the assigned value (`var y = (x = 5) + 1`).

Assignment targets must be identifiers, field accesses, or subscripts; assigning to a call result (`f() = 1`) is rejected.

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

The `else` clause is optional. Each branch introduces its own scope. `if` is a statement and does not produce a value. The condition **must be `bool`** — anything else is a runtime error (consistent with `&&`/`||`/`!`). `switch`/`case`/`default` are **not supported** — use `if`/`else if` chains instead.

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

### Mutation Rules (v2)

1. **Variables can be reassigned.** `x = ...` updates the nearest enclosing binding; a loop body can accumulate into an outer variable. Redeclaring the same name with `var` in the *same* scope is still an error (inner scopes may shadow).
2. **Object fields are writable in place.** `p.name = "x"` writes the Map entry; `p["name"] = "x"` is equivalent.
3. **Container elements are writable in place.** `lst[i] = ...` (bounds-checked) and `m["k"] = ...`. Aliases share the container: `var b = a; a[0] = 9; b[0]` → `9`. Strings are immutable.
4. **Host data is isolated.** Context values are deep-copied on injection and on return — scripts can never mutate caller-owned Go data.
5. **No `return`.** Methods are lambdas; they return their trailing expression. Mutation happens through field/element assignment and through the pure built-ins (`append`, `put`, ...) which still return new containers.

### Objects: Lambda Factory + Map Literal

Custom "types" need no dedicated syntax. A type is a **factory function** that returns a Map literal: fields are Map keys, methods are lambda values.

```
var Person = (name, age) -> {
    name: name,
    age: age,
    greet: () -> "hi " + name,              // implicit this — see below
    birthday: () -> { age = age + 1; age }, // writes the FIELD, p.age updates
    olderThan: (o) -> age > o.age
}

var p = Person("alice", 30)
p.greet()          // "hi alice"
p.birthday()       // 31
p.age              // 31 — field and method views always agree
p.olderThan(Person("bob", 20))   // true
```

#### `this` binding

`obj.method(args)` binds `this` to `obj` (the Map the method was fetched from) at the call site. `m["f"](args)` behaves identically. Consequences:

- **Explicit form**: `this.name` reads/writes the receiver's field at call time.
- **Nested lambdas** inside a method see `this` through the environment chain.
- **`obj.other()`** works — `this` is re-bound per call to the same Map.
- **Detached methods have no `this`**: `var g = p.greet; g()` — accessing `this` (or an implicitly rewritten field) inside the call is a runtime error, like detached functions in JavaScript.

#### Implicit `this` rewriting

Inside a lambda that is lexically nested in a Map literal, a bare identifier that **collides with a field name of that Map** (and is not shadowed by a parameter or local `var`) is compiled as a field access on `this` — for both reads and writes:

```
greet: () -> "hi " + name       // compiles to this.name
birthday: () -> { age = age + 1; age }   // compiles to this.age = this.age + 1
```

- **Construction position is exempt**: the field initializer `name: name` reads the factory parameter (evaluated immediately at construction).
- **Parameters and locals shadow fields**: `f: (name) -> name` uses the parameter; `var name = ...` inside a method body shadows the field.
- The rewrite makes wrong-slot writes impossible: a field-colliding bare name *always* means the field, never a captured closure variable.
- **Implicit `this` is lexical, explicit `this` is dynamic.** A lambda gets implicit field rewriting only if it is *written inside* the Map literal. A lambda defined elsewhere and stored into a Map later keeps pure closure semantics — `m.f()` still binds `this`, but the body only sees the receiver if it says `this.x` explicitly.
- **Method names shadow outer bindings too.** The rewrite zone covers *all* Map keys, including methods. If a bare name collides with a sibling method's name, it becomes `this.<method>` — use `capture` (or rename) when you mean an outer function with the same name.
- **Nested Maps resolve against the receiver.** In `{ x: 1, inner: { f: () -> x } }`, calling `M.inner.f()` rewrites `x` to `this.x` and `this` is `inner` — the missing key yields `null` (consistent with Map subscripting). Reference the outer object explicitly (`M.x`) in such cases.
- **`capture` opts out** (see [Lambda Expressions](#lambda-expressions)): inside `capture { ... }` bare names follow pure closure semantics.

#### Private state (closure objects)

Variables declared in the factory body but not stored in the Map are private; only methods can reach them:

```
var Counter = (start) -> {
    var count = start              // private — not a Map field
    {
        inc: () -> { count = count + 1; count },
        value: () -> count
    }
}
var c = Counter(0)
c.inc()      // 1
c.inc()      // 2
c.value()    // 2
c.count      // null — no such field
```

Private state usually doesn't collide with field names, so it works without `capture`. Each factory call produces independent state (its own environment layer).

> **Map bare-identifier key shorthand:** inside a Map literal, `{name: expr}` uses the identifier `name` as the string key (equivalent to `{"name": expr}`).

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

A lambda **shares variable bindings with its defining scope** (reference capture, like JS/Python/Go):

- A closure sees later modifications of the captured variable, and assignments inside the closure update the outer binding (`Env.Assign` walks the scope chain).
- Sibling closures created in the same environment layer share its slots — this is the basis of the private-state pattern.
- **Per-iteration capture**: each lambda *call* creates a fresh environment layer, so closures created in different loop iterations (or `map` callback invocations) capture independent bindings:

```
var fns = map(range(0, 3), (i) -> () -> i)
[fns[0](), fns[1](), fns[2]()]     // [0, 1, 2] — not [2, 2, 2]
```

- **Recursion** works via self-reference: `var fact = (n) -> n <= 1 ? 1 : n * fact(n - 1)`.
- **Lambda identity**: `f == f` is true for the same closure; two distinct closures with identical source are not equal.

#### `capture` pragma

`capture` is a **compile-time-only** prefix that disables implicit-`this` rewriting within its lexical range. Runtime behavior is identical to the wrapped expression.

```
var level = "outer"
var P = capture { level: "field", read: () -> level }
P.read()                            // "outer" — bare level follows closure rules

var Q = { level: "field", read: () -> level }
Q.read()                            // "field" — implicit this wins without capture
```

It can also prefix a single lambda: `capture (a) -> a + b`.

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

```
p.name              // field access (Map lookup)
lst[i]              // subscript access
f(args)             // function call
obj.method(args)    // method call (binds this to obj)
m["f"](args)        // subscript call (also binds this)
s.upper()           // string built-in method (see String Methods)
```

Subscript access `base[i]` behavior depends on the base type:
- **List:** `i` must be an int index in `[0, len)`. Out-of-range or negative indices are a runtime error. Returns the element.
- **Map:** `i` must be a string key. Returns the value, or `null` if the key is absent.
- **string:** `i` must be an int index in `[0, rune-length)`. Returns a **single-character string** (not a rune integer). Out-of-range is a runtime error.

Field and subscript positions are also valid **assignment targets** (see [Assignment Semantics](#assignment-semantics-v2)); string subscripting is read-only.

Access chains left-associatively, so `a.b.c` and `m["a"]["b"]` work as expected.

## String Methods

Strings carry a built-in method set, invoked in postfix form `s.method(args)`. Strings are immutable — every method returns a new value. All indices are **rune indices**, consistent with `s[i]`, `len(s)`, and `for ch in s` (multi-byte characters count as one).

| Method | Signature | Description |
|--------|-----------|-------------|
| `upper()` | `() -> string` | Uppercase (Unicode-aware). |
| `lower()` | `() -> string` | Lowercase (Unicode-aware). |
| `trim()` | `() -> string` | Whitespace removed from both ends. |
| `contains(sub)` | `(string) -> bool` | True if `sub` is a substring. |
| `startsWith(prefix)` | `(string) -> bool` | Prefix test. |
| `endsWith(suffix)` | `(string) -> bool` | Suffix test. |
| `indexOf(sub)` | `(string) -> int` | Rune index of the first occurrence; `-1` if absent. |
| `lastIndexOf(sub)` | `(string) -> int` | Rune index of the last occurrence; `-1` if absent. |
| `substring(start[, end])` | `(int[, int]) -> string` | Substring over `[start, end)`; `end` defaults to the length. Negative indices, out-of-range, or `start > end` are runtime errors (consistent with `s[i]`). |
| `replace(old, new)` | `(string, string) -> string` | All occurrences of `old` replaced with `new`. |
| `repeat(n)` | `(int) -> string` | Receiver repeated `n` times; negative `n` is a runtime error. |
| `reverse()` | `() -> string` | Reversed by rune. |
| `split(sep)` | `(string) -> List` | Split into a List of strings; an empty separator splits per rune. |
| `join(list)` | `(List) -> string` | Concatenate the List's string elements with the receiver as separator. |

```
"  Hello  ".trim().lower()                 // "hello" — chains naturally
"a,b,c".split(",")                         // ["a", "b", "c"]
",".join(["x", "y"])                       // "x,y"
"你好世界".indexOf("世")                      // 2 — rune index, not byte offset
"https://x".startsWith("https") ? 1 : 0    // 1 — works in ternary conditions
map("a,b".split(","), x -> x.upper())      // ["A", "B"] — combines with builtins
```

- Calling an unknown method is a runtime error (`no method '...' on string`); so are wrong argument types or arity.
- Methods attach only to strings — there is no implicit `toString` conversion; calling a method on other kinds keeps the existing `cannot call method` error.
- `join` requires every element to be a string; `join` on a non-List or with non-string elements is a runtime error.

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
| `int(v)` | `(int\|float\|string) -> int` | Convert to `int`. Float **truncates** (Go semantics: `int(12.7)` → `12`); string must be a strict decimal integer (`"12.5"` or garbage → runtime error). |
| `float(v)` | `(int\|float\|string) -> float` | Convert to `float`; string parsed (`"12.5"`, `"1e3"` → `1000.0`). |
| `string(v)` | `(int\|float\|bool\|string) -> string` | Format as string with Go's shortest representation (`string(5.0)` → `"5"`); null/containers are runtime errors — no implicit `toString`. |
| `abs(x)` | `(int\|float) -> int\|float` | Absolute value, type-preserving (`abs(-3)` → `3` as int; `abs(-3.5)` → `3.5`). |
| `round(x[, digits])` | `(number[, int]) -> float` | Round half away from zero; optional decimal digits (`round(8.247, 2)` → `8.25`). Always returns float. |
| `floor(x)` / `ceil(x)` | `(number) -> float` | Floor / ceiling (Go `math` semantics); wrap with `int(...)` when an integer index is needed. |
| `min(...)` / `max(...)` | `(...number) -> number` | Variadic scalars (`min(1, 2)` → `1`) or a single List argument (`max(prices)`); all-int stays int, any float promotes the result to float. |
| `keys(m)` | `(Map) -> List` | Keys sorted lexicographically — deterministic output despite Go's random map iteration. |
| `values(m)` | `(Map) -> List` | Values ordered by the sorted keys (pairs align with `keys`). |

All built-ins return **new** values; they never mutate their inputs. In-place mutation goes through field/element assignment — the two styles coexist (pure helpers + in-place writes), mirroring Go slices/maps.

Built-ins validate their arguments — wrong arity or argument types are runtime errors naming the builtin (`reduce expected 3 args, got 2`; `map requires a list, got int`; `range requires int bounds, got string`). `filter`/`find` callbacks must return `bool`. `removeAt` bounds-checks its index like List subscripting (`removeAt index out of range: 5 (len 2)`). `range(start, end)` with `start >= end` returns an empty List.

## Parser Generation

Goval uses ANTLR4 as the parser generator. Lexical rules and grammar rules are defined separately in `grammar/RuleExprLexer.g4` and `grammar/RuleExprParser.g4`.

Run the `generate.sh` script to generate the Go code:

```bash
./generate.sh
```

This regenerates the parser and visitor code in the `internal/ast` directory.

## Static Checks

The public `goval.Evaluate` runs the full pipeline — parsing, semantic checks, compilation, evaluation — and errors from any stage are returned as `error`:

1. **Grammar-level** — assignment targets are limited to identifier / field / subscript shapes; compound assignment, `++`/`--`, C-style three-part `for`, `switch`, `return` are rejected at parse time.
2. **Semantic checks** — the `SyntaxChecker` (in `internal/syntax`) validates the parse tree beyond what the grammar enforces: `break` and `continue` are only legal inside a `for` body.
3. **Compile-time checks** — the `eval` compiler enforces: strict assignment (bare `x = ...` requires the name to be already bound), no same-scope `var` redeclaration, and invalid assignment targets (e.g. call results).

## Embedding API

Goval is embedded via the public `goval.Evaluate` function:

```go
v, err := goval.Evaluate(source string, context map[string]interface{}) (interface{}, error)
```

- `source` is a Goval program string.
- `context` injects Go values as global variables. Supported Go types: `int`, `int64`, `float64`, `float32`, `bool`, `string`, `nil`, `[]interface{}`, `map[string]interface{}`. Any other type (e.g. `func`, structs, `int32`) is rejected with an error naming the offending key and its Go type (`context "f": unsupported context value type func(int) int`) — unsupported values are never silently converted to `null`.
- The result is a Go native value: `int64`, `float64`, `bool`, `string`, `nil`, `[]interface{}`, or `map[string]interface{}`.
- Errors from every stage are returned as `error` — never as a panic to the caller: syntax errors (rejected input, unsupported operators, invalid assignment targets), semantic errors (`break`/`continue` outside a loop), compile errors (strict-assignment violations, same-scope redeclaration), and runtime panics (e.g. division by zero, type mismatches) are all recovered and returned.

```go
v, err := goval.Evaluate("1 + 2 * 3", nil)
// v == int64(7)

ctx := map[string]interface{}{"x": float64(10)}
v, err := goval.Evaluate("x > 5", ctx)
// v == true
```

### Repeated Evaluation

For rule engines applying the same rule to many rows, compile once and run per row:

```go
p, err := goval.Compile(source, "amount", "user") // context keys fixed at compile time
v, err := p.Run(map[string]interface{}{"amount": 10, "user": user})
```

- `Compile` takes the names of the context variables; the strict-assignment check treats them as pre-bound globals (the key set is fixed at compile time).
- Each `Run` is independent: the context is deep-copied on injection, and evaluation panics are recovered into `error`.
- A key missing from the `Run` context errors on first read; extra keys become additional globals.
- `Evaluate(source, context)` keeps its signature as the compile + run convenience.

## Usage Example

```
// Rule: discount VIP users' large orders.
// Objects are lambda factories returning Map literals; methods use this / implicit fields.

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

// Build a lookup map id -> user (a plain loop works in v2).
var userMap = {}
for u in users {
    userMap[u.id] = u
}

var order = Order(500, "u2")
var user = userMap[order.userId]
var final = user.level == "vip" && order.amount > 100
        ? order.discounted(0.8)
        : order.amount
```

## Summary

Goval is a concise expression language for rule engines, built on three ideas:

- **Objects as lambda factories + Map literals** — no `struct`, no method-declaration syntax. A "type" is a function returning a Map; methods are lambdas that read the receiver through `this` (explicit or implicit).
- **Mutable locals, in-place containers, strict assignment** — variables reassign freely (loop accumulation just works); fields and elements are written in place; bare assignments require a prior binding so typos fail at compile time; host data stays isolated via deep copy.
- **Expression-oriented control flow** — `if`/`else`, `for`-`in`, `break`/`continue`. No `switch`, no `return`, no three-part `for`. Lambdas and expression blocks return their trailing expression.

### Key Features

- **Type Inference**: `var` declarations with mandatory initialization; no type annotations.
- **Lambda Factories**: custom types as functions returning Map literals, with `this` bound at method-call sites and implicit-`this` field rewriting in method bodies.
- **Reassignment & Strict Assignment**: identifier/field/subscript assignment targets; each variable bound once per scope but updatable; typos rejected at compile time.
- **Reference Capture**: closures share bindings with their defining scope; per-iteration loop capture; lambda identity equality.
- **`capture` Pragma**: compile-time opt-out of implicit-`this` rewriting for pure closure semantics.
- **Container Literals**: List `[...]` and Map `{...}` (string keys, with bare-identifier shorthand).
- **Control Flow**: `if`/`else`, `for`-`in` over List/Map/string, `break`/`continue`.
- **Expression Blocks**: `{ stmts; expr }` — trailing expression is the block's value.
- **Rich Operators**: arithmetic, comparison, logical (short-circuit), bitwise, shift, `in`, ternary — with C-like precedence.
- **String Methods**: `upper/lower/trim`, `contains/startsWith/endsWith`, `indexOf/lastIndexOf` (rune-based), `substring`, `replace/repeat/reverse`, `split/join` — chained in postfix form `s.method()`.
- **Built-in Functions**: `reduce`, `map`, `filter`, `find`, `append`, `put`, `removeAt`, `len`, `range`.
