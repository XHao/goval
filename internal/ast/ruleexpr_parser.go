// Code generated from RuleExprParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // RuleExprParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RuleExprParser struct {
	*antlr.BaseParser
}

var RuleExprParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ruleexprparserParserInit() {
	staticData := &RuleExprParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'break'", "'continue'", "'else'", "'for'", "'if'", "'in'", "'var'",
		"", "", "", "", "", "'null'", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"';'", "','", "'.'", "':'", "'='", "'>'", "'<'", "'!'", "'~'", "'?'",
		"'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'+'", "'-'", "'*'",
		"'/'", "'&'", "'|'", "'^'", "'%'", "'->'", "'<<'", "'>>'",
	}
	staticData.SymbolicNames = []string{
		"", "BREAK", "CONTINUE", "ELSE", "FOR", "IF", "IN", "VAR", "IntegerLiteral",
		"FloatingPointLiteral", "BooleanLiteral", "CharacterLiteral", "StringLiteral",
		"NullLiteral", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
		"SEMI", "COMMA", "DOT", "COLON", "ASSIGN", "GT", "LT", "BANG", "TILDE",
		"QUESTION", "EQUAL", "LE", "GE", "NOTEQUAL", "AND", "OR", "ADD", "SUB",
		"MUL", "DIV", "BITAND", "BITOR", "CARET", "MOD", "ARROW", "LSHIFT",
		"RSHIFT", "PLACEHOLDER_VAR", "Identifier", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "block", "blockStatements", "blockStatement", "localVariableDeclarationStatement",
		"localVariableDeclaration", "varVariableDeclaratorList", "varVariableDeclarator",
		"variableInitializer", "statement", "expressionStatement", "ifStatement",
		"forStatement", "breakStatement", "continueStatement", "literal", "expression",
		"assignmentExpression", "assignment", "conditionalExpression", "conditionalOrExpression",
		"conditionalAndExpression", "inclusiveOrExpression", "exclusiveOrExpression",
		"andExpression", "equalityExpression", "relationalExpression", "shiftExpression",
		"additiveExpression", "multiplicativeExpression", "unaryExpression",
		"unaryExpressionNotPlusMinus", "postfixExpression", "primary", "argumentList",
		"lambdaExpression", "lambdaParameters", "formalParameterList", "lambdaBody",
		"listLiteral", "mapLiteral", "mapEntryList", "mapEntry", "expressionList",
		"expressionBlock",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 51, 462, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 5, 0, 92, 8, 0, 10, 0, 12,
		0, 95, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 101, 8, 1, 1, 1, 1, 1, 1, 2,
		4, 2, 106, 8, 2, 11, 2, 12, 2, 107, 1, 3, 1, 3, 3, 3, 112, 8, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 123, 8, 6, 10, 6,
		12, 6, 126, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 142, 8, 9, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 154, 8, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 160, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 3, 17, 179, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 192, 8, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 200, 8, 20, 10, 20, 12, 20, 203, 9,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 211, 8, 21, 10, 21,
		12, 21, 214, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 222,
		8, 22, 10, 22, 12, 22, 225, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 5, 23, 233, 8, 23, 10, 23, 12, 23, 236, 9, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 5, 24, 244, 8, 24, 10, 24, 12, 24, 247, 9, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 258,
		8, 25, 10, 25, 12, 25, 261, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 5, 26, 281, 8, 26, 10, 26, 12, 26, 284, 9, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 295, 8, 27,
		10, 27, 12, 27, 298, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 5, 28, 309, 8, 28, 10, 28, 12, 28, 312, 9, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 5, 29, 326, 8, 29, 10, 29, 12, 29, 329, 9, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 3, 30, 336, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3,
		31, 343, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 361, 8,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 367, 8, 32, 1, 32, 5, 32, 370, 8,
		32, 10, 32, 12, 32, 373, 9, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 385, 8, 33, 1, 34, 1, 34, 1, 34, 5,
		34, 390, 8, 34, 10, 34, 12, 34, 393, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 406, 8, 36, 1,
		37, 1, 37, 1, 37, 5, 37, 411, 8, 37, 10, 37, 12, 37, 414, 9, 37, 1, 38,
		1, 38, 3, 38, 418, 8, 38, 1, 39, 1, 39, 3, 39, 422, 8, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 3, 40, 428, 8, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 5,
		41, 435, 8, 41, 10, 41, 12, 41, 438, 9, 41, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 43, 5, 43, 447, 8, 43, 10, 43, 12, 43, 450, 9, 43, 1,
		44, 1, 44, 5, 44, 454, 8, 44, 10, 44, 12, 44, 457, 9, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 0, 11, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 64, 45, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 84, 86, 88, 0, 1, 1, 0, 8, 13, 478, 0, 93, 1, 0, 0, 0,
		2, 98, 1, 0, 0, 0, 4, 105, 1, 0, 0, 0, 6, 111, 1, 0, 0, 0, 8, 113, 1, 0,
		0, 0, 10, 116, 1, 0, 0, 0, 12, 119, 1, 0, 0, 0, 14, 127, 1, 0, 0, 0, 16,
		131, 1, 0, 0, 0, 18, 141, 1, 0, 0, 0, 20, 143, 1, 0, 0, 0, 22, 146, 1,
		0, 0, 0, 24, 155, 1, 0, 0, 0, 26, 165, 1, 0, 0, 0, 28, 168, 1, 0, 0, 0,
		30, 171, 1, 0, 0, 0, 32, 173, 1, 0, 0, 0, 34, 178, 1, 0, 0, 0, 36, 180,
		1, 0, 0, 0, 38, 191, 1, 0, 0, 0, 40, 193, 1, 0, 0, 0, 42, 204, 1, 0, 0,
		0, 44, 215, 1, 0, 0, 0, 46, 226, 1, 0, 0, 0, 48, 237, 1, 0, 0, 0, 50, 248,
		1, 0, 0, 0, 52, 262, 1, 0, 0, 0, 54, 285, 1, 0, 0, 0, 56, 299, 1, 0, 0,
		0, 58, 313, 1, 0, 0, 0, 60, 335, 1, 0, 0, 0, 62, 342, 1, 0, 0, 0, 64, 344,
		1, 0, 0, 0, 66, 384, 1, 0, 0, 0, 68, 386, 1, 0, 0, 0, 70, 394, 1, 0, 0,
		0, 72, 405, 1, 0, 0, 0, 74, 407, 1, 0, 0, 0, 76, 417, 1, 0, 0, 0, 78, 419,
		1, 0, 0, 0, 80, 425, 1, 0, 0, 0, 82, 431, 1, 0, 0, 0, 84, 439, 1, 0, 0,
		0, 86, 443, 1, 0, 0, 0, 88, 451, 1, 0, 0, 0, 90, 92, 3, 18, 9, 0, 91, 90,
		1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0,
		94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 5, 0, 0, 1, 97, 1, 1, 0,
		0, 0, 98, 100, 5, 16, 0, 0, 99, 101, 3, 4, 2, 0, 100, 99, 1, 0, 0, 0, 100,
		101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 5, 17, 0, 0, 103, 3, 1,
		0, 0, 0, 104, 106, 3, 6, 3, 0, 105, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0,
		0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 5, 1, 0, 0, 0, 109,
		112, 3, 8, 4, 0, 110, 112, 3, 18, 9, 0, 111, 109, 1, 0, 0, 0, 111, 110,
		1, 0, 0, 0, 112, 7, 1, 0, 0, 0, 113, 114, 3, 10, 5, 0, 114, 115, 5, 20,
		0, 0, 115, 9, 1, 0, 0, 0, 116, 117, 5, 7, 0, 0, 117, 118, 3, 12, 6, 0,
		118, 11, 1, 0, 0, 0, 119, 124, 3, 14, 7, 0, 120, 121, 5, 21, 0, 0, 121,
		123, 3, 14, 7, 0, 122, 120, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122,
		1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 13, 1, 0, 0, 0, 126, 124, 1, 0,
		0, 0, 127, 128, 5, 48, 0, 0, 128, 129, 5, 24, 0, 0, 129, 130, 3, 16, 8,
		0, 130, 15, 1, 0, 0, 0, 131, 132, 3, 32, 16, 0, 132, 17, 1, 0, 0, 0, 133,
		142, 3, 2, 1, 0, 134, 142, 3, 22, 11, 0, 135, 142, 3, 24, 12, 0, 136, 142,
		3, 26, 13, 0, 137, 142, 3, 28, 14, 0, 138, 142, 3, 20, 10, 0, 139, 142,
		3, 8, 4, 0, 140, 142, 5, 20, 0, 0, 141, 133, 1, 0, 0, 0, 141, 134, 1, 0,
		0, 0, 141, 135, 1, 0, 0, 0, 141, 136, 1, 0, 0, 0, 141, 137, 1, 0, 0, 0,
		141, 138, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142,
		19, 1, 0, 0, 0, 143, 144, 3, 32, 16, 0, 144, 145, 5, 20, 0, 0, 145, 21,
		1, 0, 0, 0, 146, 147, 5, 5, 0, 0, 147, 148, 5, 14, 0, 0, 148, 149, 3, 32,
		16, 0, 149, 150, 5, 15, 0, 0, 150, 153, 3, 18, 9, 0, 151, 152, 5, 3, 0,
		0, 152, 154, 3, 18, 9, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		23, 1, 0, 0, 0, 155, 156, 5, 4, 0, 0, 156, 159, 5, 48, 0, 0, 157, 158,
		5, 21, 0, 0, 158, 160, 5, 48, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1,
		0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 6, 0, 0, 162, 163, 3, 32, 16,
		0, 163, 164, 3, 2, 1, 0, 164, 25, 1, 0, 0, 0, 165, 166, 5, 1, 0, 0, 166,
		167, 5, 20, 0, 0, 167, 27, 1, 0, 0, 0, 168, 169, 5, 2, 0, 0, 169, 170,
		5, 20, 0, 0, 170, 29, 1, 0, 0, 0, 171, 172, 7, 0, 0, 0, 172, 31, 1, 0,
		0, 0, 173, 174, 3, 34, 17, 0, 174, 33, 1, 0, 0, 0, 175, 179, 3, 70, 35,
		0, 176, 179, 3, 38, 19, 0, 177, 179, 3, 36, 18, 0, 178, 175, 1, 0, 0, 0,
		178, 176, 1, 0, 0, 0, 178, 177, 1, 0, 0, 0, 179, 35, 1, 0, 0, 0, 180, 181,
		5, 48, 0, 0, 181, 182, 5, 24, 0, 0, 182, 183, 3, 32, 16, 0, 183, 37, 1,
		0, 0, 0, 184, 192, 3, 40, 20, 0, 185, 186, 3, 40, 20, 0, 186, 187, 5, 29,
		0, 0, 187, 188, 3, 32, 16, 0, 188, 189, 5, 23, 0, 0, 189, 190, 3, 38, 19,
		0, 190, 192, 1, 0, 0, 0, 191, 184, 1, 0, 0, 0, 191, 185, 1, 0, 0, 0, 192,
		39, 1, 0, 0, 0, 193, 194, 6, 20, -1, 0, 194, 195, 3, 42, 21, 0, 195, 201,
		1, 0, 0, 0, 196, 197, 10, 1, 0, 0, 197, 198, 5, 35, 0, 0, 198, 200, 3,
		42, 21, 0, 199, 196, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0,
		0, 0, 201, 202, 1, 0, 0, 0, 202, 41, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0,
		204, 205, 6, 21, -1, 0, 205, 206, 3, 44, 22, 0, 206, 212, 1, 0, 0, 0, 207,
		208, 10, 1, 0, 0, 208, 209, 5, 34, 0, 0, 209, 211, 3, 44, 22, 0, 210, 207,
		1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0,
		0, 0, 213, 43, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 216, 6, 22, -1, 0,
		216, 217, 3, 46, 23, 0, 217, 223, 1, 0, 0, 0, 218, 219, 10, 1, 0, 0, 219,
		220, 5, 41, 0, 0, 220, 222, 3, 46, 23, 0, 221, 218, 1, 0, 0, 0, 222, 225,
		1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 45, 1, 0,
		0, 0, 225, 223, 1, 0, 0, 0, 226, 227, 6, 23, -1, 0, 227, 228, 3, 48, 24,
		0, 228, 234, 1, 0, 0, 0, 229, 230, 10, 1, 0, 0, 230, 231, 5, 42, 0, 0,
		231, 233, 3, 48, 24, 0, 232, 229, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234,
		232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 47, 1, 0, 0, 0, 236, 234, 1,
		0, 0, 0, 237, 238, 6, 24, -1, 0, 238, 239, 3, 50, 25, 0, 239, 245, 1, 0,
		0, 0, 240, 241, 10, 1, 0, 0, 241, 242, 5, 40, 0, 0, 242, 244, 3, 50, 25,
		0, 243, 240, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245,
		246, 1, 0, 0, 0, 246, 49, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 249, 6,
		25, -1, 0, 249, 250, 3, 52, 26, 0, 250, 259, 1, 0, 0, 0, 251, 252, 10,
		2, 0, 0, 252, 253, 5, 30, 0, 0, 253, 258, 3, 52, 26, 0, 254, 255, 10, 1,
		0, 0, 255, 256, 5, 33, 0, 0, 256, 258, 3, 52, 26, 0, 257, 251, 1, 0, 0,
		0, 257, 254, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259,
		260, 1, 0, 0, 0, 260, 51, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 263, 6,
		26, -1, 0, 263, 264, 3, 54, 27, 0, 264, 282, 1, 0, 0, 0, 265, 266, 10,
		5, 0, 0, 266, 267, 5, 26, 0, 0, 267, 281, 3, 54, 27, 0, 268, 269, 10, 4,
		0, 0, 269, 270, 5, 25, 0, 0, 270, 281, 3, 54, 27, 0, 271, 272, 10, 3, 0,
		0, 272, 273, 5, 31, 0, 0, 273, 281, 3, 54, 27, 0, 274, 275, 10, 2, 0, 0,
		275, 276, 5, 32, 0, 0, 276, 281, 3, 54, 27, 0, 277, 278, 10, 1, 0, 0, 278,
		279, 5, 6, 0, 0, 279, 281, 3, 54, 27, 0, 280, 265, 1, 0, 0, 0, 280, 268,
		1, 0, 0, 0, 280, 271, 1, 0, 0, 0, 280, 274, 1, 0, 0, 0, 280, 277, 1, 0,
		0, 0, 281, 284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0,
		283, 53, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 285, 286, 6, 27, -1, 0, 286,
		287, 3, 56, 28, 0, 287, 296, 1, 0, 0, 0, 288, 289, 10, 2, 0, 0, 289, 290,
		5, 45, 0, 0, 290, 295, 3, 56, 28, 0, 291, 292, 10, 1, 0, 0, 292, 293, 5,
		46, 0, 0, 293, 295, 3, 56, 28, 0, 294, 288, 1, 0, 0, 0, 294, 291, 1, 0,
		0, 0, 295, 298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0,
		297, 55, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 299, 300, 6, 28, -1, 0, 300,
		301, 3, 58, 29, 0, 301, 310, 1, 0, 0, 0, 302, 303, 10, 2, 0, 0, 303, 304,
		5, 36, 0, 0, 304, 309, 3, 58, 29, 0, 305, 306, 10, 1, 0, 0, 306, 307, 5,
		37, 0, 0, 307, 309, 3, 58, 29, 0, 308, 302, 1, 0, 0, 0, 308, 305, 1, 0,
		0, 0, 309, 312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0,
		311, 57, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 314, 6, 29, -1, 0, 314,
		315, 3, 60, 30, 0, 315, 327, 1, 0, 0, 0, 316, 317, 10, 3, 0, 0, 317, 318,
		5, 38, 0, 0, 318, 326, 3, 60, 30, 0, 319, 320, 10, 2, 0, 0, 320, 321, 5,
		39, 0, 0, 321, 326, 3, 60, 30, 0, 322, 323, 10, 1, 0, 0, 323, 324, 5, 43,
		0, 0, 324, 326, 3, 60, 30, 0, 325, 316, 1, 0, 0, 0, 325, 319, 1, 0, 0,
		0, 325, 322, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 327,
		328, 1, 0, 0, 0, 328, 59, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 330, 331, 5,
		36, 0, 0, 331, 336, 3, 60, 30, 0, 332, 333, 5, 37, 0, 0, 333, 336, 3, 60,
		30, 0, 334, 336, 3, 62, 31, 0, 335, 330, 1, 0, 0, 0, 335, 332, 1, 0, 0,
		0, 335, 334, 1, 0, 0, 0, 336, 61, 1, 0, 0, 0, 337, 343, 3, 64, 32, 0, 338,
		339, 5, 28, 0, 0, 339, 343, 3, 60, 30, 0, 340, 341, 5, 27, 0, 0, 341, 343,
		3, 60, 30, 0, 342, 337, 1, 0, 0, 0, 342, 338, 1, 0, 0, 0, 342, 340, 1,
		0, 0, 0, 343, 63, 1, 0, 0, 0, 344, 345, 6, 32, -1, 0, 345, 346, 3, 66,
		33, 0, 346, 371, 1, 0, 0, 0, 347, 348, 10, 4, 0, 0, 348, 349, 5, 18, 0,
		0, 349, 350, 3, 32, 16, 0, 350, 351, 5, 19, 0, 0, 351, 370, 1, 0, 0, 0,
		352, 353, 10, 3, 0, 0, 353, 354, 5, 22, 0, 0, 354, 370, 5, 48, 0, 0, 355,
		356, 10, 2, 0, 0, 356, 357, 5, 22, 0, 0, 357, 358, 5, 48, 0, 0, 358, 360,
		5, 14, 0, 0, 359, 361, 3, 68, 34, 0, 360, 359, 1, 0, 0, 0, 360, 361, 1,
		0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 370, 5, 15, 0, 0, 363, 364, 10, 1,
		0, 0, 364, 366, 5, 14, 0, 0, 365, 367, 3, 68, 34, 0, 366, 365, 1, 0, 0,
		0, 366, 367, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 370, 5, 15, 0, 0, 369,
		347, 1, 0, 0, 0, 369, 352, 1, 0, 0, 0, 369, 355, 1, 0, 0, 0, 369, 363,
		1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 372, 1, 0,
		0, 0, 372, 65, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 385, 3, 30, 15, 0,
		375, 385, 5, 47, 0, 0, 376, 377, 5, 14, 0, 0, 377, 378, 3, 32, 16, 0, 378,
		379, 5, 15, 0, 0, 379, 385, 1, 0, 0, 0, 380, 385, 5, 48, 0, 0, 381, 385,
		3, 78, 39, 0, 382, 385, 3, 80, 40, 0, 383, 385, 3, 88, 44, 0, 384, 374,
		1, 0, 0, 0, 384, 375, 1, 0, 0, 0, 384, 376, 1, 0, 0, 0, 384, 380, 1, 0,
		0, 0, 384, 381, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 383, 1, 0, 0, 0,
		385, 67, 1, 0, 0, 0, 386, 391, 3, 32, 16, 0, 387, 388, 5, 21, 0, 0, 388,
		390, 3, 32, 16, 0, 389, 387, 1, 0, 0, 0, 390, 393, 1, 0, 0, 0, 391, 389,
		1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 69, 1, 0, 0, 0, 393, 391, 1, 0,
		0, 0, 394, 395, 3, 72, 36, 0, 395, 396, 5, 44, 0, 0, 396, 397, 3, 76, 38,
		0, 397, 71, 1, 0, 0, 0, 398, 406, 5, 48, 0, 0, 399, 400, 5, 14, 0, 0, 400,
		406, 5, 15, 0, 0, 401, 402, 5, 14, 0, 0, 402, 403, 3, 74, 37, 0, 403, 404,
		5, 15, 0, 0, 404, 406, 1, 0, 0, 0, 405, 398, 1, 0, 0, 0, 405, 399, 1, 0,
		0, 0, 405, 401, 1, 0, 0, 0, 406, 73, 1, 0, 0, 0, 407, 412, 5, 48, 0, 0,
		408, 409, 5, 21, 0, 0, 409, 411, 5, 48, 0, 0, 410, 408, 1, 0, 0, 0, 411,
		414, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 75, 1,
		0, 0, 0, 414, 412, 1, 0, 0, 0, 415, 418, 3, 88, 44, 0, 416, 418, 3, 32,
		16, 0, 417, 415, 1, 0, 0, 0, 417, 416, 1, 0, 0, 0, 418, 77, 1, 0, 0, 0,
		419, 421, 5, 18, 0, 0, 420, 422, 3, 86, 43, 0, 421, 420, 1, 0, 0, 0, 421,
		422, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 5, 19, 0, 0, 424, 79,
		1, 0, 0, 0, 425, 427, 5, 16, 0, 0, 426, 428, 3, 82, 41, 0, 427, 426, 1,
		0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 430, 5, 17, 0,
		0, 430, 81, 1, 0, 0, 0, 431, 436, 3, 84, 42, 0, 432, 433, 5, 21, 0, 0,
		433, 435, 3, 84, 42, 0, 434, 432, 1, 0, 0, 0, 435, 438, 1, 0, 0, 0, 436,
		434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 83, 1, 0, 0, 0, 438, 436, 1,
		0, 0, 0, 439, 440, 3, 32, 16, 0, 440, 441, 5, 23, 0, 0, 441, 442, 3, 32,
		16, 0, 442, 85, 1, 0, 0, 0, 443, 448, 3, 32, 16, 0, 444, 445, 5, 21, 0,
		0, 445, 447, 3, 32, 16, 0, 446, 444, 1, 0, 0, 0, 447, 450, 1, 0, 0, 0,
		448, 446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 87, 1, 0, 0, 0, 450, 448,
		1, 0, 0, 0, 451, 455, 5, 16, 0, 0, 452, 454, 3, 6, 3, 0, 453, 452, 1, 0,
		0, 0, 454, 457, 1, 0, 0, 0, 455, 453, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0,
		456, 458, 1, 0, 0, 0, 457, 455, 1, 0, 0, 0, 458, 459, 3, 32, 16, 0, 459,
		460, 5, 17, 0, 0, 460, 89, 1, 0, 0, 0, 41, 93, 100, 107, 111, 124, 141,
		153, 159, 178, 191, 201, 212, 223, 234, 245, 257, 259, 280, 282, 294, 296,
		308, 310, 325, 327, 335, 342, 360, 366, 369, 371, 384, 391, 405, 412, 417,
		421, 427, 436, 448, 455,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RuleExprParserInit initializes any static state used to implement RuleExprParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRuleExprParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RuleExprParserInit() {
	staticData := &RuleExprParserParserStaticData
	staticData.once.Do(ruleexprparserParserInit)
}

// NewRuleExprParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRuleExprParser(input antlr.TokenStream) *RuleExprParser {
	RuleExprParserInit()
	this := new(RuleExprParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RuleExprParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "RuleExprParser.g4"

	return this
}

// RuleExprParser tokens.
const (
	RuleExprParserEOF                  = antlr.TokenEOF
	RuleExprParserBREAK                = 1
	RuleExprParserCONTINUE             = 2
	RuleExprParserELSE                 = 3
	RuleExprParserFOR                  = 4
	RuleExprParserIF                   = 5
	RuleExprParserIN                   = 6
	RuleExprParserVAR                  = 7
	RuleExprParserIntegerLiteral       = 8
	RuleExprParserFloatingPointLiteral = 9
	RuleExprParserBooleanLiteral       = 10
	RuleExprParserCharacterLiteral     = 11
	RuleExprParserStringLiteral        = 12
	RuleExprParserNullLiteral          = 13
	RuleExprParserLPAREN               = 14
	RuleExprParserRPAREN               = 15
	RuleExprParserLBRACE               = 16
	RuleExprParserRBRACE               = 17
	RuleExprParserLBRACK               = 18
	RuleExprParserRBRACK               = 19
	RuleExprParserSEMI                 = 20
	RuleExprParserCOMMA                = 21
	RuleExprParserDOT                  = 22
	RuleExprParserCOLON                = 23
	RuleExprParserASSIGN               = 24
	RuleExprParserGT                   = 25
	RuleExprParserLT                   = 26
	RuleExprParserBANG                 = 27
	RuleExprParserTILDE                = 28
	RuleExprParserQUESTION             = 29
	RuleExprParserEQUAL                = 30
	RuleExprParserLE                   = 31
	RuleExprParserGE                   = 32
	RuleExprParserNOTEQUAL             = 33
	RuleExprParserAND                  = 34
	RuleExprParserOR                   = 35
	RuleExprParserADD                  = 36
	RuleExprParserSUB                  = 37
	RuleExprParserMUL                  = 38
	RuleExprParserDIV                  = 39
	RuleExprParserBITAND               = 40
	RuleExprParserBITOR                = 41
	RuleExprParserCARET                = 42
	RuleExprParserMOD                  = 43
	RuleExprParserARROW                = 44
	RuleExprParserLSHIFT               = 45
	RuleExprParserRSHIFT               = 46
	RuleExprParserPLACEHOLDER_VAR      = 47
	RuleExprParserIdentifier           = 48
	RuleExprParserWS                   = 49
	RuleExprParserCOMMENT              = 50
	RuleExprParserLINE_COMMENT         = 51
)

// RuleExprParser rules.
const (
	RuleExprParserRULE_program                           = 0
	RuleExprParserRULE_block                             = 1
	RuleExprParserRULE_blockStatements                   = 2
	RuleExprParserRULE_blockStatement                    = 3
	RuleExprParserRULE_localVariableDeclarationStatement = 4
	RuleExprParserRULE_localVariableDeclaration          = 5
	RuleExprParserRULE_varVariableDeclaratorList         = 6
	RuleExprParserRULE_varVariableDeclarator             = 7
	RuleExprParserRULE_variableInitializer               = 8
	RuleExprParserRULE_statement                         = 9
	RuleExprParserRULE_expressionStatement               = 10
	RuleExprParserRULE_ifStatement                       = 11
	RuleExprParserRULE_forStatement                      = 12
	RuleExprParserRULE_breakStatement                    = 13
	RuleExprParserRULE_continueStatement                 = 14
	RuleExprParserRULE_literal                           = 15
	RuleExprParserRULE_expression                        = 16
	RuleExprParserRULE_assignmentExpression              = 17
	RuleExprParserRULE_assignment                        = 18
	RuleExprParserRULE_conditionalExpression             = 19
	RuleExprParserRULE_conditionalOrExpression           = 20
	RuleExprParserRULE_conditionalAndExpression          = 21
	RuleExprParserRULE_inclusiveOrExpression             = 22
	RuleExprParserRULE_exclusiveOrExpression             = 23
	RuleExprParserRULE_andExpression                     = 24
	RuleExprParserRULE_equalityExpression                = 25
	RuleExprParserRULE_relationalExpression              = 26
	RuleExprParserRULE_shiftExpression                   = 27
	RuleExprParserRULE_additiveExpression                = 28
	RuleExprParserRULE_multiplicativeExpression          = 29
	RuleExprParserRULE_unaryExpression                   = 30
	RuleExprParserRULE_unaryExpressionNotPlusMinus       = 31
	RuleExprParserRULE_postfixExpression                 = 32
	RuleExprParserRULE_primary                           = 33
	RuleExprParserRULE_argumentList                      = 34
	RuleExprParserRULE_lambdaExpression                  = 35
	RuleExprParserRULE_lambdaParameters                  = 36
	RuleExprParserRULE_formalParameterList               = 37
	RuleExprParserRULE_lambdaBody                        = 38
	RuleExprParserRULE_listLiteral                       = 39
	RuleExprParserRULE_mapLiteral                        = 40
	RuleExprParserRULE_mapEntryList                      = 41
	RuleExprParserRULE_mapEntry                          = 42
	RuleExprParserRULE_expressionList                    = 43
	RuleExprParserRULE_expressionBlock                   = 44
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuleExprParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RuleExprParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422419027558326) != 0 {
		{
			p.SetState(90)
			p.Statement()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(RuleExprParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	BlockStatements() IBlockStatementsContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRBRACE, 0)
}

func (s *BlockContext) BlockStatements() IBlockStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RuleExprParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(RuleExprParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422419027558326) != 0 {
		{
			p.SetState(99)
			p.BlockStatements()
		}

	}
	{
		p.SetState(102)
		p.Match(RuleExprParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockStatementsContext is an interface to support dynamic dispatch.
type IBlockStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBlockStatement() []IBlockStatementContext
	BlockStatement(i int) IBlockStatementContext

	// IsBlockStatementsContext differentiates from other interfaces.
	IsBlockStatementsContext()
}

type BlockStatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementsContext() *BlockStatementsContext {
	var p = new(BlockStatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_blockStatements
	return p
}

func InitEmptyBlockStatementsContext(p *BlockStatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_blockStatements
}

func (*BlockStatementsContext) IsBlockStatementsContext() {}

func NewBlockStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementsContext {
	var p = new(BlockStatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_blockStatements

	return p
}

func (s *BlockStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementsContext) AllBlockStatement() []IBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStatementContext); ok {
			tst[i] = t.(IBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStatementsContext) BlockStatement(i int) IBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *BlockStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterBlockStatements(s)
	}
}

func (s *BlockStatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitBlockStatements(s)
	}
}

func (s *BlockStatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitBlockStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) BlockStatements() (localctx IBlockStatementsContext) {
	localctx = NewBlockStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RuleExprParserRULE_blockStatements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422419027558326) != 0) {
		{
			p.SetState(104)
			p.BlockStatement()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LocalVariableDeclarationStatement() ILocalVariableDeclarationStatementContext
	Statement() IStatementContext

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_blockStatement
	return p
}

func InitEmptyBlockStatementContext(p *BlockStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_blockStatement
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) LocalVariableDeclarationStatement() ILocalVariableDeclarationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalVariableDeclarationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalVariableDeclarationStatementContext)
}

