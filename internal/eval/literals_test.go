package eval

import "testing"

func TestLiterals(t *testing.T) {
	t.Run("decimal_int", func(t *testing.T) {
		assertEval(t, "42", int64(42))
		assertEval(t, "0", int64(0))
	})
	t.Run("hex_octal_binary", func(t *testing.T) {
		assertEval(t, "0x1A", int64(26))
		assertEval(t, "0755", int64(493))
		assertEval(t, "0b1010", int64(10))
	})
	t.Run("int_L_suffix", func(t *testing.T) {
		assertEval(t, "123L", int64(123))
	})
	t.Run("float", func(t *testing.T) {
		assertEval(t, "3.14", 3.14)
		assertEval(t, ".5", 0.5)
		assertEval(t, "2.", 2.0)
	})
	t.Run("float_f_d_suffix", func(t *testing.T) {
		assertEval(t, "3.14f", 3.14)
		assertEval(t, "2.5d", 2.5)
	})
	t.Run("scientific", func(t *testing.T) {
		assertEval(t, "1e3", 1000.0)
		assertEval(t, "1.5E-2", 0.015)
	})
	t.Run("bool", func(t *testing.T) {
		assertEval(t, "true", true)
		assertEval(t, "false", false)
	})
	t.Run("null", func(t *testing.T) {
		assertEval(t, "null", nil)
	})
	t.Run("char_literal", func(t *testing.T) {
		assertEval(t, "'a'", int64(97))
		assertEval(t, "'\\n'", int64(10))
		assertEval(t, "'\\t'", int64(9))
		assertEval(t, "'\\''", int64(39))
		assertEval(t, "'\\\\'", int64(92)) // 转义反斜杠字符
	})
	t.Run("string", func(t *testing.T) {
		assertEval(t, `"hello"`, "hello")
	})
	t.Run("string_escape", func(t *testing.T) {
		assertEval(t, `"a\nb"`, "a\nb")
		assertEval(t, `"a\tb"`, "a\tb")
		assertEval(t, `"a\\b"`, "a\\b")
		assertEval(t, `"a\"b"`, "a\"b")
	})
	t.Run("list", func(t *testing.T) {
		assertEval(t, "[1, 2, 3]", []interface{}{int64(1), int64(2), int64(3)})
		assertEval(t, "[]", []interface{}{})
	})
	t.Run("map", func(t *testing.T) {
		assertEval(t, `{"a": 1, "b": 2}`, map[string]interface{}{"a": int64(1), "b": int64(2)})
	})
	t.Run("map_bare_identifier_key", func(t *testing.T) {
		// {name: name} 简写：key 取标识符文本，value 取同名变量
		assertEvalCtx(t, `{name: name}`, map[string]Value{"name": StringValue("alice")},
			map[string]interface{}{"name": "alice"})
	})
}
