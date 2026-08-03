package eval

import (
	"strconv"
	"strings"

	"github.com/XHao/goval/internal/ast"
)

// compiler 遍历 AST 产出闭包树，同时做单赋值检查。
type compiler struct {
	scopes []map[string]bool // 每层作用域记录已绑定变量名
}

func (c *compiler) currentScope() map[string]bool {
	return c.scopes[len(c.scopes)-1]
}

func (c *compiler) pushScope() {
	c.scopes = append(c.scopes, make(map[string]bool))
}

func (c *compiler) popScope() {
	c.scopes = c.scopes[:len(c.scopes)-1]
}

// declare 绑定变量名，返回是否重绑定（true=冲突）。
// 仅检查当前作用域。
func (c *compiler) declare(name string) bool {
	sc := c.currentScope()
	if sc[name] {
		return true // 重绑定
	}
	sc[name] = true
	return false
}

// isAlreadyDeclared 检查从当前作用域到根的所有层级，变量是否已声明。
// 用于循环体内赋值外层变量的拦截：外层变量已在父作用域 declare，
// 子作用域内的赋值（compileAssignment）应报 rebind 错误。
func (c *compiler) isAlreadyDeclared(name string) bool {
	for _, sc := range c.scopes {
		if sc[name] {
			return true
		}
	}
	return false
}

// breakSignal / continueSignal 用 Go panic 机制实现跨闭包的 break/continue 早退。
type breakSignal struct{}
type continueSignal struct{}

func (c *compiler) compileProgram(ctx ast.IProgramContext) (func(*Env) Value, error) {
	stmts := ctx.AllStatement()
	var fns []func(*Env) Value
	for _, s := range stmts {
		fn, err := c.compileStatement(s)
		if err != nil {
			return nil, err
		}
		if fn != nil {
			fns = append(fns, fn)
		}
	}
	return func(env *Env) Value {
		var last Value = NullValue()
		for _, fn := range fns {
			last = fn(env)
		}
		return last
	}, nil
}

func (c *compiler) compileStatement(ctx ast.IStatementContext) (func(*Env) Value, error) {
	switch {
	case ctx.LocalVariableDeclarationStatement() != nil:
		return c.compileVarDecl(ctx.LocalVariableDeclarationStatement())
	case ctx.ExpressionStatement() != nil:
		return c.compileExprStmt(ctx.ExpressionStatement())
	case ctx.IfStatement() != nil:
		return c.compileIf(ctx.IfStatement())
	case ctx.ForStatement() != nil:
		return c.compileFor(ctx.ForStatement())
	case ctx.BreakStatement() != nil:
		return func(env *Env) Value { panic(breakSignal{}) }, nil
	case ctx.ContinueStatement() != nil:
		return func(env *Env) Value { panic(continueSignal{}) }, nil
	case ctx.Block() != nil:
		return c.compileBlock(ctx.Block())
	default:
		return nil, nil
	}
}

func (c *compiler) compileVarDecl(ctx ast.ILocalVariableDeclarationStatementContext) (func(*Env) Value, error) {
	decl := ctx.LocalVariableDeclaration().(*ast.LocalVariableDeclarationContext)
	list := decl.VarVariableDeclaratorList().(*ast.VarVariableDeclaratorListContext)
	decls := list.AllVarVariableDeclarator()
	var fns []func(*Env) Value
	var names []string
	for _, d := range decls {
		vd := d.(*ast.VarVariableDeclaratorContext)
		name := vd.Identifier().GetText()
		if c.declare(name) {
			return nil, &CompileError{
				Line:   vd.Identifier().GetSymbol().GetLine(),
				Column: vd.Identifier().GetSymbol().GetColumn(),
				Msg:    "cannot rebind variable '" + name + "' (single assignment)",
			}
		}
		names = append(names, name)
		initFn, err := c.compileExpression(vd.VariableInitializer().(*ast.VariableInitializerContext).Expression())
		if err != nil {
			return nil, err
		}
		fns = append(fns, initFn)
	}
	return func(env *Env) Value {
		var last Value
		for i, fn := range fns {
			last = fn(env)
			env.Set(names[i], last)
		}
		return last
	}, nil
}

