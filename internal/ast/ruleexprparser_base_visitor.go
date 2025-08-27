// Code generated from RuleExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // RuleExprParser
import "github.com/antlr4-go/antlr/v4"

type BaseRuleExprParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRuleExprParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStructMemberList(ctx *StructMemberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStructField(ctx *StructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStructMethod(ctx *StructMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitMethodParameterList(ctx *MethodParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitMethodParameter(ctx *MethodParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitMethodBody(ctx *MethodBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitParamType(ctx *ParamTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitDims(ctx *DimsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitContainerType(ctx *ContainerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitVariableInitializerList(ctx *VariableInitializerListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitVariableDeclaratorList(ctx *VariableDeclaratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
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

func (v *BaseRuleExprParserVisitor) VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitExpressionBlock(ctx *ExpressionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBlockStatements(ctx *BlockStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStatementNoShortIf(ctx *StatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStatementExpression(ctx *StatementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitIfThenStatement(ctx *IfThenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitIfThenElseStatement(ctx *IfThenElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitSwitchBlock(ctx *SwitchBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitSwitchLabels(ctx *SwitchLabelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitSwitchLabel(ctx *SwitchLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBasicForStatement(ctx *BasicForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStatementExpressionList(ctx *StatementExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitEnhancedForStatement(ctx *EnhancedForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
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

func (v *BaseRuleExprParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
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

func (v *BaseRuleExprParserVisitor) VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitPreDecrementExpression(ctx *PreDecrementExpressionContext) interface{} {
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

func (v *BaseRuleExprParserVisitor) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
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

func (v *BaseRuleExprParserVisitor) VisitSetLiteral(ctx *SetLiteralContext) interface{} {
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

func (v *BaseRuleExprParserVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStructFieldInitializerList(ctx *StructFieldInitializerListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleExprParserVisitor) VisitStructFieldInitializer(ctx *StructFieldInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}
