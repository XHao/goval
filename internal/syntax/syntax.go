package syntax

import (
	"sync"

	"github.com/XHao/goval/internal/ast"
	"github.com/XHao/goval/internal/common"

	"github.com/antlr4-go/antlr/v4"
)

// SyntaxChecker provides syntax validation functionality.
// It is safe for concurrent use by multiple goroutines.
type SyntaxChecker struct {
	mu         sync.Mutex
	lastErrors []*common.SyntaxError
}

// NewSyntaxChecker creates a new SyntaxChecker.
func NewSyntaxChecker() *SyntaxChecker {
	return &SyntaxChecker{}
}

// CheckString is a convenience wrapper that validates the syntax of a source
// string and returns the parse tree.
func (c *SyntaxChecker) CheckString(source string) (ast.IProgramContext, error) {
	return c.Check(antlr.NewInputStream(source))
}

// Check validates the syntax of the input and returns the parse tree.
func (c *SyntaxChecker) Check(input antlr.CharStream) (ast.IProgramContext, error) {
	el := &SyntaxErrorListener{}

	lexer := ast.NewRuleExprLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ast.NewRuleExprParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(el)

	tree := parser.Program()

	c.mu.Lock()
	c.lastErrors = el.Errors
	c.mu.Unlock()

	if tree == nil {
		return nil, &common.SyntaxError{Line: 0, Column: 0, Message: "failed to parse input: got nil parse tree"}
	}

	if el.HasErrors() {
		return nil, el.GetFirstError()
	}

	return tree, nil
}

// CheckWithAllErrors validates syntax and returns all errors found.
func (c *SyntaxChecker) CheckWithAllErrors(input antlr.CharStream) (ast.IProgramContext, []*common.SyntaxError) {
	el := &SyntaxErrorListener{}

	lexer := ast.NewRuleExprLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ast.NewRuleExprParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(el)

	tree := parser.Program()

	c.mu.Lock()
	c.lastErrors = el.Errors
	c.mu.Unlock()

	return tree, el.GetAllErrors()
}

// IsValid quickly checks if input has valid syntax without returning parse tree.
func (c *SyntaxChecker) IsValid(input antlr.CharStream) bool {
	_, err := c.Check(input)
	return err == nil
}

// GetErrorCount returns the number of syntax errors from the last check.
func (c *SyntaxChecker) GetErrorCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.lastErrors)
}

// Reset clears errors recorded from the last check.
func (c *SyntaxChecker) Reset() {
	c.mu.Lock()
	c.lastErrors = nil
	c.mu.Unlock()
}

// SyntaxErrorListener collects ANTLR syntax errors from both the lexer and parser.
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

// HasErrors returns true if any syntax errors were encountered.
func (l *SyntaxErrorListener) HasErrors() bool {
	return len(l.Errors) > 0
}

// GetFirstError returns the first syntax error as an error, or nil if none.
func (l *SyntaxErrorListener) GetFirstError() error {
	if len(l.Errors) == 0 {
		return nil
	}
	return l.Errors[0]
}

// GetAllErrors returns all collected syntax errors.
func (l *SyntaxErrorListener) GetAllErrors() []*common.SyntaxError {
	return l.Errors
}
