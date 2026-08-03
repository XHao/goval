package eval

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/XHao/goval/internal/ast"
)

// CompileString 是测试用便捷入口：源码字符串 → 闭包树。
func CompileString(source string) (func(*Env) Value, error) {
	input := antlr.NewInputStream(source)
	lexer := ast.NewRuleExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ast.NewRuleExprParser(stream)
	tree := parser.Program()
	return Compile(tree)
}

// Compile 遍历 parse tree，产出闭包树。
// 内置函数名注册到根作用域，阻止用户用同名 var 重绑定。
func Compile(tree ast.IProgramContext) (func(*Env) Value, error) {
	c := &compiler{
		scopes: []map[string]bool{{}}, // 根作用域
	}
	// 注册内置函数名到根作用域（不计入用户变量重绑定检查）
	for name := range defaultBuiltins() {
		c.currentScope()[name] = true
	}
	fn, err := c.compileProgram(tree)
	if err != nil {
		return nil, err
	}
	return fn, nil
}

// Run 执行闭包树求值，注入内置函数值到根 Env。
func Run(fn func(*Env) Value, env *Env) Value {
	for name, v := range defaultBuiltins() {
		if _, ok := env.Lookup(name); !ok {
			env.Set(name, v)
		}
	}
	return fn(env)
}
