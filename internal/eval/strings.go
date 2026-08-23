package eval

import (
	"strings"
	"unicode/utf8"
)

// stringMethod 是 string 内建方法的实现签名：receiver + 参数列表 → 返回值。
// 参数个数与类型由各方法自行校验，错误经 evalErrorf panic 抛出（与内置函数一致）。
type stringMethod func(recv string, args []Value) Value

// stringMethods 是 string 的内建方法表，在方法调用分发处（compiler.go）接入。
// 所有索引均为 rune 索引，与 s[i] / len(s) / for ch in s 的语义一致。
var stringMethods = map[string]stringMethod{
	"upper":       methodUpper,
	"lower":       methodLower,
	"trim":        methodTrim,
	"contains":    methodContains,
	"startsWith":  methodStartsWith,
	"endsWith":    methodEndsWith,
	"indexOf":     methodIndexOf,
	"lastIndexOf": methodLastIndexOf,
	"substring":   methodSubstring,
	"replace":     methodReplace,
	"repeat":      methodRepeat,
	"reverse":     methodReverse,
	"split":       methodSplit,
	"join":        methodJoin,
}

// callStringMethod 查表调用 string 方法；方法名不存在时报运行时错误。
func callStringMethod(name string, recv Value, args []Value) Value {
	method, ok := stringMethods[name]
	if !ok {
		panic(evalErrorf(0, 0, "no method '%s' on string", name))
	}
	return method(recv.s, args)
}

// checkArity 校验参数个数在 [min, max] 区间（max < 0 表示不设上限）。
func checkArity(method string, args []Value, min, max int) {
	if len(args) < min || (max >= 0 && len(args) > max) {
		panic(evalErrorf(0, 0, "%s() wrong number of arguments: got %d", method, len(args)))
	}
}

// wantStringArg 校验第 i 个参数为 string 并返回其值。
func wantStringArg(method string, args []Value, i int) string {
	if i >= len(args) || !args[i].IsString() {
		panic(evalErrorf(0, 0, "%s() expects a string argument, got %s", method, argKind(args, i)))
	}
	return args[i].s
}

// wantIntArg 校验第 i 个参数为 int 并返回其值。
func wantIntArg(method string, args []Value, i int) int64 {
	if i >= len(args) || !args[i].IsInt() {
		panic(evalErrorf(0, 0, "%s() expects an int argument, got %s", method, argKind(args, i)))
	}
	return args[i].i
}

// argKind 返回第 i 个参数的类型名；缺参时为 "missing"。
func argKind(args []Value, i int) string {
	if i >= len(args) {
		return "missing"
	}
	return kindName(args[i])
}

// byteToRuneIndex 把字节偏移换算成 rune 索引（调用方保证 byteIdx 在合法范围内）。
func byteToRuneIndex(s string, byteIdx int) int64 {
	return int64(utf8.RuneCountInString(s[:byteIdx]))
}

func methodUpper(recv string, args []Value) Value {
	checkArity("upper", args, 0, 0)
	return StringValue(strings.ToUpper(recv))
}

func methodLower(recv string, args []Value) Value {
	checkArity("lower", args, 0, 0)
	return StringValue(strings.ToLower(recv))
}

func methodTrim(recv string, args []Value) Value {
	checkArity("trim", args, 0, 0)
	return StringValue(strings.TrimSpace(recv))
}

func methodContains(recv string, args []Value) Value {
	checkArity("contains", args, 1, 1)
	return BoolValue(strings.Contains(recv, wantStringArg("contains", args, 0)))
}

func methodStartsWith(recv string, args []Value) Value {
	checkArity("startsWith", args, 1, 1)
	return BoolValue(strings.HasPrefix(recv, wantStringArg("startsWith", args, 0)))
}

func methodEndsWith(recv string, args []Value) Value {
	checkArity("endsWith", args, 1, 1)
	return BoolValue(strings.HasSuffix(recv, wantStringArg("endsWith", args, 0)))
}

func methodIndexOf(recv string, args []Value) Value {
	checkArity("indexOf", args, 1, 1)
	sub := wantStringArg("indexOf", args, 0)
	byteIdx := strings.Index(recv, sub)
	if byteIdx < 0 {
		return IntValue(-1)
	}
	return IntValue(byteToRuneIndex(recv, byteIdx))
}

func methodLastIndexOf(recv string, args []Value) Value {
	checkArity("lastIndexOf", args, 1, 1)
	sub := wantStringArg("lastIndexOf", args, 0)
	byteIdx := strings.LastIndex(recv, sub)
	if byteIdx < 0 {
		return IntValue(-1)
	}
	return IntValue(byteToRuneIndex(recv, byteIdx))
}

// methodSubstring 返回 [start, end) 的子串（rune 索引）；end 缺省为串尾。
// 负数、越界、start > end 均为运行时错误（与 s[i] 的严格风格一致）。
func methodSubstring(recv string, args []Value) Value {
	checkArity("substring", args, 1, 2)
	r := []rune(recv)
	n := len(r)
	start := wantIntArg("substring", args, 0)
	end := int64(n)
	if len(args) == 2 {
		end = wantIntArg("substring", args, 1)
	}
	if start < 0 || end > int64(n) || start > end {
		panic(evalErrorf(0, 0, "substring index out of range: start %d, end %d (len %d)", start, end, n))
	}
	return StringValue(string(r[start:end]))
}

func methodReplace(recv string, args []Value) Value {
	checkArity("replace", args, 2, 2)
	old := wantStringArg("replace", args, 0)
	sub := wantStringArg("replace", args, 1)
	return StringValue(strings.ReplaceAll(recv, old, sub))
}

func methodRepeat(recv string, args []Value) Value {
	checkArity("repeat", args, 1, 1)
	n := wantIntArg("repeat", args, 0)
	if n < 0 {
		panic(evalErrorf(0, 0, "repeat() negative count: %d", n))
	}
	return StringValue(strings.Repeat(recv, int(n)))
}

func methodReverse(recv string, args []Value) Value {
	checkArity("reverse", args, 0, 0)
	r := []rune(recv)
	out := make([]rune, len(r))
	for i, c := range r {
		out[len(r)-1-i] = c
	}
	return StringValue(string(out))
}

// methodSplit 按分隔符拆分为 List；空分隔符按 rune 拆分（与 strings.Split 一致）。
func methodSplit(recv string, args []Value) Value {
	checkArity("split", args, 1, 1)
	sep := wantStringArg("split", args, 0)
	parts := strings.Split(recv, sep)
	lst := make([]Value, len(parts))
	for i, p := range parts {
		lst[i] = StringValue(p)
	}
	return ListValue(lst)
}

// methodJoin 以 receiver 为分隔符拼接 List；元素必须全部是 string。
func methodJoin(recv string, args []Value) Value {
	checkArity("join", args, 1, 1)
	if !args[0].IsList() {
		panic(evalErrorf(0, 0, "join() expects a list argument, got %s", kindName(args[0])))
	}
	parts := make([]string, len(args[0].list))
	for i, item := range args[0].list {
		if !item.IsString() {
			panic(evalErrorf(0, 0, "join() expects list elements to be strings, got %s", kindName(item)))
		}
		parts[i] = item.s
	}
	return StringValue(strings.Join(parts, recv))
}
