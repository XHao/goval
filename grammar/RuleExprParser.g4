/**
 * Goval Expression Language Grammar
 * 
 * A lightweight expression language designed for embedding in Go applications.
 * Focuses on expressions, simple functions, and seamless Go integration.
 */

// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar RuleExprParser;

options {
    tokenVocab = RuleExprLexer;
}

// ============================================================================
// Program entry and top-level structure
// ============================================================================

program
    : (structDeclaration | statement)+ EOF
    ;

// Struct declaration: using "type field" format with method support
structDeclaration
    : STRUCT Identifier '{' structMemberList? '}'
    ;

structMemberList
    : structMember (',' structMember)*
    ;

structMember
    : structField
    | structMethod
    ;

structField
    : type Identifier
    ;

// Struct method: methodName(param1 type1, param2 type2) -> returnType { body }
structMethod
    : Identifier '(' methodParameterList? ')' ('->' type)? methodBody
    ;

methodParameterList
    : methodParameter (',' methodParameter)*
    ;

methodParameter
    : type Identifier
    ;

methodBody
    : block
    ;

// ============================================================================
// Literal definitions
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
// Type system (core)
// ============================================================================

primitiveType
    : BYTE | SHORT | INT | LONG | CHAR                          // Integer types
    | FLOAT | DOUBLE                                            // Floating point types
    | BOOLEAN                                                   // Boolean type
    | STRING                                                    // String type
    ;

// Simplified type system
type
    : primitiveType
    | Identifier              // Support Integer, String etc as identifier types
    | primitiveType dims      // Basic type arrays
    | Identifier dims         // Identifier type arrays
    | functionType
    | containerType
    ;

// Add paramType to break indirect left recursion
// paramType is a subset of type, but does not include functionType
paramType
    : primitiveType
    | Identifier
    | primitiveType dims
    | Identifier dims
    | containerType
    ;

dims
    : '[' ']' ('[' ']')*
    ;

// Function type: (T1, T2) -> R or (T1 name1, T2 name2) -> R
functionType
    : '(' (functionParameter (',' functionParameter)*)? ')' '->' type
    ;

// Use paramType in functionParameter
functionParameter
    : paramType Identifier?                                  // Support optional parameter names
    ;

// Container types: List<T>, Map<K,V>, Set<T>
containerType
    : LIST '<' type '>'                                 // List<int>, List<string>
    | MAP '<' type ',' type '>'                         // Map<string, int>
    | SET '<' type '>'                                  // Set<int>
    ;

// ============================================================================
// Variables and declarations
// ============================================================================

arrayInitializer
    : '{' '}'                                           // Empty array
    | '{' variableInitializerList ','? '}'              // Non-empty array, optional trailing comma
    ;

variableInitializerList
    : variableInitializer (',' variableInitializer)*
    ;

variableDeclaratorList
    : variableDeclarator (',' variableDeclarator)*
    ;

variableDeclarator
    : variableDeclaratorId ('=' variableInitializer)?
    ;

// Enhanced variable declarator ID, supports deferred inference
variableDeclaratorId
    : Identifier dims?
    ;

variableInitializer
    : expression
    | arrayInitializer
    ;

// Simplified variable declaration
localVariableDeclaration
    : type variableDeclaratorList                       // Static type declaration: int a = 1, b, c = 3;
    | VAR varVariableDeclaratorList                     // Type inference, supports list and mandatory initialization
    ;

// Add rule to enforce var declaration must have initialization
varVariableDeclaratorList
    : varVariableDeclarator (',' varVariableDeclarator)*
    ;

varVariableDeclarator
    : variableDeclaratorId '=' variableInitializer // Note: the '=' is mandatory here
    ;

localVariableDeclarationStatement
    : localVariableDeclaration ';'
    ;

// ============================================================================
// Statements and control flow
// ============================================================================

block
    : '{' blockStatements? '}'
    | expressionBlock                                   // Expression block
    ;

