package eval

func addValues(l, r Value) Value {
	if l.IsInt() && r.IsInt() {
		return IntValue(l.i + r.i)
	}
	if l.IsFloat() || r.IsFloat() {
		return FloatValue(toFloat(l) + toFloat(r))
	}
	if l.IsString() && r.IsString() {
		return StringValue(l.s + r.s)
	}
	panic(evalErrorf(0, 0, "cannot add %s and %s", kindName(l), kindName(r)))
}

func subValues(l, r Value) Value {
	if l.IsInt() && r.IsInt() {
		return IntValue(l.i - r.i)
	}
	if l.IsFloat() || r.IsFloat() {
		return FloatValue(toFloat(l) - toFloat(r))
	}
	panic(evalErrorf(0, 0, "cannot subtract %s from %s", kindName(r), kindName(l)))
}

func mulValues(l, r Value) Value {
	if l.IsInt() && r.IsInt() {
		return IntValue(l.i * r.i)
	}
	if l.IsFloat() || r.IsFloat() {
		return FloatValue(toFloat(l) * toFloat(r))
	}
	panic(evalErrorf(0, 0, "cannot multiply %s and %s", kindName(l), kindName(r)))
}

func divValues(l, r Value) Value {
	if l.IsInt() && r.IsInt() {
		if r.i == 0 {
			panic(evalErrorf(0, 0, "division by zero"))
		}
		return IntValue(l.i / r.i)
	}
	if l.IsFloat() || r.IsFloat() {
		// 浮点除零同样报错而非返回 ±Inf/NaN——静默的非有限值会流入
		// 后续比较与规则结果，与语言"显式失败"的整体取向相悖。
		if toFloat(r) == 0 {
			panic(evalErrorf(0, 0, "division by zero"))
		}
		return FloatValue(toFloat(l) / toFloat(r))
	}
	panic(evalErrorf(0, 0, "cannot divide %s by %s", kindName(l), kindName(r)))
}

func modValues(l, r Value) Value {
	if l.IsInt() && r.IsInt() {
		if r.i == 0 {
			panic(evalErrorf(0, 0, "modulo by zero"))
		}
		return IntValue(l.i % r.i)
	}
	panic(evalErrorf(0, 0, "%% requires int operands, got %s and %s", kindName(l), kindName(r)))
}

func eqValues(l, r Value) bool {
	// 数值跨 kind 按值比较（与 < > 的混型比较一致）：
	// Go JSON 数字一律注入为 float64，`x == 100` 必须按值命中而非静默为 false。
	if l.IsInt() && r.IsInt() {
		return l.i == r.i
	}
	if isNumeric(l) && isNumeric(r) {
		return toFloat(l) == toFloat(r)
	}
	if l.kind != r.kind {
		return false
	}
	switch l.kind {
	case kindBool:
		return l.b == r.b
	case kindString:
		return l.s == r.s
	case kindNull:
		return true
	case kindLambda:
		// 同一闭包恒等：f == f 为 true；不同闭包即使源码相同也不等。
		return l.fn == r.fn
	case kindList:
		// 深度相等：长度一致且逐元素 eqValues。
		if len(l.list) != len(r.list) {
			return false
		}
		for i := range l.list {
			if !eqValues(l.list[i], r.list[i]) {
				return false
			}
		}
		return true
	case kindMap:
		if len(l.m) != len(r.m) {
			return false
		}
		for k, v := range l.m {
			rv, ok := r.m[k]
			if !ok || !eqValues(v, rv) {
				return false
			}
		}
		return true
	}
	return false
}

func isNumeric(v Value) bool { return v.IsInt() || v.IsFloat() }

func ltValues(l, r Value) bool {
	if l.IsInt() && r.IsInt() {
		return l.i < r.i
	}
	if l.IsFloat() || r.IsFloat() {
		return toFloat(l) < toFloat(r)
	}
	if l.IsString() && r.IsString() {
		return l.s < r.s
	}
	panic(evalErrorf(0, 0, "cannot compare %s and %s", kindName(l), kindName(r)))
}

// inValues checks if l is contained in r (List/Map/string).
func inValues(l, r Value) bool {
	switch {
	case r.IsList():
		for _, item := range r.list {
			if eqValues(l, item) {
				return true
			}
		}
		return false
	case r.IsMap():
		if l.IsString() {
			_, ok := r.m[l.s]
			return ok
		}
		return false
	case r.IsString():
		if l.IsString() {
			return containsSubstring(r.s, l.s)
		}
		return false
	}
	panic(evalErrorf(0, 0, "in requires list, map, or string on right, got %s", kindName(r)))
}

func containsSubstring(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func toFloat(v Value) float64 {
	if v.IsInt() {
		return float64(v.i)
	}
	if v.IsFloat() {
		return v.f
	}
	return 0
}

func kindName(v Value) string {
	switch v.kind {
	case kindInt:
		return "int"
	case kindFloat:
		return "float"
	case kindBool:
		return "bool"
	case kindString:
		return "string"
	case kindNull:
		return "null"
	case kindList:
		return "list"
	case kindMap:
		return "map"
	case kindLambda:
		return "lambda"
	case kindBuiltin:
		return "builtin"
	}
	return "unknown"
}
