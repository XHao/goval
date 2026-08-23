package eval

import "testing"

// TestV2ClosureSemantics 钉死 v2 引用捕获语义：闭包与定义作用域共享变量绑定。
func TestV2ClosureSemantics(t *testing.T) {
	t.Run("outer_write_inner_read", func(t *testing.T) {
		// 外部赋值后，闭包读到新值
		assertEval(t, `var n = 1; var f = () -> n; n = 2; f()`, int64(2))
	})
	t.Run("inner_write_outer_read", func(t *testing.T) {
		// 闭包内赋值修改外层绑定
		assertEval(t, `var n = 1; var f = () -> { n = 5; n }; f(); n`, int64(5))
	})
	t.Run("sibling_closures_share_binding", func(t *testing.T) {
		// 兄弟闭包捕获同一环境层，透过一个的赋值另一个可见
		assertEval(t, `var x = 0
		               var set = (v) -> { x = v; x }
		               var get = () -> x
		               set(42)
		               get()`, int64(42))
	})
	t.Run("assignment_is_expression", func(t *testing.T) {
		// 赋值表达式返回所赋的值
		assertEval(t, `var x = 0; var y = (x = 5) + 1; x + y`, int64(11))
	})
	t.Run("param_assign_local_only", func(t *testing.T) {
		// 参数槽位可赋值，但每次调用新建层，不跨调用持久、不影响外层
		assertEval(t, `var f = (x) -> { x = x + 1; x }
		               var a = f(1)
		               var b = f(1)
		               a + b`, int64(4))
	})
	t.Run("param_shadows_outer_completely", func(t *testing.T) {
		// 遮蔽是全遮蔽：参数赋值不碰外层同名变量
		assertEval(t, `var x = 1
		               var f = (x) -> { x = 99; x }
		               f(1)
		               x`, int64(1))
	})
	t.Run("recursion_via_self_reference", func(t *testing.T) {
		assertEval(t, `var fact = (n) -> n <= 1 ? 1 : n * fact(n - 1); fact(5)`, int64(120))
	})
	t.Run("lambda_identity_equality", func(t *testing.T) {
		// D2：同一闭包恒等；不同闭包即使源码相同也不等
		assertEval(t, `var f = () -> 1; f == f`, true)
		assertEval(t, `(() -> 1) == (() -> 1)`, false)
	})
}

// TestV2PerIterationCapture 钉死 D1：循环/多次调用中创建的闭包各持其层（每轮独立绑定）。
func TestV2PerIterationCapture(t *testing.T) {
	t.Run("map_callback", func(t *testing.T) {
		assertEval(t, `var fns = map(range(0, 3), (i) -> () -> i)
		               [fns[0](), fns[1](), fns[2]()]`,
			[]interface{}{int64(0), int64(1), int64(2)})
	})
	t.Run("for_loop_body", func(t *testing.T) {
		// for 每轮 NewEnv：循环体内创建的闭包捕获各自迭代的循环变量
		assertEval(t, `var fs = []
		               for i in [10, 20] { fs = append(fs, () -> i) }
		               [fs[0](), fs[1]()]`,
			[]interface{}{int64(10), int64(20)})
	})
	t.Run("for_accumulate", func(t *testing.T) {
		// v2 核心解锁：循环体内累加外部变量
		assertEval(t, `var s = 0; for x in [1, 2, 3] { s = s + x }; s`, int64(6))
	})
}

// TestV2ThisInjection 验证方法调用点的 this 绑定。
func TestV2ThisInjection(t *testing.T) {
	t.Run("explicit_this", func(t *testing.T) {
		assertEval(t, `var P = (n) -> { n: n, g: () -> this.n }; P("a").g()`, "a")
	})
	t.Run("this_field_write", func(t *testing.T) {
		// this.name = v 写入实例字段，字段视图立即更新
		assertEval(t, `var P = (n) -> { n: n, set: (v) -> { this.n = v; this.n } }
		               var p = P("a")
		               p.set("b")
		               p.n`, "b")
	})
	t.Run("method_calls_method", func(t *testing.T) {
		assertEval(t, `var P = (n) -> {
		                  n: n,
		                  label: () -> "<" + this.n + ">",
		                  wrap: () -> "[" + this.label() + "]"
		               }
		               P("x").wrap()`, "[<x>]")
	})
	t.Run("nested_lambda_sees_this", func(t *testing.T) {
		// this 是调用环境里的普通绑定，嵌套 lambda 沿链可见
		assertEval(t, `var P = (n) -> { n: n, f: () -> { var g = () -> this.n; g() } }
		               P("deep").f()`, "deep")
	})
	t.Run("subscript_call_binds_this", func(t *testing.T) {
		// m["f"]() 与 m.f() 语义对齐，同样注入接收者
		assertEval(t, `var m = { n: "k", f: () -> this.n }; m["f"]()`, "k")
	})
	t.Run("detached_method_no_this", func(t *testing.T) {
		// 分离方法没有 this：改写后的字段访问报错
		assertEvalError(t, `var P = (n) -> { n: n, g: () -> n }
		                    var p = P("a")
		                    var g = p.g
		                    g()`, "'this' is not bound")
	})
	t.Run("this_literal_unbound", func(t *testing.T) {
		assertEvalError(t, `this`, "'this' is not bound")
	})
}

