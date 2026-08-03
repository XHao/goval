// Code generated from RuleExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // RuleExprParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RuleExprParser.
type RuleExprParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RuleExprParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RuleExprParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by RuleExprParser#blockStatements.
	VisitBlockStatements(ctx *BlockStatementsContext) interface{}

	// Visit a parse tree produced by RuleExprParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#localVariableDeclarationStatement.
	VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#localVariableDeclaration.
	VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{}

	// Visit a parse tree produced by RuleExprParser#varVariableDeclaratorList.
	VisitVarVariableDeclaratorList(ctx *VarVariableDeclaratorListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#varVariableDeclarator.
	VisitVarVariableDeclarator(ctx *VarVariableDeclaratorContext) interface{}

	// Visit a parse tree produced by RuleExprParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by RuleExprParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by RuleExprParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by RuleExprParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#conditionalOrExpression.
	VisitConditionalOrExpression(ctx *ConditionalOrExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#conditionalAndExpression.
	VisitConditionalAndExpression(ctx *ConditionalAndExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#inclusiveOrExpression.
	VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#exclusiveOrExpression.
	VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#unaryExpressionNotPlusMinus.
	VisitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) interface{}

	// Visit a parse tree produced by RuleExprParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by RuleExprParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#lambdaExpression.
	VisitLambdaExpression(ctx *LambdaExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#lambdaParameters.
	VisitLambdaParameters(ctx *LambdaParametersContext) interface{}

	// Visit a parse tree produced by RuleExprParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#lambdaBody.
	VisitLambdaBody(ctx *LambdaBodyContext) interface{}

	// Visit a parse tree produced by RuleExprParser#listLiteral.
	VisitListLiteral(ctx *ListLiteralContext) interface{}

	// Visit a parse tree produced by RuleExprParser#mapLiteral.
	VisitMapLiteral(ctx *MapLiteralContext) interface{}

	// Visit a parse tree produced by RuleExprParser#mapEntryList.
	VisitMapEntryList(ctx *MapEntryListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#mapEntry.
	VisitMapEntry(ctx *MapEntryContext) interface{}

	// Visit a parse tree produced by RuleExprParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#expressionBlock.
	VisitExpressionBlock(ctx *ExpressionBlockContext) interface{}
}
