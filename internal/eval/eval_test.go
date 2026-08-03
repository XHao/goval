package eval

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_Constructors(t *testing.T) {
	assert.Equal(t, kindInt, IntValue(42).kind)
	assert.Equal(t, int64(42), IntValue(42).i)

	assert.Equal(t, kindFloat, FloatValue(3.14).kind)
	assert.Equal(t, 3.14, FloatValue(3.14).f)

	assert.Equal(t, kindBool, BoolValue(true).kind)
	assert.True(t, BoolValue(true).b)

	assert.Equal(t, kindString, StringValue("hi").kind)
	assert.Equal(t, "hi", StringValue("hi").s)

	assert.Equal(t, kindNull, NullValue().kind)

	lst := ListValue([]Value{IntValue(1), IntValue(2)})
	assert.Equal(t, kindList, lst.kind)
	assert.Equal(t, 2, len(lst.list))

	m := MapValue(map[string]Value{"k": IntValue(1)})
	assert.Equal(t, kindMap, m.kind)
	assert.Equal(t, int64(1), m.m["k"].i)
}

func TestEnv_LookupAndDefine(t *testing.T) {
	root := NewRootEnv()
	e1 := root.Define("x", IntValue(1))

	// 定义后可查到
	v, ok := e1.Lookup("x")
	assert.True(t, ok)
	assert.Equal(t, int64(1), v.i)

	// 根 Env 不受影响（不可变）
	_, ok = root.Lookup("x")
	assert.False(t, ok)

	// 子作用域链
	e2 := NewEnv(e1).Define("y", IntValue(2))
	v, ok = e2.Lookup("x") // 从父链查到
	assert.True(t, ok)
	assert.Equal(t, int64(1), v.i)
	v, ok = e2.Lookup("y")
	assert.True(t, ok)
	assert.Equal(t, int64(2), v.i)
}

func valueToString(v Value) string {
	switch {
	case v.IsInt():
		return strconv.FormatInt(v.i, 10)
	case v.IsFloat():
		return strconv.FormatFloat(v.f, 'f', -1, 64)
	case v.IsBool():
		return strconv.FormatBool(v.b)
	case v.IsString():
		return v.s
	case v.IsNull():
		return "null"
	default:
		return "<value>"
	}
}

func TestCompile_Literals(t *testing.T) {
	cases := []struct{ src, expect string }{
		{"42", "42"},
		{"3.14", "3.14"},
		{`"hello"`, "hello"},
		{"true", "true"},
		{"null", "null"},
	}
	for _, c := range cases {
		fn, err := CompileString(c.src)
		assert.NoError(t, err, c.src)
		v := Run(fn, NewRootEnv())
		assert.Equal(t, c.expect, valueToString(v), c.src)
	}
}

func TestCompile_VariableBindAndRef(t *testing.T) {
	fn, err := CompileString("var x = 10; x")
	assert.NoError(t, err)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, int64(10), v.i)
}

func TestCompile_SingleAssignmentRebind(t *testing.T) {
	// 单赋值：同作用域内重绑定应编译报错
	_, err := CompileString("var x = 1; var x = 2")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rebind")
}

func TestCompile_Arithmetic(t *testing.T) {
	cases := []struct{ src string; expect int64 }{
		{"1 + 2", 3},
		{"10 - 3", 7},
		{"4 * 5", 20},
		{"20 / 4", 5},
		{"17 % 5", 2},
		{"1 + 2 * 3", 7},
		{"(1 + 2) * 3", 9},
	}
	for _, c := range cases {
		fn, err := CompileString(c.src)
		assert.NoError(t, err, c.src)
		v := Run(fn, NewRootEnv())
		assert.Equal(t, c.expect, v.i, c.src)
	}
}

func TestCompile_Comparison(t *testing.T) {
	fn, _ := CompileString("3 > 2")
	assert.True(t, Run(fn, NewRootEnv()).b)
	fn, _ = CompileString("1 == 1")
	assert.True(t, Run(fn, NewRootEnv()).b)
	fn, _ = CompileString("1 != 2")
	assert.True(t, Run(fn, NewRootEnv()).b)
}

func TestCompile_Logic(t *testing.T) {
	fn, _ := CompileString("true && false")
	assert.False(t, Run(fn, NewRootEnv()).b)
	fn, _ = CompileString("true || false")
	assert.True(t, Run(fn, NewRootEnv()).b)
	fn, _ = CompileString("!false")
	assert.True(t, Run(fn, NewRootEnv()).b)
}

func TestCompile_Ternary(t *testing.T) {
	fn, _ := CompileString("3 > 2 ? 100 : 200")
	assert.Equal(t, int64(100), Run(fn, NewRootEnv()).i)
	fn, _ = CompileString("1 > 2 ? 100 : 200")
	assert.Equal(t, int64(200), Run(fn, NewRootEnv()).i)
}

