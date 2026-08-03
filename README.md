# Goval

[![CI](https://github.com/XHao/goval/workflows/CI/badge.svg)](https://github.com/XHao/goval/actions/workflows/ci.yml)
[![CodeQL](https://github.com/XHao/goval/workflows/CodeQL/badge.svg)](https://github.com/XHao/goval/actions/workflows/codeql.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/XHao/goval)](https://goreportcard.com/report/github.com/XHao/goval)
[![codecov](https://codecov.io/gh/XHao/goval/branch/main/graph/badge.svg)](https://codecov.io/gh/XHao/goval)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight expression language designed for embedding in Go applications, with a focus on rule engine scenarios. Compiles to closure trees for fast repeated evaluation.

## Features

- 🚀 **Closure-Tree Evaluator**: Source compiles to `func(*Env) Value` once, evaluates many times — ideal for rule engines that apply the same rule to large datasets
- 🔒 **Fully Immutable**: Single-assignment variables, immutable objects and containers — rules are pure functions, safe for concurrent evaluation
- 🏗️ **Objects via Lambda Factories**: No `struct` keyword — objects are Maps returned by lambda factories, with methods as closures capturing construction parameters
- 🔍 **Lambda & Closures**: First-class lambda expressions with lexical closure capture
- 📦 **Containers**: List and Map literals (`[1, 2, 3]`, `{"key": value}`), with immutable operations via builtins (`append`, `put`, `reduce`, `map`, `filter`, `find`)
- ⚡ **Minimal Syntax**: `if/else`, `for-in`, `var` — only the keywords a rule engine needs
- 🔧 **Go Integration**: `Evaluate(source, context)` API with automatic Go ↔ goval value conversion

## Quick Start

```bash
go get github.com/XHao/goval
```

```go
package main

import (
    "fmt"
    "github.com/XHao/goval"
)

func main() {
    // Basic arithmetic
    result, _ := goval.Evaluate("1 + 2 * 3", nil)
    fmt.Println(result) // 7

    // Variables and ternary
    result, _ = goval.Evaluate("x > 100 ? \"high\" : \"low\"",
        map[string]interface{}{"x": int64(150)})
    fmt.Println(result) // high

    // Objects via lambda factory
    result, _ = goval.Evaluate(`
        var Person = (name, level) -> {
            name: name,
            level: level,
            discount: (rate) -> rate
        };
        var p = Person("alice", "vip");
        p.level == "vip" ? p.discount(0.8) : 1.0
    `, nil)
    fmt.Println(result) // 0.8

    // Built-in functions
    result, _ = goval.Evaluate("reduce([1, 2, 3, 4], 0, (acc, x) -> acc + x)", nil)
    fmt.Println(result) // 10
}
```

## Language Overview

### Variables (single assignment)

```goval
var x = 10        // bind once, cannot rebind
var s = "hello"   // type inferred from initializer
```

### Operators

Arithmetic (`+ - * / %`), comparison (`< > <= >= == != in`), logical (`&& || !`), bitwise (`& | ^ ~ << >>`), ternary (`? :`).

### Lambda & Objects

```goval
// Lambda with closure capture
var adder = (base) -> (x) -> base + x
adder(10)(5)  // 15

// Object = lambda factory + Map literal
var Order = (amount, userId) -> {
    amount: amount,
    userId: userId,
    discounted: (rate) -> amount * rate
}
var o = Order(500, "u1")
o.discounted(0.8)  // 400
```

### Control Flow

```goval
// if / else
if (x > 0) { ... } else { ... }

// for-in (traversal only; no rebinding of outer variables)
for item in lst { ... }
for k, v in map { ... }
```

### Built-in Functions

| Function | Description |
|----------|-------------|
| `reduce(lst, init, fn)` | Aggregate elements |
| `map(lst, fn)` | Transform elements |
| `filter(lst, fn)` | Keep matching elements |
| `find(lst, fn)` | First matching element, or null |
| `append(lst, item)` | New list with item added |
| `put(map, key, value)` | New map with entry added |
| `removeAt(lst, index)` | New list with index removed |
| `len(v)` | Length of list/string/map |
| `range(start, end)` | Integer list `[start, end)` |

All container operations return new values — originals are never mutated.

## Documentation

- [Syntax Documentation](Goval_Syntax_Documentation.md)
- [Design Spec](docs/superpowers/specs/2026-08-03-syntax-simplification-design.md)

## Architecture

```
Source → [ANTLR4 Parser] → Parse Tree → [Compiler] → Closure Tree (func(*Env) Value)
                                                              ↓
                                                    [Evaluate] → Value
```

- **Parser**: ANTLR4-generated lexer/parser (`grammar/`, `internal/ast/`)
- **Syntax Checker**: `internal/syntax` — parse + semantic validation (single-assignment, break/continue scope)
- **Evaluator**: `internal/eval` — closure-tree compiler, immutable Env, built-in functions
- **Public API**: `pkg/goval` — `Evaluate(source, context)` with Go value conversion

## License

MIT
