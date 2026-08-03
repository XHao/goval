package eval

import "fmt"

// CompileError 是编译期错误（单赋值检查、语法问题等）。
type CompileError struct {
	Line   int
	Column int
	Msg    string
}

func (e *CompileError) Error() string {
	return fmt.Sprintf("compile error at line %d, column %d: %s", e.Line, e.Column, e.Msg)
}

// EvalError 是求值期错误（类型不匹配、未定义变量、索引越界等）。
type EvalError struct {
	Line   int
	Column int
	Msg    string
}

func (e *EvalError) Error() string {
	return fmt.Sprintf("eval error at line %d, column %d: %s", e.Line, e.Column, e.Msg)
}

func evalErrorf(line, col int, format string, args ...interface{}) *EvalError {
	return &EvalError{Line: line, Column: col, Msg: fmt.Sprintf(format, args...)}
}
