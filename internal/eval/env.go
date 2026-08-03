package eval

// Env 是不可变作用域链。单赋值语义下，Define 返回新节点，不修改原 Env。
type Env struct {
	vars  map[string]Value
	parent *Env
}

func NewEnv(parent *Env) *Env {
	return &Env{vars: make(map[string]Value), parent: parent}
}

// NewRootEnv 创建无父节点的根 Env。
func NewRootEnv() *Env {
	return NewEnv(nil)
}

// Lookup 沿父链查找变量。
func (e *Env) Lookup(name string) (Value, bool) {
	if v, ok := e.vars[name]; ok {
		return v, true
	}
	if e.parent != nil {
		return e.parent.Lookup(name)
	}
	return Value{}, false
}

// Define 在当前层绑定变量，返回包含该绑定的新 Env 节点。
// 注意：单赋值下调用方需保证 name 未在当前层绑定过（编译期检查）。
func (e *Env) Define(name string, v Value) *Env {
	vars := make(map[string]Value, len(e.vars)+1)
	for k, val := range e.vars {
		vars[k] = val
	}
	vars[name] = v
	return &Env{vars: vars, parent: e.parent}
}

// Set 直接写入当前层，给顶层语句序列用。
// 单赋值编译期检查保证安全：同一变量不会被二次赋值。
func (e *Env) Set(name string, v Value) {
	e.vars[name] = v
}
