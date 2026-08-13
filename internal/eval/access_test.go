package eval

import "testing"

func TestAccess(t *testing.T) {
	t.Run("list_index", func(t *testing.T) {
		assertEval(t, "[10, 20, 30][0]", int64(10))
		assertEval(t, "[10, 20, 30][2]", int64(30))
	})
	t.Run("list_index_out_of_range", func(t *testing.T) {
		assertEvalError(t, "[10, 20, 30][5]", "out of range")
	})
	t.Run("list_negative_index", func(t *testing.T) {
		assertEvalError(t, "[10, 20, 30][-1]", "out of range")
	})
	t.Run("map_index", func(t *testing.T) {
		assertEval(t, `{"key": 42}["key"]`, int64(42))
	})
	t.Run("map_index_missing_returns_null", func(t *testing.T) {
		assertEval(t, `{"key": 42}["x"]`, nil)
	})
	t.Run("field_access", func(t *testing.T) {
		assertEval(t, `{"name": "alice"}.name`, "alice")
	})
	t.Run("string_index", func(t *testing.T) {
		// 字符串下标返回单字符 string（非 rune int）
		assertEval(t, `"hello"[1]`, "e")
		assertEval(t, `"hello"[0]`, "h")
	})
	t.Run("string_index_out_of_range", func(t *testing.T) {
		assertEvalError(t, `"hi"[5]`, "out of range")
	})
	t.Run("function_call", func(t *testing.T) {
		assertEval(t, "((x) -> x * 2)(21)", int64(42))
	})
	t.Run("method_call", func(t *testing.T) {
		src := `var o = {val: 10, double: () -> o.val * 2}; o.double()`
		assertEval(t, src, int64(20))
	})
	t.Run("chained_access", func(t *testing.T) {
		assertEval(t, `{"a": {"b": {"c": 7}}}.a.b.c`, int64(7))
	})
	t.Run("nested_index", func(t *testing.T) {
		assertEval(t, `{"a": [1, 2, 3]}["a"][1]`, int64(2))
	})
}
