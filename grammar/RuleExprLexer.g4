/**
 * Goval Expression Language Lexer (simplified)
 * 面向规则引擎：不可变对象 + lambda + 最小控制流
 */

// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar RuleExprLexer;

// Keywords（仅保留精简后所需）
BREAK    : 'break';
CONTINUE : 'continue';
ELSE     : 'else';
FOR      : 'for';
IF       : 'if';
IN       : 'in';
VAR      : 'var';

// Integer Literals
IntegerLiteral
    : DecimalIntegerLiteral
    | HexIntegerLiteral
    | OctalIntegerLiteral
    | BinaryIntegerLiteral
    ;

fragment DecimalIntegerLiteral
    : '0' [lL]?
    | [1-9] [0-9]* [lL]?
    ;

fragment HexIntegerLiteral
    : '0' [xX] [0-9a-fA-F]+ [lL]?
    ;

fragment OctalIntegerLiteral
    : '0' [0-7]+ [lL]?
    ;

fragment BinaryIntegerLiteral
    : '0' [bB] [01]+ [lL]?
    ;

// Floating-Point Literals
FloatingPointLiteral
    : [0-9]+ '.' [0-9]* [fFdD]?
    | '.' [0-9]+ [fFdD]?
    | [0-9]+ [eE] [+-]? [0-9]+ [fFdD]?
    | [0-9]+ '.' [0-9]* [eE] [+-]? [0-9]+ [fFdD]?
    | '.' [0-9]+ [eE] [+-]? [0-9]+ [fFdD]?
    | [0-9]+ [fFdD]
    ;

// Boolean Literals
BooleanLiteral: 'true' | 'false';

// Character Literals
CharacterLiteral
    : '\'' (EscapeSequence | ~['\\\r\n]) '\''
    ;

// String Literals
StringLiteral
    : '"' (EscapeSequence | ~["\\\r\n])* '"'
    ;

// Escape Sequences
fragment EscapeSequence
    : '\\' [btnfr"'\\]
    | '\\' [0-7] [0-7]?
    | '\\' [0-3] [0-7] [0-7]
    | UnicodeEscape
    ;

fragment UnicodeEscape
    : '\\' 'u' [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F]
    ;

// The Null Literal
NullLiteral: 'null';

// Separators
LPAREN : '(';
RPAREN : ')';
LBRACE : '{';
RBRACE : '}';
LBRACK : '[';
RBRACK : ']';
SEMI   : ';';
COMMA  : ',';
DOT    : '.';
COLON  : ':';

// Operators
ASSIGN   : '=';
GT       : '>';
LT       : '<';
BANG     : '!';
TILDE    : '~';
QUESTION : '?';
EQUAL    : '==';
LE       : '<=';
GE       : '>=';
NOTEQUAL : '!=';
AND      : '&&';
OR       : '||';
ADD      : '+';
SUB      : '-';
MUL      : '*';
DIV      : '/';
BITAND   : '&';
BITOR    : '|';
CARET    : '^';
MOD      : '%';
ARROW    : '->';
LSHIFT   : '<<';
RSHIFT   : '>>';

// Identifiers
Identifier: IdentifierStart IdentifierPart*;

fragment IdentifierStart:
    [a-zA-Z_$]
    | [À-ÿ]
    | [Ā-ſ]
    | [ƀ-ɏ]
    | [一-鿿]
    | [㐀-䶿]
    | [가-힯]
    | [぀-ゟ]
    | [゠-ヿ]
;

fragment IdentifierPart:
    IdentifierStart
    | [0-9]
    | [̀-ͯ]
    | [‌-‍]
;

// Whitespace and comments
WS: [ \t\r\n]+ -> skip;
COMMENT: '/*' .*? '*/' -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;
