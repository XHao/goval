package eval

import (
	"github.com/XHao/goval/internal/ast"
	"github.com/antlr4-go/antlr/v4"
)

// CompileString 是测试用便捷入口：源码字符串 → 闭包树。
// globals 追加到根作用域的已知名字（如 context 注入的变量名），
// 供严格赋值检查在编译期识别可更新的绑定。
func CompileString(source string, globals ...string) (func(*Env) Value, error) {
	input := antlr.NewInputStream(source)
	lexer := ast.NewRuleExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ast.NewRuleExprParser(stream)
	tree := parser.Program()
	return Compile(tree, globals...)
}

// Compile 遍历 parse tree，产出闭包树。
// 内置函数名与 rootGlobals 注册到根作用域：裸赋值的目标必须已绑定，
// 防止拼写错误静默创建新变量（v2 严格赋值）。
func Compile(tree ast.IProgramContext, rootGlobals ...string) (func(*Env) Value, error) {
	c := &compiler{
		scopes: []scopeInfo{{names: map[string]bool{}, isLambda: false}}, // 根作用域
	}
	for name := range defaultBuiltins() {
		c.currentScope()[name] = true
	}
	for _, name := range rootGlobals {
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
