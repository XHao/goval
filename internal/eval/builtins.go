package eval

import (
	"sort"
	"unicode/utf8"
)

// BuiltinFunc 是内置函数的 Go 实现签名。
type BuiltinFunc func(args []Value) Value

// defaultBuiltins 返回默认内置函数表。
func defaultBuiltins() map[string]Value {
	return map[string]Value{
		"reduce":   builtinReduce(),
		"map":      builtinMap(),
		"filter":   builtinFilter(),
		"find":     builtinFind(),
		"append":   builtinAppend(),
		"put":      builtinPut(),
		"removeAt": builtinRemoveAt(),
		"len":      builtinLen(),
		"range":    builtinRange(),
		"int":      builtinInt(),
		"float":    builtinFloat(),
		"string":   builtinString(),
		"abs":      builtinAbs(),
		"round":    builtinRound(),
		"floor":    builtinFloor(),
		"ceil":     builtinCeil(),
		"min":      minMax("min", func(cur, cand float64) bool { return cand < cur }),
		"max":      minMax("max", func(cur, cand float64) bool { return cand > cur }),
		"keys":     builtinKeys(),
		"values":   builtinValues(),
	}
}

func builtin(name string, fn BuiltinFunc) Value {
	_ = name
	return Value{kind: kindBuiltin, builtin: fn}
}

// wantArgs 校验内建函数参数个数，错误格式与 callValueWithReceiver 一致。
func wantArgs(name string, args []Value, n int) {
	if len(args) != n {
		panic(evalErrorf(0, 0, "%s expected %d args, got %d", name, n, len(args)))
	}
}

// want 校验内建函数参数类型，错误信息带函数名与实际类型。
func want(name, what string, ok bool, v Value) {
	if !ok {
		panic(evalErrorf(0, 0, "%s requires %s, got %s", name, what, kindName(v)))
	}
}

// callLambda 调用 lambda，供内置函数使用。
// 新建 callEnv（parent 指向 l.env），绑定参数，调 l.body。
func callLambda(l *Lambda, args []Value) Value {
	if len(args) != len(l.params) {
		panic(evalErrorf(0, 0, "expected %d args, got %d", len(l.params), len(args)))
	}
	callEnv := NewEnv(l.env)
	for i, p := range l.params {
		callEnv.Set(p, args[i])
	}
	return l.body(callEnv)
}

// callValueWithReceiver 以方法调用语义调用 lambda：在参数层同时绑定 this。
// builtin 函数不接收 this。供 obj.method(args) 编译路径使用。
func callValueWithReceiver(callee Value, recv Value, args []Value) Value {
	if callee.IsBuiltin() {
		return callee.builtin(args)
	}
	if !callee.IsLambda() {
		panic(evalErrorf(0, 0, "cannot call %s", kindName(callee)))
	}
	l := callee.fn
	if len(args) != len(l.params) {
		panic(evalErrorf(0, 0, "expected %d args, got %d", len(l.params), len(args)))
	}
	callEnv := NewEnv(l.env)
	callEnv.Set("this", recv)
	for i, p := range l.params {
		callEnv.Set(p, args[i])
	}
	return l.body(callEnv)
}

func builtinReduce() Value {
	return builtin("reduce", func(args []Value) Value {
		wantArgs("reduce", args, 3)
		want("reduce", "a list", args[0].IsList(), args[0])
		want("reduce", "a lambda", args[2].IsLambda(), args[2])
		lst := args[0]
		acc := args[1]
		l := args[2].fn
		for _, item := range lst.list {
			acc = callLambda(l, []Value{acc, item})
		}
		return acc
	})
}

func builtinMap() Value {
	return builtin("map", func(args []Value) Value {
		wantArgs("map", args, 2)
		want("map", "a list", args[0].IsList(), args[0])
		want("map", "a lambda", args[1].IsLambda(), args[1])
		lst := args[0]
		l := args[1].fn
		result := make([]Value, len(lst.list))
		for i, item := range lst.list {
			result[i] = callLambda(l, []Value{item})
		}
		return ListValue(result)
	})
}

func builtinFilter() Value {
	return builtin("filter", func(args []Value) Value {
		wantArgs("filter", args, 2)
		want("filter", "a list", args[0].IsList(), args[0])
		want("filter", "a lambda", args[1].IsLambda(), args[1])
		lst := args[0]
		l := args[1].fn
		var result []Value
		for _, item := range lst.list {
			res := callLambda(l, []Value{item})
			if !res.IsBool() {
				panic(evalErrorf(0, 0, "filter callback must return bool, got %s", kindName(res)))
			}
			if res.b {
				result = append(result, item)
			}
		}
		return ListValue(result)
	})
}