func (s *BlockStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RuleExprParserRULE_blockStatement)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.LocalVariableDeclarationStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Statement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILocalVariableDeclarationStatementContext is an interface to support dynamic dispatch.
type ILocalVariableDeclarationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LocalVariableDeclaration() ILocalVariableDeclarationContext
	SEMI() antlr.TerminalNode

	// IsLocalVariableDeclarationStatementContext differentiates from other interfaces.
	IsLocalVariableDeclarationStatementContext()
}

type LocalVariableDeclarationStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalVariableDeclarationStatementContext() *LocalVariableDeclarationStatementContext {
	var p = new(LocalVariableDeclarationStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_localVariableDeclarationStatement
	return p
}

func InitEmptyLocalVariableDeclarationStatementContext(p *LocalVariableDeclarationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_localVariableDeclarationStatement
}

func (*LocalVariableDeclarationStatementContext) IsLocalVariableDeclarationStatementContext() {}

func NewLocalVariableDeclarationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalVariableDeclarationStatementContext {
	var p = new(LocalVariableDeclarationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_localVariableDeclarationStatement

	return p
}

func (s *LocalVariableDeclarationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalVariableDeclarationStatementContext) LocalVariableDeclaration() ILocalVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalVariableDeclarationContext)
}

func (s *LocalVariableDeclarationStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RuleExprParserSEMI, 0)
}

