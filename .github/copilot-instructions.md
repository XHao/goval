# Goval - AI Coding Agent Instructions

This document provides guidance for AI coding agents working on the Goval codebase.

## Project Overview

Goval is a lightweight, statically typed expression language designed to be embedded in Go applications. It features a custom syntax, parsed by ANTLR, and is designed for high performance and safety.

## Core Architecture

The core of Goval is its ANTLR4-based parser and the Go code that implements the language's semantics.

- **Grammar Definitions**: The language syntax is defined in `grammar/RuleExprLexer.g4` and `grammar/RuleExprParser.g4`. These files are the source of truth for the language's syntax.
- **Parser Generation**: The Go parser code is generated from the grammar files. Any changes to the `.g4` files require regenerating the parser.
- **Generated Code**: The generated parser, lexer, visitor, and listener are located in `internal/ast`. **Do not edit these files directly.**
- **Syntax Checking**: `internal/syntax/syntax.go` provides the `SyntaxChecker`, which uses the generated parser to validate Goval code.
- **Semantic Analysis**: The `internal/semantic` directory is responsible for semantic checks, such as type checking.
- **Runtime**: The `internal/runtime` directory will handle the execution of the parsed expressions.

## Development Workflow

### Regenerating the ANTLR Parser

When you modify the grammar files (`.g4`), you must regenerate the parser. This is a critical step.

A shell script is provided to automate this process:

```sh
# From the 'grammar' directory
./generate.sh
```

This script uses `antlr4` to generate the necessary Go files in the `internal/ast` directory. Make sure you have ANTLR4 and Java installed and configured.

### Testing

Tests are located in files ending with `_test.go` and are written using the standard `testing` package and `github.com/stretchr/testify/assert`.

To run all tests:

```sh
go test ./...
```

When adding new features or fixing bugs, please add corresponding tests. For example, when modifying the syntax, add test cases to `internal/syntax/syntax_test.go` to cover both valid and invalid syntax.

## Coding Conventions

- **Error Handling**: Errors are handled explicitly. Syntax errors are collected in `SyntaxErrorListener` in `internal/syntax/syntax.go`.
- **Dependencies**: The project uses Go modules for dependency management.
- **Formatting**: Code should be formatted with `gofmt`.
