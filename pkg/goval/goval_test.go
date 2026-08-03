package goval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate_Arithmetic(t *testing.T) {
	v, err := Evaluate("1 + 2 * 3", nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(7), v)
}

func TestEvaluate_Context(t *testing.T) {
	ctx := map[string]interface{}{"x": float64(10)}
	v, err := Evaluate("x > 5", ctx)
	assert.NoError(t, err)
	assert.Equal(t, true, v)
}

func TestEvaluate_ObjectFactory(t *testing.T) {
	src := `var Person = (name) -> { name: name, greet: () -> "hi " + name };
	        Person("bob").greet()`
	v, err := Evaluate(src, nil)
	assert.NoError(t, err)
	assert.Equal(t, "hi bob", v)
}

func TestEvaluate_Reduce(t *testing.T) {
	v, err := Evaluate("reduce([1,2,3], 0, (acc, x) -> acc + x)", nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(6), v)
}

func TestEvaluate_Error(t *testing.T) {
	// 1 + + 2 实际上被解析为 1 + (+2) = 3（一元正号），不会报错。
	// 用除以零触发求值期 EvalError，验证 panic 被 recover 转为 error。
	_, err := Evaluate("1 / 0", nil)
	assert.Error(t, err)
}
