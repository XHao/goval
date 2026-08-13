package eval

import "testing"

func TestControlFlow(t *testing.T) {
	t.Run("if_statement", func(t *testing.T) {
		// if 语句不返回值；if 内只能新建变量（单赋值下不能重赋值外部变量）
		assertEval(t, `var x = 5; if (x > 3) { var hit = 100 }; x`, int64(5))
	})
	t.Run("if_else", func(t *testing.T) {
		// 各分支独立作用域，var 合法
		src := `var x = 1; var out = 0;
		        if (x > 3) { var hit = 100 } else { var hit = 200 };
		        out`
		assertEval(t, src, int64(0))
	})
	t.Run("if_else_if_chain", func(t *testing.T) {
		// if/else if/else 链：进入命中分支的局部 var 合法
		src := `var x = 2; var out = 0;
		        if (x == 1) { var h = 10 } else if (x == 2) { var h = 20 } else { var h = 30 };
		        out`
		assertEval(t, src, int64(0))
	})
	t.Run("if_else_if_block_value", func(t *testing.T) {
		// 表达式块作为语句返回值：if/else if 链进入命中分支
		assertEval(t, "if (false) { 1 } else if (true) { 2 } else { 3 }", int64(2))
		assertEval(t, "if (false) { 1 } else if (false) { 2 } else { 3 }", int64(3))
		assertEval(t, "if (true) { 1 } else if (true) { 2 } else { 3 }", int64(1))
	})
	t.Run("if_rebind_rejected", func(t *testing.T) {
		// 单赋值：if 内重赋值外部变量应编译报错
		assertEvalError(t, `var x = 5; var a = 0; if (x > 3) { a = 100 }`, "rebind")
	})

	// for 循环：goval 单赋值语义下，循环体内不能重赋值外部变量，
	// 因此无法在循环体内累加结果到外部。for 的行为验证依赖 reduce/map/filter
	// （它们内部实现累加）。这里验证 for 的语法合法性、迭代绑定、局部作用域。
	t.Run("for_in_list_local_var", func(t *testing.T) {
		// 循环体内声明局部变量合法；每轮新作用域
		assertEval(t, `for x in [1, 2, 3] { var local = x }; 0`, int64(0))
	})
	t.Run("for_in_string_local_var", func(t *testing.T) {
		// for x in string：迭代为单字符 string，绑定到局部
		assertEval(t, `for ch in "abc" { var c = ch }; 0`, int64(0))
	})
	t.Run("for_in_map_key", func(t *testing.T) {
		// for k in map：k 绑定为 key string
		assertEval(t, `for k in {"a": 1, "b": 2} { var key = k }; 0`, int64(0))
	})
	t.Run("for_kv_in_map", func(t *testing.T) {
		// for k, v in map：k=key string, v=value(int)，两者均可绑定到局部。
		// v + 1 不报错即证明 v 绑定的是 value(int) 而非 key(string)。
		assertEval(t, `for k, v in {"a": 1, "b": 2} { var kk = k; var sum = v + 1 }; 0`, int64(0))
	})
	t.Run("for_rebind_rejected", func(t *testing.T) {
		// 循环体内重赋值外部变量应编译报错
		assertEvalError(t, "var s = 0; for x in [1,2] { s = s + x }", "rebind")
	})
	t.Run("for_iteration_via_reduce", func(t *testing.T) {
		// 间接验证 for-in 遍历顺序：用 reduce 累加（reduce 内部处理累加）
		assertEval(t, "reduce([1, 2, 3, 4], 0, (acc, x) -> acc + x)", int64(10))
	})
	t.Run("for_iteration_via_map", func(t *testing.T) {
		// 间接验证遍历：map 变换每个元素
		assertEval(t, "map([1, 2, 3], x -> x * 10)", []interface{}{int64(10), int64(20), int64(30)})
	})

	// break/continue：验证在循环体内合法、不 panic。
	// 单赋值下无法验证 break 的「提前终止」副作用（无法累加到外部），
	// 故仅验证语法合法性与求值不报错。
	t.Run("break_compiles", func(t *testing.T) {
		assertEval(t, `for x in [1, 2, 3] { if (x == 2) { break } }; 0`, int64(0))
	})
	t.Run("continue_compiles", func(t *testing.T) {
		assertEval(t, `for x in [1, 2, 3] { if (x == 2) { continue }; var s = x }; 0`, int64(0))
	})
	t.Run("break_nested_inner_only", func(t *testing.T) {
		// break 只中断内层：外层继续迭代，不 panic
		assertEval(t, `for x in [1, 2] { for y in [1, 2, 3] { if (y == 2) { break } } }; 0`, int64(0))
	})
	t.Run("continue_nested_inner_only", func(t *testing.T) {
		assertEval(t, `for x in [1, 2] { for y in [1, 2, 3] { if (y == 2) { continue } } }; 0`, int64(0))
	})

	t.Run("expression_block", func(t *testing.T) {
		assertEval(t, "{ var x = 1; x + 1 }", int64(2))
	})
	t.Run("expression_block_nested", func(t *testing.T) {
		assertEval(t, "var r = { var t = 5; t * 2 }; r", int64(10))
	})
	t.Run("statement_block_scope", func(t *testing.T) {
		// 块内 var 不泄漏到外层
		assertEval(t, "var x = 1; { var y = 2 }; x", int64(1))
	})
}
