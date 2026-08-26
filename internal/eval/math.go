package eval

import "math"

// 数学内建：abs 保型；round half away from zero（可选小数位）；
// floor/ceil 返回 float；min/max 变参标量或单个 List。

func builtinAbs() Value {
	return builtin("abs", func(args []Value) Value {
		wantArgs("abs", args, 1)
		v := args[0]
		if v.IsInt() {
			if v.i < 0 {
				return IntValue(-v.i)
			}
			return v
		}
		if v.IsFloat() {
			return FloatValue(math.Abs(v.f))
		}
		panic(evalErrorf(0, 0, "abs requires numeric arg, got %s", kindName(v)))
	})
}

func builtinRound() Value {
	return builtin("round", func(args []Value) Value {
		if len(args) != 1 && len(args) != 2 {
			panic(evalErrorf(0, 0, "round expected 1 or 2 args, got %d", len(args)))
		}
		x := args[0]
		if !x.IsInt() && !x.IsFloat() {
			panic(evalErrorf(0, 0, "round requires numeric arg, got %s", kindName(x)))
		}
		f := toFloat(x)
		digits := 0
		if len(args) == 2 {
			if !args[1].IsInt() {
				panic(evalErrorf(0, 0, "round requires int digits, got %s", kindName(args[1])))
			}
			digits = int(args[1].i)
			if digits < 0 {
				panic(evalErrorf(0, 0, "round digits must be >= 0, got %d", digits))
			}
		}
		r := math.Round(f)
		if digits > 0 {
			p := math.Pow(10, float64(digits))
			r = math.Round(f*p) / p
		}
		return FloatValue(r)
	})
}

func builtinFloor() Value {
	return builtin("floor", func(args []Value) Value {
		wantArgs("floor", args, 1)
		v := args[0]
		if !v.IsInt() && !v.IsFloat() {
			panic(evalErrorf(0, 0, "floor requires numeric arg, got %s", kindName(v)))
		}
		return FloatValue(math.Floor(toFloat(v)))
	})
}

func builtinCeil() Value {
	return builtin("ceil", func(args []Value) Value {
		wantArgs("ceil", args, 1)
		v := args[0]
		if !v.IsInt() && !v.IsFloat() {
			panic(evalErrorf(0, 0, "ceil requires numeric arg, got %s", kindName(v)))
		}
		return FloatValue(math.Ceil(toFloat(v)))
	})
}

// minMax 实现 min/max：至少 2 个数值参数，或恰好 1 个 List 参数（作用于其元素）。
// 全 int 保持 int；混入 float 提升为 float。
func minMax(name string, better func(cur, cand float64) bool) Value {
	return builtin(name, func(args []Value) Value {
		var items []Value
		switch {
		case len(args) == 1 && args[0].IsList():
			items = args[0].list
			if len(items) == 0 {
				panic(evalErrorf(0, 0, "%s of empty list", name))
			}
		case len(args) >= 2:
			items = args
		default:
			panic(evalErrorf(0, 0, "%s expected a list or at least 2 args, got %d", name, len(args)))
		}
		best := items[0]
		if !best.IsInt() && !best.IsFloat() {
			panic(evalErrorf(0, 0, "%s requires numeric args, got %s", name, kindName(best)))
		}
		hasFloat := best.IsFloat()
		for _, it := range items[1:] {
			if !it.IsInt() && !it.IsFloat() {
				panic(evalErrorf(0, 0, "%s requires numeric args, got %s", name, kindName(it)))
			}
			hasFloat = hasFloat || it.IsFloat()
			if better(toFloat(best), toFloat(it)) {
				best = it
			}
		}
		if hasFloat {
			return FloatValue(toFloat(best))
		}
		return best
	})
}