// TestV2ImplicitThis 验证隐式 this 改写：方法体内撞字段名的裸名编译为 this.field。
func TestV2ImplicitThis(t *testing.T) {
	t.Run("read_omits_this", func(t *testing.T) {
		// 旧惯用法 greet: () -> name 现在经改写得到字段语义
		assertEval(t, `var Person = (name) -> { name: name, greet: () -> "hi " + name }
		               var p = Person("alice")
		               p.greet()`, "hi alice")
	})
	t.Run("write_hits_field_not_closure", func(t *testing.T) {
		// birthday：裸 age 赋值改写为 this.age 赋值，字段与方法视图一致
		assertEval(t, `var Person = (name, age) -> {
		                  name: name,
		                  age: age,
		                  birthday: () -> { age = age + 1; age }
		               }
		               var p = Person("alice", 30)
		               var newAge = p.birthday()
		               [newAge, p.age]`,
			[]interface{}{int64(31), int64(31)})
	})
	t.Run("construction_position_not_rewritten", func(t *testing.T) {
		// 构造位置 name: name 立即求值，读参数，不改写
		assertEval(t, `var n = "outer"; var m = { n: n }; m.n`, "outer")
	})
	t.Run("param_shadows_field", func(t *testing.T) {
		// 方法参数与字段同名：参数遮蔽，不改写
		assertEval(t, `var P = (n) -> { n: "field", f: (n) -> n }; P(0).f("param")`, "param")
	})
	t.Run("local_shadows_field", func(t *testing.T) {
		// 方法内局部 var 与字段同名：局部遮蔽
		assertEval(t, `var P = (n) -> { n: "field", f: () -> { var n = "local"; n } }
		               P(0).f()`, "local")
	})
	t.Run("method_with_args_on_this", func(t *testing.T) {
		assertEval(t, `var P = (name, age) -> {
		                  name: name, age: age,
		                  older: (o) -> age > o.age
		               }
		               P("a", 30).older(P("b", 20))`, true)
	})
	t.Run("nested_map_dispatch", func(t *testing.T) {
		// 嵌套 Map：inner.f() 调用时 this=inner，改写读 inner 的字段
		assertEval(t, `var M = { inner: { v: 1, f: () -> v }, g: () -> inner.f() }
		               M.g()`, int64(1))
	})
	t.Run("plain_lambda_unaffected", func(t *testing.T) {
		// 非 Map 语境的闭包照常捕获，不受改写影响
		assertEval(t, `var n = 7; var f = () -> n; f()`, int64(7))
	})
}

// TestV2Capture 验证 capture pragma：关闭词法范围内的隐式 this 改写，回归纯闭包语义。
func TestV2Capture(t *testing.T) {
	t.Run("map_capture_restores_closure", func(t *testing.T) {
		// capture：n 不改写 → 链查找 → 外层绑定
		assertEval(t, `var n = "outer"
		               var P = capture { n: "field", f: () -> n }
		               P.f()`, "outer")
	})
	t.Run("without_capture_field_wins", func(t *testing.T) {
		// 对照组：不写 capture，字段胜出
		assertEval(t, `var n = "outer"
		               var P = { n: "field", f: () -> n }
		               P.f()`, "field")
	})
	t.Run("capture_lambda_identity", func(t *testing.T) {
		assertEval(t, `var f = capture (x) -> x; f(3)`, int64(3))
	})
	t.Run("explicit_this_still_works_under_capture", func(t *testing.T) {
		assertEval(t, `var P = capture { n: "field", f: () -> this.n }; P.f()`, "field")
	})
	t.Run("counter_private_state_no_capture_needed", func(t *testing.T) {
		// 私有状态：闭包变量不撞字段名，无需 capture
		assertEval(t, `var Counter = (start) -> {
		                  var count = start
		                  {
		                     inc: () -> { count = count + 1; count },
		                     value: () -> count
		                  }
		               }
		               var c = Counter(0)
		               c.inc()
		               c.inc()
		               [c.value(), c.count]`,
			[]interface{}{int64(2), nil})
	})
	t.Run("counter_instances_independent", func(t *testing.T) {
		assertEval(t, `var Counter = (start) -> {
		                  var count = start
		                  { inc: () -> { count = count + 1; count } }
		               }
		               var a = Counter(0)
		               var b = Counter(10)
		               a.inc()
		               [a.inc(), b.inc()]`,
			[]interface{}{int64(2), int64(11)})
	})
	t.Run("namespace_self_reference", func(t *testing.T) {
		assertEval(t, `var Math2 = {
		                  square: (x) -> x * x,
		                  quad: (x) -> Math2.square(x) + Math2.square(x)
		               }
		               Math2.quad(3)`, int64(18))
	})
}