expressionBlock
    : '{' blockStatements expression '}'               // Statements + last expression as return value
    ;

blockStatements
    : blockStatement+
    ;

blockStatement
    : localVariableDeclarationStatement
    | statement
    ;

statement
    : statementWithoutTrailingSubstatement
    | ifThenStatement
    | ifThenElseStatement
    | forStatement
    ;

statementNoShortIf
    : statementWithoutTrailingSubstatement
    | ifThenElseStatementNoShortIf
    | forStatementNoShortIf
    ;

statementWithoutTrailingSubstatement
    : block
    | emptyStatement_
    | expressionStatement
    | localVariableDeclarationStatement
    | switchStatement
    | breakStatement
    | continueStatement
    | returnStatement
    ;

emptyStatement_
    : ';'
    ;

expressionStatement
    : statementExpression ';'
    ;

statementExpression
    : assignment
    | preIncrementExpression
    | preDecrementExpression
    | postfixExpression
    ;

// if statements
ifThenStatement
    : 'if' '(' expression ')' statement
    ;

ifThenElseStatement
    : 'if' '(' expression ')' statementNoShortIf 'else' statement
    ;

ifThenElseStatementNoShortIf
    : 'if' '(' expression ')' statementNoShortIf 'else' statementNoShortIf
    ;

// switch statements
switchStatement
    : 'switch' '(' expression ')' switchBlock
    ;

switchBlock
    : '{' switchBlockStatementGroup* switchLabel* '}'
    ;

switchBlockStatementGroup
    : switchLabels blockStatements
    ;

switchLabels
    : switchLabel switchLabel*
    ;

switchLabel
    : 'case' expression ':'
    | 'default' ':'
    ;

// for statements
forStatement
    : basicForStatement
    | enhancedForStatement
    ;

forStatementNoShortIf
    : basicForStatementNoShortIf
    | enhancedForStatementNoShortIf
    ;

basicForStatement
    : 'for' '(' forInit? ';' expression? ';' forUpdate? ')' statement
    ;

basicForStatementNoShortIf
    : 'for' '(' forInit? ';' expression? ';' forUpdate? ')' statementNoShortIf
    ;

forInit
    : statementExpressionList
    | localVariableDeclaration
    ;

forUpdate
    : statementExpressionList
    ;

statementExpressionList
    : statementExpression (',' statementExpression)*
    ;

enhancedForStatement
    : 'for' '(' type Identifier ':' expression ')' statement
    ;

enhancedForStatementNoShortIf
    : 'for' '(' type variableDeclaratorId ':' expression ')' statementNoShortIf
    ;

// Control statements
breakStatement
    : 'break' ';'
    ;

continueStatement
    : 'continue' ';'
    ;

returnStatement
    : 'return' expression? ';'
    ;

// ============================================================================
// Expressions (by precedence from low to high)
// ============================================================================

expression
    : assignmentExpression
    ;

assignmentExpression
    : lambdaExpression
    | conditionalExpression
    | assignment
    ;

assignment
    : postfixExpression assignmentOperator expression
    ;

assignmentOperator
    : '='
    | '*='
    | '/='
    | '%='
    | '+='
    | '-='
    | '<<='
    | '>>='
    | '&='
    | '^='
    | '|='
    ;

conditionalExpression
    : conditionalOrExpression
    | conditionalOrExpression '?' expression ':' conditionalExpression
    ;

conditionalOrExpression
    : conditionalAndExpression
    | conditionalOrExpression '||' conditionalAndExpression
    ;

conditionalAndExpression
    : inclusiveOrExpression
    | conditionalAndExpression '&&' inclusiveOrExpression
    ;

inclusiveOrExpression
    : exclusiveOrExpression
    | inclusiveOrExpression '|' exclusiveOrExpression
    ;

exclusiveOrExpression
    : andExpression
    | exclusiveOrExpression '^' andExpression
    ;

