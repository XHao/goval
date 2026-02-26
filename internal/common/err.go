package common

import "fmt"

// SyntaxError represents a syntax error encountered during parsing.
type SyntaxError struct {
	Line    int
	Column  int
	Message string
}

// NewSyntaxError creates a new syntax error
func NewSyntaxError(line, column int, message string) *SyntaxError {
	return &SyntaxError{
		Line:    line,
		Column:  column,
		Message: message,
	}
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %d, column %d: %s", e.Line, e.Column, e.Message)
}

// SemanticError represents a semantic error encountered during semantic analysis.
type SemanticError struct {
	Msg    string
	Line   int // line number (1-based)
	Column int // column offset (0-based)
}

func (e *SemanticError) Error() string {
	return fmt.Sprintf("semantic error at line %d, column %d: %s", e.Line, e.Column, e.Msg)
}
