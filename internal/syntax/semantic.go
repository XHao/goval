package syntax

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/XHao/goval/internal/ast"
	"github.com/XHao/goval/internal/common"
)

// checkSemantics performs post-parse semantic validation:
//   - break/continue 只能在 for 体内
//   - lambda 块体 / expressionBlock 末尾必须是表达式
// 返回第一个发现的语义错误，或 nil。
func checkSemantics(tree ast.IProgramContext) error {
	v := &semanticVisitor{
		BaseRuleExprParserVisitor: &ast.BaseRuleExprParserVisitor{},
	}
	v.Visit(tree)
	if v.firstErr == nil {
		return nil
	}
	return v.firstErr
}

// semanticVisitor 走遍 AST 收集语义错误。
type semanticVisitor struct {
	*ast.BaseRuleExprParserVisitor
	loopDepth int
	firstErr  *common.SemanticError
}

// Visit overrides the base Visit to pass the semanticVisitor itself
// (not the embedded base) to Accept, so that overridden Visit* methods
// are dispatched correctly.
func (v *semanticVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *semanticVisitor) report(line, col int, msg string) {
	if v.firstErr == nil {
		v.firstErr = &common.SemanticError{Line: line, Column: col, Msg: msg}
	}
}

// VisitProgram traverses all top-level statements.
func (v *semanticVisitor) VisitProgram(ctx *ast.ProgramContext) interface{} {
	if v.firstErr != nil {
		return nil
	}
	for _, s := range ctx.AllStatement() {
		v.Visit(s)
		if v.firstErr != nil {
			return nil
		}
	}
	return nil
}

// VisitStatement dispatches to the specific statement sub-visitor.
func (v *semanticVisitor) VisitStatement(ctx *ast.StatementContext) interface{} {
	if v.firstErr != nil {
		return nil
	}
	switch {
	case ctx.ForStatement() != nil:
		v.Visit(ctx.ForStatement())
	case ctx.IfStatement() != nil:
		v.Visit(ctx.IfStatement())
	case ctx.Block() != nil:
		v.Visit(ctx.Block())
	case ctx.BreakStatement() != nil:
		v.Visit(ctx.BreakStatement())
	case ctx.ContinueStatement() != nil:
		v.Visit(ctx.ContinueStatement())
	}
	// ExpressionStatement and LocalVariableDeclarationStatement do not
	// contain break/continue at the statement level.
	return nil
}

// VisitForStatement increments loopDepth, visits the block body, then decrements.
func (v *semanticVisitor) VisitForStatement(ctx *ast.ForStatementContext) interface{} {
	if v.firstErr != nil {
		return nil
	}
	v.loopDepth++
	if block := ctx.Block(); block != nil {
		v.Visit(block)
	}
	v.loopDepth--
	return nil
}

// VisitBreakStatement reports a semantic error if break is outside a for loop.
func (v *semanticVisitor) VisitBreakStatement(ctx *ast.BreakStatementContext) interface{} {
	if v.loopDepth == 0 {
		tok := ctx.BREAK().GetSymbol()
		v.report(tok.GetLine(), tok.GetColumn(), "'break' is not allowed outside a for loop")
	}
	return nil
}

// VisitContinueStatement reports a semantic error if continue is outside a for loop.
func (v *semanticVisitor) VisitContinueStatement(ctx *ast.ContinueStatementContext) interface{} {
	if v.loopDepth == 0 {
		tok := ctx.CONTINUE().GetSymbol()
		v.report(tok.GetLine(), tok.GetColumn(), "'continue' is not allowed outside a for loop")
	}
	return nil
}

// VisitIfStatement traverses the if/else statement bodies.
func (v *semanticVisitor) VisitIfStatement(ctx *ast.IfStatementContext) interface{} {
	if v.firstErr != nil {
		return nil
	}
	for _, s := range ctx.AllStatement() {
		v.Visit(s)
		if v.firstErr != nil {
			return nil
		}
	}
	return nil
}

// VisitBlock traverses the block statements.
func (v *semanticVisitor) VisitBlock(ctx *ast.BlockContext) interface{} {
	if v.firstErr != nil {
		return nil
	}
	if bs := ctx.BlockStatements(); bs != nil {
		v.Visit(bs)
	}
	return nil
}

// VisitBlockStatements traverses all block statements.
func (v *semanticVisitor) VisitBlockStatements(ctx *ast.BlockStatementsContext) interface{} {
	if v.firstErr != nil {
		return nil
	}
	for _, bs := range ctx.AllBlockStatement() {
		v.Visit(bs)
		if v.firstErr != nil {
			return nil
		}
	}
	return nil
}

// VisitBlockStatement traverses the statement within a block statement.
func (v *semanticVisitor) VisitBlockStatement(ctx *ast.BlockStatementContext) interface{} {
	if v.firstErr != nil {
		return nil
	}
	if s := ctx.Statement(); s != nil {
		v.Visit(s)
	}
	return nil
}
