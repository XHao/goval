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

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *BaseRuleExprParserListener) EnterStructDeclaration(ctx *StructDeclarationContext) {}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *BaseRuleExprParserListener) ExitStructDeclaration(ctx *StructDeclarationContext) {}

// EnterStructMemberList is called when production structMemberList is entered.
func (s *BaseRuleExprParserListener) EnterStructMemberList(ctx *StructMemberListContext) {}

// ExitStructMemberList is called when production structMemberList is exited.
func (s *BaseRuleExprParserListener) ExitStructMemberList(ctx *StructMemberListContext) {}

// EnterStructMember is called when production structMember is entered.
func (s *BaseRuleExprParserListener) EnterStructMember(ctx *StructMemberContext) {}

// ExitStructMember is called when production structMember is exited.
func (s *BaseRuleExprParserListener) ExitStructMember(ctx *StructMemberContext) {}

// EnterStructField is called when production structField is entered.
func (s *BaseRuleExprParserListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BaseRuleExprParserListener) ExitStructField(ctx *StructFieldContext) {}

// EnterStructMethod is called when production structMethod is entered.
func (s *BaseRuleExprParserListener) EnterStructMethod(ctx *StructMethodContext) {}

// ExitStructMethod is called when production structMethod is exited.
func (s *BaseRuleExprParserListener) ExitStructMethod(ctx *StructMethodContext) {}

// EnterMethodParameterList is called when production methodParameterList is entered.
func (s *BaseRuleExprParserListener) EnterMethodParameterList(ctx *MethodParameterListContext) {}

// ExitMethodParameterList is called when production methodParameterList is exited.
func (s *BaseRuleExprParserListener) ExitMethodParameterList(ctx *MethodParameterListContext) {}

// EnterMethodParameter is called when production methodParameter is entered.
func (s *BaseRuleExprParserListener) EnterMethodParameter(ctx *MethodParameterContext) {}

