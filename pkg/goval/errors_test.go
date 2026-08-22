package goval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEvaluateErrorPropagation 验证非法输入在各阶段都返回 error，绝不 panic 到调用方。
func TestEvaluateErrorPropagation(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		// v2：字段/下标赋值合法（见 TestV2Mutation）；此处只保留仍被拒绝的输入。
		// 解析错误：不支持的运算符（复合赋值/自增）
		{"compound_assignment", "var x = 1; x += 1"},
		{"increment", "var x = 1; x++"},
		// 解析错误：C 式三段 for
		{"c_style_for", "for (i = 0; i < 10; i = i + 1) {}"},
		// 解析错误：残缺/无法识别的输入
		{"trailing_operator", "1 +"},
		{"unclosed_paren", "(1"},
		{"unrecognized_token", "@@@"},
		// 语义错误：break/continue 只允许出现在 for 体内
		{"break_outside_loop", "break"},
		{"continue_outside_loop", "continue"},
		// 编译错误：v2 严格赋值——裸赋值要求目标已声明
		{"assign_undeclared", "y = 2"},
		{"assign_typo", "var userId = 1; usreId = 2"},
		// 编译错误：非法左值（调用结果不可赋值）
		{"assign_to_call_result", "var f = () -> 1; f() = 2"},
		// 运行时错误：分离方法无 this
		{"detached_method_this", `var P = (n) -> { n: n, g: () -> n }; var p = P("a"); var g = p.g; g()`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// 若 Evaluate panic，测试直接失败（不会被 recover 掩盖）
			_, err := Evaluate(tc.src, nil)
			assert.Error(t, err, "src: %s", tc.src)
		})
	}
}

// TestEvaluateNoPanicOnAnyInput 抽查一组历史上会导致编译期 panic 穿透的输入。
func TestEvaluateNoPanicOnAnyInput(t *testing.T) {
	for _, src := range []string{"", " ", ";", "1 }", "} else {", "if (", "for x in", "-> x", "a..b"} {
		assert.NotPanics(t, func() {
			_, _ = Evaluate(src, nil)
		}, "src: %s", src)
	}
}

// TestLambdaBlockCompoundTrailingExpression 回归：块体尾表达式为复合表达式时必须整体作为返回值。
// 修复前 { var t = 1; t + 1 } 被贪婪解析成语句 t + 尾表达式 +1，求值为 1 而非 2。
func TestLambdaBlockCompoundTrailingExpression(t *testing.T) {
	t.Run("plus", func(t *testing.T) {
		assertEvaluate(t, `var f = () -> { var t = 1; t + 1 }; f()`, nil, int64(2))
	})
	t.Run("minus", func(t *testing.T) {
		assertEvaluate(t, `var f = (x) -> { var t = x * 2; t - 1 }; f(5)`, nil, int64(9))
	})
	t.Run("simple_trailing_still_works", func(t *testing.T) {
		assertEvaluate(t, `var f = (x, y) -> { var sum = x + y; sum }; f(3, 4)`, nil, int64(7))
	})
	t.Run("block_as_primary", func(t *testing.T) {
		assertEvaluate(t, "{ var x = 10; var y = 20; x + y }", nil, int64(30))
	})
	t.Run("nested_block", func(t *testing.T) {
		assertEvaluate(t, "{ var a = { var b = 2; b * 3 }; a + 1 }", nil, int64(7))
	})
	t.Run("trailing_ternary_and_call", func(t *testing.T) {
		assertEvaluate(t, `var f = (n) -> { var t = n; t > 0 ? t * 2 : -t }; f(3)`, nil, int64(6))
	})
}