// TestV2Mutation 验证通用左值：字段/下标赋值与别名共享。
func TestV2Mutation(t *testing.T) {
	t.Run("field_assign", func(t *testing.T) {
		assertEval(t, `var p = { n: 1 }; p.n = 2; p.n`, int64(2))
	})
	t.Run("nested_field_assign", func(t *testing.T) {
		assertEval(t, `var a = { b: { c: 1 } }; a.b.c = 2; a.b.c`, int64(2))
	})
	t.Run("list_index_assign", func(t *testing.T) {
		assertEval(t, `var l = [1, 2]; l[0] = 9; l`, []interface{}{int64(9), int64(2)})
	})
	t.Run("map_key_assign", func(t *testing.T) {
		assertEval(t, `var m = {"k": 1}; m["k"] = 5; m["k"]`, int64(5))
	})
	t.Run("list_alias_shares_container", func(t *testing.T) {
		// 闭包/变量别名共享同一容器：原地修改互相可见
		assertEval(t, `var a = [1]; var b = a; a[0] = 9; b[0]`, int64(9))
	})
	t.Run("list_index_out_of_range", func(t *testing.T) {
		assertEvalError(t, `var l = [1]; l[1] = 2`, "out of range")
	})
	t.Run("string_immutable", func(t *testing.T) {
		assertEvalError(t, `var s = "ab"; s[0] = 1`, "cannot index-assign")
	})
	t.Run("field_assign_on_non_map", func(t *testing.T) {
		assertEvalError(t, `var x = 1; x.f = 2`, "cannot assign field")
	})
	t.Run("method_mutates_then_other_method_reads", func(t *testing.T) {
		// 改写字段后，另一个方法立即看到新值
		assertEval(t, `var Acc = (bal) -> {
		                  bal: bal,
		                  deposit: (amt) -> { bal = bal + amt; bal },
		                  balance: () -> bal
		               }
		               var a = Acc(100)
		               a.deposit(50)
		               a.balance()`, int64(150))
	})
}

// TestV2StrictAssignment 验证 D5 严格赋值：裸赋值要求目标已绑定。
func TestV2StrictAssignment(t *testing.T) {
	t.Run("context_global_assignable", func(t *testing.T) {
		// context 注入的变量可赋值更新（键名传入编译器根作用域）
		assertEvalCtx(t, `x = x + 1; x`, map[string]Value{"x": IntValue(10)}, int64(11))
	})
	t.Run("builtin_assignable", func(t *testing.T) {
		// 内置函数名在根作用域，赋值合法（覆盖）
		assertEval(t, `len = 5; len`, int64(5))
	})
	t.Run("if_condition_must_be_bool", func(t *testing.T) {
		// 非 bool 条件显式报错，不再静默当 false（与 &&/||/! 一致）
		assertEvalError(t, `if (1) { 2 } else { 3 }`, "if condition must be bool")
		assertEvalError(t, `var s = "x"; if (s) { 1 }`, "if condition must be bool")
	})
	t.Run("ternary_else_branch_assignment", func(t *testing.T) {
		// colon 分支放宽为完整 expression：else 分支裸赋值生效
		assertEval(t, `var x = 0; var y = 0
		               false ? x = 1 : y = 2
		               x + y`, int64(2))
	})
	t.Run("ternary_right_assoc_preserved", func(t *testing.T) {
		assertEval(t, `false ? 1 : true ? 2 : 3`, int64(2))
	})
	t.Run("ternary_in_map_value_unchanged", func(t *testing.T) {
		assertEval(t, `{ k: true ? 1 : 2 }.k`, int64(1))
	})
}

func TestTernaryConditionMustBeBool(t *testing.T) {
	// 三元条件与 if 条件一致：非 bool 显式报错，不静默当 false
	assertEvalError(t, `1 ? "a" : "b"`, "ternary condition must be bool")
	assertEvalError(t, `var s = "x"; s ? 1 : 2`, "ternary condition must be bool")
	assertEvalError(t, `var z = null; z ? 1 : 2`, "ternary condition must be bool")
	assertEval(t, `true ? 1 : 2`, int64(1))
}

func TestCallInTernaryCondition(t *testing.T) {
	// 带参调用/下标出现在三元条件中不得被拆成两条语句（f | (1)?1:2 形态）
	assertEval(t, `len([1, 2]) > 1 ? 1 : 2`, int64(1))
	assertEval(t, `var m = {"k": true}; m["k"] ? 1 : 2`, int64(1))
	assertEval(t, `var l = [true]; l[0] ? 1 : 2`, int64(1))
	assertEval(t, `var f = (x) -> x > 0; f(1) ? 10 : 20`, int64(10))
	assertEval(t, `var p = {f: (x) -> x > 0}; p.f(1) ? 10 : 20`, int64(10))
}
