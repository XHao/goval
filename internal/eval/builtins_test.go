package eval

import "testing"

func TestBuiltins(t *testing.T) {
	t.Run("reduce", func(t *testing.T) {
		assertEval(t, "reduce([1,2,3], 0, (acc, x) -> acc + x)", int64(6))
		assertEval(t, "reduce([], 0, (acc, x) -> acc + x)", int64(0))
		assertEval(t, "reduce([1,2,3], 10, (acc, x) -> acc + x)", int64(16))
	})
	t.Run("map", func(t *testing.T) {
		assertEval(t, "map([1,2,3], x -> x * 2)", []interface{}{int64(2), int64(4), int64(6)})
		assertEval(t, "map([], x -> x)", []interface{}{})
	})
	t.Run("filter", func(t *testing.T) {
		assertEval(t, "filter([1,2,3,4], x -> x > 2)", []interface{}{int64(3), int64(4)})
		assertEval(t, "filter([1,2,3], x -> x > 9)", []interface{}{})
	})
	t.Run("find", func(t *testing.T) {
		assertEval(t, "find([1,2,3], x -> x == 2)", int64(2))
		assertEval(t, "find([1,2,3], x -> x == 9)", nil) // 未命中返回 null
	})
	t.Run("append", func(t *testing.T) {
		assertEval(t, "append([1,2], 3)", []interface{}{int64(1), int64(2), int64(3)})
		assertEval(t, "append([], 1)", []interface{}{int64(1)})
	})
	t.Run("put", func(t *testing.T) {
		assertEval(t, `put({"a": 1}, "b", 2)`, map[string]interface{}{"a": int64(1), "b": int64(2)})
	})
	t.Run("removeAt", func(t *testing.T) {
		assertEval(t, "removeAt([1,2,3], 1)", []interface{}{int64(1), int64(3)})
		assertEval(t, "removeAt([1,2,3], 0)", []interface{}{int64(2), int64(3)})
		assertEval(t, "removeAt([1,2,3], 2)", []interface{}{int64(1), int64(2)})
	})
	t.Run("len", func(t *testing.T) {
		assertEval(t, "len([1,2,3])", int64(3))
		assertEval(t, `len("hello")`, int64(5))
		assertEval(t, `len({"a":1,"b":2})`, int64(2))
		assertEval(t, "len([])", int64(0))
		assertEval(t, `len("")`, int64(0))
	})
	t.Run("range", func(t *testing.T) {
		assertEval(t, "range(1, 4)", []interface{}{int64(1), int64(2), int64(3)})
		assertEval(t, "range(0, 0)", []interface{}{})
		assertEval(t, "range(5, 8)", []interface{}{int64(5), int64(6), int64(7)})
	})
}
