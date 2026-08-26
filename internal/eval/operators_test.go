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

// TestEqualitySemantics 回归：== 此前按 kind 严格比较，
// 1 == 1.0 为 false，且 JSON 数字注入为 float64 后 x == 100 静默不命中；
// List/Map 无深度相等语义，[1,2] == [1,2] 恒为 false。
func TestEqualitySemantics(t *testing.T) {
	t.Run("numeric_cross_kind", func(t *testing.T) {
		assertEval(t, "1 == 1.0", true)
		assertEval(t, "1.0 == 1", true)
		assertEval(t, "1 != 1.0", false)
		assertEval(t, "1 == 1.5", false)
		assertEval(t, "0 == 0.0", true)
		assertEval(t, "100 == 1e2", true)
	})
	t.Run("numeric_from_context", func(t *testing.T) {
		// Go JSON 反序列化的数字一律是 float64：规则 `amount == 100` 必须按值命中
		assertEvalCtx(t, "x == 1", map[string]Value{"x": FloatValue(1)}, true)
		assertEvalCtx(t, "x == 100", map[string]Value{"x": FloatValue(100)}, true)
		assertEvalCtx(t, "x != 100", map[string]Value{"x": FloatValue(100)}, false)
		assertEvalCtx(t, "x == 100.5", map[string]Value{"x": FloatValue(100.5)}, true)
	})
	t.Run("cross_kind_still_false", func(t *testing.T) {
		assertEval(t, `1 == "1"`, false) // 数值与字符串不隐式转换
		assertEval(t, "true == 1", false)
		assertEval(t, "null == 0", false)
		assertEval(t, "null == null", true)
		assertEval(t, "null != 0", true)
	})
	t.Run("list_deep_equality", func(t *testing.T) {
		assertEval(t, "[1, 2] == [1, 2]", true)
		assertEval(t, "[1, 2] != [1, 2]", false)
		assertEval(t, "[1, 2] == [1, 3]", false)
		assertEval(t, "[1] == [1, 2]", false)
		assertEval(t, "[] == []", true)
		assertEval(t, `[1, "a"] == [1.0, "a"]`, true) // 元素级数值相等
		assertEval(t, `[[1], [2]] == [[1.0], [2.0]]`, true)
		assertEval(t, `[1, 2] == "ab"`, false)
	})
	t.Run("map_deep_equality", func(t *testing.T) {
		assertEval(t, `{"a": 1} == {"a": 1}`, true)
		assertEval(t, `{"a": 1} != {"a": 1}`, false)
		assertEval(t, `{"a": 1} == {"a": 2}`, false)
		assertEval(t, `{"a": 1} == {"b": 1}`, false)
		assertEval(t, `{"a": 1} == {"a": 1, "b": 2}`, false)
		assertEval(t, `{} == {}`, true)
		assertEval(t, `{"a": [1, {"b": 2}]} == {"a": [1.0, {"b": 2.0}]}`, true)
	})
	t.Run("in_list_numeric_cross_kind", func(t *testing.T) {
		assertEval(t, "1 in [1.0, 2.0]", true) // in 复用 eqValues
		assertEval(t, "3 in [1.0, 2.0]", false)
	})
}