func (s *LocalVariableDeclarationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalVariableDeclarationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalVariableDeclarationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterLocalVariableDeclarationStatement(s)
	}
}

func (s *LocalVariableDeclarationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitLocalVariableDeclarationStatement(s)
	}
}

func (s *LocalVariableDeclarationStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitLocalVariableDeclarationStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) LocalVariableDeclarationStatement() (localctx ILocalVariableDeclarationStatementContext) {
	localctx = NewLocalVariableDeclarationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RuleExprParserRULE_localVariableDeclarationStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.LocalVariableDeclaration()
	}
	{
		p.SetState(114)
		p.Match(RuleExprParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILocalVariableDeclarationContext is an interface to support dynamic dispatch.
type ILocalVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	VarVariableDeclaratorList() IVarVariableDeclaratorListContext

	// IsLocalVariableDeclarationContext differentiates from other interfaces.
	IsLocalVariableDeclarationContext()
}

type LocalVariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalVariableDeclarationContext() *LocalVariableDeclarationContext {
	var p = new(LocalVariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_localVariableDeclaration
	return p
}

func InitEmptyLocalVariableDeclarationContext(p *LocalVariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_localVariableDeclaration
}

func (*LocalVariableDeclarationContext) IsLocalVariableDeclarationContext() {}

func NewLocalVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalVariableDeclarationContext {
	var p = new(LocalVariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_localVariableDeclaration

	return p
}

func (s *LocalVariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalVariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(RuleExprParserVAR, 0)
}

func (s *LocalVariableDeclarationContext) VarVariableDeclaratorList() IVarVariableDeclaratorListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarVariableDeclaratorListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarVariableDeclaratorListContext)
}

func (s *LocalVariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalVariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalVariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterLocalVariableDeclaration(s)
	}
}

