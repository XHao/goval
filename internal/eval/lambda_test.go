package eval

import "testing"

func TestLambda(t *testing.T) {
	t.Run("multi_param", func(t *testing.T) {
		assertEval(t, "((a, b) -> a + b)(3, 4)", int64(7))
	})
	t.Run("single_param_no_paren", func(t *testing.T) {
		assertEval(t, "var inc = x -> x + 1; inc(10)", int64(11))
	})
	t.Run("zero_param", func(t *testing.T) {
		assertEval(t, "var f = () -> 42; f()", int64(42))
	})
	t.Run("block_body", func(t *testing.T) {
		// lambda 块体：执行语句后返回尾表达式（单变量形式）
		assertEval(t, "((x) -> { var t = x * 2; t })(5)", int64(10))
		assertEval(t, "((x) -> { var t = x + 1; t })(5)", int64(6))
	})
	t.Run("block_body_compound_tail", func(t *testing.T) {
		// 块体尾表达式为复合运算时必须整体作为返回值
		// （依赖 expressionBlock 的非贪婪语句循环，回归钉死）
		assertEval(t, "((x) -> { var t = x; t + 1 })(5)", int64(6))
	})
	t.Run("closure_capture", func(t *testing.T) {
		// 闭包捕获定义点环境
		src := `var adder = (base) -> (x) -> base + x;
		        adder(10)(5)`
		assertEval(t, src, int64(15))
	})
	t.Run("closure_capture_value", func(t *testing.T) {
		// 捕获定义点的值
		src := `var n = 1;
		        var f = () -> n;
		        f()`
		assertEval(t, src, int64(1))
	})
	t.Run("higher_order_param", func(t *testing.T) {
		// lambda 作为参数
		assertEval(t, "((f, x) -> f(x))((y) -> y * 3, 5)", int64(15))
	})
	t.Run("higher_order_return", func(t *testing.T) {
		// lambda 作为返回值
		src := `var mk = (n) -> (x) -> x * n;
		        mk(4)(2)`
		assertEval(t, src, int64(8))
	})
	t.Run("object_factory", func(t *testing.T) {
		src := `var Person = (name) -> {
		            name: name,
		            greet: () -> "hi " + name
		        };
		        var p = Person("alice");
		        p.greet()`
		assertEval(t, src, "hi alice")
	})
	t.Run("currying", func(t *testing.T) {
		// 柯里化：多次捕获
		src := `var add = (a) -> (b) -> (c) -> a + b + c;
		        add(1)(2)(3)`
		assertEval(t, src, int64(6))
	})
}
