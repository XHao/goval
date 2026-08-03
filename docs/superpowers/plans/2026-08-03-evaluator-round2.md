# goval 求值器第二轮 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 补全求值器缺口 + 修复 deferred minors：移除 PLACEHOLDER_VAR、实现顶层 expressionBlock、消除分号噪音、逻辑运算符类型校验、Map key 类型检查、evalErrorf 真实行号。

**Architecture:** grammar 层改动（删 PLACEHOLDER_VAR、分号可选）→ 重新生成 AST → 求值器适配（compilePrimary 增加 expressionBlock、逻辑运算符校验、Map key 校验、行号传递）。

**Tech Stack:** Go 1.22，antlr4-go/antlr/v4，testify

## Global Constraints

- Go 1.22，模块路径 `github.com/XHao/goval`
- 砍掉 PLACEHOLDER_VAR：lexer 删规则、parser 删 primary 分支、重新生成 AST
- 分号可选：`expressionStatement` 和 `localVariableDeclarationStatement` 的 SEMI 改为可选（`SEMI?`），消除 ANTLR `missing ';'` 警告
- 顶层 expressionBlock：compilePrimary 增加 expressionBlock 分支，在新作用域编译 blockStatements + 尾表达式
- 逻辑运算符 `&& || !` 操作数必须 IsBool，否则运行时报 EvalError
- Map key 必须是 string（StringValue），否则运行时报 EvalError
- evalErrorf 携带真实 line/col（从 AST 节点的 GetStart() 获取）
- 现有测试不能回归（36+ 个测试全绿）
- SyntaxChecker 对外 API 不变

---

## File Structure

- **Modify:** `grammar/RuleExprLexer.g4` — 删除 PLACEHOLDER_VAR 规则
- **Modify:** `grammar/RuleExprParser.g4` — 删除 primary 的 PLACEHOLDER_VAR 分支；SEMI 改可选
- **Regenerate:** `internal/ast/*.go` — generate.sh 重新生成
- **Modify:** `internal/eval/compiler.go` — compilePrimary 增加 expressionBlock；逻辑运算符校验；Map key 校验；evalErrorf 行号
- **Modify:** `internal/eval/eval_test.go` — 新增测试覆盖缺口和修复
- **Modify:** `internal/syntax/syntax_test.go` — 移除 PLACEHOLDER_VAR 相关测试（如有）

---

### Task 1: 移除 PLACEHOLDER_VAR + 分号可选 + 重新生成 AST

**Files:**
- Modify: `grammar/RuleExprLexer.g4`
- Modify: `grammar/RuleExprParser.g4`
- Regenerate: `internal/ast/*.go`

**Interfaces:**
- Produces: 精简后的 grammar（无 PLACEHOLDER_VAR、分号可选）+ 重新生成的 AST。供 Task 2-3 使用。

- [ ] **Step 1: 删除 lexer 的 PLACEHOLDER_VAR 规则**

在 `grammar/RuleExprLexer.g4` 中删除这一行：
```
PLACEHOLDER_VAR : '#' Identifier '#' ;
```

- [ ] **Step 2: 删除 parser primary 的 PLACEHOLDER_VAR 分支**

在 `grammar/RuleExprParser.g4` 的 `primary` 规则中删除：
```
    | PLACEHOLDER_VAR
```

- [ ] **Step 3: 分号改可选**

在 `grammar/RuleExprParser.g4` 中：
- `localVariableDeclarationStatement: localVariableDeclaration SEMI` → `localVariableDeclaration SEMI?`
- `expressionStatement: expression SEMI` → `expression SEMI?`

注意：`breakStatement: BREAK SEMI` 和 `continueStatement: CONTINUE SEMI` 的分号也改为可选（`SEMI?`），保持一致。

- [ ] **Step 4: 重新生成 AST**

```bash
rm -f internal/ast/ruleexpr_lexer.go internal/ast/ruleexpr_parser.go internal/ast/ruleexprparser_base_listener.go internal/ast/ruleexprparser_base_visitor.go internal/ast/ruleexprparser_listener.go internal/ast/ruleexprparser_visitor.go internal/ast/RuleExprLexer.interp internal/ast/RuleExprLexer.tokens internal/ast/RuleExprParser.interp internal/ast/RuleExprParser.tokens
cd grammar && ./generate.sh
```

验证：`go build ./internal/ast/` 通过。

- [ ] **Step 5: 修复 syntax.go 编译（如有）**

`internal/syntax/syntax.go` 可能不引用 PLACEHOLDER_VAR，但检查 `go build ./internal/syntax/` 是否通过。若失败则修复。

- [ ] **Step 6: 验证现有测试**

Run: `go test ./internal/syntax/ ./internal/eval/`
Expected: 可能有个别测试因分号变化或 PLACEHOLDER_VAR 移除断言变化。记录失败项，在 Task 3 统一修复测试。

- [ ] **Step 7: Commit**

