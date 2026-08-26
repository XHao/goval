package eval

import "testing"

// TestConvertBuiltins 类型转换内建：严格语义——不做隐式转换，失败即运行时错误。
func TestConvertBuiltins(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assertEval(t, "int(5)", int64(5))
		assertEval(t, "int(12.7)", int64(12)) // 截断，Go 语义
		assertEval(t, "int(-12.7)", int64(-12))
		assertEval(t, `int("12")`, int64(12))
		assertEval(t, `int("-12")`, int64(-12))
		assertEvalError(t, `int("12.5")`, "cannot parse")
		assertEvalError(t, `int("")`, "cannot parse")
		assertEvalError(t, `int("abc")`, "cannot parse")
		assertEvalError(t, "int(true)", "requires")
		assertEvalError(t, "int(null)", "requires")
		assertEvalError(t, "int([1])", "requires")
		assertEvalError(t, "int()", "expected 1 args")
	})
	t.Run("float", func(t *testing.T) {
		assertEval(t, "float(5)", 5.0)
		assertEval(t, "float(12.5)", 12.5)
		assertEval(t, `float("12.5")`, 12.5)
		assertEval(t, `float("1e3")`, 1000.0)
		assertEval(t, `float("-0.5")`, -0.5)
		assertEvalError(t, `float("abc")`, "cannot parse")
		assertEvalError(t, `float("")`, "cannot parse")
		assertEvalError(t, "float(true)", "requires")
		assertEvalError(t, "float(null)", "requires")
	})
	t.Run("string", func(t *testing.T) {
		assertEval(t, `string(5)`, "5")
		assertEval(t, `string(5.0)`, "5") // Go 最短格式化
		assertEval(t, `string(12.5)`, "12.5")
		assertEval(t, `string(true)`, "true")
		assertEval(t, `string(false)`, "false")
		assertEval(t, `string("a")`, "a")
		assertEvalError(t, "string(null)", "requires")
		assertEvalError(t, "string([1])", "requires")
		assertEvalError(t, `string({"a": 1})`, "requires")
	})
	t.Run("compose", func(t *testing.T) {
		assertEval(t, `int("12") + 3`, int64(15))
		assertEval(t, `"共" + string(3) + " 条"`, "共3 条")
		assertEval(t, `float("1.5") * 2`, 3.0)
	})
}
