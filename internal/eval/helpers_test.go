package eval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// evalResult 编译并求值 src，返回 Value 与 error（求值期 panic 转为 error）。
func evalResult(t *testing.T, src string, ctx map[string]Value) (Value, error) {
	t.Helper()
	fn, err := CompileString(src)
	if err != nil {
		return Value{}, err
	}
	env := NewRootEnv()
	for name, val := range ctx {
		env.Set(name, val)
	}
	defer func() {
		if r := recover(); r != nil {
			err = toEvalErr(r)
		}
	}()
	return Run(fn, env), err
}

func toEvalErr(r interface{}) error {
	if ee, ok := r.(*EvalError); ok {
		return ee
	}
	return fmt.Errorf("%v", r)
}

// wantToValue 把断言期望值（Go 原生类型）转为 eval.Value，用于对比。
func wantToValue(v interface{}) Value {
	switch val := v.(type) {
	case int:
		return IntValue(int64(val))
	case int64:
		return IntValue(val)
	case float64:
		return FloatValue(val)
	case bool:
		return BoolValue(val)
	case string:
		return StringValue(val)
	case nil:
		return NullValue()
	case []interface{}:
		lst := make([]Value, len(val))
		for i, item := range val {
			lst[i] = wantToValue(item)
		}
		return ListValue(lst)
	case map[string]interface{}:
		m := make(map[string]Value)
		for k, item := range val {
			m[k] = wantToValue(item)
		}
		return MapValue(m)
	}
	panic(fmt.Sprintf("unsupported want type: %T", v))
}

// assertEval 断言 src 求值成功且结果与 want 相等。
func assertEval(t *testing.T, src string, want interface{}) {
	t.Helper()
	v, err := evalResult(t, src, nil)
	if !assert.NoError(t, err, "src: %s", src) {
		return
	}
	assert.Equal(t, wantToValue(want), v, "src: %s", src)
}

// assertEvalCtx 断言带上下文求值成功。
func assertEvalCtx(t *testing.T, src string, ctx map[string]Value, want interface{}) {
	t.Helper()
	v, err := evalResult(t, src, ctx)
	if !assert.NoError(t, err, "src: %s", src) {
		return
	}
	assert.Equal(t, wantToValue(want), v, "src: %s", src)
}

// assertEvalError 断言 src 编译或求值报错；可选匹配错误子串。
func assertEvalError(t *testing.T, src string, wantErrSub ...string) {
	t.Helper()
	_, err := evalResult(t, src, nil)
	if !assert.Error(t, err, "src: %s", src) {
		return
	}
	for _, sub := range wantErrSub {
		assert.Contains(t, err.Error(), sub, "src: %s", src)
	}
}