```bash
git add grammar/ internal/ast/ internal/syntax/
git commit -m "refactor(grammar): remove PLACEHOLDER_VAR, make semicolons optional

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

### Task 2: 求值器适配——顶层 expressionBlock + 逻辑运算符校验 + Map key 校验 + 行号

**Files:**
- Modify: `internal/eval/compiler.go`
- Modify: `internal/eval/eval_test.go`

**Interfaces:**
- Consumes: Task 1 重新生成的 AST（无 PLACEHOLDER_VAR）。
- Produces: compilePrimary 支持 expressionBlock；逻辑运算符 IsBool 校验；Map key string 校验；evalErrorf 携带真实行号。

- [ ] **Step 1: 写失败测试 — 顶层 expressionBlock**

追加到 `eval_test.go`：

```go
func TestCompile_ExpressionBlock(t *testing.T) {
	fn, err := CompileString("{ var x = 1; x + 1 }")
	assert.NoError(t, err)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, int64(2), v.i)
}

func TestCompile_ExpressionBlockNested(t *testing.T) {
	// 块表达式作为子表达式
	fn, err := CompileString("var r = { var t = 5; t * 2 }; r")
	assert.NoError(t, err)
	v := Run(fn, NewRootEnv())
	assert.Equal(t, int64(10), v.i)
}
```

- [ ] **Step 2: 写失败测试 — 逻辑运算符类型校验**

```go
func TestCompile_LogicTypeCheck(t *testing.T) {
	// 非 bool 操作数应报错
	fn, err := CompileString("1 || 2")
	assert.NoError(t, err) // 编译期不报错（类型在运行时）
	_, err = catchEvalPanic(fn)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "bool")
}

func TestCompile_LogicTypeCheckAnd(t *testing.T) {
	fn, _ := CompileString("\"a\" && true")
	_, err := catchEvalPanic(fn)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "bool")
}
```

`catchEvalPanic` 是测试辅助：

```go
func catchEvalPanic(fn func(*Env) Value) (Value, error) {
	defer func() { recover() }()
	v := fn(NewRootEnv())
	return v, nil
}
```

注意：上面的 catchEvalPanic 不捕获 error——需要用 recover 转 error。修正为：

```go
func catchEvalPanic(fn func(*Env) Value) (v Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			if ee, ok := r.(*EvalError); ok {
				err = ee
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	v = fn(NewRootEnv())
	return
}
```

（需 import "fmt"）

- [ ] **Step 3: 写失败测试 — Map key 类型校验**

```go
func TestCompile_MapKeyNonString(t *testing.T) {
	// 非字符串 key（如数字）应报错
	fn, err := CompileString("{1: \"a\"}")
	assert.NoError(t, err)
	_, err = catchEvalPanic(fn)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "string")
}
```

- [ ] **Step 4: 运行测试确认失败**

Run: `go test ./internal/eval/ -run "TestCompile_ExpressionBlock|TestCompile_LogicType|TestCompile_MapKey"`
Expected: FAIL（expressionBlock 报 unsupported primary；逻辑/Map 校验未实现）

- [ ] **Step 5: 实现 compilePrimary 的 expressionBlock 分支**

在 `compilePrimary` 中，在 `return nil, &CompileError{...}` 之前增加：

```go
	if ctx.ExpressionBlock() != nil {
		return c.compileExpressionBlock(ctx.ExpressionBlock().(*ast.ExpressionBlockContext))
	}
```

实现 `compileExpressionBlock`（与 lambda 块体逻辑相同，但在新作用域）：

```go
func (c *compiler) compileExpressionBlock(ctx *ast.ExpressionBlockContext) (func(*Env) Value, error) {
	c.pushScope()
	var stmtFns []func(*Env) Value
	for _, bs := range ctx.AllBlockStatement() {
		fn, err := c.compileBlockStatement(bs)
		if err != nil { c.popScope(); return nil, err }
		if fn != nil { stmtFns = append(stmtFns, fn) }
	}
	exprFn, err := c.compileExpression(ctx.Expression())
	c.popScope()
	if err != nil { return nil, err }
	return func(env *Env) Value {
		// 块表达式在新作用域——但 Env 不可变，顶层用 Set
		// 块内的 var 声明需要新作用域可见
		blockEnv := NewEnv(env)
		for _, fn := range stmtFns {
			fn(blockEnv)
		}
		return exprFn(blockEnv)
	}, nil
}
```

注意：块表达式内的 var 声明需要在新作用域内可见，但 compileVarDecl 用 `env.Set` 直接写入。这里 blockEnv 是新 Env，var 声明写入 blockEnv，不影响外层 env。但 compileVarDecl 编译时用 pushScope/popScope 管理作用域——需确保编译期作用域与运行期 Env 一致。

**关键**：compileExpressionBlock 已 pushScope，compileBlockStatement 内的 compileVarDecl 会在当前 pushScope 的作用域内 declare。运行时 blockEnv 是新 Env，var 声明的 Set 写入 blockEnv。exprFn 在 blockEnv 上求值，能查到 blockEnv 的变量。逻辑自洽。

- [ ] **Step 6: 实现逻辑运算符 IsBool 校验**

在 `compileConditionalOr` 和 `compileConditionalAnd` 的返回闭包中，增加操作数类型检查。修改 `compileConditionalOr`：

```go
	return func(env *Env) Value {
		l := leftFn(env)
		if !l.IsBool() {
			panic(evalErrorf(0, 0, "|| requires bool, got %s", kindName(l)))
		}
		if l.b {
			return BoolValue(true)
		}
		r := rightFn(env)
		if !r.IsBool() {
			panic(evalErrorf(0, 0, "|| requires bool, got %s", kindName(r)))
		}
		return BoolValue(r.b)
	}, nil