func TestCompile_Unary(t *testing.T) {
	fn, _ := CompileString("-5")
	assert.Equal(t, int64(-5), Run(fn, NewRootEnv()).i)
	fn, _ = CompileString("-3.14")
	assert.Equal(t, -3.14, Run(fn, NewRootEnv()).f)
}

func TestCompile_ListLiteral(t *testing.T) {
	fn, _ := CompileString("[1, 2, 3]")
	v := Run(fn, NewRootEnv())
	assert.True(t, v.IsList())
	assert.Equal(t, 3, len(v.list))
	assert.Equal(t, int64(2), v.list[1].i)
}

func TestCompile_ListIndex(t *testing.T) {
	fn, _ := CompileString("[10, 20, 30][1]")
	assert.Equal(t, int64(20), Run(fn, NewRootEnv()).i)
}

func TestCompile_MapLiteral(t *testing.T) {
	fn, _ := CompileString(`{"a": 1, "b": 2}`)
	v := Run(fn, NewRootEnv())
	assert.True(t, v.IsMap())
	assert.Equal(t, int64(1), v.m["a"].i)
}

func TestCompile_MapAccess(t *testing.T) {
	fn, _ := CompileString(`{"key": 42}["key"]`)
	assert.Equal(t, int64(42), Run(fn, NewRootEnv()).i)
}

func TestCompile_FieldAccess(t *testing.T) {
	// Map 字段访问用 . 语法
	fn, _ := CompileString(`{"name": "alice"}.name`)
	assert.Equal(t, "alice", Run(fn, NewRootEnv()).s)
}

func TestCompile_Lambda(t *testing.T) {
	fn, _ := CompileString("(x) -> x + 1")
	lambda := Run(fn, NewRootEnv())
	assert.True(t, lambda.IsLambda())
}

func TestCompile_LambdaCall(t *testing.T) {
	fn, _ := CompileString("((x) -> x * 2)(21)")
	assert.Equal(t, int64(42), Run(fn, NewRootEnv()).i)
}

func TestCompile_LambdaClosure(t *testing.T) {
	src := `var adder = (base) -> (x) -> base + x;
	        adder(10)(5)`
	fn, _ := CompileString(src)
	assert.Equal(t, int64(15), Run(fn, NewRootEnv()).i)
}

func TestCompile_ObjectFactory(t *testing.T) {
	src := `var Person = (name) -> {
		name: name,
		greet: () -> "hi " + name
	};
	var p = Person("alice");
	p.greet()`
	fn, _ := CompileString(src)
	assert.Equal(t, "hi alice", Run(fn, NewRootEnv()).s)
}

func TestCompile_MethodCall(t *testing.T) {
	src := `var o = {val: 10, double: () -> o.val * 2};
	        o.double()`
	// 注意：o 在 lambda 内引用——单赋值下 o 已绑定，闭包捕获可见
	fn, _ := CompileString(src)
	assert.Equal(t, int64(20), Run(fn, NewRootEnv()).i)
}

func TestCompile_IfElse(t *testing.T) {
	src := `var x = 5;
	        var a = 0;
	        if (x > 3) { a = 100; }`
	// 注意：单赋值下 a 不能在 if 内重新赋值
	_, err := CompileString(src)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rebind")
}

func TestCompile_IfStatement(t *testing.T) {
	// if 语句内 var 新绑定合法；if 不返回值
	src := `var x = 5;
	        if (x > 3) { var hit = 100; };
	        x`
	fn, err := CompileString(src)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), Run(fn, NewRootEnv()).i)
}

func TestCompile_IfElseBranch(t *testing.T) {
	// if/else 分支选择，各分支独立作用域
	src := `var x = 1;
	        if (x > 3) { var hit = 100; } else { var hit = 200; };
	        x`
	fn, err := CompileString(src)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), Run(fn, NewRootEnv()).i)
}

func TestCompile_ForBreak(t *testing.T) {
	// for-in 遍历，break 早退。循环体内 var 声明局部变量。
	src := `var result = 0;
	        for x in [1, 2, 3, 4, 5] {
	            var hit = x;
	        };
	        result`
	// result 始终 0（for 不修改外层）
	fn, err := CompileString(src)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), Run(fn, NewRootEnv()).i)
}

func TestCompile_ForNoRebind(t *testing.T) {
	// 循环体内重新赋值外层变量应编译报错
	_, err := CompileString("var s = 0; for x in [1,2] { s = s + x; }")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rebind")
}

func TestCompile_ForLocalVar(t *testing.T) {
	// 循环体内 var 声明局部变量（每轮新作用域）合法
	src := `for x in [1, 2, 3] { var local = x; }; 0`
	fn, err := CompileString(src)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), Run(fn, NewRootEnv()).i)
}

