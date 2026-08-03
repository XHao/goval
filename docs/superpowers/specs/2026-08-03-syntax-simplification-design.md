# goval 语法精简设计

日期：2026-08-03
修订：2026-08-03（变量改为单赋值；新增求值器架构决策）
状态：语法部分已实现；求值器待实现

## 背景与动机

goval 当前的 grammar（`grammar/RuleExprParser.g4`、`RuleExprLexer.g4`）定义了一门类 Groovy 的完整脚本语言：struct + 方法声明、function type 显式签名、List/Map/Set 容器类型、switch、for 三段式、break/continue/return 等。现状是**只有语法分析器（`internal/syntax`）能工作，没有求值器**。在如此丰富的语法上直接实现求值器，复杂度过高、收益不匹配。

本设计精简语法，使其与"规则引擎"这一核心场景对齐，降低求值器实现成本，同时保留规则引擎真正需要的能力。

## 定位

面向**规则引擎**的可嵌入表达式语言。在纯表达式求值基础上，引入**不可变对象、lambda、最小控制流**。不追求通用脚本能力，砍掉一切可变状态语法。

## 核心机制

### 对象 = lambda 工厂 + Map 字面量

自定义类型不需要专门语法。一个"类型"就是一个返回 Map 的工厂函数，字段是 Map 的键，方法是闭包值：

```
Person = (name, age) -> {
    name: name,
    age: age,
    greet: () -> "hi " + name     // 闭包捕获构造参数，无需 this
}
p = Person("alice", 30)
p.greet()                          // "hi alice"
```

- 访问 `p.name` 即 Map 查找；`p.greet()` 即取出 lambda 再调用。语法上本就是 `.标识符` 与 `.标识符()`，统一，无需专门的方法调用规则。
- 多次调用工厂产生互不影响的独立对象（各自闭包、各自 Map）。
- 方法靠闭包捕获**构造参数**，因此**不需要 `this`**。

### 全不可变（含单赋值变量）

对象属性、List/Map 元素、**变量绑定**一律不可写。一旦绑定，值固定。

- **单赋值（不可重绑定）**：`var x = expr` 绑定后，`x` 在其作用域内不可再次赋值或重新声明。这把"不可变"从对象扩展到一切，消除闭包捕获的歧义（值快照与引用捕获在无重新绑定下可观察等价），并使求值器并发安全天然成立（编译结果可在多 goroutine 间零成本共享）。
- **代价（已接受）**：循环累加不能用"重新赋值"实现，改用 `reduce` 等内置函数式聚合。
- **容器修改**不开放写语法，通过内置函数以"返回新容器"的不可变风格完成：

```
lst2 = append(lst, x)          // 返回新 List，原 lst 不变
m2 = put(m, "k", v)            // 返回新 Map
lst2 = removeAt(lst, 0)
```

### 无 this / 无 return

- 无 `this`：方法不引用当前对象本身，只靠闭包捕获。
- 无 `return`：lambda 块体 `{ stmts; expr }` 以**最后一个表达式**作为返回值。块体末尾必须是表达式。

## 保留的能力

- **字面量**：int / long / float / double / bool / char / string / null
- **变量**：`var x = ...` 单赋值声明（绑定后不可重绑定）；赋值 `x = ...` 仅在首次绑定时合法（等价于 `var` 的简写形式，同作用域内禁止重复）
- **运算符**：算术 / 比较 / 逻辑 / 位 / 三元 `?:`
- **lambda**：`(params) -> expr` 或 `(params) -> { stmts; expr }`，闭包捕获
- **Map 字面量**：`{ k: v, ... }`（兼作对象载体）
- **List 字面量**：`[a, b, c]`
- **访问**：`p.name` / `lst[i]` / `f(args)` —— 全只读
- **控制流**：`if/else`、`for x in lst {}` / `for k, v in m {}`、`break`/`continue`（for-in 仅用于遍历/过滤/早退，不承担累加）
- **内置函数**：容器操作（append/put/removeAt 等）+ 聚合（reduce/map/filter —— 单赋值下累加的必需手段）+ 规则引擎常用函数 + 后期补 Set 能力（setOf/unique 等，不进语法）

## 砍掉的能力

| 砍掉 | 理由 |
|---|---|
| `struct` / `this` / 方法声明语法 / methodParameterList | lambda 工厂 + Map 替代 |
| `functionType` 显式类型签名 `(T1,T2)->R` 作为类型注解 | lambda 字面量即用即得 |
| `Set` 容器（语法层） | 规则引擎少用；后期内置函数补 |
| 显式类型声明语法（`type` / `dims` / `primitiveType` / `containerType` 作为注解） | 靠 `var` 推断 |
| `switch` | `if/else if` 替代，少一个关键字 |
| for 三段式 `for(init;cond;update)` | 不可变语义下别扭；`range` + for-in 替代 |
| `return` | 表达式语言，块体末表达式返回 |
| 所有写左值（`.field =` / `[i] =`） | 全不可变铁律 |
| `PLACEHOLDER_VAR` `#name#` | 与 context 注入机制重复；外部数据统一通过 `Evaluate(source, context)` 注入根 Env，普通 Identifier 查找取值，无需特殊占位符语法 |

