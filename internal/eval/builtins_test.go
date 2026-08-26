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
		// start >= end 一律返回空列表：此前 range(5, 0) 因 make 负容量 panic
		assertEval(t, "range(5, 0)", []interface{}{})
		assertEval(t, "range(-3, -5)", []interface{}{})
		assertEval(t, "range(-3, 0)", []interface{}{int64(-3), int64(-2), int64(-1)})
	})
	t.Run("keys_values", func(t *testing.T) {
		// 按键排序，输出确定性（Go map 迭代序随机）
		assertEval(t, `keys({"b": 1, "a": 2})`, []interface{}{"a", "b"})
		assertEval(t, `keys({})`, []interface{}{})
		assertEval(t, `values({"b": 1, "a": 2})`, []interface{}{int64(2), int64(1)}) // 值随 key 排序对应
		assertEval(t, `values({})`, []interface{}{})
		assertEvalError(t, `keys([1])`, "map")
		assertEvalError(t, `values("x")`, "map")
		assertEvalError(t, `keys()`, "expected 1 args")
	})
}

// TestBuiltinArgumentValidation 回归：内建函数此前对参数类型/个数不做校验——
// filter 回调返回非 bool 被静默当 false、reduce 传非 List 静默返回初值、
// put 非字符串键静默写入空键，len() 则以索引越界 panic 报错。
func TestBuiltinArgumentValidation(t *testing.T) {
	t.Run("callback_must_return_bool", func(t *testing.T) {
		assertEvalError(t, "filter([1, 2], x -> x * 2)", "bool")
		assertEvalError(t, "find([1, 2], x -> x + 1)", "bool")
	})
	t.Run("callback_must_be_lambda", func(t *testing.T) {
		assertEvalError(t, "reduce([1], 0, 5)", "lambda")
		assertEvalError(t, "map([1], 5)", "lambda")
	})
	t.Run("lambda_arity_checked", func(t *testing.T) {
		assertEvalError(t, "map([1, 2], (a, b) -> a)", "expected 2 args")
	})
	t.Run("requires_list", func(t *testing.T) {
		assertEvalError(t, "reduce(5, 0, (a, x) -> a + x)", "list")
		assertEvalError(t, "map(5, x -> x)", "list")
		assertEvalError(t, "filter(5, x -> true)", "list")
		assertEvalError(t, "find(5, x -> true)", "list")
		assertEvalError(t, "append(5, 1)", "list")
		assertEvalError(t, "removeAt(5, 0)", "list")
	})
	t.Run("requires_map_and_string_key", func(t *testing.T) {
		assertEvalError(t, `put([1], "k", 1)`, "map")
		assertEvalError(t, "put({}, 1, 1)", "string key")
	})
	t.Run("requires_int", func(t *testing.T) {
		assertEvalError(t, `removeAt([1, 2], "x")`, "int index")
		assertEvalError(t, `range("a", 2)`, "int")
		assertEvalError(t, "range(1, 2.5)", "int")
	})
	t.Run("removeAt_bounds", func(t *testing.T) {
		// 越界索引此前静默返回原列表、空列表则 makeslice panic
		assertEvalError(t, "removeAt([1, 2], 5)", "out of range")
		assertEvalError(t, "removeAt([1, 2], -1)", "out of range")
		assertEvalError(t, "removeAt([], 0)", "out of range")
	})
	t.Run("arity_checked", func(t *testing.T) {
		assertEvalError(t, "len()", "expected 1 args")
		assertEvalError(t, "append([1])", "expected 2 args")
		assertEvalError(t, `put({}, "k")`, "expected 3 args")
	})
}