func TestCompile_ForBreakEarly(t *testing.T) {
	// for-in 配合 break 早退，循环体内 var 声明局部变量
	src := `for x in [1, 2, 3, 4, 5] {
	            if (x == 3) { var found = x; break; }
	        }; 0`
	fn, err := CompileString(src)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), Run(fn, NewRootEnv()).i)
}

func TestCompile_ForContinue(t *testing.T) {
	// for-in 配合 continue 跳过当轮
	src := `for x in [1, 2, 3, 4, 5] {
	            if (x == 3) { continue; }
	            var skip = x;
	        }; 0`
	fn, err := CompileString(src)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), Run(fn, NewRootEnv()).i)
}

func TestCompile_ForInString(t *testing.T) {
	// for-in 遍历 string
	src := `for ch in "abc" { var c = ch; }; 0`
	fn, err := CompileString(src)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), Run(fn, NewRootEnv()).i)
}

func TestCompile_ForInMap(t *testing.T) {
	// for-in 遍历 Map（key）
	src := `for k in {"a": 1, "b": 2} { var key = k; }; 0`
	fn, err := CompileString(src)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), Run(fn, NewRootEnv()).i)
}

func TestBuiltin_Reduce(t *testing.T) {
	src := `var sum = reduce([1, 2, 3, 4], 0, (acc, x) -> acc + x); sum`
	fn, _ := CompileString(src)
	assert.Equal(t, int64(10), Run(fn, NewRootEnv()).i)
}

func TestBuiltin_Map(t *testing.T) {
	src := `map([1, 2, 3], (x) -> x * 2)`
	fn, _ := CompileString(src)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, int64(2), v.list[0].i)
	assert.Equal(t, int64(6), v.list[2].i)
}

func TestBuiltin_Filter(t *testing.T) {
	src := `filter([1, 2, 3, 4, 5], (x) -> x > 2)`
	fn, _ := CompileString(src)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, 3, len(v.list))
	assert.Equal(t, int64(3), v.list[0].i)
}

func TestBuiltin_Find(t *testing.T) {
	src := `find([1, 2, 3, 4], (x) -> x > 2)`
	fn, _ := CompileString(src)
	assert.Equal(t, int64(3), Run(fn, NewRootEnv()).i)
}

func TestBuiltin_Append(t *testing.T) {
	src := `append([1, 2], 3)`
	fn, _ := CompileString(src)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, 3, len(v.list))
	assert.Equal(t, int64(3), v.list[2].i)
}

func TestBuiltin_Put(t *testing.T) {
	src := `put({"a": 1}, "b", 2)`
	fn, _ := CompileString(src)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, int64(2), v.m["b"].i)
}

func TestBuiltin_Len(t *testing.T) {
	fn, _ := CompileString("len([1, 2, 3])")
	assert.Equal(t, int64(3), Run(fn, NewRootEnv()).i)
	fn, _ = CompileString(`len("hello")`)
	assert.Equal(t, int64(5), Run(fn, NewRootEnv()).i)
	// 非 ASCII：rune 计数，与 rune-based 下标访问一致
	fn, _ = CompileString(`len("héllo")`)
	assert.Equal(t, int64(5), Run(fn, NewRootEnv()).i)
}

func TestBuiltin_Range(t *testing.T) {
	fn, _ := CompileString("range(1, 4)")
	v := Run(fn, NewRootEnv())
	assert.Equal(t, 3, len(v.list))
	assert.Equal(t, int64(1), v.list[0].i)
	assert.Equal(t, int64(3), v.list[2].i)
}

// catchEvalPanic runs fn and converts any panic (EvalError or other) into an error.
func catchEvalPanic(fn func(*Env) Value) (v Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			if ee, ok := r.(*EvalError); ok {
				err = ee
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	v = fn(NewRootEnv())
	return
}

func TestCompile_ExpressionBlock(t *testing.T) {
	fn, err := CompileString("{ var x = 1; x + 1 }")
	assert.NoError(t, err)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, int64(2), v.i)
}

func TestCompile_ExpressionBlockNested(t *testing.T) {
	// 块表达式作为子表达式
	fn, err := CompileString("var r = { var t = 5; t * 2 }; r")
	assert.NoError(t, err)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, int64(10), v.i)
}

func TestCompile_LogicTypeCheck(t *testing.T) {
	// 非 bool 操作数应报错
	fn, err := CompileString("1 || 2")
	assert.NoError(t, err) // 编译期不报错（类型在运行时）
	_, err = catchEvalPanic(fn)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "bool")
}

func TestCompile_LogicTypeCheckAnd(t *testing.T) {
	fn, _ := CompileString("\"a\" && true")
	_, err := catchEvalPanic(fn)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "bool")
}

func TestCompile_MapKeyNonString(t *testing.T) {
	// 非字符串 key（如数字）应报错
	fn, err := CompileString("{1: \"a\"}")
	assert.NoError(t, err)
	_, err = catchEvalPanic(fn)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "string")
}
