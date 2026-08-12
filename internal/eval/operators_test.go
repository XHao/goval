package eval

import "testing"

func TestOperators(t *testing.T) {
	t.Run("arithmetic_int", func(t *testing.T) {
		assertEval(t, "1 + 2", int64(3))
		assertEval(t, "10 - 3", int64(7))
		assertEval(t, "4 * 5", int64(20))
		assertEval(t, "20 / 4", int64(5))
		assertEval(t, "17 % 5", int64(2))
	})
	t.Run("arithmetic_float", func(t *testing.T) {
		assertEval(t, "1.5 + 2.5", 4.0)
		assertEval(t, "10.0 / 4.0", 2.5)
		assertEval(t, "3.5 - 1.0", 2.5)
	})
	t.Run("string_concat", func(t *testing.T) {
		assertEval(t, `"foo" + "bar"`, "foobar")
	})
	t.Run("comparison", func(t *testing.T) {
		assertEval(t, "3 > 2", true)
		assertEval(t, "3 < 2", false)
		assertEval(t, "2 >= 2", true)
		assertEval(t, "2 <= 1", false)
		assertEval(t, "1 == 1", true)
		assertEval(t, "1 != 2", true)
		assertEval(t, `"a" < "b"`, true)
		assertEval(t, "1.5 > 1.0", true)
	})
	t.Run("logic", func(t *testing.T) {
		assertEval(t, "true && false", false)
		assertEval(t, "true || false", true)
		assertEval(t, "!false", true)
		assertEval(t, "!true", false)
	})
	t.Run("logic_short_circuit", func(t *testing.T) {
		// false && x：x 不应求值（用除零验证短路，若求值则报错）
		assertEval(t, "false && (1 / 0 == 0)", false)
		// true || x：x 不应求值
		assertEval(t, "true || (1 / 0 == 0)", true)
	})
	t.Run("unary", func(t *testing.T) {
		assertEval(t, "-5", int64(-5))
		assertEval(t, "+5", int64(5))
		assertEval(t, "-3.14", -3.14)
	})
	t.Run("bitwise", func(t *testing.T) {
		assertEval(t, "0xF0 & 0x0F", int64(0))
		assertEval(t, "0xF0 | 0x0F", int64(255))
		assertEval(t, "0xFF ^ 0x0F", int64(240))
		assertEval(t, "~0", int64(-1))
	})
	t.Run("shift", func(t *testing.T) {
		assertEval(t, "1 << 4", int64(16))
		assertEval(t, "256 >> 2", int64(64))
	})
	t.Run("in_list", func(t *testing.T) {
		assertEval(t, "2 in [1, 2, 3]", true)
		assertEval(t, "5 in [1, 2, 3]", false)
	})
	t.Run("in_map", func(t *testing.T) {
		assertEval(t, `"k" in {"k": 1}`, true)
		assertEval(t, `"x" in {"k": 1}`, false)
	})
	t.Run("in_string", func(t *testing.T) {
		assertEval(t, `"ll" in "hello"`, true)
		assertEval(t, `"xx" in "hello"`, false)
	})
}