// ExitMethodParameter is called when production methodParameter is exited.
func (s *BaseRuleExprParserListener) ExitMethodParameter(ctx *MethodParameterContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseRuleExprParserListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseRuleExprParserListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseRuleExprParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseRuleExprParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseRuleExprParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseRuleExprParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterType is called when production type is entered.
func (s *BaseRuleExprParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseRuleExprParserListener) ExitType(ctx *TypeContext) {}

// EnterParamType is called when production paramType is entered.
func (s *BaseRuleExprParserListener) EnterParamType(ctx *ParamTypeContext) {}

// ExitParamType is called when production paramType is exited.
func (s *BaseRuleExprParserListener) ExitParamType(ctx *ParamTypeContext) {}

// EnterDims is called when production dims is entered.
func (s *BaseRuleExprParserListener) EnterDims(ctx *DimsContext) {}

// ExitDims is called when production dims is exited.
func (s *BaseRuleExprParserListener) ExitDims(ctx *DimsContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseRuleExprParserListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseRuleExprParserListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterFunctionParameter is called when production functionParameter is entered.
func (s *BaseRuleExprParserListener) EnterFunctionParameter(ctx *FunctionParameterContext) {}

// ExitFunctionParameter is called when production functionParameter is exited.
func (s *BaseRuleExprParserListener) ExitFunctionParameter(ctx *FunctionParameterContext) {}

// EnterContainerType is called when production containerType is entered.
func (s *BaseRuleExprParserListener) EnterContainerType(ctx *ContainerTypeContext) {}

// ExitContainerType is called when production containerType is exited.
func (s *BaseRuleExprParserListener) ExitContainerType(ctx *ContainerTypeContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseRuleExprParserListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseRuleExprParserListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterVariableInitializerList is called when production variableInitializerList is entered.
func (s *BaseRuleExprParserListener) EnterVariableInitializerList(ctx *VariableInitializerListContext) {
}

// ExitVariableInitializerList is called when production variableInitializerList is exited.
func (s *BaseRuleExprParserListener) ExitVariableInitializerList(ctx *VariableInitializerListContext) {
}

// EnterVariableDeclaratorList is called when production variableDeclaratorList is entered.
func (s *BaseRuleExprParserListener) EnterVariableDeclaratorList(ctx *VariableDeclaratorListContext) {
}

// ExitVariableDeclaratorList is called when production variableDeclaratorList is exited.
func (s *BaseRuleExprParserListener) ExitVariableDeclaratorList(ctx *VariableDeclaratorListContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseRuleExprParserListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseRuleExprParserListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseRuleExprParserListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseRuleExprParserListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseRuleExprParserListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseRuleExprParserListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

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

// EnterLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is entered.
func (s *BaseRuleExprParserListener) EnterLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// ExitLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is exited.
func (s *BaseRuleExprParserListener) ExitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// EnterBlock is called when production block is entered.
func (s *BaseRuleExprParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseRuleExprParserListener) ExitBlock(ctx *BlockContext) {}

// EnterExpressionBlock is called when production expressionBlock is entered.
func (s *BaseRuleExprParserListener) EnterExpressionBlock(ctx *ExpressionBlockContext) {}

// ExitExpressionBlock is called when production expressionBlock is exited.
func (s *BaseRuleExprParserListener) ExitExpressionBlock(ctx *ExpressionBlockContext) {}

// EnterBlockStatements is called when production blockStatements is entered.
func (s *BaseRuleExprParserListener) EnterBlockStatements(ctx *BlockStatementsContext) {}

// ExitBlockStatements is called when production blockStatements is exited.
func (s *BaseRuleExprParserListener) ExitBlockStatements(ctx *BlockStatementsContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseRuleExprParserListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseRuleExprParserListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseRuleExprParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRuleExprParserListener) ExitStatement(ctx *StatementContext) {}

// EnterStatementNoShortIf is called when production statementNoShortIf is entered.
func (s *BaseRuleExprParserListener) EnterStatementNoShortIf(ctx *StatementNoShortIfContext) {}

// ExitStatementNoShortIf is called when production statementNoShortIf is exited.
func (s *BaseRuleExprParserListener) ExitStatementNoShortIf(ctx *StatementNoShortIfContext) {}

// EnterStatementWithoutTrailingSubstatement is called when production statementWithoutTrailingSubstatement is entered.
func (s *BaseRuleExprParserListener) EnterStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) {
}

// ExitStatementWithoutTrailingSubstatement is called when production statementWithoutTrailingSubstatement is exited.
func (s *BaseRuleExprParserListener) ExitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) {
}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseRuleExprParserListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseRuleExprParserListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseRuleExprParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseRuleExprParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterStatementExpression is called when production statementExpression is entered.
func (s *BaseRuleExprParserListener) EnterStatementExpression(ctx *StatementExpressionContext) {}

// ExitStatementExpression is called when production statementExpression is exited.
func (s *BaseRuleExprParserListener) ExitStatementExpression(ctx *StatementExpressionContext) {}

// EnterIfThenStatement is called when production ifThenStatement is entered.
func (s *BaseRuleExprParserListener) EnterIfThenStatement(ctx *IfThenStatementContext) {}

// ExitIfThenStatement is called when production ifThenStatement is exited.
func (s *BaseRuleExprParserListener) ExitIfThenStatement(ctx *IfThenStatementContext) {}

// EnterIfThenElseStatement is called when production ifThenElseStatement is entered.
func (s *BaseRuleExprParserListener) EnterIfThenElseStatement(ctx *IfThenElseStatementContext) {}

// ExitIfThenElseStatement is called when production ifThenElseStatement is exited.
func (s *BaseRuleExprParserListener) ExitIfThenElseStatement(ctx *IfThenElseStatementContext) {}

// EnterIfThenElseStatementNoShortIf is called when production ifThenElseStatementNoShortIf is entered.
func (s *BaseRuleExprParserListener) EnterIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) {
}

// ExitIfThenElseStatementNoShortIf is called when production ifThenElseStatementNoShortIf is exited.
func (s *BaseRuleExprParserListener) ExitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) {
}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseRuleExprParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseRuleExprParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchBlock is called when production switchBlock is entered.
func (s *BaseRuleExprParserListener) EnterSwitchBlock(ctx *SwitchBlockContext) {}

// ExitSwitchBlock is called when production switchBlock is exited.
func (s *BaseRuleExprParserListener) ExitSwitchBlock(ctx *SwitchBlockContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BaseRuleExprParserListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BaseRuleExprParserListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// EnterSwitchLabels is called when production switchLabels is entered.
func (s *BaseRuleExprParserListener) EnterSwitchLabels(ctx *SwitchLabelsContext) {}

// ExitSwitchLabels is called when production switchLabels is exited.
func (s *BaseRuleExprParserListener) ExitSwitchLabels(ctx *SwitchLabelsContext) {}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseRuleExprParserListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseRuleExprParserListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseRuleExprParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseRuleExprParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForStatementNoShortIf is called when production forStatementNoShortIf is entered.
func (s *BaseRuleExprParserListener) EnterForStatementNoShortIf(ctx *ForStatementNoShortIfContext) {}

// ExitForStatementNoShortIf is called when production forStatementNoShortIf is exited.
func (s *BaseRuleExprParserListener) ExitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) {}