func (s *LocalVariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitLocalVariableDeclaration(s)
	}
}

func (s *LocalVariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitLocalVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) LocalVariableDeclaration() (localctx ILocalVariableDeclarationContext) {
	localctx = NewLocalVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RuleExprParserRULE_localVariableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(RuleExprParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.VarVariableDeclaratorList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarVariableDeclaratorListContext is an interface to support dynamic dispatch.
type IVarVariableDeclaratorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVarVariableDeclarator() []IVarVariableDeclaratorContext
	VarVariableDeclarator(i int) IVarVariableDeclaratorContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVarVariableDeclaratorListContext differentiates from other interfaces.
	IsVarVariableDeclaratorListContext()
}

type VarVariableDeclaratorListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarVariableDeclaratorListContext() *VarVariableDeclaratorListContext {
	var p = new(VarVariableDeclaratorListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_varVariableDeclaratorList
	return p
}

func InitEmptyVarVariableDeclaratorListContext(p *VarVariableDeclaratorListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_varVariableDeclaratorList
}

func (*VarVariableDeclaratorListContext) IsVarVariableDeclaratorListContext() {}

func NewVarVariableDeclaratorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarVariableDeclaratorListContext {
	var p = new(VarVariableDeclaratorListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_varVariableDeclaratorList

	return p
}

func (s *VarVariableDeclaratorListContext) GetParser() antlr.Parser { return s.parser }

func (s *VarVariableDeclaratorListContext) AllVarVariableDeclarator() []IVarVariableDeclaratorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarVariableDeclaratorContext); ok {
			len++
		}
	}

	tst := make([]IVarVariableDeclaratorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarVariableDeclaratorContext); ok {
			tst[i] = t.(IVarVariableDeclaratorContext)
			i++
		}
	}

	return tst
}

func (s *VarVariableDeclaratorListContext) VarVariableDeclarator(i int) IVarVariableDeclaratorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarVariableDeclaratorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarVariableDeclaratorContext)
}

func (s *VarVariableDeclaratorListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RuleExprParserCOMMA)
}

func (s *VarVariableDeclaratorListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RuleExprParserCOMMA, i)
}

func (s *VarVariableDeclaratorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarVariableDeclaratorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarVariableDeclaratorListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterVarVariableDeclaratorList(s)
	}
}

func (s *VarVariableDeclaratorListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitVarVariableDeclaratorList(s)
	}
}

func (s *VarVariableDeclaratorListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitVarVariableDeclaratorList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) VarVariableDeclaratorList() (localctx IVarVariableDeclaratorListContext) {
	localctx = NewVarVariableDeclaratorListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RuleExprParserRULE_varVariableDeclaratorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.VarVariableDeclarator()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RuleExprParserCOMMA {
		{
			p.SetState(120)
			p.Match(RuleExprParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.VarVariableDeclarator()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVarVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	VariableInitializer() IVariableInitializerContext

	// IsVarVariableDeclaratorContext differentiates from other interfaces.
	IsVarVariableDeclaratorContext()
}

type VarVariableDeclaratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarVariableDeclaratorContext() *VarVariableDeclaratorContext {
	var p = new(VarVariableDeclaratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_varVariableDeclarator
	return p
}

func InitEmptyVarVariableDeclaratorContext(p *VarVariableDeclaratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_varVariableDeclarator
}

func (*VarVariableDeclaratorContext) IsVarVariableDeclaratorContext() {}

func NewVarVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarVariableDeclaratorContext {
	var p = new(VarVariableDeclaratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_varVariableDeclarator

	return p
}

func (s *VarVariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VarVariableDeclaratorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIdentifier, 0)
}

func (s *VarVariableDeclaratorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserASSIGN, 0)
}

func (s *VarVariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VarVariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarVariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarVariableDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterVarVariableDeclarator(s)
	}
}

func (s *VarVariableDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitVarVariableDeclarator(s)
	}
}

func (s *VarVariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitVarVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) VarVariableDeclarator() (localctx IVarVariableDeclaratorContext) {
	localctx = NewVarVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RuleExprParserRULE_varVariableDeclarator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(RuleExprParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(RuleExprParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.VariableInitializer()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_variableInitializer
	return p
}

func InitEmptyVariableInitializerContext(p *VariableInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_variableInitializer
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RuleExprParserRULE_variableInitializer)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	IfStatement() IIfStatementContext
	ForStatement() IForStatementContext
	BreakStatement() IBreakStatementContext
	ContinueStatement() IContinueStatementContext
	ExpressionStatement() IExpressionStatementContext
	LocalVariableDeclarationStatement() ILocalVariableDeclarationStatementContext
	SEMI() antlr.TerminalNode

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) LocalVariableDeclarationStatement() ILocalVariableDeclarationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalVariableDeclarationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalVariableDeclarationStatementContext)
}

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RuleExprParserSEMI, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RuleExprParserRULE_statement)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
			p.IfStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(135)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.BreakStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(137)
			p.ContinueStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(138)
			p.ExpressionStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(139)
			p.LocalVariableDeclarationStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(140)
			p.Match(RuleExprParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	SEMI() antlr.TerminalNode

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_expressionStatement
	return p
}

func InitEmptyExpressionStatementContext(p *ExpressionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_expressionStatement
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RuleExprParserSEMI, 0)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RuleExprParserRULE_expressionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Expression()
	}
	{
		p.SetState(144)
		p.Match(RuleExprParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	ELSE() antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIF, 0)
}

