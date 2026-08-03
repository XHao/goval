package eval

type valueKind int

const (
	kindNull valueKind = iota
	kindInt
	kindFloat
	kindBool
	kindString
	kindList
	kindMap
	kindLambda
	kindBuiltin
)

// Value 是 goval 的运行时值，用标签联合表示，避免接口装箱开销。
type Value struct {
	kind    valueKind
	i       int64
	f       float64
	b       bool
	s       string
	list    []Value
	m       map[string]Value
	fn      *Lambda
	builtin BuiltinFunc
}

// Lambda 是闭包：参数名列表 + 编译后的求值函数 + 定义点 Env。
type Lambda struct {
	params []string
	body   func(*Env) Value
	env    *Env
}

func IntValue(i int64) Value            { return Value{kind: kindInt, i: i} }
func FloatValue(f float64) Value        { return Value{kind: kindFloat, f: f} }
func BoolValue(b bool) Value            { return Value{kind: kindBool, b: b} }
func StringValue(s string) Value        { return Value{kind: kindString, s: s} }
func NullValue() Value                  { return Value{kind: kindNull} }
func ListValue(l []Value) Value         { return Value{kind: kindList, list: l} }
func MapValue(m map[string]Value) Value { return Value{kind: kindMap, m: m} }
func LambdaValue(l *Lambda) Value       { return Value{kind: kindLambda, fn: l} }
func BuiltinValue(fn BuiltinFunc) Value { return Value{kind: kindBuiltin, builtin: fn} }

func (v Value) IsInt() bool     { return v.kind == kindInt }
func (v Value) IsFloat() bool   { return v.kind == kindFloat }
func (v Value) IsBool() bool    { return v.kind == kindBool }
func (v Value) IsString() bool  { return v.kind == kindString }
func (v Value) IsNull() bool    { return v.kind == kindNull }
func (v Value) IsList() bool    { return v.kind == kindList }
func (v Value) IsMap() bool     { return v.kind == kindMap }
func (v Value) IsLambda() bool  { return v.kind == kindLambda }
func (v Value) IsBuiltin() bool { return v.kind == kindBuiltin }

// 导出 getter，供 pkg/goval 在包外访问 Value 的小写字段。
func (v Value) I() int64              { return v.i }
func (v Value) F() float64            { return v.f }
func (v Value) B() bool               { return v.b }
func (v Value) S() string             { return v.s }
func (v Value) List() []Value         { return v.list }
func (v Value) Map() map[string]Value { return v.m }
