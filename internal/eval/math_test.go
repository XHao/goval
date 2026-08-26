package eval

import "testing"

// TestMathBuiltins 数学内建：abs 保型；round half away from zero（可选小数位）；
// floor/ceil 返回 float；min/max 变参标量或单个 List。
func TestMathBuiltins(t *testing.T) {
	t.Run("abs", func(t *testing.T) {
		assertEval(t, "abs(-3)", int64(3)) // 保型：int → int
		assertEval(t, "abs(3)", int64(3))
		assertEval(t, "abs(-3.5)", 3.5)
		assertEval(t, "abs(0)", int64(0))
		assertEvalError(t, `abs("x")`, "requires")
		assertEvalError(t, "abs()", "expected 1 args")
	})
	t.Run("round", func(t *testing.T) {
		assertEval(t, "round(2.5)", 3.0) // half away from zero
		assertEval(t, "round(-2.5)", -3.0)
		assertEval(t, "round(2.4)", 2.0)
		assertEval(t, "round(5)", 5.0) // 返回 float
		assertEval(t, "round(8.247, 2)", 8.25)
		assertEval(t, "round(8.247, 0)", 8.0)
		assertEval(t, "round(1234.567, 1)", 1234.6)
		assertEvalError(t, "round(2.5, -1)", "digits")
		assertEvalError(t, `round("x")`, "requires")
		assertEvalError(t, "round(1, 2, 3)", "expected 1 or 2 args")
	})
	t.Run("floor_ceil", func(t *testing.T) {
		assertEval(t, "floor(2.7)", 2.0)
		assertEval(t, "ceil(2.1)", 3.0)
		assertEval(t, "floor(-2.5)", -3.0)
		assertEval(t, "ceil(-2.5)", -2.0)
		assertEval(t, "floor(5)", 5.0)
		assertEvalError(t, `floor("x")`, "requires")
	})
	t.Run("min_max", func(t *testing.T) {
		assertEval(t, "min(1, 2)", int64(1))
		assertEval(t, "max(1, 2)", int64(2))
		assertEval(t, "min(3, 1, 2)", int64(1))
		assertEval(t, "min(1, 2.5)", 1.0) // 混 float 提升为 float
		assertEval(t, "max(2, 2.5)", 2.5)
		assertEval(t, "min([3, 1, 2])", int64(1)) // 单个 List 参数作用于元素
		assertEval(t, "max([3, 1, 2])", int64(3))
		assertEval(t, "min([-1.5, -2.5])", -2.5)
		assertEvalError(t, "min()", "at least 2")
		assertEvalError(t, "min(5)", "at least 2") // 单标量不是合法形态
		assertEvalError(t, "min([])", "empty list")
		assertEvalError(t, `min("a", 1)`, "numeric")
		assertEvalError(t, `min([1, "a"])`, "numeric")
	})
}
