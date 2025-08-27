// Code generated from RuleExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // RuleExprParser
import "github.com/antlr4-go/antlr/v4"

// RuleExprParserListener is a complete listener for a parse tree produced by RuleExprParser.
type RuleExprParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStructDeclaration is called when entering the structDeclaration production.
	EnterStructDeclaration(c *StructDeclarationContext)

	// EnterStructMemberList is called when entering the structMemberList production.
	EnterStructMemberList(c *StructMemberListContext)

	// EnterStructMember is called when entering the structMember production.
	EnterStructMember(c *StructMemberContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterStructMethod is called when entering the structMethod production.
	EnterStructMethod(c *StructMethodContext)

	// EnterMethodParameterList is called when entering the methodParameterList production.
	EnterMethodParameterList(c *MethodParameterListContext)

	// EnterMethodParameter is called when entering the methodParameter production.
	EnterMethodParameter(c *MethodParameterContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterParamType is called when entering the paramType production.
	EnterParamType(c *ParamTypeContext)

	// EnterDims is called when entering the dims production.
	EnterDims(c *DimsContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterFunctionParameter is called when entering the functionParameter production.
	EnterFunctionParameter(c *FunctionParameterContext)

	// EnterContainerType is called when entering the containerType production.
	EnterContainerType(c *ContainerTypeContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterVariableInitializerList is called when entering the variableInitializerList production.
	EnterVariableInitializerList(c *VariableInitializerListContext)

	// EnterVariableDeclaratorList is called when entering the variableDeclaratorList production.
	EnterVariableDeclaratorList(c *VariableDeclaratorListContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableDeclaratorId is called when entering the variableDeclaratorId production.
	EnterVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterVarVariableDeclaratorList is called when entering the varVariableDeclaratorList production.
	EnterVarVariableDeclaratorList(c *VarVariableDeclaratorListContext)

	// EnterVarVariableDeclarator is called when entering the varVariableDeclarator production.
	EnterVarVariableDeclarator(c *VarVariableDeclaratorContext)

	// EnterLocalVariableDeclarationStatement is called when entering the localVariableDeclarationStatement production.
	EnterLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterExpressionBlock is called when entering the expressionBlock production.
	EnterExpressionBlock(c *ExpressionBlockContext)

	// EnterBlockStatements is called when entering the blockStatements production.
	EnterBlockStatements(c *BlockStatementsContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStatementNoShortIf is called when entering the statementNoShortIf production.
	EnterStatementNoShortIf(c *StatementNoShortIfContext)

	// EnterStatementWithoutTrailingSubstatement is called when entering the statementWithoutTrailingSubstatement production.
	EnterStatementWithoutTrailingSubstatement(c *StatementWithoutTrailingSubstatementContext)

	// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
	EnterEmptyStatement_(c *EmptyStatement_Context)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterStatementExpression is called when entering the statementExpression production.
	EnterStatementExpression(c *StatementExpressionContext)

	// EnterIfThenStatement is called when entering the ifThenStatement production.
	EnterIfThenStatement(c *IfThenStatementContext)

	// EnterIfThenElseStatement is called when entering the ifThenElseStatement production.
	EnterIfThenElseStatement(c *IfThenElseStatementContext)

	// EnterIfThenElseStatementNoShortIf is called when entering the ifThenElseStatementNoShortIf production.
	EnterIfThenElseStatementNoShortIf(c *IfThenElseStatementNoShortIfContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchBlock is called when entering the switchBlock production.
	EnterSwitchBlock(c *SwitchBlockContext)

	// EnterSwitchBlockStatementGroup is called when entering the switchBlockStatementGroup production.
	EnterSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// EnterSwitchLabels is called when entering the switchLabels production.
	EnterSwitchLabels(c *SwitchLabelsContext)

	// EnterSwitchLabel is called when entering the switchLabel production.
	EnterSwitchLabel(c *SwitchLabelContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForStatementNoShortIf is called when entering the forStatementNoShortIf production.
	EnterForStatementNoShortIf(c *ForStatementNoShortIfContext)

	// EnterBasicForStatement is called when entering the basicForStatement production.
	EnterBasicForStatement(c *BasicForStatementContext)

	// EnterBasicForStatementNoShortIf is called when entering the basicForStatementNoShortIf production.
	EnterBasicForStatementNoShortIf(c *BasicForStatementNoShortIfContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterStatementExpressionList is called when entering the statementExpressionList production.
	EnterStatementExpressionList(c *StatementExpressionListContext)

	// EnterEnhancedForStatement is called when entering the enhancedForStatement production.
	EnterEnhancedForStatement(c *EnhancedForStatementContext)

	// EnterEnhancedForStatementNoShortIf is called when entering the enhancedForStatementNoShortIf production.
	EnterEnhancedForStatementNoShortIf(c *EnhancedForStatementNoShortIfContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterConditionalOrExpression is called when entering the conditionalOrExpression production.
	EnterConditionalOrExpression(c *ConditionalOrExpressionContext)

	// EnterConditionalAndExpression is called when entering the conditionalAndExpression production.
	EnterConditionalAndExpression(c *ConditionalAndExpressionContext)

	// EnterInclusiveOrExpression is called when entering the inclusiveOrExpression production.
	EnterInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// EnterExclusiveOrExpression is called when entering the exclusiveOrExpression production.
	EnterExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterPreIncrementExpression is called when entering the preIncrementExpression production.
	EnterPreIncrementExpression(c *PreIncrementExpressionContext)

	// EnterPreDecrementExpression is called when entering the preDecrementExpression production.
	EnterPreDecrementExpression(c *PreDecrementExpressionContext)

	// EnterUnaryExpressionNotPlusMinus is called when entering the unaryExpressionNotPlusMinus production.
	EnterUnaryExpressionNotPlusMinus(c *UnaryExpressionNotPlusMinusContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterLambdaExpression is called when entering the lambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterLambdaParameters is called when entering the lambdaParameters production.
	EnterLambdaParameters(c *LambdaParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterLambdaBody is called when entering the lambdaBody production.
	EnterLambdaBody(c *LambdaBodyContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterSetLiteral is called when entering the setLiteral production.
	EnterSetLiteral(c *SetLiteralContext)

	// EnterMapEntryList is called when entering the mapEntryList production.
	EnterMapEntryList(c *MapEntryListContext)

	// EnterMapEntry is called when entering the mapEntry production.
	EnterMapEntry(c *MapEntryContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterStructLiteral is called when entering the structLiteral production.
	EnterStructLiteral(c *StructLiteralContext)

	// EnterStructFieldInitializerList is called when entering the structFieldInitializerList production.
	EnterStructFieldInitializerList(c *StructFieldInitializerListContext)

	// EnterStructFieldInitializer is called when entering the structFieldInitializer production.
	EnterStructFieldInitializer(c *StructFieldInitializerContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStructDeclaration is called when exiting the structDeclaration production.
	ExitStructDeclaration(c *StructDeclarationContext)

	// ExitStructMemberList is called when exiting the structMemberList production.
	ExitStructMemberList(c *StructMemberListContext)

	// ExitStructMember is called when exiting the structMember production.
	ExitStructMember(c *StructMemberContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitStructMethod is called when exiting the structMethod production.
	ExitStructMethod(c *StructMethodContext)

	// ExitMethodParameterList is called when exiting the methodParameterList production.
	ExitMethodParameterList(c *MethodParameterListContext)

	// ExitMethodParameter is called when exiting the methodParameter production.
	ExitMethodParameter(c *MethodParameterContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitParamType is called when exiting the paramType production.
	ExitParamType(c *ParamTypeContext)

	// ExitDims is called when exiting the dims production.
	ExitDims(c *DimsContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitFunctionParameter is called when exiting the functionParameter production.
	ExitFunctionParameter(c *FunctionParameterContext)

	// ExitContainerType is called when exiting the containerType production.
	ExitContainerType(c *ContainerTypeContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitVariableInitializerList is called when exiting the variableInitializerList production.
	ExitVariableInitializerList(c *VariableInitializerListContext)

	// ExitVariableDeclaratorList is called when exiting the variableDeclaratorList production.
	ExitVariableDeclaratorList(c *VariableDeclaratorListContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableDeclaratorId is called when exiting the variableDeclaratorId production.
	ExitVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitVarVariableDeclaratorList is called when exiting the varVariableDeclaratorList production.
	ExitVarVariableDeclaratorList(c *VarVariableDeclaratorListContext)

	// ExitVarVariableDeclarator is called when exiting the varVariableDeclarator production.
	ExitVarVariableDeclarator(c *VarVariableDeclaratorContext)

	// ExitLocalVariableDeclarationStatement is called when exiting the localVariableDeclarationStatement production.
	ExitLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitExpressionBlock is called when exiting the expressionBlock production.
	ExitExpressionBlock(c *ExpressionBlockContext)

	// ExitBlockStatements is called when exiting the blockStatements production.
	ExitBlockStatements(c *BlockStatementsContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStatementNoShortIf is called when exiting the statementNoShortIf production.
	ExitStatementNoShortIf(c *StatementNoShortIfContext)

	// ExitStatementWithoutTrailingSubstatement is called when exiting the statementWithoutTrailingSubstatement production.
	ExitStatementWithoutTrailingSubstatement(c *StatementWithoutTrailingSubstatementContext)

	// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
	ExitEmptyStatement_(c *EmptyStatement_Context)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitStatementExpression is called when exiting the statementExpression production.
	ExitStatementExpression(c *StatementExpressionContext)

	// ExitIfThenStatement is called when exiting the ifThenStatement production.
	ExitIfThenStatement(c *IfThenStatementContext)

	// ExitIfThenElseStatement is called when exiting the ifThenElseStatement production.
	ExitIfThenElseStatement(c *IfThenElseStatementContext)

	// ExitIfThenElseStatementNoShortIf is called when exiting the ifThenElseStatementNoShortIf production.
	ExitIfThenElseStatementNoShortIf(c *IfThenElseStatementNoShortIfContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchBlock is called when exiting the switchBlock production.
	ExitSwitchBlock(c *SwitchBlockContext)

	// ExitSwitchBlockStatementGroup is called when exiting the switchBlockStatementGroup production.
	ExitSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// ExitSwitchLabels is called when exiting the switchLabels production.
	ExitSwitchLabels(c *SwitchLabelsContext)

	// ExitSwitchLabel is called when exiting the switchLabel production.
	ExitSwitchLabel(c *SwitchLabelContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForStatementNoShortIf is called when exiting the forStatementNoShortIf production.
	ExitForStatementNoShortIf(c *ForStatementNoShortIfContext)

	// ExitBasicForStatement is called when exiting the basicForStatement production.
	ExitBasicForStatement(c *BasicForStatementContext)

	// ExitBasicForStatementNoShortIf is called when exiting the basicForStatementNoShortIf production.
	ExitBasicForStatementNoShortIf(c *BasicForStatementNoShortIfContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitStatementExpressionList is called when exiting the statementExpressionList production.
	ExitStatementExpressionList(c *StatementExpressionListContext)

	// ExitEnhancedForStatement is called when exiting the enhancedForStatement production.
	ExitEnhancedForStatement(c *EnhancedForStatementContext)

	// ExitEnhancedForStatementNoShortIf is called when exiting the enhancedForStatementNoShortIf production.
	ExitEnhancedForStatementNoShortIf(c *EnhancedForStatementNoShortIfContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitConditionalOrExpression is called when exiting the conditionalOrExpression production.
	ExitConditionalOrExpression(c *ConditionalOrExpressionContext)

	// ExitConditionalAndExpression is called when exiting the conditionalAndExpression production.
	ExitConditionalAndExpression(c *ConditionalAndExpressionContext)

	// ExitInclusiveOrExpression is called when exiting the inclusiveOrExpression production.
	ExitInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// ExitExclusiveOrExpression is called when exiting the exclusiveOrExpression production.
	ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitPreIncrementExpression is called when exiting the preIncrementExpression production.
	ExitPreIncrementExpression(c *PreIncrementExpressionContext)

	// ExitPreDecrementExpression is called when exiting the preDecrementExpression production.
	ExitPreDecrementExpression(c *PreDecrementExpressionContext)

	// ExitUnaryExpressionNotPlusMinus is called when exiting the unaryExpressionNotPlusMinus production.
	ExitUnaryExpressionNotPlusMinus(c *UnaryExpressionNotPlusMinusContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitLambdaExpression is called when exiting the lambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitLambdaParameters is called when exiting the lambdaParameters production.
	ExitLambdaParameters(c *LambdaParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitLambdaBody is called when exiting the lambdaBody production.
	ExitLambdaBody(c *LambdaBodyContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitSetLiteral is called when exiting the setLiteral production.
	ExitSetLiteral(c *SetLiteralContext)

	// ExitMapEntryList is called when exiting the mapEntryList production.
	ExitMapEntryList(c *MapEntryListContext)

	// ExitMapEntry is called when exiting the mapEntry production.
	ExitMapEntry(c *MapEntryContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitStructLiteral is called when exiting the structLiteral production.
	ExitStructLiteral(c *StructLiteralContext)

	// ExitStructFieldInitializerList is called when exiting the structFieldInitializerList production.
	ExitStructFieldInitializerList(c *StructFieldInitializerListContext)

	// ExitStructFieldInitializer is called when exiting the structFieldInitializer production.
	ExitStructFieldInitializer(c *StructFieldInitializerContext)
}
