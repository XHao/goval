package eval

import "testing"

// 大小写转换

func TestStringUpper(t *testing.T) {
	assertEval(t, `"hello".upper()`, "HELLO")
	assertEval(t, `"héllo".upper()`, "HÉLLO")
	assertEval(t, `"中文a".upper()`, "中文A")
	assertEval(t, `"".upper()`, "")
}

func TestStringLower(t *testing.T) {
	assertEval(t, `"WORLD".lower()`, "world")
	assertEval(t, `"WÉRLD".lower()`, "wérld")
	assertEval(t, `"中文A".lower()`, "中文a")
}

// 修剪

func TestStringTrim(t *testing.T) {
	assertEval(t, `"  hi  ".trim()`, "hi")
	assertEval(t, `"\t hi \n".trim()`, "hi")
	assertEval(t, `"hi".trim()`, "hi")
	assertEval(t, `"   ".trim()`, "")
}

// 查找判断

func TestStringContains(t *testing.T) {
	assertEval(t, `"hello".contains("ell")`, true)
	assertEval(t, `"hello".contains("xyz")`, false)
	assertEval(t, `"你好世界".contains("世界")`, true)
	assertEval(t, `"hello".contains("")`, true)
}

func TestStringStartsWith(t *testing.T) {
	assertEval(t, `"hello".startsWith("he")`, true)
	assertEval(t, `"hello".startsWith("lo")`, false)
	assertEval(t, `"你好".startsWith("你")`, true)
}

func TestStringEndsWith(t *testing.T) {
	assertEval(t, `"hello".endsWith("lo")`, true)
	assertEval(t, `"hello".endsWith("he")`, false)
	assertEval(t, `"你好".endsWith("好")`, true)
}

// 索引查找 —— rune 语义

func TestStringIndexOf(t *testing.T) {
	assertEval(t, `"hello".indexOf("l")`, 2)
	assertEval(t, `"hello".indexOf("z")`, -1)
	// rune 索引而非字节索引：世 是第 3 个字符（索引 2；字节索引会是 6）
	assertEval(t, `"你好世界".indexOf("世")`, 2)
	assertEval(t, `"hello".indexOf("hello")`, 0)
}

func TestStringLastIndexOf(t *testing.T) {
	assertEval(t, `"hello".lastIndexOf("l")`, 3)
	assertEval(t, `"hello".lastIndexOf("z")`, -1)
	assertEval(t, `"你好世界".lastIndexOf("好")`, 1) // rune 索引；字节索引会是 3
}

// 截取 —— rune 语义，越界报错

func TestStringSubstring(t *testing.T) {
	assertEval(t, `"hello".substring(1)`, "ello")
	assertEval(t, `"hello".substring(1, 3)`, "el")
	assertEval(t, `"hello".substring(0, 5)`, "hello")
	assertEval(t, `"你好世界".substring(1, 3)`, "好世")
	assertEval(t, `"hello".substring(5)`, "")
	assertEval(t, `"hello".substring(5, 5)`, "")
}

func TestStringSubstringErrors(t *testing.T) {
	assertEvalError(t, `"hello".substring(-1)`, "out of range")
	assertEvalError(t, `"hello".substring(6)`, "out of range")
	assertEvalError(t, `"hello".substring(0, 6)`, "out of range")
	assertEvalError(t, `"hello".substring(3, 1)`, "start")
}

// 变换

func TestStringReplace(t *testing.T) {
	assertEval(t, `"a-b-c".replace("-", "+")`, "a+b+c")
	assertEval(t, `"aaa".replace("a", "b")`, "bbb") // 全部替换
	assertEval(t, `"abc".replace("xyz", "1")`, "abc")
	assertEval(t, `"a-b".replace("", "+")`, "+a+-+b+")
}

func TestStringRepeat(t *testing.T) {
	assertEval(t, `"ab".repeat(3)`, "ababab")
	assertEval(t, `"ab".repeat(0)`, "")
	assertEval(t, `"你".repeat(2)`, "你你")
	assertEvalError(t, `"ab".repeat(-1)`, "negative")
}

func TestStringReverse(t *testing.T) {
	assertEval(t, `"abc".reverse()`, "cba")
	assertEval(t, `"你好".reverse()`, "好你")
	assertEval(t, `"".reverse()`, "")
}

// 分合

func TestStringSplit(t *testing.T) {
	assertEval(t, `"a,b,c".split(",")`, []interface{}{"a", "b", "c"})
	assertEval(t, `",a,".split(",")`, []interface{}{"", "a", ""})
	assertEval(t, `"abc".split("")`, []interface{}{"a", "b", "c"}) // 空分隔符按 rune 拆
	assertEval(t, `"你 好".split(" ")`, []interface{}{"你", "好"})
	assertEval(t, `"abc".split("xyz")`, []interface{}{"abc"})
}

func TestStringJoin(t *testing.T) {
	assertEval(t, `",".join(["a", "b", "c"])`, "a,b,c")
	assertEval(t, `",".join(["a"])`, "a")
	assertEval(t, `",".join([])`, "")
	assertEval(t, `"-".join(["你", "好"])`, "你-好")
}

// 链式调用与组合

func TestStringChaining(t *testing.T) {
	assertEval(t, `"  Hello  ".trim().lower()`, "hello")
	assertEval(t, `("a" + "b").upper()`, "AB")
	assertEval(t, `map("a,b".split(","), x -> x.upper())`, []interface{}{"A", "B"})
}

// string 方法出现在三元条件中（依赖解析器不拆语句 + 方法分发）
func TestStringMethodInTernary(t *testing.T) {
	assertEvalCtx(t, `level.startsWith("v") ? "vip" : "normal"`,
		map[string]Value{"level": StringValue("vip")}, "vip")
	assertEval(t, `"vip".startsWith("v") ? 1 : 2`, 1)
	assertEval(t, `len("a,b".split(",")) > 1 ? 1 : 2`, 1)
}

// 错误路径

func TestStringMethodErrors(t *testing.T) {
	assertEvalError(t, `"abc".noSuch()`, "no method 'noSuch' on string")
	assertEvalError(t, `"abc".contains(1)`, "contains")
	assertEvalError(t, `"abc".upper(1)`, "upper")
	assertEvalError(t, `",".join([1, 2])`, "join")
	assertEvalError(t, `"abc".substring("x")`, "substring")
	assertEvalError(t, `"abc".repeat("x")`, "repeat")
}

func TestStringMethodOnNonString(t *testing.T) {
	assertEvalError(t, `var n = 1; n.upper()`, "cannot call method 'upper' on int")
	assertEvalError(t, `var z = null; z.upper()`, "cannot call method 'upper' on null")
}