func builtinFind() Value {
	return builtin("find", func(args []Value) Value {
		wantArgs("find", args, 2)
		want("find", "a list", args[0].IsList(), args[0])
		want("find", "a lambda", args[1].IsLambda(), args[1])
		lst := args[0]
		l := args[1].fn
		for _, item := range lst.list {
			res := callLambda(l, []Value{item})
			if !res.IsBool() {
				panic(evalErrorf(0, 0, "find callback must return bool, got %s", kindName(res)))
			}
			if res.b {
				return item
			}
		}
		return NullValue()
	})
}

func builtinAppend() Value {
	return builtin("append", func(args []Value) Value {
		wantArgs("append", args, 2)
		want("append", "a list", args[0].IsList(), args[0])
		lst := args[0]
		newLst := make([]Value, len(lst.list)+1)
		copy(newLst, lst.list)
		newLst[len(lst.list)] = args[1]
		return ListValue(newLst)
	})
}

func builtinPut() Value {
	return builtin("put", func(args []Value) Value {
		wantArgs("put", args, 3)
		want("put", "a map", args[0].IsMap(), args[0])
		want("put", "a string key", args[1].IsString(), args[1])
		m := args[0]
		newM := make(map[string]Value, len(m.m)+1)
		for k, v := range m.m {
			newM[k] = v
		}
		newM[args[1].s] = args[2]
		return MapValue(newM)
	})
}

func builtinRemoveAt() Value {
	return builtin("removeAt", func(args []Value) Value {
		wantArgs("removeAt", args, 2)
		want("removeAt", "a list", args[0].IsList(), args[0])
		want("removeAt", "an int index", args[1].IsInt(), args[1])
		lst := args[0]
		idx := args[1].i
		if idx < 0 || idx >= int64(len(lst.list)) {
			panic(evalErrorf(0, 0, "removeAt index out of range: %d (len %d)", idx, len(lst.list)))
		}
		newLst := make([]Value, 0, len(lst.list)-1)
		for i, item := range lst.list {
			if i != int(idx) {
				newLst = append(newLst, item)
			}
		}
		return ListValue(newLst)
	})
}

func builtinLen() Value {
	return builtin("len", func(args []Value) Value {
		wantArgs("len", args, 1)
		v := args[0]
		if v.IsList() {
			return IntValue(int64(len(v.list)))
		}
		if v.IsString() {
			return IntValue(int64(utf8.RuneCountInString(v.s)))
		}
		if v.IsMap() {
			return IntValue(int64(len(v.m)))
		}
		return IntValue(0)
	})
}

func builtinRange() Value {
	return builtin("range", func(args []Value) Value {
		wantArgs("range", args, 2)
		want("range", "int bounds", args[0].IsInt(), args[0])
		want("range", "int bounds", args[1].IsInt(), args[1])
		start := args[0].i
		end := args[1].i
		// start >= end 一律返回空列表（避免 make 负容量 panic）
		if end <= start {
			return ListValue([]Value{})
		}
		result := make([]Value, 0, end-start)
		for i := start; i < end; i++ {
			result = append(result, IntValue(i))
		}
		return ListValue(result)
	})
}

// sortedKeys 返回 Map 键的字典序排序。Go map 迭代序随机，
// keys/values 必须给确定性输出（规则引擎可复现）。
func sortedKeys(m map[string]Value) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}

func builtinKeys() Value {
	return builtin("keys", func(args []Value) Value {
		wantArgs("keys", args, 1)
		want("keys", "a map", args[0].IsMap(), args[0])
		ks := sortedKeys(args[0].m)
		lst := make([]Value, len(ks))
		for i, k := range ks {
			lst[i] = StringValue(k)
		}
		return ListValue(lst)
	})
}

func builtinValues() Value {
	return builtin("values", func(args []Value) Value {
		wantArgs("values", args, 1)
		want("values", "a map", args[0].IsMap(), args[0])
		ks := sortedKeys(args[0].m)
		lst := make([]Value, len(ks))
		for i, k := range ks {
			lst[i] = args[0].m[k]
		}
		return ListValue(lst)
	})
}
