package goval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// assertEvaluate 断言公开 Evaluate 求值成功且结果相等。
func assertEvaluate(t *testing.T, src string, ctx map[string]interface{}, want interface{}) {
	t.Helper()
	v, err := Evaluate(src, ctx)
	if !assert.NoError(t, err, "src: %s", src) {
		return
	}
	assert.Equal(t, want, v, "src: %s", src)
}

// assertEvaluateError 断言公开 Evaluate 报错。
func assertEvaluateError(t *testing.T, src string, ctx map[string]interface{}) {
	t.Helper()
	_, err := Evaluate(src, ctx)
	assert.Error(t, err, "src: %s", src)
}
