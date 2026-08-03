// Package goval 提供公开的表达式求值 API。
//
// 用法：
//
//	v, err := goval.Evaluate("1 + 2 * 3", nil)
//	// v == int64(7)
//
//	context := map[string]interface{}{"x": float64(10)}
//	v, err := goval.Evaluate("x > 5", context)
//	// v == true
package goval

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/XHao/goval/internal/ast"
	"github.com/XHao/goval/internal/eval"
)

// Evaluate 编译并求值 source 表达式。context 中的值作为全局变量注入。
// 返回 Go 原生值（int64/float64/bool/string/nil/[]interface{}/map[string]interface{}）。
// 求值期发生的 panic 会被 recover 转为 error 返回，不会 panic 到调用方。
func Evaluate(source string, context map[string]interface{}) (result interface{}, err error) {
	input := antlr.NewInputStream(source)
	lexer := ast.NewRuleExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ast.NewRuleExprParser(stream)
	tree := parser.Program()

	fn, err := eval.Compile(tree)
	if err != nil {
		return nil, err
	}

	env := eval.NewRootEnv()
	for name, val := range context {
		env.Set(name, toGovalValue(val))
	}

	defer func() {
		if r := recover(); r != nil {
			if ee, ok := r.(*eval.EvalError); ok {
				err = ee
			} else {
				err = fmt.Errorf("%v", r)
			}
			result = nil
		}
	}()
	val := eval.Run(fn, env)
	return toGoValue(val), nil
}

// toGovalValue 将 Go 原生值转为 goval Value。
func toGovalValue(v interface{}) eval.Value {
	switch val := v.(type) {
	case int:
		return eval.IntValue(int64(val))
	case int64:
		return eval.IntValue(val)
	case float64:
		return eval.FloatValue(val)
	case float32:
		return eval.FloatValue(float64(val))
	case bool:
		return eval.BoolValue(val)
	case string:
		return eval.StringValue(val)
	case nil:
		return eval.NullValue()
	case []interface{}:
		lst := make([]eval.Value, len(val))
		for i, item := range val {
			lst[i] = toGovalValue(item)
		}
		return eval.ListValue(lst)
	case map[string]interface{}:
		m := make(map[string]eval.Value)
		for k, item := range val {
			m[k] = toGovalValue(item)
		}
		return eval.MapValue(m)
	}
	return eval.NullValue()
}

// toGoValue 将 goval Value 转为 Go 原生值。
func toGoValue(v eval.Value) interface{} {
	switch {
	case v.IsInt():
		return v.I()
	case v.IsFloat():
		return v.F()
	case v.IsBool():
		return v.B()
	case v.IsString():
		return v.S()
	case v.IsNull():
		return nil
	case v.IsList():
		lst := make([]interface{}, len(v.List()))
		for i, item := range v.List() {
			lst[i] = toGoValue(item)
		}
		return lst
	case v.IsMap():
		m := make(map[string]interface{})
		for k, item := range v.Map() {
			m[k] = toGoValue(item)
		}
		return m
	}
	return nil
}
