// Code generated from RuleExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // RuleExprParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RuleExprParser.
type RuleExprParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RuleExprParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RuleExprParser#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by RuleExprParser#structMemberList.
	VisitStructMemberList(ctx *StructMemberListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#structMember.
	VisitStructMember(ctx *StructMemberContext) interface{}

	// Visit a parse tree produced by RuleExprParser#structField.
	VisitStructField(ctx *StructFieldContext) interface{}

	// Visit a parse tree produced by RuleExprParser#structMethod.
	VisitStructMethod(ctx *StructMethodContext) interface{}

	// Visit a parse tree produced by RuleExprParser#methodParameterList.
	VisitMethodParameterList(ctx *MethodParameterListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#methodParameter.
	VisitMethodParameter(ctx *MethodParameterContext) interface{}

	// Visit a parse tree produced by RuleExprParser#methodBody.
	VisitMethodBody(ctx *MethodBodyContext) interface{}

	// Visit a parse tree produced by RuleExprParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by RuleExprParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by RuleExprParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by RuleExprParser#paramType.
	VisitParamType(ctx *ParamTypeContext) interface{}

	// Visit a parse tree produced by RuleExprParser#dims.
	VisitDims(ctx *DimsContext) interface{}

	// Visit a parse tree produced by RuleExprParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by RuleExprParser#functionParameter.
	VisitFunctionParameter(ctx *FunctionParameterContext) interface{}

	// Visit a parse tree produced by RuleExprParser#containerType.
	VisitContainerType(ctx *ContainerTypeContext) interface{}

	// Visit a parse tree produced by RuleExprParser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by RuleExprParser#variableInitializerList.
	VisitVariableInitializerList(ctx *VariableInitializerListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#variableDeclaratorList.
	VisitVariableDeclaratorList(ctx *VariableDeclaratorListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by RuleExprParser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{}

	// Visit a parse tree produced by RuleExprParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by RuleExprParser#localVariableDeclaration.
	VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{}

	// Visit a parse tree produced by RuleExprParser#varVariableDeclaratorList.
	VisitVarVariableDeclaratorList(ctx *VarVariableDeclaratorListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#varVariableDeclarator.
	VisitVarVariableDeclarator(ctx *VarVariableDeclaratorContext) interface{}

	// Visit a parse tree produced by RuleExprParser#localVariableDeclarationStatement.
	VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by RuleExprParser#expressionBlock.
	VisitExpressionBlock(ctx *ExpressionBlockContext) interface{}

	// Visit a parse tree produced by RuleExprParser#blockStatements.
	VisitBlockStatements(ctx *BlockStatementsContext) interface{}

	// Visit a parse tree produced by RuleExprParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#statementNoShortIf.
	VisitStatementNoShortIf(ctx *StatementNoShortIfContext) interface{}

	// Visit a parse tree produced by RuleExprParser#statementWithoutTrailingSubstatement.
	VisitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by RuleExprParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#statementExpression.
	VisitStatementExpression(ctx *StatementExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#ifThenStatement.
	VisitIfThenStatement(ctx *IfThenStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#ifThenElseStatement.
	VisitIfThenElseStatement(ctx *IfThenElseStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#ifThenElseStatementNoShortIf.
	VisitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by RuleExprParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#switchBlock.
	VisitSwitchBlock(ctx *SwitchBlockContext) interface{}

	// Visit a parse tree produced by RuleExprParser#switchBlockStatementGroup.
	VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{}

	// Visit a parse tree produced by RuleExprParser#switchLabels.
	VisitSwitchLabels(ctx *SwitchLabelsContext) interface{}

	// Visit a parse tree produced by RuleExprParser#switchLabel.
	VisitSwitchLabel(ctx *SwitchLabelContext) interface{}

	// Visit a parse tree produced by RuleExprParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#forStatementNoShortIf.
	VisitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by RuleExprParser#basicForStatement.
	VisitBasicForStatement(ctx *BasicForStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#basicForStatementNoShortIf.
	VisitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by RuleExprParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by RuleExprParser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by RuleExprParser#statementExpressionList.
	VisitStatementExpressionList(ctx *StatementExpressionListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#enhancedForStatement.
	VisitEnhancedForStatement(ctx *EnhancedForStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#enhancedForStatementNoShortIf.
	VisitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by RuleExprParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by RuleExprParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by RuleExprParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

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

	// Visit a parse tree produced by RuleExprParser#preIncrementExpression.
	VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{}

	// Visit a parse tree produced by RuleExprParser#preDecrementExpression.
	VisitPreDecrementExpression(ctx *PreDecrementExpressionContext) interface{}

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

	// Visit a parse tree produced by RuleExprParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by RuleExprParser#lambdaBody.
	VisitLambdaBody(ctx *LambdaBodyContext) interface{}

	// Visit a parse tree produced by RuleExprParser#listLiteral.
	VisitListLiteral(ctx *ListLiteralContext) interface{}

	// Visit a parse tree produced by RuleExprParser#mapLiteral.
	VisitMapLiteral(ctx *MapLiteralContext) interface{}

	// Visit a parse tree produced by RuleExprParser#setLiteral.
	VisitSetLiteral(ctx *SetLiteralContext) interface{}

	// Visit a parse tree produced by RuleExprParser#mapEntryList.
	VisitMapEntryList(ctx *MapEntryListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#mapEntry.
	VisitMapEntry(ctx *MapEntryContext) interface{}

	// Visit a parse tree produced by RuleExprParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#structLiteral.
	VisitStructLiteral(ctx *StructLiteralContext) interface{}

	// Visit a parse tree produced by RuleExprParser#structFieldInitializerList.
	VisitStructFieldInitializerList(ctx *StructFieldInitializerListContext) interface{}

	// Visit a parse tree produced by RuleExprParser#structFieldInitializer.
	VisitStructFieldInitializer(ctx *StructFieldInitializerContext) interface{}
}
