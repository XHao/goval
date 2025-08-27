package pkg

import (
	"github.com/XHao/goval/internal/ast"
	"github.com/XHao/goval/internal/syntax"

	"github.com/antlr4-go/antlr/v4"
)

type RuleExprEngine struct {
	syntaxChecker *syntax.SyntaxChecker
	tree          ast.IProgramContext
}

// NewRuleExprEngine creates a new RuleExprEngine with the given input.
func NewRuleExprEngine(input antlr.CharStream) (*RuleExprEngine, error) {
	syntaxChecker := syntax.NewSyntaxChecker()

	engine := &RuleExprEngine{
		syntaxChecker: syntaxChecker,
	}

	if err := engine.checkSyntax(input); err != nil {
		return nil, err
	}
	return engine, nil
}

func (e *RuleExprEngine) checkSyntax(input antlr.CharStream) error {
	tree, err := e.syntaxChecker.Check(input)
	if err != nil {
		return err
	}
	e.tree = tree
	return nil
}

// GetParseTree returns the parsed syntax tree
func (e *RuleExprEngine) GetParseTree() ast.IProgramContext {
	return e.tree
}
