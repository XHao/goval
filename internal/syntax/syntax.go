package syntax

import (
	"github.com/XHao/goval/internal/ast"
	"github.com/XHao/goval/internal/common"

	"github.com/antlr4-go/antlr/v4"
)

// SyntaxChecker provides syntax validation functionality
type SyntaxChecker struct {
	errorListener *SyntaxErrorListener
}

// NewSyntaxChecker creates a new syntax checker
func NewSyntaxChecker() *SyntaxChecker {
	return &SyntaxChecker{
		errorListener: &SyntaxErrorListener{},
	}
}

// Check validates the syntax of the input and returns the parse tree
func (c *SyntaxChecker) Check(input antlr.CharStream) (ast.IProgramContext, error) {
	lexer := ast.NewRuleExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ast.NewRuleExprParser(stream)

	// Reset error listener for new parsing
	c.errorListener.Errors = nil
	parser.RemoveErrorListeners()
	parser.AddErrorListener(c.errorListener)

	tree := parser.Program()
	if tree == nil {
		return nil, &common.SyntaxError{Line: 0, Column: 0, Message: "failed to parse input: got nil parse tree"}
	}

	if c.errorListener.HasErrors() {
		return nil, c.errorListener.GetFirstError()
	}

	return tree, nil
}

// CheckWithAllErrors validates syntax and returns all errors found
func (c *SyntaxChecker) CheckWithAllErrors(input antlr.CharStream) (ast.IProgramContext, []*common.SyntaxError) {
	lexer := ast.NewRuleExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ast.NewRuleExprParser(stream)

	// Reset error listener for new parsing
	c.errorListener.Errors = nil
	parser.RemoveErrorListeners()
	parser.AddErrorListener(c.errorListener)

	tree := parser.Program()
	return tree, c.errorListener.GetAllErrors()
}

// IsValid quickly checks if input has valid syntax without returning parse tree
func (c *SyntaxChecker) IsValid(input antlr.CharStream) bool {
	_, err := c.Check(input)
	return err == nil
}

// GetErrorCount returns the number of syntax errors from the last check
func (c *SyntaxChecker) GetErrorCount() int {
	return len(c.errorListener.Errors)
}

// Reset clears any previous errors
func (c *SyntaxChecker) Reset() {
	c.errorListener.Errors = nil
}

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []*common.SyntaxError
}

func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, &common.SyntaxError{
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

// HasErrors returns true if any syntax errors were encountered
func (l *SyntaxErrorListener) HasErrors() bool {
	return len(l.Errors) > 0
}

// GetFirstError returns the first syntax error, or nil if no errors
func (l *SyntaxErrorListener) GetFirstError() error {
	if len(l.Errors) == 0 {
		return nil
	}
	return l.Errors[0]
}

// GetAllErrors returns all syntax errors
func (l *SyntaxErrorListener) GetAllErrors() []*common.SyntaxError {
	return l.Errors
}
