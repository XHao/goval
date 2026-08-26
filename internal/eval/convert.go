package eval

import "strconv"

// 类型转换内建：严格语义——不做隐式转换，失败即运行时错误。

func builtinInt() Value {
	return builtin("int", func(args []Value) Value {
		wantArgs("int", args, 1)
		v := args[0]
		switch {
		case v.IsInt():
			return v
		case v.IsFloat():
			return IntValue(int64(v.f)) // 截断，Go 语义
		case v.IsString():
			n, err := strconv.ParseInt(v.s, 10, 64)
			if err != nil {
				panic(evalErrorf(0, 0, "int cannot parse %q", v.s))
			}
			return IntValue(n)
		}
		panic(evalErrorf(0, 0, "int requires int, float or string, got %s", kindName(v)))
	})
}

func builtinFloat() Value {
	return builtin("float", func(args []Value) Value {
		wantArgs("float", args, 1)
		v := args[0]
		switch {
		case v.IsFloat():
			return v
		case v.IsInt():
			return FloatValue(float64(v.i))
		case v.IsString():
			f, err := strconv.ParseFloat(v.s, 64)
			if err != nil {
				panic(evalErrorf(0, 0, "float cannot parse %q", v.s))
			}
			return FloatValue(f)
		}
		panic(evalErrorf(0, 0, "float requires int, float or string, got %s", kindName(v)))
	})
}

func builtinString() Value {
	return builtin("string", func(args []Value) Value {
		wantArgs("string", args, 1)
		v := args[0]
		switch {
		case v.IsString():
			return v
		case v.IsInt():
			return StringValue(strconv.FormatInt(v.i, 10))
		case v.IsFloat():
			// 'g' + -1：最短往返表示，string(5.0) → "5"
			return StringValue(strconv.FormatFloat(v.f, 'g', -1, 64))
		case v.IsBool():
			return StringValue(strconv.FormatBool(v.b))
		}
		panic(evalErrorf(0, 0, "string requires int, float, bool or string, got %s", kindName(v)))
	})
}