func (c *compiler) compileExpression(ctx ast.IExpressionContext) (func(*Env) Value, error) {
	return c.compileAssignmentExpr(ctx.(*ast.ExpressionContext).AssignmentExpression().(*ast.AssignmentExpressionContext))
}

func (c *compiler) compileAssignmentExpr(ctx *ast.AssignmentExpressionContext) (func(*Env) Value, error) {
	if ctx.LambdaExpression() != nil {
		return c.compileLambda(ctx.LambdaExpression().(*ast.LambdaExpressionContext))
	}
	if ctx.Assignment() != nil {
		// 赋值（单赋值下等价于声明，编译期检查）
		return c.compileAssignment(ctx.Assignment().(*ast.AssignmentContext))
	}
	// conditionalExpression
	return c.compileConditional(ctx.ConditionalExpression().(*ast.ConditionalExpressionContext))
}

func (c *compiler) compileAssignment(ctx *ast.AssignmentContext) (func(*Env) Value, error) {
	name := ctx.Identifier().GetText()
	// 单赋值检查：如果变量在任何外层作用域已声明，赋值即为重绑定，编译报错。
	// 这拦截了循环体内对外层变量的赋值（如 for { s = s + x }）。
	if c.isAlreadyDeclared(name) {
		return nil, &CompileError{
			Line:   ctx.Identifier().GetSymbol().GetLine(),
			Column: ctx.Identifier().GetSymbol().GetColumn(),
			Msg:    "cannot rebind variable '" + name + "' (single assignment)",
		}
	}
	c.declare(name) // 在当前作用域标记（isAlreadyDeclared 已确认未声明，不会冲突）
	valFn, err := c.compileExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	return func(env *Env) Value {
		v := valFn(env)
		env.Set(name, v)
		return v
	}, nil
}

func (c *compiler) compileConditional(ctx *ast.ConditionalExpressionContext) (func(*Env) Value, error) {
	if ctx.QUESTION() != nil {
		// 三元 a ? b : c
		condFn, err := c.compileConditionalOr(ctx.ConditionalOrExpression().(*ast.ConditionalOrExpressionContext))
		if err != nil {
			return nil, err
		}
		thenFn, err := c.compileExpression(ctx.Expression())
		if err != nil {
			return nil, err
		}
		elseFn, err := c.compileConditional(ctx.ConditionalExpression().(*ast.ConditionalExpressionContext))
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			if condFn(env).b {
				return thenFn(env)
			}
			return elseFn(env)
		}, nil
	}
	return c.compileConditionalOr(ctx.ConditionalOrExpression().(*ast.ConditionalOrExpressionContext))
}

