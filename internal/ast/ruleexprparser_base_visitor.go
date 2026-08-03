// Code generated from RuleExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // RuleExprParser
import "github.com/antlr4-go/antlr/v4"

type BaseRuleExprParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRuleExprParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBlockStatements(ctx *BlockStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitVarVariableDeclaratorList(ctx *VarVariableDeclaratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitVarVariableDeclarator(ctx *VarVariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitConditionalOrExpression(ctx *ConditionalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitConditionalAndExpression(ctx *ConditionalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitLambdaExpression(ctx *LambdaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitLambdaParameters(ctx *LambdaParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitLambdaBody(ctx *LambdaBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitListLiteral(ctx *ListLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitMapEntryList(ctx *MapEntryListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitMapEntry(ctx *MapEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitExpressionBlock(ctx *ExpressionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}
