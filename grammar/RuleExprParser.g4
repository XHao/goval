/**
 * Goval Expression Language Grammar (simplified)
 * 面向规则引擎：不可变对象(lambda 工厂+Map) + lambda + if/else + for-in
 */

// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar RuleExprParser;

options {
    tokenVocab = RuleExprLexer;
}

// ============================================================================
// Program entry
// ============================================================================

program
    : statement* EOF
    ;

// ============================================================================
// Statements
// ============================================================================

block
    : '{' blockStatements? '}'
    ;

blockStatements
    : blockStatement+
    ;

blockStatement
    : localVariableDeclarationStatement
    | statement
    ;

localVariableDeclarationStatement
    : localVariableDeclaration SEMI?
    ;

localVariableDeclaration
    : VAR varVariableDeclaratorList          // var x = 1, y = 2
    ;

varVariableDeclaratorList
    : varVariableDeclarator (COMMA varVariableDeclarator)*
    ;

varVariableDeclarator
    : Identifier ASSIGN variableInitializer   // '=' 强制初始化
    ;

variableInitializer
    : expression
    ;

statement
    : block
    | ifStatement
    | forStatement
    | breakStatement
    | continueStatement
    | expressionStatement
    | localVariableDeclarationStatement
    | SEMI
    ;

expressionStatement
    : expression SEMI?
    ;

// if / else if / else
ifStatement
    : IF LPAREN expression RPAREN statement (ELSE statement)?
    ;

// for-in only: for x in expr {}  |  for k, v in expr {}
forStatement
    : FOR Identifier (COMMA Identifier)? IN expression block
    ;

breakStatement
    : BREAK SEMI?
    ;

continueStatement
    : CONTINUE SEMI?
    ;

// ============================================================================
// Literals
// ============================================================================

literal
    : IntegerLiteral
    | FloatingPointLiteral
    | BooleanLiteral
    | CharacterLiteral
    | StringLiteral
    | NullLiteral
    ;

// ============================================================================
// Expressions (by precedence, low to high)
// ============================================================================

expression
    : assignmentExpression
    ;

// 不可变铁律：左值只能是标识符
assignmentExpression
    : lambdaExpression
    | conditionalExpression
    | assignment
    ;

assignment
    : Identifier ASSIGN expression
    ;

conditionalExpression
    : conditionalOrExpression
    | conditionalOrExpression QUESTION expression COLON conditionalExpression
    ;

conditionalOrExpression
    : conditionalAndExpression
    | conditionalOrExpression OR conditionalAndExpression
    ;

conditionalAndExpression
    : inclusiveOrExpression
    | conditionalAndExpression AND inclusiveOrExpression
    ;

inclusiveOrExpression
    : exclusiveOrExpression
    | inclusiveOrExpression BITOR exclusiveOrExpression
    ;

exclusiveOrExpression
    : andExpression
    | exclusiveOrExpression CARET andExpression
    ;

andExpression
    : equalityExpression
    | andExpression BITAND equalityExpression
    ;

equalityExpression
    : relationalExpression
    | equalityExpression EQUAL relationalExpression
    | equalityExpression NOTEQUAL relationalExpression
    ;

relationalExpression
    : shiftExpression
    | relationalExpression LT shiftExpression
    | relationalExpression GT shiftExpression
    | relationalExpression LE shiftExpression
    | relationalExpression GE shiftExpression
    | relationalExpression IN shiftExpression
    ;

shiftExpression
    : additiveExpression
    | shiftExpression LSHIFT additiveExpression
    | shiftExpression RSHIFT additiveExpression
    ;

additiveExpression
    : multiplicativeExpression
    | additiveExpression ADD multiplicativeExpression
    | additiveExpression SUB multiplicativeExpression
    ;

multiplicativeExpression
    : unaryExpression
    | multiplicativeExpression MUL unaryExpression
    | multiplicativeExpression DIV unaryExpression
    | multiplicativeExpression MOD unaryExpression
    ;

unaryExpression
    : ADD unaryExpression
    | SUB unaryExpression
    | unaryExpressionNotPlusMinus
    ;

unaryExpressionNotPlusMinus
    : postfixExpression
    | TILDE unaryExpression
    | BANG unaryExpression
    ;

postfixExpression
    : primary
    | postfixExpression LBRACK expression RBRACK              // 下标访问（只读）
    | postfixExpression DOT Identifier                        // 字段访问（只读）
    | postfixExpression DOT Identifier LPAREN argumentList? RPAREN  // 方法调用
    | postfixExpression LPAREN argumentList? RPAREN           // 函数调用
    ;

primary
    : literal
    | LPAREN expression RPAREN
    | Identifier
    | listLiteral
    | mapLiteral
    | expressionBlock
    ;

argumentList
    : expression (COMMA expression)*
    ;

// ============================================================================
// Lambda
// ============================================================================

lambdaExpression
    : lambdaParameters ARROW lambdaBody
    ;

lambdaParameters
    : Identifier
    | LPAREN RPAREN
    | LPAREN formalParameterList RPAREN
    ;

formalParameterList
    : Identifier (COMMA Identifier)*          // 仅标识符，无类型注解
    ;

lambdaBody
    : expressionBlock
    | expression
    ;

// ============================================================================
// Container literals
// ============================================================================

listLiteral
    : LBRACK expressionList? RBRACK
    ;

mapLiteral
    : LBRACE mapEntryList? RBRACE
    ;

mapEntryList
    : mapEntry (COMMA mapEntry)*
    ;

mapEntry
    : expression COLON expression
    ;

expressionList
    : expression (COMMA expression)*
    ;

// ============================================================================
// Expression block（lambda 块体 / 顶层块）
// ============================================================================

expressionBlock
    : LBRACE blockStatement* expression RBRACE    // 末尾必须是表达式（返回值）
    ;