func (s *IfStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLPAREN, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRPAREN, 0)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RuleExprParserRULE_ifStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(RuleExprParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(RuleExprParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Expression()
	}
	{
		p.SetState(149)
		p.Match(RuleExprParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Statement()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(151)
			p.Match(RuleExprParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Statement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	IN() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext
	COMMA() antlr.TerminalNode

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(RuleExprParserFOR, 0)
}

func (s *ForStatementContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(RuleExprParserIdentifier)
}

func (s *ForStatementContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(RuleExprParserIdentifier, i)
}

func (s *ForStatementContext) IN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIN, 0)
}

func (s *ForStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(RuleExprParserCOMMA, 0)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RuleExprParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(RuleExprParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(RuleExprParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RuleExprParserCOMMA {
		{
			p.SetState(157)
			p.Match(RuleExprParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(RuleExprParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(161)
		p.Match(RuleExprParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Expression()
	}
	{
		p.SetState(163)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_breakStatement
	return p
}

func InitEmptyBreakStatementContext(p *BreakStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_breakStatement
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(RuleExprParserBREAK, 0)
}

func (s *BreakStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RuleExprParserSEMI, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RuleExprParserRULE_breakStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(RuleExprParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(RuleExprParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_continueStatement
	return p
}

func InitEmptyContinueStatementContext(p *ContinueStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_continueStatement
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserCONTINUE, 0)
}

func (s *ContinueStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RuleExprParserSEMI, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RuleExprParserRULE_continueStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(RuleExprParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(RuleExprParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerLiteral() antlr.TerminalNode
	FloatingPointLiteral() antlr.TerminalNode
	BooleanLiteral() antlr.TerminalNode
	CharacterLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	NullLiteral() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIntegerLiteral, 0)
}

func (s *LiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(RuleExprParserFloatingPointLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(RuleExprParserBooleanLiteral, 0)
}

func (s *LiteralContext) CharacterLiteral() antlr.TerminalNode {
	return s.GetToken(RuleExprParserCharacterLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(RuleExprParserStringLiteral, 0)
}

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(RuleExprParserNullLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RuleExprParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16128) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignmentExpression() IAssignmentExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RuleExprParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.AssignmentExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LambdaExpression() ILambdaExpressionContext
	ConditionalExpression() IConditionalExpressionContext
	Assignment() IAssignmentContext

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_assignmentExpression
	return p
}

func InitEmptyAssignmentExpressionContext(p *AssignmentExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_assignmentExpression
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) LambdaExpression() ILambdaExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaExpressionContext)
}

func (s *AssignmentExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *AssignmentExpressionContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitAssignmentExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RuleExprParserRULE_assignmentExpression)
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.LambdaExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.ConditionalExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.Assignment()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIdentifier, 0)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RuleExprParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(RuleExprParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(RuleExprParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalExpressionContext is an interface to support dynamic dispatch.
type IConditionalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalOrExpression() IConditionalOrExpressionContext
	QUESTION() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	ConditionalExpression() IConditionalExpressionContext

	// IsConditionalExpressionContext differentiates from other interfaces.
	IsConditionalExpressionContext()
}

type ConditionalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalExpressionContext() *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_conditionalExpression
	return p
}

func InitEmptyConditionalExpressionContext(p *ConditionalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_conditionalExpression
}

func (*ConditionalExpressionContext) IsConditionalExpressionContext() {}

func NewConditionalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_conditionalExpression

	return p
}

func (s *ConditionalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExpressionContext) ConditionalOrExpression() IConditionalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalOrExpressionContext)
}

func (s *ConditionalExpressionContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(RuleExprParserQUESTION, 0)
}

func (s *ConditionalExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(RuleExprParserCOLON, 0)
}

func (s *ConditionalExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *ConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitConditionalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ConditionalExpression() (localctx IConditionalExpressionContext) {
	localctx = NewConditionalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RuleExprParserRULE_conditionalExpression)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.conditionalOrExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.conditionalOrExpression(0)
		}
		{
			p.SetState(186)
			p.Match(RuleExprParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Expression()
		}
		{
			p.SetState(188)
			p.Match(RuleExprParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.ConditionalExpression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalOrExpressionContext is an interface to support dynamic dispatch.
type IConditionalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalAndExpression() IConditionalAndExpressionContext
	ConditionalOrExpression() IConditionalOrExpressionContext
	OR() antlr.TerminalNode

	// IsConditionalOrExpressionContext differentiates from other interfaces.
	IsConditionalOrExpressionContext()
}

type ConditionalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalOrExpressionContext() *ConditionalOrExpressionContext {
	var p = new(ConditionalOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_conditionalOrExpression
	return p
}

func InitEmptyConditionalOrExpressionContext(p *ConditionalOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_conditionalOrExpression
}

func (*ConditionalOrExpressionContext) IsConditionalOrExpressionContext() {}

func NewConditionalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalOrExpressionContext {
	var p = new(ConditionalOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_conditionalOrExpression

	return p
}

func (s *ConditionalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalOrExpressionContext) ConditionalAndExpression() IConditionalAndExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalAndExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalAndExpressionContext)
}

func (s *ConditionalOrExpressionContext) ConditionalOrExpression() IConditionalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalOrExpressionContext)
}

func (s *ConditionalOrExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(RuleExprParserOR, 0)
}

func (s *ConditionalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterConditionalOrExpression(s)
	}
}

func (s *ConditionalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitConditionalOrExpression(s)
	}
}

func (s *ConditionalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitConditionalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ConditionalOrExpression() (localctx IConditionalOrExpressionContext) {
	return p.conditionalOrExpression(0)
}

func (p *RuleExprParser) conditionalOrExpression(_p int) (localctx IConditionalOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewConditionalOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionalOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, RuleExprParserRULE_conditionalOrExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.conditionalAndExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConditionalOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_conditionalOrExpression)
			p.SetState(196)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(197)
				p.Match(RuleExprParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(198)
				p.conditionalAndExpression(0)
			}

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalAndExpressionContext is an interface to support dynamic dispatch.
type IConditionalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InclusiveOrExpression() IInclusiveOrExpressionContext
	ConditionalAndExpression() IConditionalAndExpressionContext
	AND() antlr.TerminalNode

	// IsConditionalAndExpressionContext differentiates from other interfaces.
	IsConditionalAndExpressionContext()
}

type ConditionalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalAndExpressionContext() *ConditionalAndExpressionContext {
	var p = new(ConditionalAndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_conditionalAndExpression
	return p
}

func InitEmptyConditionalAndExpressionContext(p *ConditionalAndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_conditionalAndExpression
}

func (*ConditionalAndExpressionContext) IsConditionalAndExpressionContext() {}

func NewConditionalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalAndExpressionContext {
	var p = new(ConditionalAndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_conditionalAndExpression

	return p
}

func (s *ConditionalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalAndExpressionContext) InclusiveOrExpression() IInclusiveOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInclusiveOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInclusiveOrExpressionContext)
}

func (s *ConditionalAndExpressionContext) ConditionalAndExpression() IConditionalAndExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalAndExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalAndExpressionContext)
}

func (s *ConditionalAndExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(RuleExprParserAND, 0)
}

func (s *ConditionalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterConditionalAndExpression(s)
	}
}

func (s *ConditionalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitConditionalAndExpression(s)
	}
}

func (s *ConditionalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitConditionalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ConditionalAndExpression() (localctx IConditionalAndExpressionContext) {
	return p.conditionalAndExpression(0)
}

func (p *RuleExprParser) conditionalAndExpression(_p int) (localctx IConditionalAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewConditionalAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionalAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, RuleExprParserRULE_conditionalAndExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.inclusiveOrExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConditionalAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_conditionalAndExpression)
			p.SetState(207)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(208)
				p.Match(RuleExprParserAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(209)
				p.inclusiveOrExpression(0)
			}

		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInclusiveOrExpressionContext is an interface to support dynamic dispatch.
type IInclusiveOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExclusiveOrExpression() IExclusiveOrExpressionContext
	InclusiveOrExpression() IInclusiveOrExpressionContext
	BITOR() antlr.TerminalNode

	// IsInclusiveOrExpressionContext differentiates from other interfaces.
	IsInclusiveOrExpressionContext()
}

type InclusiveOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclusiveOrExpressionContext() *InclusiveOrExpressionContext {
	var p = new(InclusiveOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_inclusiveOrExpression
	return p
}

func InitEmptyInclusiveOrExpressionContext(p *InclusiveOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_inclusiveOrExpression
}

func (*InclusiveOrExpressionContext) IsInclusiveOrExpressionContext() {}

func NewInclusiveOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InclusiveOrExpressionContext {
	var p = new(InclusiveOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_inclusiveOrExpression

	return p
}

func (s *InclusiveOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InclusiveOrExpressionContext) ExclusiveOrExpression() IExclusiveOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExclusiveOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExclusiveOrExpressionContext)
}

func (s *InclusiveOrExpressionContext) InclusiveOrExpression() IInclusiveOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInclusiveOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInclusiveOrExpressionContext)
}

func (s *InclusiveOrExpressionContext) BITOR() antlr.TerminalNode {
	return s.GetToken(RuleExprParserBITOR, 0)
}

func (s *InclusiveOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InclusiveOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InclusiveOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterInclusiveOrExpression(s)
	}
}

func (s *InclusiveOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitInclusiveOrExpression(s)
	}
}

func (s *InclusiveOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitInclusiveOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) InclusiveOrExpression() (localctx IInclusiveOrExpressionContext) {
	return p.inclusiveOrExpression(0)
}

func (p *RuleExprParser) inclusiveOrExpression(_p int) (localctx IInclusiveOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewInclusiveOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IInclusiveOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, RuleExprParserRULE_inclusiveOrExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.exclusiveOrExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewInclusiveOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_inclusiveOrExpression)
			p.SetState(218)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(219)
				p.Match(RuleExprParserBITOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(220)
				p.exclusiveOrExpression(0)
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExclusiveOrExpressionContext is an interface to support dynamic dispatch.
type IExclusiveOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AndExpression() IAndExpressionContext
	ExclusiveOrExpression() IExclusiveOrExpressionContext
	CARET() antlr.TerminalNode

	// IsExclusiveOrExpressionContext differentiates from other interfaces.
	IsExclusiveOrExpressionContext()
}

type ExclusiveOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExclusiveOrExpressionContext() *ExclusiveOrExpressionContext {
	var p = new(ExclusiveOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_exclusiveOrExpression
	return p
}

func InitEmptyExclusiveOrExpressionContext(p *ExclusiveOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_exclusiveOrExpression
}

func (*ExclusiveOrExpressionContext) IsExclusiveOrExpressionContext() {}

func NewExclusiveOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusiveOrExpressionContext {
	var p = new(ExclusiveOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_exclusiveOrExpression

	return p
}

func (s *ExclusiveOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExclusiveOrExpressionContext) AndExpression() IAndExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *ExclusiveOrExpressionContext) ExclusiveOrExpression() IExclusiveOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExclusiveOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExclusiveOrExpressionContext)
}

func (s *ExclusiveOrExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(RuleExprParserCARET, 0)
}

func (s *ExclusiveOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusiveOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExclusiveOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterExclusiveOrExpression(s)
	}
}

func (s *ExclusiveOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitExclusiveOrExpression(s)
	}
}

func (s *ExclusiveOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitExclusiveOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ExclusiveOrExpression() (localctx IExclusiveOrExpressionContext) {
	return p.exclusiveOrExpression(0)
}

func (p *RuleExprParser) exclusiveOrExpression(_p int) (localctx IExclusiveOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExclusiveOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExclusiveOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, RuleExprParserRULE_exclusiveOrExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.andExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExclusiveOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_exclusiveOrExpression)
			p.SetState(229)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(230)
				p.Match(RuleExprParserCARET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(231)
				p.andExpression(0)
			}

		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualityExpression() IEqualityExpressionContext
	AndExpression() IAndExpressionContext
	BITAND() antlr.TerminalNode

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_andExpression
	return p
}

func InitEmptyAndExpressionContext(p *AndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_andExpression
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *AndExpressionContext) AndExpression() IAndExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *AndExpressionContext) BITAND() antlr.TerminalNode {
	return s.GetToken(RuleExprParserBITAND, 0)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) AndExpression() (localctx IAndExpressionContext) {
	return p.andExpression(0)
}

func (p *RuleExprParser) andExpression(_p int) (localctx IAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, RuleExprParserRULE_andExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.equalityExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_andExpression)
			p.SetState(240)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(241)
				p.Match(RuleExprParserBITAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(242)
				p.equalityExpression(0)
			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationalExpression() IRelationalExpressionContext
	EqualityExpression() IEqualityExpressionContext
	EQUAL() antlr.TerminalNode
	NOTEQUAL() antlr.TerminalNode

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_equalityExpression
	return p
}

func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_equalityExpression
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *EqualityExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(RuleExprParserEQUAL, 0)
}

func (s *EqualityExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(RuleExprParserNOTEQUAL, 0)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	return p.equalityExpression(0)
}

func (p *RuleExprParser) equalityExpression(_p int) (localctx IEqualityExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqualityExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, RuleExprParserRULE_equalityExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.relationalExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqualityExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_equalityExpression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(252)
					p.Match(RuleExprParserEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(253)
					p.relationalExpression(0)
				}

			case 2:
				localctx = NewEqualityExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_equalityExpression)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(255)
					p.Match(RuleExprParserNOTEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(256)
					p.relationalExpression(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ShiftExpression() IShiftExpressionContext
	RelationalExpression() IRelationalExpressionContext
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode
	IN() antlr.TerminalNode

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_relationalExpression
	return p
}

func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_relationalExpression
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) ShiftExpression() IShiftExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftExpressionContext)
}

func (s *RelationalExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *RelationalExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLT, 0)
}

func (s *RelationalExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(RuleExprParserGT, 0)
}

func (s *RelationalExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLE, 0)
}

func (s *RelationalExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserGE, 0)
}

func (s *RelationalExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIN, 0)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	return p.relationalExpression(0)
}

func (p *RuleExprParser) relationalExpression(_p int) (localctx IRelationalExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationalExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, RuleExprParserRULE_relationalExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.shiftExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_relationalExpression)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(266)
					p.Match(RuleExprParserLT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(267)
					p.shiftExpression(0)
				}

			case 2:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_relationalExpression)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(269)
					p.Match(RuleExprParserGT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(270)
					p.shiftExpression(0)
				}

			case 3:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_relationalExpression)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(272)
					p.Match(RuleExprParserLE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(273)
					p.shiftExpression(0)
				}

			case 4:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_relationalExpression)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(275)
					p.Match(RuleExprParserGE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(276)
					p.shiftExpression(0)
				}

			case 5:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_relationalExpression)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(278)
					p.Match(RuleExprParserIN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(279)
					p.shiftExpression(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShiftExpressionContext is an interface to support dynamic dispatch.
type IShiftExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AdditiveExpression() IAdditiveExpressionContext
	ShiftExpression() IShiftExpressionContext
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode

	// IsShiftExpressionContext differentiates from other interfaces.
	IsShiftExpressionContext()
}

type ShiftExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftExpressionContext() *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_shiftExpression
	return p
}

func InitEmptyShiftExpressionContext(p *ShiftExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_shiftExpression
}

func (*ShiftExpressionContext) IsShiftExpressionContext() {}

func NewShiftExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_shiftExpression

	return p
}

func (s *ShiftExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *ShiftExpressionContext) ShiftExpression() IShiftExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftExpressionContext)
}

func (s *ShiftExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLSHIFT, 0)
}

func (s *ShiftExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRSHIFT, 0)
}

func (s *ShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterShiftExpression(s)
	}
}

func (s *ShiftExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitShiftExpression(s)
	}
}

func (s *ShiftExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitShiftExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ShiftExpression() (localctx IShiftExpressionContext) {
	return p.shiftExpression(0)
}

func (p *RuleExprParser) shiftExpression(_p int) (localctx IShiftExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewShiftExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IShiftExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, RuleExprParserRULE_shiftExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.additiveExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewShiftExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_shiftExpression)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(289)
					p.Match(RuleExprParserLSHIFT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(290)
					p.additiveExpression(0)
				}

			case 2:
				localctx = NewShiftExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_shiftExpression)
				p.SetState(291)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(292)
					p.Match(RuleExprParserRSHIFT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(293)
					p.additiveExpression(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MultiplicativeExpression() IMultiplicativeExpressionContext
	AdditiveExpression() IAdditiveExpressionContext
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_additiveExpression
	return p
}

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_additiveExpression
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *AdditiveExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(RuleExprParserADD, 0)
}

func (s *AdditiveExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(RuleExprParserSUB, 0)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	return p.additiveExpression(0)
}

func (p *RuleExprParser) additiveExpression(_p int) (localctx IAdditiveExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, RuleExprParserRULE_additiveExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.multiplicativeExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(308)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAdditiveExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_additiveExpression)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(303)
					p.Match(RuleExprParserADD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(304)
					p.multiplicativeExpression(0)
				}

			case 2:
				localctx = NewAdditiveExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_additiveExpression)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(306)
					p.Match(RuleExprParserSUB)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(307)
					p.multiplicativeExpression(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	MultiplicativeExpression() IMultiplicativeExpressionContext
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_multiplicativeExpression
	return p
}

func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_multiplicativeExpression
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *MultiplicativeExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(RuleExprParserMUL, 0)
}

func (s *MultiplicativeExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(RuleExprParserDIV, 0)
}

func (s *MultiplicativeExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(RuleExprParserMOD, 0)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	return p.multiplicativeExpression(0)
}

func (p *RuleExprParser) multiplicativeExpression(_p int) (localctx IMultiplicativeExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicativeExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, RuleExprParserRULE_multiplicativeExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.UnaryExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_multiplicativeExpression)
				p.SetState(316)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(317)
					p.Match(RuleExprParserMUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(318)
					p.UnaryExpression()
				}

			case 2:
				localctx = NewMultiplicativeExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_multiplicativeExpression)
				p.SetState(319)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(320)
					p.Match(RuleExprParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(321)
					p.UnaryExpression()
				}

			case 3:
				localctx = NewMultiplicativeExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_multiplicativeExpression)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(323)
					p.Match(RuleExprParserMOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(324)
					p.UnaryExpression()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	UnaryExpression() IUnaryExpressionContext
	SUB() antlr.TerminalNode
	UnaryExpressionNotPlusMinus() IUnaryExpressionNotPlusMinusContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(RuleExprParserADD, 0)
}

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(RuleExprParserSUB, 0)
}

func (s *UnaryExpressionContext) UnaryExpressionNotPlusMinus() IUnaryExpressionNotPlusMinusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionNotPlusMinusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionNotPlusMinusContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, RuleExprParserRULE_unaryExpression)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RuleExprParserADD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(RuleExprParserADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.UnaryExpression()
		}

	case RuleExprParserSUB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Match(RuleExprParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.UnaryExpression()
		}

	case RuleExprParserIntegerLiteral, RuleExprParserFloatingPointLiteral, RuleExprParserBooleanLiteral, RuleExprParserCharacterLiteral, RuleExprParserStringLiteral, RuleExprParserNullLiteral, RuleExprParserLPAREN, RuleExprParserLBRACE, RuleExprParserLBRACK, RuleExprParserBANG, RuleExprParserTILDE, RuleExprParserPLACEHOLDER_VAR, RuleExprParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)
			p.UnaryExpressionNotPlusMinus()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionNotPlusMinusContext is an interface to support dynamic dispatch.
type IUnaryExpressionNotPlusMinusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PostfixExpression() IPostfixExpressionContext
	TILDE() antlr.TerminalNode
	UnaryExpression() IUnaryExpressionContext
	BANG() antlr.TerminalNode

	// IsUnaryExpressionNotPlusMinusContext differentiates from other interfaces.
	IsUnaryExpressionNotPlusMinusContext()
}

type UnaryExpressionNotPlusMinusContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionNotPlusMinusContext() *UnaryExpressionNotPlusMinusContext {
	var p = new(UnaryExpressionNotPlusMinusContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_unaryExpressionNotPlusMinus
	return p
}

func InitEmptyUnaryExpressionNotPlusMinusContext(p *UnaryExpressionNotPlusMinusContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_unaryExpressionNotPlusMinus
}

func (*UnaryExpressionNotPlusMinusContext) IsUnaryExpressionNotPlusMinusContext() {}

func NewUnaryExpressionNotPlusMinusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionNotPlusMinusContext {
	var p = new(UnaryExpressionNotPlusMinusContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_unaryExpressionNotPlusMinus

	return p
}

func (s *UnaryExpressionNotPlusMinusContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionNotPlusMinusContext) PostfixExpression() IPostfixExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *UnaryExpressionNotPlusMinusContext) TILDE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserTILDE, 0)
}

func (s *UnaryExpressionNotPlusMinusContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionNotPlusMinusContext) BANG() antlr.TerminalNode {
	return s.GetToken(RuleExprParserBANG, 0)
}

func (s *UnaryExpressionNotPlusMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionNotPlusMinusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionNotPlusMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterUnaryExpressionNotPlusMinus(s)
	}
}

func (s *UnaryExpressionNotPlusMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitUnaryExpressionNotPlusMinus(s)
	}
}

func (s *UnaryExpressionNotPlusMinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitUnaryExpressionNotPlusMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) UnaryExpressionNotPlusMinus() (localctx IUnaryExpressionNotPlusMinusContext) {
	localctx = NewUnaryExpressionNotPlusMinusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, RuleExprParserRULE_unaryExpressionNotPlusMinus)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RuleExprParserIntegerLiteral, RuleExprParserFloatingPointLiteral, RuleExprParserBooleanLiteral, RuleExprParserCharacterLiteral, RuleExprParserStringLiteral, RuleExprParserNullLiteral, RuleExprParserLPAREN, RuleExprParserLBRACE, RuleExprParserLBRACK, RuleExprParserPLACEHOLDER_VAR, RuleExprParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(337)
			p.postfixExpression(0)
		}

	case RuleExprParserTILDE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(338)
			p.Match(RuleExprParserTILDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.UnaryExpression()
		}

	case RuleExprParserBANG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(340)
			p.Match(RuleExprParserBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.UnaryExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixExpressionContext is an interface to support dynamic dispatch.
type IPostfixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	PostfixExpression() IPostfixExpressionContext
	LBRACK() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACK() antlr.TerminalNode
	DOT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsPostfixExpressionContext differentiates from other interfaces.
	IsPostfixExpressionContext()
}

type PostfixExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixExpressionContext() *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_postfixExpression
	return p
}

func InitEmptyPostfixExpressionContext(p *PostfixExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_postfixExpression
}

func (*PostfixExpressionContext) IsPostfixExpressionContext() {}

func NewPostfixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_postfixExpression

	return p
}

func (s *PostfixExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixExpressionContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PostfixExpressionContext) PostfixExpression() IPostfixExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *PostfixExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLBRACK, 0)
}

func (s *PostfixExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PostfixExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRBRACK, 0)
}

func (s *PostfixExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(RuleExprParserDOT, 0)
}

func (s *PostfixExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIdentifier, 0)
}

func (s *PostfixExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLPAREN, 0)
}

func (s *PostfixExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRPAREN, 0)
}

func (s *PostfixExpressionContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *PostfixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitPostfixExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) PostfixExpression() (localctx IPostfixExpressionContext) {
	return p.postfixExpression(0)
}

func (p *RuleExprParser) postfixExpression(_p int) (localctx IPostfixExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPostfixExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPostfixExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 64
	p.EnterRecursionRule(localctx, 64, RuleExprParserRULE_postfixExpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Primary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(369)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_postfixExpression)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(348)
					p.Match(RuleExprParserLBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(349)
					p.Expression()
				}
				{
					p.SetState(350)
					p.Match(RuleExprParserRBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_postfixExpression)
				p.SetState(352)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(353)
					p.Match(RuleExprParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(354)
					p.Match(RuleExprParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_postfixExpression)
				p.SetState(355)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(356)
					p.Match(RuleExprParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(357)
					p.Match(RuleExprParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(358)
					p.Match(RuleExprParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(360)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422419026509568) != 0 {
					{
						p.SetState(359)
						p.ArgumentList()
					}

				}
				{
					p.SetState(362)
					p.Match(RuleExprParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 4:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RuleExprParserRULE_postfixExpression)
				p.SetState(363)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(364)
					p.Match(RuleExprParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(366)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422419026509568) != 0 {
					{
						p.SetState(365)
						p.ArgumentList()
					}

				}
				{
					p.SetState(368)
					p.Match(RuleExprParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	PLACEHOLDER_VAR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	ListLiteral() IListLiteralContext
	MapLiteral() IMapLiteralContext
	ExpressionBlock() IExpressionBlockContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) PLACEHOLDER_VAR() antlr.TerminalNode {
	return s.GetToken(RuleExprParserPLACEHOLDER_VAR, 0)
}

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRPAREN, 0)
}

func (s *PrimaryContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIdentifier, 0)
}

func (s *PrimaryContext) ListLiteral() IListLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListLiteralContext)
}

func (s *PrimaryContext) MapLiteral() IMapLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapLiteralContext)
}

func (s *PrimaryContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, RuleExprParserRULE_primary)
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Match(RuleExprParserPLACEHOLDER_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(376)
			p.Match(RuleExprParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Expression()
		}
		{
			p.SetState(378)
			p.Match(RuleExprParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(380)
			p.Match(RuleExprParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(381)
			p.ListLiteral()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(382)
			p.MapLiteral()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(383)
			p.ExpressionBlock()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RuleExprParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RuleExprParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, RuleExprParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Expression()
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RuleExprParserCOMMA {
		{
			p.SetState(387)
			p.Match(RuleExprParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Expression()
		}

		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILambdaExpressionContext is an interface to support dynamic dispatch.
type ILambdaExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LambdaParameters() ILambdaParametersContext
	ARROW() antlr.TerminalNode
	LambdaBody() ILambdaBodyContext

	// IsLambdaExpressionContext differentiates from other interfaces.
	IsLambdaExpressionContext()
}

type LambdaExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaExpressionContext() *LambdaExpressionContext {
	var p = new(LambdaExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_lambdaExpression
	return p
}

func InitEmptyLambdaExpressionContext(p *LambdaExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_lambdaExpression
}

func (*LambdaExpressionContext) IsLambdaExpressionContext() {}

func NewLambdaExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaExpressionContext {
	var p = new(LambdaExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_lambdaExpression

	return p
}

func (s *LambdaExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaExpressionContext) LambdaParameters() ILambdaParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaParametersContext)
}

func (s *LambdaExpressionContext) ARROW() antlr.TerminalNode {
	return s.GetToken(RuleExprParserARROW, 0)
}

func (s *LambdaExpressionContext) LambdaBody() ILambdaBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaBodyContext)
}

func (s *LambdaExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterLambdaExpression(s)
	}
}

func (s *LambdaExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitLambdaExpression(s)
	}
}

func (s *LambdaExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitLambdaExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) LambdaExpression() (localctx ILambdaExpressionContext) {
	localctx = NewLambdaExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, RuleExprParserRULE_lambdaExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.LambdaParameters()
	}
	{
		p.SetState(395)
		p.Match(RuleExprParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.LambdaBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILambdaParametersContext is an interface to support dynamic dispatch.
type ILambdaParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	FormalParameterList() IFormalParameterListContext

	// IsLambdaParametersContext differentiates from other interfaces.
	IsLambdaParametersContext()
}

type LambdaParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaParametersContext() *LambdaParametersContext {
	var p = new(LambdaParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_lambdaParameters
	return p
}

func InitEmptyLambdaParametersContext(p *LambdaParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_lambdaParameters
}

func (*LambdaParametersContext) IsLambdaParametersContext() {}

func NewLambdaParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaParametersContext {
	var p = new(LambdaParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_lambdaParameters

	return p
}

func (s *LambdaParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaParametersContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RuleExprParserIdentifier, 0)
}

func (s *LambdaParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLPAREN, 0)
}

func (s *LambdaParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRPAREN, 0)
}

func (s *LambdaParametersContext) FormalParameterList() IFormalParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *LambdaParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterLambdaParameters(s)
	}
}

func (s *LambdaParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitLambdaParameters(s)
	}
}

