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

	"github.com/XHao/goval/internal/eval"
	"github.com/XHao/goval/internal/syntax"
)

// Program 是编译后的 goval 程序：编译一次，可对不同 context 反复求值。
type Program struct {
	fn func(*eval.Env) eval.Value
}

// Compile 编译 source 产出 Program。contextKeys 是求值时将注入的 context
// 变量名集合——编译期的严格赋值检查以此识别可赋值的绑定（键集合在编译期固定）。
func Compile(source string, contextKeys ...string) (*Program, error) {
	tree, err := syntax.NewSyntaxChecker().CheckString(source)
	if err != nil {
		return nil, err
	}
	fn, err := eval.Compile(tree, contextKeys...)
	if err != nil {
		return nil, err
	}
	return &Program{fn: fn}, nil
}

// Run 注入 context 并求值。每次调用相互独立：context 深拷贝注入，
// 不支持的 context 值类型报错；求值期 panic 被 recover 转为 error。
// 编译期键集合之外的键注入为额外全局变量；缺失的键在首次读取时报错。
func (p *Program) Run(context map[string]interface{}) (result interface{}, err error) {
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

	env := eval.NewRootEnv()
	for name, val := range context {
		gv, cerr := toGovalValue(val)
		if cerr != nil {
			return nil, fmt.Errorf("context %q: %w", name, cerr)
		}
		env.Set(name, gv)
	}
	return toGoValue(eval.Run(p.fn, env)), nil
}

// Evaluate 编译并求值 source 表达式——Compile + Run 的便捷组合。
// context 中的值作为全局变量注入；重复求值同一规则时用 Compile/Run 省去重复编译。
// 返回 Go 原生值（int64/float64/bool/string/nil/[]interface{}/map[string]interface{}）。
// 解析/语义/编译/求值任一阶段的错误都转为 error 返回；
// 求值期发生的 panic 会被 recover 转为 error 返回，不会 panic 到调用方。
func Evaluate(source string, context map[string]interface{}) (interface{}, error) {
	keys := make([]string, 0, len(context))
	for name := range context {
		keys = append(keys, name)
	}
	p, err := Compile(source, keys...)
	if err != nil {
		return nil, err
	}
	return p.Run(context)
}

// toGovalValue 将 Go 原生值转为 goval Value；不支持的类型返回错误而非静默转 null。
func toGovalValue(v interface{}) (eval.Value, error) {
	switch val := v.(type) {
	case int:
		return eval.IntValue(int64(val)), nil
	case int64:
		return eval.IntValue(val), nil
	case float64:
		return eval.FloatValue(val), nil
	case float32:
		return eval.FloatValue(float64(val)), nil
	case bool:
		return eval.BoolValue(val), nil
	case string:
		return eval.StringValue(val), nil
	case nil:
		return eval.NullValue(), nil
	case []interface{}:
		lst := make([]eval.Value, len(val))
		for i, item := range val {
			iv, err := toGovalValue(item)
			if err != nil {
				return eval.Value{}, err
			}
			lst[i] = iv
		}
		return eval.ListValue(lst), nil
	case map[string]interface{}:
		m := make(map[string]eval.Value)
		for k, item := range val {
			iv, err := toGovalValue(item)
			if err != nil {
				return eval.Value{}, err
			}
			m[k] = iv
		}
		return eval.MapValue(m), nil
	}
	return eval.Value{}, fmt.Errorf("unsupported context value type %T", v)
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