## 静态检查规则（语法/语义层）

1. **左值检查**：赋值左边只能是标识符。`.field =`、`[i] =` 直接报错。
2. **单赋值检查**：同一作用域内，变量绑定后不可再次赋值或重新声明。`var x = 1; x = 2` 和 `var x = 1; var x = 2` 均报错。内层作用域可 shadow（遮蔽）外层同名变量，但遮蔽不算重绑定。
3. **for-in 目标检查**：`in` 后面必须是可遍历类型（List / Map / string）。
4. **break/continue 只在 for 体内合法**。
5. **lambda 块体末尾必须是表达式**（作为返回值）。

## 完整示例

```
// 规则：VIP 用户大额订单打八折
Order = (amount, userId) -> {
    amount: amount,
    userId: userId,
    discounted: (rate) -> amount * rate
}
User = (id, level) -> { id: id, level: level }

users = [
    User("u1", "gold"),
    User("u2", "vip")
]
// 单赋值下，累加用 reduce 而非 for 重新赋值
userMap = reduce(users, {}, (acc, u) -> put(acc, u.id, u))

order = Order(500, "u2")
user = userMap[order.userId]
final = user.level == "vip" && order.amount > 100
        ? order.discounted(0.8)
        : order.amount
```

for-in 仍可用于遍历/过滤/早退（无需累加的场景）：

```
// 找第一个 VIP 用户（早退）
found = null
for u in users {
    if (u.level == "vip") {
        found = u        // 首次绑定，合法
        break
    }
}
```

注意：上例中 `found` 在 if 块内首次绑定。单赋值语义下，"首次绑定"可在声明作用域的任意子作用域发生，但同一变量不可绑定两次。若循环体可能执行多次，编译器需保证绑定只发生一次（或求值器在重复绑定时报运行时错误）——这要求语义检查区分"保证单次执行的路径"与"可能多次的循环体"。简化起见，**循环体内禁止绑定外层变量**（如上例写法非法），累加一律走 reduce。

## 对现有 grammar 的影响（概要）

精简后需要重写 `RuleExprParser.g4` 与 `RuleExprLexer.g4`：
- 删除 `structDeclaration` / `structMemberList` / `structField` / `structMethod` / `methodParameterList` / `methodParameter` / `methodBody`
- 删除 `functionType` / `functionParameter` / `paramType`（类型注解体系）
- 删除 `primitiveType` / `type` / `dims` / `containerType` 作为显式类型声明的用途（保留容器字面量语法）
- 删除 `switchStatement` 及相关
- 简化 `forStatement` 为仅 for-in 形态
- 删除 `returnStatement`
- 赋值规则收紧：左值限定为标识符
- `localVariableDeclaration` 简化为 `var` 推断 + 标识符赋值

`internal/syntax` 的 `SyntaxChecker` 接口保持不变，底层换用精简后的 grammar 重新生成 AST。后续求值器建立在精简后的 AST 之上。

## 求值器架构决策

### 选型：闭包树（Closure-based Compilation）

求值器采用两阶段管线：

```
源码 → [编译器] → 闭包树 func(*Env) Value → [求值] → Value
```

编译器遍历 parse tree，每个节点产出一个 `func(*Env) Value`。求值时只调根闭包，不再碰 AST。

**选择理由**（对比树遍历 / 字节码 VM）：
- 规则引擎典型用法是"编译一次、对海量数据反复求值"——闭包树编译结果可缓存复用，性能远胜树遍历，接近字节码 VM。
- 无需设计指令集和 VM，实现成本远低于字节码方案。
- 调试比字节码容易（仍是 Go 代码的调用栈）。

### Value 表示：标签联合

```go
type Value struct {
    kind valueKind  // intKind/floatKind/boolKind/stringKind/nullKind/listKind/mapKind/lambdaKind
    i    int64
    f    float64
    s    string
    list []Value
    m    map[string]Value  // 或有序结构
    fn   *lambda
}
```

固定类型集，标签联合避免接口装箱与分发开销，利于规则引擎密集算术运算。

### Env：不可变作用域链

单赋值语义下，Env 是 append-only 的不可变链表（每层作用域是父 Env 的新节点），无需可变 map。闭包捕获的是定义点 Env 的引用；由于变量不可重绑定，捕获值与捕获绑定可观察等价，无需区分。

**并发安全**：编译结果（闭包树）与 Env 均不可变，可在多 goroutine 间零成本共享，规则引擎并发求值无需同步。

### 内置函数

作为求值器可调用的 Go 函数表注册，首版必需：
- 聚合：`reduce`/`map`/`filter`（单赋值下聚合的必需手段）
- 查找：`find`（返回第一个满足条件的元素；循环体内禁止绑定外层变量，"找第一个"场景走此函数）
- 容器操作：`append`/`put`/`removeAt`（不可变容器操作，返回新容器）