func (s *LambdaParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitLambdaParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) LambdaParameters() (localctx ILambdaParametersContext) {
	localctx = NewLambdaParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, RuleExprParserRULE_lambdaParameters)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.Match(RuleExprParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.Match(RuleExprParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(RuleExprParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(401)
			p.Match(RuleExprParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.FormalParameterList()
		}
		{
			p.SetState(403)
			p.Match(RuleExprParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_formalParameterList
	return p
}

func InitEmptyFormalParameterListContext(p *FormalParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_formalParameterList
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(RuleExprParserIdentifier)
}

func (s *FormalParameterListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(RuleExprParserIdentifier, i)
}

func (s *FormalParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RuleExprParserCOMMA)
}

func (s *FormalParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RuleExprParserCOMMA, i)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitFormalParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) FormalParameterList() (localctx IFormalParameterListContext) {
	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, RuleExprParserRULE_formalParameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.Match(RuleExprParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RuleExprParserCOMMA {
		{
			p.SetState(408)
			p.Match(RuleExprParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(RuleExprParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILambdaBodyContext is an interface to support dynamic dispatch.
type ILambdaBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionBlock() IExpressionBlockContext
	Expression() IExpressionContext

	// IsLambdaBodyContext differentiates from other interfaces.
	IsLambdaBodyContext()
}

type LambdaBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaBodyContext() *LambdaBodyContext {
	var p = new(LambdaBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_lambdaBody
	return p
}

func InitEmptyLambdaBodyContext(p *LambdaBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_lambdaBody
}

func (*LambdaBodyContext) IsLambdaBodyContext() {}

func NewLambdaBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaBodyContext {
	var p = new(LambdaBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_lambdaBody

	return p
}

func (s *LambdaBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaBodyContext) ExpressionBlock() IExpressionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionBlockContext)
}

func (s *LambdaBodyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LambdaBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterLambdaBody(s)
	}
}

func (s *LambdaBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitLambdaBody(s)
	}
}

func (s *LambdaBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitLambdaBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) LambdaBody() (localctx ILambdaBodyContext) {
	localctx = NewLambdaBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, RuleExprParserRULE_lambdaBody)
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)
			p.ExpressionBlock()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListLiteralContext is an interface to support dynamic dispatch.
type IListLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	ExpressionList() IExpressionListContext

	// IsListLiteralContext differentiates from other interfaces.
	IsListLiteralContext()
}

type ListLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListLiteralContext() *ListLiteralContext {
	var p = new(ListLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_listLiteral
	return p
}

func InitEmptyListLiteralContext(p *ListLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_listLiteral
}

func (*ListLiteralContext) IsListLiteralContext() {}

func NewListLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListLiteralContext {
	var p = new(ListLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_listLiteral

	return p
}

func (s *ListLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ListLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLBRACK, 0)
}

func (s *ListLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRBRACK, 0)
}

func (s *ListLiteralContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ListLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterListLiteral(s)
	}
}

func (s *ListLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitListLiteral(s)
	}
}

func (s *ListLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitListLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ListLiteral() (localctx IListLiteralContext) {
	localctx = NewListLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, RuleExprParserRULE_listLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(RuleExprParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422419026509568) != 0 {
		{
			p.SetState(420)
			p.ExpressionList()
		}

	}
	{
		p.SetState(423)
		p.Match(RuleExprParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapLiteralContext is an interface to support dynamic dispatch.
type IMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	MapEntryList() IMapEntryListContext

	// IsMapLiteralContext differentiates from other interfaces.
	IsMapLiteralContext()
}

type MapLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapLiteralContext() *MapLiteralContext {
	var p = new(MapLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_mapLiteral
	return p
}

func InitEmptyMapLiteralContext(p *MapLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_mapLiteral
}

func (*MapLiteralContext) IsMapLiteralContext() {}

func NewMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralContext {
	var p = new(MapLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_mapLiteral

	return p
}

func (s *MapLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *MapLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLBRACE, 0)
}

func (s *MapLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRBRACE, 0)
}

func (s *MapLiteralContext) MapEntryList() IMapEntryListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapEntryListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapEntryListContext)
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterMapLiteral(s)
	}
}

func (s *MapLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitMapLiteral(s)
	}
}

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitMapLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) MapLiteral() (localctx IMapLiteralContext) {
	localctx = NewMapLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, RuleExprParserRULE_mapLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(RuleExprParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422419026509568) != 0 {
		{
			p.SetState(426)
			p.MapEntryList()
		}

	}
	{
		p.SetState(429)
		p.Match(RuleExprParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapEntryListContext is an interface to support dynamic dispatch.
type IMapEntryListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMapEntry() []IMapEntryContext
	MapEntry(i int) IMapEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMapEntryListContext differentiates from other interfaces.
	IsMapEntryListContext()
}

type MapEntryListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapEntryListContext() *MapEntryListContext {
	var p = new(MapEntryListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_mapEntryList
	return p
}

func InitEmptyMapEntryListContext(p *MapEntryListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_mapEntryList
}

func (*MapEntryListContext) IsMapEntryListContext() {}

func NewMapEntryListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryListContext {
	var p = new(MapEntryListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_mapEntryList

	return p
}

func (s *MapEntryListContext) GetParser() antlr.Parser { return s.parser }

func (s *MapEntryListContext) AllMapEntry() []IMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapEntryContext); ok {
			len++
		}
	}

	tst := make([]IMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapEntryContext); ok {
			tst[i] = t.(IMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *MapEntryListContext) MapEntry(i int) IMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapEntryContext)
}

func (s *MapEntryListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RuleExprParserCOMMA)
}

func (s *MapEntryListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RuleExprParserCOMMA, i)
}

func (s *MapEntryListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapEntryListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapEntryListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterMapEntryList(s)
	}
}

func (s *MapEntryListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitMapEntryList(s)
	}
}

func (s *MapEntryListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitMapEntryList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) MapEntryList() (localctx IMapEntryListContext) {
	localctx = NewMapEntryListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, RuleExprParserRULE_mapEntryList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.MapEntry()
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RuleExprParserCOMMA {
		{
			p.SetState(432)
			p.Match(RuleExprParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.MapEntry()
		}

		p.SetState(438)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapEntryContext is an interface to support dynamic dispatch.
type IMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsMapEntryContext differentiates from other interfaces.
	IsMapEntryContext()
}

type MapEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapEntryContext() *MapEntryContext {
	var p = new(MapEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_mapEntry
	return p
}

func InitEmptyMapEntryContext(p *MapEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_mapEntry
}

func (*MapEntryContext) IsMapEntryContext() {}

func NewMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryContext {
	var p = new(MapEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_mapEntry

	return p
}

func (s *MapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *MapEntryContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MapEntryContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MapEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(RuleExprParserCOLON, 0)
}

func (s *MapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterMapEntry(s)
	}
}

func (s *MapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitMapEntry(s)
	}
}

func (s *MapEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitMapEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) MapEntry() (localctx IMapEntryContext) {
	localctx = NewMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, RuleExprParserRULE_mapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.Expression()
	}
	{
		p.SetState(440)
		p.Match(RuleExprParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RuleExprParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RuleExprParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, RuleExprParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.Expression()
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RuleExprParserCOMMA {
		{
			p.SetState(444)
			p.Match(RuleExprParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Expression()
		}

		p.SetState(450)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionBlockContext is an interface to support dynamic dispatch.
type IExpressionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACE() antlr.TerminalNode
	AllBlockStatement() []IBlockStatementContext
	BlockStatement(i int) IBlockStatementContext

	// IsExpressionBlockContext differentiates from other interfaces.
	IsExpressionBlockContext()
}

type ExpressionBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionBlockContext() *ExpressionBlockContext {
	var p = new(ExpressionBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_expressionBlock
	return p
}

func InitEmptyExpressionBlockContext(p *ExpressionBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RuleExprParserRULE_expressionBlock
}

func (*ExpressionBlockContext) IsExpressionBlockContext() {}

func NewExpressionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionBlockContext {
	var p = new(ExpressionBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleExprParserRULE_expressionBlock

	return p
}

func (s *ExpressionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserLBRACE, 0)
}

func (s *ExpressionBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RuleExprParserRBRACE, 0)
}

func (s *ExpressionBlockContext) AllBlockStatement() []IBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStatementContext); ok {
			tst[i] = t.(IBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionBlockContext) BlockStatement(i int) IBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ExpressionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.EnterExpressionBlock(s)
	}
}

func (s *ExpressionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleExprParserListener); ok {
		listenerT.ExitExpressionBlock(s)
	}
}

func (s *ExpressionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleExprParserVisitor:
		return t.VisitExpressionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleExprParser) ExpressionBlock() (localctx IExpressionBlockContext) {
	localctx = NewExpressionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, RuleExprParserRULE_expressionBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(RuleExprParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(452)
				p.BlockStatement()
			}

		}
		p.SetState(457)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(458)
		p.Expression()
	}
	{
		p.SetState(459)
		p.Match(RuleExprParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *RuleExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 20:
		var t *ConditionalOrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ConditionalOrExpressionContext)
		}
		return p.ConditionalOrExpression_Sempred(t, predIndex)

	case 21:
		var t *ConditionalAndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ConditionalAndExpressionContext)
		}
		return p.ConditionalAndExpression_Sempred(t, predIndex)

	case 22:
		var t *InclusiveOrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*InclusiveOrExpressionContext)
		}
		return p.InclusiveOrExpression_Sempred(t, predIndex)

	case 23:
		var t *ExclusiveOrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExclusiveOrExpressionContext)
		}
		return p.ExclusiveOrExpression_Sempred(t, predIndex)

	case 24:
		var t *AndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AndExpressionContext)
		}
		return p.AndExpression_Sempred(t, predIndex)

	case 25:
		var t *EqualityExpressionContext = nil
		if localctx != nil {
			t = localctx.(*EqualityExpressionContext)
		}
		return p.EqualityExpression_Sempred(t, predIndex)

	case 26:
		var t *RelationalExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelationalExpressionContext)
		}
		return p.RelationalExpression_Sempred(t, predIndex)

	case 27:
		var t *ShiftExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ShiftExpressionContext)
		}
		return p.ShiftExpression_Sempred(t, predIndex)

	case 28:
		var t *AdditiveExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveExpressionContext)
		}
		return p.AdditiveExpression_Sempred(t, predIndex)

	case 29:
		var t *MultiplicativeExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeExpressionContext)
		}
		return p.MultiplicativeExpression_Sempred(t, predIndex)

	case 32:
		var t *PostfixExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PostfixExpressionContext)
		}
		return p.PostfixExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RuleExprParser) ConditionalOrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) ConditionalAndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) InclusiveOrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) ExclusiveOrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) AndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) EqualityExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) RelationalExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) ShiftExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) AdditiveExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) MultiplicativeExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 16:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleExprParser) PostfixExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 19:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 20:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 22:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
