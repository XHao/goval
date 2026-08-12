package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// assertValid 断言 src 通过语法+语义检查。
func assertValid(t *testing.T, src string) {
	t.Helper()
	_, err := NewSyntaxChecker().CheckString(src)
	assert.NoError(t, err, "应可解析: %s", src)
}

// assertInvalid 断言 src 报错（语法或语义错误）。
func assertInvalid(t *testing.T, src string) {
	t.Helper()
	_, err := NewSyntaxChecker().CheckString(src)
	assert.Error(t, err, "应被拒绝: %s", src)
}
