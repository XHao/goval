// Code generated from RuleExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // RuleExprParser
import "github.com/antlr4-go/antlr/v4"

// BaseRuleExprParserListener is a complete listener for a parse tree produced by RuleExprParser.
type BaseRuleExprParserListener struct{}

var _ RuleExprParserListener = &BaseRuleExprParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRuleExprParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRuleExprParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRuleExprParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRuleExprParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseRuleExprParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseRuleExprParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseRuleExprParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseRuleExprParserListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatements is called when production blockStatements is entered.
func (s *BaseRuleExprParserListener) EnterBlockStatements(ctx *BlockStatementsContext) {}

// ExitBlockStatements is called when production blockStatements is exited.
func (s *BaseRuleExprParserListener) ExitBlockStatements(ctx *BlockStatementsContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseRuleExprParserListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseRuleExprParserListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is entered.
func (s *BaseRuleExprParserListener) EnterLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// ExitLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is exited.
func (s *BaseRuleExprParserListener) ExitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseRuleExprParserListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseRuleExprParserListener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
}

// EnterVarVariableDeclaratorList is called when production varVariableDeclaratorList is entered.
func (s *BaseRuleExprParserListener) EnterVarVariableDeclaratorList(ctx *VarVariableDeclaratorListContext) {
}

// ExitVarVariableDeclaratorList is called when production varVariableDeclaratorList is exited.
func (s *BaseRuleExprParserListener) ExitVarVariableDeclaratorList(ctx *VarVariableDeclaratorListContext) {
}

// EnterVarVariableDeclarator is called when production varVariableDeclarator is entered.
func (s *BaseRuleExprParserListener) EnterVarVariableDeclarator(ctx *VarVariableDeclaratorContext) {}

// ExitVarVariableDeclarator is called when production varVariableDeclarator is exited.
func (s *BaseRuleExprParserListener) ExitVarVariableDeclarator(ctx *VarVariableDeclaratorContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseRuleExprParserListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseRuleExprParserListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseRuleExprParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRuleExprParserListener) ExitStatement(ctx *StatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseRuleExprParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseRuleExprParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseRuleExprParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseRuleExprParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseRuleExprParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseRuleExprParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseRuleExprParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseRuleExprParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseRuleExprParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseRuleExprParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseRuleExprParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseRuleExprParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRuleExprParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRuleExprParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseRuleExprParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseRuleExprParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseRuleExprParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseRuleExprParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseRuleExprParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseRuleExprParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterConditionalOrExpression is called when production conditionalOrExpression is entered.
func (s *BaseRuleExprParserListener) EnterConditionalOrExpression(ctx *ConditionalOrExpressionContext) {
}

// ExitConditionalOrExpression is called when production conditionalOrExpression is exited.
func (s *BaseRuleExprParserListener) ExitConditionalOrExpression(ctx *ConditionalOrExpressionContext) {
}

// EnterConditionalAndExpression is called when production conditionalAndExpression is entered.
func (s *BaseRuleExprParserListener) EnterConditionalAndExpression(ctx *ConditionalAndExpressionContext) {
}

// ExitConditionalAndExpression is called when production conditionalAndExpression is exited.
func (s *BaseRuleExprParserListener) ExitConditionalAndExpression(ctx *ConditionalAndExpressionContext) {
}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BaseRuleExprParserListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BaseRuleExprParserListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BaseRuleExprParserListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BaseRuleExprParserListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseRuleExprParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseRuleExprParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseRuleExprParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseRuleExprParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseRuleExprParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseRuleExprParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseRuleExprParserListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseRuleExprParserListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseRuleExprParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseRuleExprParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseRuleExprParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseRuleExprParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseRuleExprParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseRuleExprParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryExpressionNotPlusMinus is called when production unaryExpressionNotPlusMinus is entered.
func (s *BaseRuleExprParserListener) EnterUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) {
}

// ExitUnaryExpressionNotPlusMinus is called when production unaryExpressionNotPlusMinus is exited.
func (s *BaseRuleExprParserListener) ExitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) {
}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseRuleExprParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseRuleExprParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseRuleExprParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseRuleExprParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseRuleExprParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseRuleExprParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseRuleExprParserListener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseRuleExprParserListener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *BaseRuleExprParserListener) EnterLambdaParameters(ctx *LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *BaseRuleExprParserListener) ExitLambdaParameters(ctx *LambdaParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseRuleExprParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseRuleExprParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterLambdaBody is called when production lambdaBody is entered.
func (s *BaseRuleExprParserListener) EnterLambdaBody(ctx *LambdaBodyContext) {}

// ExitLambdaBody is called when production lambdaBody is exited.
func (s *BaseRuleExprParserListener) ExitLambdaBody(ctx *LambdaBodyContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseRuleExprParserListener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseRuleExprParserListener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseRuleExprParserListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseRuleExprParserListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterMapEntryList is called when production mapEntryList is entered.
func (s *BaseRuleExprParserListener) EnterMapEntryList(ctx *MapEntryListContext) {}

// ExitMapEntryList is called when production mapEntryList is exited.
func (s *BaseRuleExprParserListener) ExitMapEntryList(ctx *MapEntryListContext) {}

// EnterMapEntry is called when production mapEntry is entered.
func (s *BaseRuleExprParserListener) EnterMapEntry(ctx *MapEntryContext) {}

// ExitMapEntry is called when production mapEntry is exited.
func (s *BaseRuleExprParserListener) ExitMapEntry(ctx *MapEntryContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseRuleExprParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseRuleExprParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpressionBlock is called when production expressionBlock is entered.
func (s *BaseRuleExprParserListener) EnterExpressionBlock(ctx *ExpressionBlockContext) {}

// ExitExpressionBlock is called when production expressionBlock is exited.
func (s *BaseRuleExprParserListener) ExitExpressionBlock(ctx *ExpressionBlockContext) {}