```

同样修改 `compileConditionalAnd`。`!` 运算符在 compileUnaryNotPlusMinus 的 BANG 分支也加校验。

- [ ] **Step 7: 实现 Map key 类型校验**

在 `compileMapLiteral` 的返回闭包中，`k := keyFns[i](env)` 后增加：

```go
			k := keyFns[i](env)
			if !k.IsString() {
				panic(evalErrorf(0, 0, "map key must be string, got %s", kindName(k)))
			}
			m[k.s] = valFns[i](env)
```

- [ ] **Step 8: 实现 evalErrorf 真实行号**

在编译时捕获 AST 节点的 line/col，传入闭包。这需要改 evalErrorf 的调用方式——编译时记录行号，运行时用。

对于逻辑运算符和 Map key 校验，行号从对应的 AST context 获取：

```go
	tok := ctx.GetStart()
	line, col := tok.GetLine(), tok.GetColumn()
	return func(env *Env) Value {
		// ...
		panic(evalErrorf(line, col, "|| requires bool, got %s", kindName(l)))
		// ...
	}, nil
```

对其他 evalErrorf 调用（ops.go 的 addValues 等），由于它们是运行时辅助函数无 AST 上下文，保持 0,0 或改为接收 line/col 参数（后者改动大，本任务只改有 AST 上下文的调用点）。

- [ ] **Step 9: 运行新测试确认通过**

Run: `go test ./internal/eval/ -run "TestCompile_ExpressionBlock|TestCompile_LogicType|TestCompile_MapKey"`
Expected: PASS

- [ ] **Step 10: 运行全量测试确认无回归**

Run: `go test ./internal/eval/`
Expected: 全部 PASS

- [ ] **Step 11: Commit**

```bash
git add internal/eval/
git commit -m "feat(eval): expressionBlock primary, logic bool check, map key check, real line numbers

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

### Task 3: 修复测试 + 全量验证

**Files:**
- Modify: `internal/syntax/syntax_test.go`（移除 PLACEHOLDER_VAR 测试，如有）
- Modify: `internal/eval/eval_test.go`（修复因分号/PLACEHOLDER_VAR 变化的测试）
- Modify: `pkg/goval/example_test.go`（确认 example 仍通过）

**Interfaces:**
- Consumes: Task 1-2 的改动。

- [ ] **Step 1: 检查并修复 syntax 测试**

Run: `go test ./internal/syntax/ -v 2>&1 | grep -E "FAIL|PASS|placeholder|PLACEHOLDER"`
如有 PLACEHOLDER_VAR 相关测试失败，移除或更新断言。

- [ ] **Step 2: 检查并修复 eval 测试**

Run: `go test ./internal/eval/ -v 2>&1 | grep -E "FAIL|PLACEHOLDER"`
如有失败，修复断言。

- [ ] **Step 3: 确认 example 测试通过**

Run: `go test ./pkg/goval/ -run Example -v`
Expected: 全部 PASS，且**无 `missing ';'` 警告**（分号可选后应消失）。

- [ ] **Step 4: 全量验证**

Run: `go build ./... && go vet ./... && go test ./...`
Expected: 全部通过。`go vet` 的 unreachable-code 警告在 ANTLR 生成文件中是固有的，忽略。

- [ ] **Step 5: Commit**

```bash
git add internal/syntax/ internal/eval/ pkg/goval/
git commit -m "test: fix tests for PLACEHOLDER_VAR removal and optional semicolons

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Self-Review 结论

- **Spec 覆盖**：PLACEHOLDER_VAR 移除（Task 1）+ expressionBlock（Task 2）+ 分号噪音（Task 1）+ 逻辑校验（Task 2）+ Map key 校验（Task 2）+ 行号（Task 2）。覆盖全部 3 个缺口 + 高优先级 deferred minors。
- **占位符**：无。所有步骤含具体代码。
- **类型一致**：evalErrorf 签名不变（line, col int, format, args）；compileExpressionBlock 复用 compileBlockStatement。
- **风险**：分号可选可能改变 ANTLR 的 parse tree 结构（SEMI? 产生可选节点），需验证 syntax checker 和 compiler 不依赖 SEMI 节点的存在。
