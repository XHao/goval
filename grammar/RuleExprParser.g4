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

// statement 备选排序：expressionStatement 在 block 之前。语句级 {} 同时匹配
// 空 block 与空 Map 字面量，歧义按备选顺序消解——排在前面的赢。若 block 在前，
// `{}["k"]` 会断成两条语句（空 block + 列表字面量）静默返回 ["k"]，且裸 `{}`
// 求值为 null 而非空 Map。优先按表达式解析后：`{}` 是 Map 字面量；无尾表达式的
// `{ var x = 1 }` 仍只有 block 可匹配，不受影响。if/for 等关键字开头的语句
// 无法作为表达式解析，自然落到各自备选。
statement
    : expressionStatement
    | block
    | ifStatement
    | forStatement
    | breakStatement
    | continueStatement
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

// 赋值：左值为 postfix 表达式（标识符 / 字段 / 下标），形状在编译期校验
assignmentExpression
    : lambdaExpression
    | conditionalExpression
    | assignment
    ;

// 赋值左值的合法形状：标识符 / base.field / base[index]。
// 收窄成显式三形而非泛化 postfixExpression：否则 `l[0] = 9` 会歧义解析成
// 语句 `l` + 语句 `[0] = 9`（[0] 被当列表字面量左值）。裸字面量左值
// 在语法层即不可解析，语句边界歧义随之消失。
assignment
    : Identifier ASSIGN expression
    | postfixExpression DOT Identifier ASSIGN expression
    | postfixExpression LBRACK expression RBRACK ASSIGN expression
    ;

// 三元：colon 分支放宽为完整 expression，允许 else 分支裸写赋值
// （then 分支本就是 expression）。右结合经 expression → conditionalExpression 保持。
conditionalExpression
    : conditionalOrExpression (QUESTION expression COLON expression)?
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
    | THIS
    | CAPTURE mapLiteral      // capture { ... }：关闭隐式 this 改写（纯编译期 pragma）
    | CAPTURE lambdaExpression // capture (a) -> ...：同上，作用于单个 lambda
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

// 非贪婪循环：语句尽可能少消费，把 token 留给尾表达式。
// 否则 { var t = 1; t + 1 } 会被贪婪解析成语句 t + 尾表达式 +1（一元正号），
// 导致求值结果错误。非贪婪下两种切分都可行时优先少消费语句。
expressionBlock
    : LBRACE blockStatement*? expression RBRACE    // 末尾必须是表达式（返回值）
    ;