// EnterBasicForStatement is called when production basicForStatement is entered.
func (s *BaseRuleExprParserListener) EnterBasicForStatement(ctx *BasicForStatementContext) {}

// ExitBasicForStatement is called when production basicForStatement is exited.
func (s *BaseRuleExprParserListener) ExitBasicForStatement(ctx *BasicForStatementContext) {}

// EnterBasicForStatementNoShortIf is called when production basicForStatementNoShortIf is entered.
func (s *BaseRuleExprParserListener) EnterBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) {
}

// ExitBasicForStatementNoShortIf is called when production basicForStatementNoShortIf is exited.
func (s *BaseRuleExprParserListener) ExitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) {
}

// EnterForInit is called when production forInit is entered.
func (s *BaseRuleExprParserListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseRuleExprParserListener) ExitForInit(ctx *ForInitContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseRuleExprParserListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseRuleExprParserListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterStatementExpressionList is called when production statementExpressionList is entered.
func (s *BaseRuleExprParserListener) EnterStatementExpressionList(ctx *StatementExpressionListContext) {
}

// ExitStatementExpressionList is called when production statementExpressionList is exited.
func (s *BaseRuleExprParserListener) ExitStatementExpressionList(ctx *StatementExpressionListContext) {
}

// EnterEnhancedForStatement is called when production enhancedForStatement is entered.
func (s *BaseRuleExprParserListener) EnterEnhancedForStatement(ctx *EnhancedForStatementContext) {}

// ExitEnhancedForStatement is called when production enhancedForStatement is exited.
func (s *BaseRuleExprParserListener) ExitEnhancedForStatement(ctx *EnhancedForStatementContext) {}

// EnterEnhancedForStatementNoShortIf is called when production enhancedForStatementNoShortIf is entered.
func (s *BaseRuleExprParserListener) EnterEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) {
}

// ExitEnhancedForStatementNoShortIf is called when production enhancedForStatementNoShortIf is exited.
func (s *BaseRuleExprParserListener) ExitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) {
}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseRuleExprParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseRuleExprParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseRuleExprParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseRuleExprParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseRuleExprParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseRuleExprParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

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

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseRuleExprParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseRuleExprParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

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

// EnterPreIncrementExpression is called when production preIncrementExpression is entered.
func (s *BaseRuleExprParserListener) EnterPreIncrementExpression(ctx *PreIncrementExpressionContext) {
}

// ExitPreIncrementExpression is called when production preIncrementExpression is exited.
func (s *BaseRuleExprParserListener) ExitPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// EnterPreDecrementExpression is called when production preDecrementExpression is entered.
func (s *BaseRuleExprParserListener) EnterPreDecrementExpression(ctx *PreDecrementExpressionContext) {
}

// ExitPreDecrementExpression is called when production preDecrementExpression is exited.
func (s *BaseRuleExprParserListener) ExitPreDecrementExpression(ctx *PreDecrementExpressionContext) {}

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

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseRuleExprParserListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseRuleExprParserListener) ExitFormalParameter(ctx *FormalParameterContext) {}

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

// EnterSetLiteral is called when production setLiteral is entered.
func (s *BaseRuleExprParserListener) EnterSetLiteral(ctx *SetLiteralContext) {}

// ExitSetLiteral is called when production setLiteral is exited.
func (s *BaseRuleExprParserListener) ExitSetLiteral(ctx *SetLiteralContext) {}

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

// EnterStructLiteral is called when production structLiteral is entered.
func (s *BaseRuleExprParserListener) EnterStructLiteral(ctx *StructLiteralContext) {}

// ExitStructLiteral is called when production structLiteral is exited.
func (s *BaseRuleExprParserListener) ExitStructLiteral(ctx *StructLiteralContext) {}

// EnterStructFieldInitializerList is called when production structFieldInitializerList is entered.
func (s *BaseRuleExprParserListener) EnterStructFieldInitializerList(ctx *StructFieldInitializerListContext) {
}

// ExitStructFieldInitializerList is called when production structFieldInitializerList is exited.
func (s *BaseRuleExprParserListener) ExitStructFieldInitializerList(ctx *StructFieldInitializerListContext) {
}

// EnterStructFieldInitializer is called when production structFieldInitializer is entered.
func (s *BaseRuleExprParserListener) EnterStructFieldInitializer(ctx *StructFieldInitializerContext) {
}

// ExitStructFieldInitializer is called when production structFieldInitializer is exited.
func (s *BaseRuleExprParserListener) ExitStructFieldInitializer(ctx *StructFieldInitializerContext) {}
