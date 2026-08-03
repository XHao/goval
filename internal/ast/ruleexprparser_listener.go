// Code generated from RuleExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // RuleExprParser
import "github.com/antlr4-go/antlr/v4"

// RuleExprParserListener is a complete listener for a parse tree produced by RuleExprParser.
type RuleExprParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStatements is called when entering the blockStatements production.
	EnterBlockStatements(c *BlockStatementsContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterLocalVariableDeclarationStatement is called when entering the localVariableDeclarationStatement production.
	EnterLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterVarVariableDeclaratorList is called when entering the varVariableDeclaratorList production.
	EnterVarVariableDeclaratorList(c *VarVariableDeclaratorListContext)

	// EnterVarVariableDeclarator is called when entering the varVariableDeclarator production.
	EnterVarVariableDeclarator(c *VarVariableDeclaratorContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

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

	// EnterLambdaBody is called when entering the lambdaBody production.
	EnterLambdaBody(c *LambdaBodyContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterMapEntryList is called when entering the mapEntryList production.
	EnterMapEntryList(c *MapEntryListContext)

	// EnterMapEntry is called when entering the mapEntry production.
	EnterMapEntry(c *MapEntryContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterExpressionBlock is called when entering the expressionBlock production.
	EnterExpressionBlock(c *ExpressionBlockContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStatements is called when exiting the blockStatements production.
	ExitBlockStatements(c *BlockStatementsContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitLocalVariableDeclarationStatement is called when exiting the localVariableDeclarationStatement production.
	ExitLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitVarVariableDeclaratorList is called when exiting the varVariableDeclaratorList production.
	ExitVarVariableDeclaratorList(c *VarVariableDeclaratorListContext)

	// ExitVarVariableDeclarator is called when exiting the varVariableDeclarator production.
	ExitVarVariableDeclarator(c *VarVariableDeclaratorContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

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

	// ExitLambdaBody is called when exiting the lambdaBody production.
	ExitLambdaBody(c *LambdaBodyContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitMapEntryList is called when exiting the mapEntryList production.
	ExitMapEntryList(c *MapEntryListContext)

	// ExitMapEntry is called when exiting the mapEntry production.
	ExitMapEntry(c *MapEntryContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpressionBlock is called when exiting the expressionBlock production.
	ExitExpressionBlock(c *ExpressionBlockContext)
}