andExpression
    : equalityExpression
    | andExpression '&' equalityExpression
    ;

equalityExpression
    : relationalExpression
    | equalityExpression '==' relationalExpression
    | equalityExpression '!=' relationalExpression
    ;

relationalExpression
    : shiftExpression
    | relationalExpression '<' shiftExpression
    | relationalExpression '>' shiftExpression
    | relationalExpression '<=' shiftExpression
    | relationalExpression '>=' shiftExpression
    | relationalExpression 'in' shiftExpression
    ;

shiftExpression
    : additiveExpression
    | shiftExpression LSHIFT additiveExpression
    | shiftExpression RSHIFT additiveExpression
    ;

additiveExpression
    : multiplicativeExpression
    | additiveExpression '+' multiplicativeExpression
    | additiveExpression '-' multiplicativeExpression
    ;

multiplicativeExpression
    : unaryExpression
    | multiplicativeExpression '*' unaryExpression
    | multiplicativeExpression '/' unaryExpression
    | multiplicativeExpression '%' unaryExpression
    ;

unaryExpression
    : preIncrementExpression
    | preDecrementExpression
    | '+' unaryExpression
    | '-' unaryExpression
    | unaryExpressionNotPlusMinus
    ;

preIncrementExpression
    : '++' unaryExpression
    ;

preDecrementExpression
    : '--' unaryExpression
    ;

unaryExpressionNotPlusMinus
    : postfixExpression
    | '~' unaryExpression
    | '!' unaryExpression
    ;

postfixExpression
    : primary
    | postfixExpression '[' expression ']'              // Array access
    | postfixExpression '.' Identifier                  // Field access
    | postfixExpression '.' Identifier '(' argumentList? ')' // Method call
    | postfixExpression '(' argumentList? ')'           // Function call
    | postfixExpression '++'                            // Post-increment
    | postfixExpression '--'                            // Post-decrement
    ;

primary
    : literal
    | PLACEHOLDER_VAR
    | '(' expression ')'
    | Identifier
    | THIS                                              // 'this' keyword for method context
    | listLiteral
    | mapLiteral
    | setLiteral
    | structLiteral
    | expressionBlock                                   // Support expression block as primary
    ;

argumentList
    : expression (',' expression)*
    ;

// ============================================================================
// Lambda expressions
// ============================================================================

lambdaExpression
    : lambdaParameters '->' lambdaBody              // Unified lambda syntax
    ;

lambdaParameters
    : Identifier                                    // a
    | '(' ')'                                       // ()
    | '(' formalParameterList ')'                   // (a), (a,b), (int a), (int a, int b)
    ;

formalParameterList
    : formalParameter (',' formalParameter)*
    ;

formalParameter
    : type Identifier  // Explicitly require type and identifier
    | Identifier       // Or just identifier for type inference
    ;

lambdaBody
    : type block
    | block
    | expression
    ;

// ============================================================================
// Container and struct literals
// ============================================================================

// Container literals
listLiteral
    : LIST '{' expressionList? '}'                      // List{1, 2, 3}
    | '[' expressionList? ']'                           // [1, 2, 3] syntax sugar
    ;

mapLiteral
    : MAP '{' mapEntryList? '}'                         // Map{"key1": value1, "key2": value2}
    ;

setLiteral
    : SET '{' expressionList? '}'                       // Set{1, 2, 3}
    ;

mapEntryList
    : mapEntry (',' mapEntry)*
    ;

mapEntry
    : expression ':' expression                         // key: value
    ;

expressionList
    : expression (',' expression)*
    ;

// Struct literal: Point{x: 1.0, y: 2.0}
structLiteral
    : Identifier '{' structFieldInitializerList? '}'
    ;

structFieldInitializerList
    : structFieldInitializer (',' structFieldInitializer)*
    ;

structFieldInitializer
    : Identifier ':' expression
    ;