package eval

import "unicode/utf8"

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
	}
}

func builtin(name string, fn BuiltinFunc) Value {
	_ = name
	return Value{kind: kindBuiltin, builtin: fn}
}

// callLambda 调用 lambda，供内置函数使用。
// 新建 callEnv（parent 指向 l.env），绑定参数，调 l.body。
func callLambda(l *Lambda, args []Value) Value {
	callEnv := NewEnv(l.env)
	for i, p := range l.params {
		callEnv.Set(p, args[i])
	}
	return l.body(callEnv)
}

func builtinReduce() Value {
	return builtin("reduce", func(args []Value) Value {
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
		lst := args[0]
		l := args[1].fn
		var result []Value
		for _, item := range lst.list {
			if callLambda(l, []Value{item}).b {
				result = append(result, item)
			}
		}
		return ListValue(result)
	})
}

func builtinFind() Value {
	return builtin("find", func(args []Value) Value {
		lst := args[0]
		l := args[1].fn
		for _, item := range lst.list {
			if callLambda(l, []Value{item}).b {
				return item
			}
		}
		return NullValue()
	})
}

func builtinAppend() Value {
	return builtin("append", func(args []Value) Value {
		lst := args[0]
		newLst := make([]Value, len(lst.list)+1)
		copy(newLst, lst.list)
		newLst[len(lst.list)] = args[1]
		return ListValue(newLst)
	})
}

func builtinPut() Value {
	return builtin("put", func(args []Value) Value {
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
		lst := args[0]
		idx := args[1].i
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
		start := args[0].i
		end := args[1].i
		result := make([]Value, 0, end-start)
		for i := start; i < end; i++ {
			result = append(result, IntValue(i))
		}
		return ListValue(result)
	})
}
