package eval

import (
	"fmt"
	"strings"
	"testing"
)

// assertEvalErrorAt 断言 src 求值报错，且错误信息含 "line N" 与指定子串。
func assertEvalErrorAt(t *testing.T, src string, line int, msgSub string) {
	t.Helper()
	_, err := evalResult(t, src, nil)
	if err == nil {
		t.Fatalf("src: %s\nwant error, got nil", src)
	}
	if !strings.Contains(err.Error(), fmt.Sprintf("line %d", line)) {
		t.Errorf("src: %s\nwant error at line %d, got %q", src, line, err.Error())
	}
	if !strings.Contains(err.Error(), msgSub) {
		t.Errorf("src: %s\nwant error containing %q, got %q", src, msgSub, err.Error())
	}
}

// TestErrorPositions 运行时错误应携带出错节点的行列（此前一律 line 0, column 0）。
func TestErrorPositions(t *testing.T) {
	t.Run("add_type_error", func(t *testing.T) {
		assertEvalErrorAt(t, "var a = 1\na + \"x\"", 2, "cannot add")
	})
	t.Run("field_on_null", func(t *testing.T) {
		assertEvalErrorAt(t, "var m = {}\nm.a.b", 2, "cannot access field")
	})
	t.Run("index_out_of_range", func(t *testing.T) {
		assertEvalErrorAt(t, "var lst = [1]\nlst[5]", 2, "out of range")
	})
	t.Run("division_by_zero", func(t *testing.T) {
		assertEvalErrorAt(t, "var a = 6\na / 0", 2, "division by zero")
	})
	t.Run("comparison_type_error", func(t *testing.T) {
		assertEvalErrorAt(t, "var a = 1\n\"x\" < a", 2, "cannot compare")
	})
	t.Run("builtin_arg_error", func(t *testing.T) {
		assertEvalErrorAt(t, "var s = \"x\"\nvar n = 1\nabs(s)", 3, "abs requires")
	})
	t.Run("method_missing", func(t *testing.T) {
		assertEvalErrorAt(t, "var p = {}\np.foo()", 2, "no method 'foo'")
	})
	t.Run("string_method_arg_error", func(t *testing.T) {
		assertEvalErrorAt(t, "var s = \"a\"\ns.repeat(\"x\")", 2, "repeat")
	})
	t.Run("innermost_lambda_body", func(t *testing.T) {
		// 错误定位到 lambda 体内出错的表达式（第 2 行），不是调用点（第 3 行）
		assertEvalErrorAt(t, "var x = 1\nvar f = (a) -> a * \"s\"\nf(x)", 2, "cannot multiply")
	})
	t.Run("assignment_index_error", func(t *testing.T) {
		assertEvalErrorAt(t, "var lst = [1]\nlst[3] = 9", 2, "out of range")
	})
	t.Run("if_condition_regression", func(t *testing.T) {
		// compileIf 已带位置，回归确认不被破坏
		assertEvalErrorAt(t, "if (1) {}", 1, "must be bool")
	})
}
