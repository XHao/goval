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
	t.Run("redeclare_same_scope", func(t *testing.T) {
		// 同作用域 var 重复声明应编译报错
		assertEvalError(t, "var x = 1; var x = 2", "already declared")
	})
	t.Run("reassign_updates_value", func(t *testing.T) {
		// v2：赋值更新已声明变量
		assertEval(t, "var x = 1; x = 2; x", int64(2))
	})
	t.Run("assign_undeclared_rejected", func(t *testing.T) {
		// v2 严格赋值：裸赋值要求目标已声明（防拼写错误静默建新变量）
		assertEvalError(t, "x = 2", "undefined variable 'x'")
	})
	t.Run("assign_undeclared_typo_rejected", func(t *testing.T) {
		assertEvalError(t, "var userId = 1; usreId = 2", "undefined variable 'usreId'")
	})
	t.Run("undeclared_variable", func(t *testing.T) {
		// 引用未声明变量
		assertEvalError(t, "undefinedVar")
	})
}
