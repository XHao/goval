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
	if l.kind != r.kind {
		return false
	}
	switch l.kind {
	case kindInt:
		return l.i == r.i
	case kindFloat:
		return l.f == r.f
	case kindBool:
		return l.b == r.b
	case kindString:
		return l.s == r.s
	case kindNull:
		return true
	}
	return false
}

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
