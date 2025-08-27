/**
 * Goval Expression Language Lexer
 * 
 * Lexical analysis for a lightweight expression language
 * designed for embedding in Go applications.
 */

// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar RuleExprLexer;

// LEXER

// Keywords
BOOLEAN      : 'boolean';
BREAK        : 'break';
BYTE         : 'byte';
CASE         : 'case';
CHAR         : 'char';
CONTINUE     : 'continue';
DEFAULT      : 'default';
DOUBLE       : 'double';
ELSE         : 'else';
FLOAT        : 'float';
FOR          : 'for';
IF           : 'if';
IN           : 'in';
INT          : 'int';
LIST         : 'List';
LONG         : 'long';
MAP          : 'Map';
RETURN       : 'return';
SET          : 'Set';
SHORT        : 'short';
STRING       : 'string';
STRUCT       : 'struct';
SWITCH       : 'switch';
THIS         : 'this';
VAR          : 'var';

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
    : DecimalFloatingPointLiteral
    ;

fragment DecimalFloatingPointLiteral
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

// Operators

ASSIGN     : '=';
GT         : '>';
LT         : '<';
BANG       : '!';
TILDE      : '~';
QUESTION   : '?';
COLON      : ':';
EQUAL      : '==';
LE         : '<=';
GE         : '>=';
NOTEQUAL   : '!=';
AND        : '&&';
OR         : '||';
INC        : '++';
DEC        : '--';
ADD        : '+';
SUB        : '-';
MUL        : '*';
DIV        : '/';
BITAND     : '&';
BITOR      : '|';
CARET      : '^';
MOD        : '%';
ARROW      : '->';

ADD_ASSIGN     : '+=';
SUB_ASSIGN     : '-=';
MUL_ASSIGN     : '*=';
DIV_ASSIGN     : '/=';
AND_ASSIGN     : '&=';
OR_ASSIGN      : '|=';
XOR_ASSIGN     : '^=';
MOD_ASSIGN     : '%=';
LSHIFT_ASSIGN  : '<<=';
RSHIFT_ASSIGN  : '>>=';

// Shift operators (simplified)
LSHIFT         : '<<';
RSHIFT         : '>>';

PLACEHOLDER_VAR : '#' Identifier '#' ;

// Identifiers (must appear after all keywords in the grammar)
Identifier: IdentifierStart IdentifierPart*;

// Simplified identifier rules, supporting basic ASCII characters and common Unicode ranges
fragment IdentifierStart:
    [a-zA-Z_$]                          // Basic ASCII letters, underscore, dollar sign
    | [\u00A0-\u00FF]                   // Extended ASCII (Latin-1 Supplement)
    | [\u0100-\u017F]                   // Latin Extended-A
    | [\u0180-\u024F]                   // Latin Extended-B
    | [\u4E00-\u9FFF]                   // CJK Unified Ideographs
    | [\u3400-\u4DBF]                   // CJK Extension A
    | [\uAC00-\uD7AF]                   // Hangul Syllables
    | [\u3040-\u309F]                   // Hiragana
    | [\u30A0-\u30FF]                   // Katakana
;

fragment IdentifierPart:
    IdentifierStart
    | [0-9]                             // Digits
    | [\u0300-\u036F]                   // Combining Diacritical Marks
    | [\u200C-\u200D]                   // Zero Width Non-Joiner and Joiner
;

//
// Whitespace and comments
//

WS: [ \t\r\n\u000C]+ -> skip;

COMMENT: '/*' .*? '*/' -> skip;

LINE_COMMENT: '//' ~[\r\n]* -> skip;