func (c *compiler) compileConditionalOr(ctx *ast.ConditionalOrExpressionContext) (func(*Env) Value, error) {
	if ctx.OR() != nil {
		// 短路 ||
		leftFn, err := c.compileConditionalOr(ctx.ConditionalOrExpression().(*ast.ConditionalOrExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileConditionalAnd(ctx.ConditionalAndExpression().(*ast.ConditionalAndExpressionContext))
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			if leftFn(env).b {
				return BoolValue(true)
			}
			return BoolValue(rightFn(env).b)
		}, nil
	}
	return c.compileConditionalAnd(ctx.ConditionalAndExpression().(*ast.ConditionalAndExpressionContext))
}

func (c *compiler) compileConditionalAnd(ctx *ast.ConditionalAndExpressionContext) (func(*Env) Value, error) {
	if ctx.AND() != nil {
		// 短路 &&
		leftFn, err := c.compileConditionalAnd(ctx.ConditionalAndExpression().(*ast.ConditionalAndExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileInclusiveOr(ctx.InclusiveOrExpression().(*ast.InclusiveOrExpressionContext))
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			if !leftFn(env).b {
				return BoolValue(false)
			}
			return BoolValue(rightFn(env).b)
		}, nil
	}
	return c.compileInclusiveOr(ctx.InclusiveOrExpression().(*ast.InclusiveOrExpressionContext))
}

func (c *compiler) compileInclusiveOr(ctx *ast.InclusiveOrExpressionContext) (func(*Env) Value, error) {
	if ctx.BITOR() != nil {
		leftFn, err := c.compileInclusiveOr(ctx.InclusiveOrExpression().(*ast.InclusiveOrExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileExclusiveOr(ctx.ExclusiveOrExpression().(*ast.ExclusiveOrExpressionContext))
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			l, r := leftFn(env), rightFn(env)
			if l.IsInt() && r.IsInt() {
				return IntValue(l.i | r.i)
			}
			panic(evalErrorf(0, 0, "| requires int operands, got %s and %s", kindName(l), kindName(r)))
		}, nil
	}
	return c.compileExclusiveOr(ctx.ExclusiveOrExpression().(*ast.ExclusiveOrExpressionContext))
}

func (c *compiler) compileExclusiveOr(ctx *ast.ExclusiveOrExpressionContext) (func(*Env) Value, error) {
	if ctx.CARET() != nil {
		leftFn, err := c.compileExclusiveOr(ctx.ExclusiveOrExpression().(*ast.ExclusiveOrExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileAnd(ctx.AndExpression().(*ast.AndExpressionContext))
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			l, r := leftFn(env), rightFn(env)
			if l.IsInt() && r.IsInt() {
				return IntValue(l.i ^ r.i)
			}
			panic(evalErrorf(0, 0, "^ requires int operands, got %s and %s", kindName(l), kindName(r)))
		}, nil
	}
	return c.compileAnd(ctx.AndExpression().(*ast.AndExpressionContext))
}

func (c *compiler) compileAnd(ctx *ast.AndExpressionContext) (func(*Env) Value, error) {
	if ctx.BITAND() != nil {
		leftFn, err := c.compileAnd(ctx.AndExpression().(*ast.AndExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileEquality(ctx.EqualityExpression().(*ast.EqualityExpressionContext))
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			l, r := leftFn(env), rightFn(env)
			if l.IsInt() && r.IsInt() {
				return IntValue(l.i & r.i)
			}
			panic(evalErrorf(0, 0, "& requires int operands, got %s and %s", kindName(l), kindName(r)))
		}, nil
	}
	return c.compileEquality(ctx.EqualityExpression().(*ast.EqualityExpressionContext))
}

func (c *compiler) compileEquality(ctx *ast.EqualityExpressionContext) (func(*Env) Value, error) {
	if ctx.EQUAL() != nil || ctx.NOTEQUAL() != nil {
		leftFn, err := c.compileEquality(ctx.EqualityExpression().(*ast.EqualityExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileRelational(ctx.RelationalExpression().(*ast.RelationalExpressionContext))
		if err != nil {
			return nil, err
		}
		isEqual := ctx.EQUAL() != nil
		return func(env *Env) Value {
			l, r := leftFn(env), rightFn(env)
			if isEqual {
				return BoolValue(eqValues(l, r))
			}
			return BoolValue(!eqValues(l, r))
		}, nil
	}
	return c.compileRelational(ctx.RelationalExpression().(*ast.RelationalExpressionContext))
}

func (c *compiler) compileRelational(ctx *ast.RelationalExpressionContext) (func(*Env) Value, error) {
	if ctx.LT() != nil || ctx.GT() != nil || ctx.LE() != nil || ctx.GE() != nil || ctx.IN() != nil {
		leftFn, err := c.compileRelational(ctx.RelationalExpression().(*ast.RelationalExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileShift(ctx.ShiftExpression().(*ast.ShiftExpressionContext))
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			l, r := leftFn(env), rightFn(env)
			switch {
			case ctx.LT() != nil:
				return BoolValue(ltValues(l, r))
			case ctx.GT() != nil:
				return BoolValue(ltValues(r, l))
			case ctx.LE() != nil:
				return BoolValue(!ltValues(r, l))
			case ctx.GE() != nil:
				return BoolValue(!ltValues(l, r))
			case ctx.IN() != nil:
				return BoolValue(inValues(l, r))
			}
			return BoolValue(false)
		}, nil
	}
	return c.compileShift(ctx.ShiftExpression().(*ast.ShiftExpressionContext))
}

func (c *compiler) compileShift(ctx *ast.ShiftExpressionContext) (func(*Env) Value, error) {
	if ctx.LSHIFT() != nil || ctx.RSHIFT() != nil {
		leftFn, err := c.compileShift(ctx.ShiftExpression().(*ast.ShiftExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileAdditive(ctx.AdditiveExpression().(*ast.AdditiveExpressionContext))
		if err != nil {
			return nil, err
		}
		isLeft := ctx.LSHIFT() != nil
		return func(env *Env) Value {
			l, r := leftFn(env), rightFn(env)
			if l.IsInt() && r.IsInt() {
				if isLeft {
					return IntValue(l.i << uint(r.i))
				}
				return IntValue(l.i >> uint(r.i))
			}
			panic(evalErrorf(0, 0, "shift requires int operands, got %s and %s", kindName(l), kindName(r)))
		}, nil
	}
	return c.compileAdditive(ctx.AdditiveExpression().(*ast.AdditiveExpressionContext))
}

func (c *compiler) compileAdditive(ctx *ast.AdditiveExpressionContext) (func(*Env) Value, error) {
	if ctx.ADD() != nil || ctx.SUB() != nil {
		leftFn, err := c.compileAdditive(ctx.AdditiveExpression().(*ast.AdditiveExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileMultiplicative(ctx.MultiplicativeExpression().(*ast.MultiplicativeExpressionContext))
		if err != nil {
			return nil, err
		}
		isAdd := ctx.ADD() != nil
		return func(env *Env) Value {
			l, r := leftFn(env), rightFn(env)
			if isAdd {
				return addValues(l, r)
			}
			return subValues(l, r)
		}, nil
	}
	return c.compileMultiplicative(ctx.MultiplicativeExpression().(*ast.MultiplicativeExpressionContext))
}

func (c *compiler) compileMultiplicative(ctx *ast.MultiplicativeExpressionContext) (func(*Env) Value, error) {
	if ctx.MUL() != nil || ctx.DIV() != nil || ctx.MOD() != nil {
		leftFn, err := c.compileMultiplicative(ctx.MultiplicativeExpression().(*ast.MultiplicativeExpressionContext))
		if err != nil {
			return nil, err
		}
		rightFn, err := c.compileUnary(ctx.UnaryExpression().(*ast.UnaryExpressionContext))
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			l, r := leftFn(env), rightFn(env)
			switch {
			case ctx.MUL() != nil:
				return mulValues(l, r)
			case ctx.DIV() != nil:
				return divValues(l, r)
			case ctx.MOD() != nil:
				return modValues(l, r)
			}
			panic(evalErrorf(0, 0, "unreachable multiplicative"))
		}, nil
	}
	return c.compileUnary(ctx.UnaryExpression().(*ast.UnaryExpressionContext))
}

func (c *compiler) compileUnary(ctx *ast.UnaryExpressionContext) (func(*Env) Value, error) {
	if ctx.ADD() != nil || ctx.SUB() != nil {
		innerFn, err := c.compileUnary(ctx.UnaryExpression().(*ast.UnaryExpressionContext))
		if err != nil {
			return nil, err
		}
		isSub := ctx.SUB() != nil
		return func(env *Env) Value {
			v := innerFn(env)
			if isSub {
				if v.IsInt() {
					return IntValue(-v.i)
				}
				if v.IsFloat() {
					return FloatValue(-v.f)
				}
				panic(evalErrorf(0, 0, "unary - requires numeric operand, got %s", kindName(v)))
			}
			// unary +: no-op for numeric
			if v.IsInt() || v.IsFloat() {
				return v
			}
			panic(evalErrorf(0, 0, "unary + requires numeric operand, got %s", kindName(v)))
		}, nil
	}
	// unaryExpressionNotPlusMinus
	un := ctx.UnaryExpressionNotPlusMinus().(*ast.UnaryExpressionNotPlusMinusContext)
	if un.BANG() != nil || un.TILDE() != nil {
		innerFn, err := c.compileUnary(un.UnaryExpression().(*ast.UnaryExpressionContext))
		if err != nil {
			return nil, err
		}
		isBang := un.BANG() != nil
		return func(env *Env) Value {
			v := innerFn(env)
			if isBang {
				if v.IsBool() {
					return BoolValue(!v.b)
				}
				panic(evalErrorf(0, 0, "! requires bool operand, got %s", kindName(v)))
			}
			// ~: bitwise NOT
			if v.IsInt() {
				return IntValue(^v.i)
			}
			panic(evalErrorf(0, 0, "~ requires int operand, got %s", kindName(v)))
		}, nil
	}
	return c.compilePostfix(un.PostfixExpression().(*ast.PostfixExpressionContext))
}

func (c *compiler) compilePostfix(ctx *ast.PostfixExpressionContext) (func(*Env) Value, error) {
	// base case：纯 primary，无后续 [] / . / () 操作
	if ctx.Primary() != nil {
		return c.compilePrimary(ctx.Primary().(*ast.PrimaryContext))
	}
	// 递归编译左操作数（postfix 是左结合链）
	leftFn, err := c.compilePostfix(ctx.PostfixExpression().(*ast.PostfixExpressionContext))
	if err != nil {
		return nil, err
	}

	switch {
	case ctx.LBRACK() != nil: // 下标访问 base[index]
		idxFn, err := c.compileExpression(ctx.Expression())
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			base := leftFn(env)
			idx := idxFn(env)
			switch {
			case base.IsList():
				i := int(idx.i)
				if i < 0 || i >= len(base.list) {
					panic(evalErrorf(0, 0, "list index out of range: %d (len %d)", i, len(base.list)))
				}
				return base.list[i]
			case base.IsMap():
				v, ok := base.m[idx.s]
				if !ok {
					return NullValue()
				}
				return v
			case base.IsString():
				i := int(idx.i)
				r := []rune(base.s)
				if i < 0 || i >= len(r) {
					panic(evalErrorf(0, 0, "string index out of range: %d (len %d)", i, len(r)))
				}
				return StringValue(string(r[i]))
			}
			panic(evalErrorf(0, 0, "cannot index %s", kindName(base)))
		}, nil
	case ctx.DOT() != nil && ctx.LPAREN() != nil: // 方法调用 obj.method(args)
		name := ctx.Identifier().GetText()
		argFns, err := c.compileArgList(ctx.ArgumentList())
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			base := leftFn(env)
			args := make([]Value, len(argFns))
			for i, fn := range argFns {
				args[i] = fn(env)
			}
			if base.IsMap() {
				method, ok := base.m[name]
				if !ok {
					panic(evalErrorf(0, 0, "no method '%s' on map", name))
				}
				return callValue(method, args)
			}
			panic(evalErrorf(0, 0, "cannot call method '%s' on %s", name, kindName(base)))
		}, nil
	case ctx.DOT() != nil: // 属性访问 base.field
		name := ctx.Identifier().GetText()
		return func(env *Env) Value {
			base := leftFn(env)
			if base.IsMap() {
				v, ok := base.m[name]
				if !ok {
					return NullValue()
				}
				return v
			}
			panic(evalErrorf(0, 0, "cannot access field '%s' on %s", name, kindName(base)))
		}, nil
	case ctx.LPAREN() != nil: // 函数调用 f(args)
		argFns, err := c.compileArgList(ctx.ArgumentList())
		if err != nil {
			return nil, err
		}
		return func(env *Env) Value {
			callee := leftFn(env)
			args := make([]Value, len(argFns))
			for i, fn := range argFns {
				args[i] = fn(env)
			}
			return callValue(callee, args)
		}, nil
	}
	return nil, &CompileError{Msg: "unsupported postfix expression"}
}

func (c *compiler) compilePrimary(ctx *ast.PrimaryContext) (func(*Env) Value, error) {
	if ctx.Literal() != nil {
		return c.compileLiteral(ctx.Literal().(*ast.LiteralContext))
	}
	if ctx.Identifier() != nil {
		name := ctx.Identifier().GetText()
		return func(env *Env) Value {
			v, ok := env.Lookup(name)
			if !ok {
				panic(evalErrorf(0, 0, "undefined variable: %s", name))
			}
			return v
		}, nil
	}
	// 括号表达式 (expr)
	if ctx.LPAREN() != nil {
		return c.compileExpression(ctx.Expression())
	}
	// list/map 字面量
	if ctx.ListLiteral() != nil {
		return c.compileListLiteral(ctx.ListLiteral().(*ast.ListLiteralContext))
	}
	if ctx.MapLiteral() != nil {
		return c.compileMapLiteral(ctx.MapLiteral().(*ast.MapLiteralContext))
	}
	return nil, &CompileError{Msg: "unsupported primary"}
}

func (c *compiler) compileListLiteral(ctx *ast.ListLiteralContext) (func(*Env) Value, error) {
	var elemFns []func(*Env) Value
	if ctx.ExpressionList() != nil {
		el := ctx.ExpressionList().(*ast.ExpressionListContext)
		for _, e := range el.AllExpression() {
			fn, err := c.compileExpression(e)
			if err != nil {
				return nil, err
			}
			elemFns = append(elemFns, fn)
		}
	}
	return func(env *Env) Value {
		lst := make([]Value, len(elemFns))
		for i, fn := range elemFns {
			lst[i] = fn(env)
		}
		return ListValue(lst)
	}, nil
}

func (c *compiler) compileMapLiteral(ctx *ast.MapLiteralContext) (func(*Env) Value, error) {
	var keyFns, valFns []func(*Env) Value
	if ctx.MapEntryList() != nil {
		ml := ctx.MapEntryList().(*ast.MapEntryListContext)
		for _, entry := range ml.AllMapEntry() {
			me := entry.(*ast.MapEntryContext)
			exprs := me.AllExpression()
			// Bare identifier as map key → string key (JS-like shorthand)
			keyText := exprs[0].GetText()
			if isIdentifierKey(keyText) {
				k := keyText
				keyFns = append(keyFns, func(env *Env) Value { return StringValue(k) })
			} else {
				kFn, err := c.compileExpression(exprs[0])
				if err != nil {
					return nil, err
				}
				keyFns = append(keyFns, kFn)
			}
			vFn, err := c.compileExpression(exprs[1])
			if err != nil {
				return nil, err
			}
			valFns = append(valFns, vFn)
		}
	}
	return func(env *Env) Value {
		m := make(map[string]Value)
		for i := range keyFns {
			k := keyFns[i](env)
			m[k.s] = valFns[i](env)
		}
		return MapValue(m)
	}, nil
}

// isIdentifierKey 判断文本是否为合法标识符（用于 map 字面量的 bare identifier key 简写）。
func isIdentifierKey(text string) bool {
	if len(text) == 0 {
		return false
	}
	ch := text[0]
	if !(ch == '_' || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
		return false
	}
	for i := 1; i < len(text); i++ {
		c := text[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
			return false
		}
	}
	return true
}

func (c *compiler) compileLiteral(ctx *ast.LiteralContext) (func(*Env) Value, error) {
	tok := ctx.GetStart()
	switch {
	case ctx.IntegerLiteral() != nil:
		text := ctx.IntegerLiteral().GetText()
		// 去掉 l/L 后缀
		clean := strings.TrimRight(text, "lL")
		n, err := strconv.ParseInt(clean, 0, 64)
		if err != nil {
			return nil, &CompileError{Line: tok.GetLine(), Column: tok.GetColumn(), Msg: "invalid integer: " + text}
		}
		return func(env *Env) Value { return IntValue(n) }, nil
	case ctx.FloatingPointLiteral() != nil:
		text := ctx.FloatingPointLiteral().GetText()
		clean := strings.TrimRight(text, "fFdD")
		f, err := strconv.ParseFloat(clean, 64)
		if err != nil {
			return nil, &CompileError{Line: tok.GetLine(), Column: tok.GetColumn(), Msg: "invalid float: " + text}
		}
		return func(env *Env) Value { return FloatValue(f) }, nil
	case ctx.BooleanLiteral() != nil:
		b := ctx.BooleanLiteral().GetText() == "true"
		return func(env *Env) Value { return BoolValue(b) }, nil
	case ctx.StringLiteral() != nil:
		s := ctx.StringLiteral().GetText()
		// 去掉引号，处理转义（简化：只去引号）
		s = s[1 : len(s)-1]
		s = unescapeString(s)
		return func(env *Env) Value { return StringValue(s) }, nil
	case ctx.CharacterLiteral() != nil:
		s := ctx.CharacterLiteral().GetText()
		// 'a' → int
		s = s[1 : len(s)-1]
		s = unescapeString(s)
		r := []rune(s)
		var code int64
		if len(r) > 0 {
			code = int64(r[0])
		}
		return func(env *Env) Value { return IntValue(code) }, nil
	case ctx.NullLiteral() != nil:
		return func(env *Env) Value { return NullValue() }, nil
	}
	return nil, &CompileError{Msg: "unknown literal"}
}

func unescapeString(s string) string {
	r := strings.NewReplacer(
		"\\n", "\n", "\\t", "\t", "\\r", "\r",
		`\"`, `"`, `\\`, `\`,
		`\'`, "'", `\0`, "\x00",
	)
	return r.Replace(s)
}

func (c *compiler) compileExprStmt(ctx ast.IExpressionStatementContext) (func(*Env) Value, error) {
	return c.compileExpression(ctx.(*ast.ExpressionStatementContext).Expression())
}

// compileLambda 编译 lambda 表达式，产出闭包。
// 闭包捕获定义点 Env（运行时传入的 env 参数）。
func (c *compiler) compileLambda(ctx *ast.LambdaExpressionContext) (func(*Env) Value, error) {
	params, err := c.compileLambdaParams(ctx.LambdaParameters().(*ast.LambdaParametersContext))
	if err != nil {
		return nil, err
	}

	// lambda 体在新的作用域编译，参数名在新作用域内 declare
	c.pushScope()
	for _, p := range params {
		c.declare(p)
	}
	bodyFn, err := c.compileLambdaBody(ctx.LambdaBody().(*ast.LambdaBodyContext))
	c.popScope()
	if err != nil {
		return nil, err
	}

	return func(env *Env) Value {
		return LambdaValue(&Lambda{
			params: params,
			body:   bodyFn,
			env:    env, // 闭包捕获定义点 Env
		})
	}, nil
}

func (c *compiler) compileLambdaParams(ctx *ast.LambdaParametersContext) ([]string, error) {
	if ctx.Identifier() != nil {
		return []string{ctx.Identifier().GetText()}, nil
	}
	if ctx.FormalParameterList() != nil {
		fpl := ctx.FormalParameterList().(*ast.FormalParameterListContext)
		var params []string
		for _, id := range fpl.AllIdentifier() {
			params = append(params, id.GetText())
		}
		return params, nil
	}
	return nil, nil // () 无参数
}

func (c *compiler) compileLambdaBody(ctx *ast.LambdaBodyContext) (func(*Env) Value, error) {
	if ctx.Expression() != nil {
		return c.compileExpression(ctx.Expression())
	}
	// expressionBlock: { blockStatement* expression }
	eb := ctx.ExpressionBlock().(*ast.ExpressionBlockContext)
	var stmtFns []func(*Env) Value
	for _, bs := range eb.AllBlockStatement() {
		fn, err := c.compileBlockStatement(bs)
		if err != nil {
			return nil, err
		}
		if fn != nil {
			stmtFns = append(stmtFns, fn)
		}
	}
	exprFn, err := c.compileExpression(eb.Expression())
	if err != nil {
		return nil, err
	}
	return func(env *Env) Value {
		for _, fn := range stmtFns {
			fn(env)
		}
		return exprFn(env)
	}, nil
}

// compileIf 编译 if/else 语句。if 本身不产生值（返回 NullValue）。
// then/else 分支是 statement，可能是 block 或单语句。
func (c *compiler) compileIf(ctx ast.IIfStatementContext) (func(*Env) Value, error) {
	condFn, err := c.compileExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	stmts := ctx.AllStatement()
	thenFn, err := c.compileStatement(stmts[0])
	if err != nil {
		return nil, err
	}
	var elseFn func(*Env) Value
	if len(stmts) > 1 {
		elseFn, err = c.compileStatement(stmts[1])
		if err != nil {
			return nil, err
		}
	}
	return func(env *Env) Value {
		if condFn(env).b {
			if thenFn != nil {
				return thenFn(env)
			}
		} else if elseFn != nil {
			return elseFn(env)
		}
		return NullValue()
	}, nil
}

// compileFor 编译 for-in 语句。遍历 List/Map/string。
// 循环体在新的子作用域编译（pushScope），循环变量在子作用域内 declare。
// break/continue 用 panic 信号机制：continue 在每轮内 recover（跳过当轮），
// break 传播到循环外层 recover（终止循环）。
func (c *compiler) compileFor(ctx ast.IForStatementContext) (func(*Env) Value, error) {
	ids := ctx.AllIdentifier()
	varIterFn, err := c.compileExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	// 循环体在新的子作用域编译
	c.pushScope()
	// 声明循环变量（在循环体作用域内）
	for _, id := range ids {
		c.declare(id.GetText())
	}
	bodyFn, err := c.compileBlock(ctx.Block())
	c.popScope()
	if err != nil {
		return nil, err
	}

	varNames := make([]string, len(ids))
	for i, id := range ids {
		varNames[i] = id.GetText()
	}

	return func(env *Env) Value {
		iter := varIterFn(env)
		var items []Value
		if iter.IsList() {
			items = iter.list
		} else if iter.IsMap() {
			for k := range iter.m {
				items = append(items, StringValue(k))
			}
		} else if iter.IsString() {
			for _, r := range iter.s {
				items = append(items, StringValue(string(r)))
			}
		}

		breaked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					if _, ok := r.(breakSignal); ok {
						breaked = true
						return // break: 终止循环
					}
					panic(r) // 其他 panic 重新抛出
				}
			}()
			for _, item := range items {
				func() {
					defer func() {
						if r := recover(); r != nil {
							if _, ok := r.(continueSignal); ok {
								return // continue: 跳过本轮
							}
							panic(r) // break 或其他：传播
						}
					}()
					loopEnv := NewEnv(env)
					loopEnv.Set(varNames[0], item)
					if len(varNames) > 1 && iter.IsMap() {
						loopEnv.Set(varNames[1], iter.m[item.s])
					}
					if bodyFn != nil {
						bodyFn(loopEnv)
					}
				}()
			}
		}()
		_ = breaked
		return NullValue()
	}, nil
}

// compileBlock 编译 block（{ blockStatements }），在新的子作用域中编译。
// 这保证 if/else 各分支、for 循环体等有独立作用域。
func (c *compiler) compileBlock(ctx ast.IBlockContext) (func(*Env) Value, error) {
	if ctx.BlockStatements() == nil {
		return nil, nil
	}
	c.pushScope()
	fn, err := c.compileBlockStatements(ctx.BlockStatements())
	c.popScope()
	if err != nil {
		return nil, err
	}
	return fn, nil
}

// compileBlockStatements 编译 blockStatements 序列，返回顺序执行的闭包。
func (c *compiler) compileBlockStatements(ctx ast.IBlockStatementsContext) (func(*Env) Value, error) {
	var fns []func(*Env) Value
	for _, bs := range ctx.AllBlockStatement() {
		fn, err := c.compileBlockStatement(bs)
		if err != nil {
			return nil, err
		}
		if fn != nil {
			fns = append(fns, fn)
		}
	}
	return func(env *Env) Value {
		var last Value = NullValue()
		for _, fn := range fns {
			last = fn(env)
		}
		return last
	}, nil
}

// compileBlockStatement 编译 blockStatement（localVariableDeclarationStatement | statement）。
func (c *compiler) compileBlockStatement(ctx ast.IBlockStatementContext) (func(*Env) Value, error) {
	bs := ctx.(*ast.BlockStatementContext)
	if bs.LocalVariableDeclarationStatement() != nil {
		return c.compileVarDecl(bs.LocalVariableDeclarationStatement())
	}
	if bs.Statement() != nil {
		return c.compileStatement(bs.Statement())
	}
	return nil, nil
}

// compileArgList 编译参数列表，返回参数求值函数切片。
func (c *compiler) compileArgList(ctx ast.IArgumentListContext) ([]func(*Env) Value, error) {
	if ctx == nil {
		return nil, nil
	}
	al := ctx.(*ast.ArgumentListContext)
	var argFns []func(*Env) Value
	for _, a := range al.AllExpression() {
		fn, err := c.compileExpression(a)
		if err != nil {
			return nil, err
		}
		argFns = append(argFns, fn)
	}
	return argFns, nil
}

// callValue 调用 lambda 或 builtin。
// 对于 lambda：新建 callEnv，parent 指向 lambda.env（闭包捕获的 Env），绑定参数后执行 body。
// 对于 builtin：直接调用 Go 函数。
func callValue(callee Value, args []Value) Value {
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
	return callLambda(l, args)
}
