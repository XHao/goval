package eval

import "testing"

func TestErrors(t *testing.T) {
	t.Run("divide_by_zero", func(t *testing.T) {
		assertEvalError(t, "1 / 0", "zero")
	})
	t.Run("modulo_by_zero", func(t *testing.T) {
		assertEvalError(t, "1 % 0", "zero")
	})
	t.Run("type_mismatch_add", func(t *testing.T) {
		// "a" + 1：类型不匹配
		assertEvalError(t, `"a" + 1`)
	})
	t.Run("type_mismatch_compare", func(t *testing.T) {
		// true < 1：比较类型不匹配
		assertEvalError(t, "true < 1")
	})
	t.Run("logic_non_bool", func(t *testing.T) {
		// !1：非 bool
		assertEvalError(t, "!1", "bool")
	})
	t.Run("bitwise_non_int", func(t *testing.T) {
		// ~true：非 int
		assertEvalError(t, "~true", "int")
	})
	t.Run("rebind_same_scope", func(t *testing.T) {
		// 单赋值：同作用域重绑定应编译报错
		assertEvalError(t, "var x = 1; var x = 2", "rebind")
	})
	// 字段/下标赋值拒绝见 syntax 包测试：
	// CompileString 用 ANTLR 默认错误策略，parse 错误只打印不返回 error，
	// 故无法在 eval 层断言；syntax 包的 SyntaxChecker 有 error listener 可捕获。
	t.Run("reassign_rejected", func(t *testing.T) {
		// 单赋值：x = expr 重绑定外部变量
		assertEvalError(t, "var x = 1; x = 2", "rebind")
	})
	t.Run("undeclared_variable", func(t *testing.T) {
		// 引用未声明变量
		assertEvalError(t, "undefinedVar")
	})
}
