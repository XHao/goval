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
	Msg   string
	Start int // start token index or position
	End   int // end token index or position
}

func (e *SemanticError) Error() string {
	return fmt.Sprintf("semantic error [%d:%d]: %s", e.Start, e.End, e.Msg)
}
