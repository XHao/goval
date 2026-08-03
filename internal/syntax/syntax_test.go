package syntax

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

func TestBasic_Literals(t *testing.T) {
	cases := []string{
		"42;",
		"0x1A;",
		"3.14;",
		`"hello";`,
		"'a';",
		"true;",
		"null;",
	}
	for _, src := range cases {
		_, err := NewSyntaxChecker().CheckString(ErrAt(src))
		assert.NoError(t, err, src)
	}
}

func TestBasic_VariableDeclaration(t *testing.T) {
	cases := []string{
		"var x = 1;",
		"var x = 1, y = 2;",
		"var s = \"hi\";",
	}
	for _, src := range cases {
		assert.True(t, NewSyntaxChecker().CheckStringValid(src), src)
	}
}

func TestBasic_Assignment(t *testing.T) {
	// 标识符赋值合法
	assert.True(t, NewSyntaxChecker().CheckStringValid("x = 1;"))
}

func TestBasic_Immutable_RejectFieldWrite(t *testing.T) {
	// p.name = x 应被拒绝（grammar 层：左值非标识符 → 解析失败）
	err := NewSyntaxChecker().CheckStringErr("p.name = 1;")
	assert.Error(t, err)
}

func TestBasic_Immutable_RejectIndexWrite(t *testing.T) {
	err := NewSyntaxChecker().CheckStringErr("lst[0] = 1;")
	assert.Error(t, err)
}

func TestBasic_Operators(t *testing.T) {
	cases := []string{
		"1 + 2 * 3;",
		"a && b || c;",
		"x == y ? 1 : 0;",
		"1 << 2;",
		"!true;",
	}
	for _, src := range cases {
		assert.True(t, NewSyntaxChecker().CheckStringValid(src), src)
	}
}

func TestLambda_BasicAndBlock(t *testing.T) {
	cases := []string{
		"x -> x + 1;",
		"(a, b) -> a + b;",
		"() -> 42;",
		"(x) -> { var y = x * 2; y + 1 };",
	}
	for _, src := range cases {
		assert.True(t, NewSyntaxChecker().CheckStringValid(src), src)
	}
}

func TestObject_LambdaFactory(t *testing.T) {
	src := `Person = (name, age) -> {
		name: name,
		age: age,
		greet: () -> "hi " + name
	};
	p = Person("alice", 30);
	p.greet();`
	assert.True(t, NewSyntaxChecker().CheckStringValid(src))
}

func TestContainer_ListAndMap(t *testing.T) {
	cases := []string{
		"[1, 2, 3];",
		"[];",
		`{"k": 1, "v": 2};`,
		"{};",
		"m[\"k\"];",
		"lst[0];",
	}
	for _, src := range cases {
		assert.True(t, NewSyntaxChecker().CheckStringValid(src), src)
	}
}

func TestControl_IfElse(t *testing.T) {
	src := "if (x > 0) { 1; } else { 0; };"
	assert.True(t, NewSyntaxChecker().CheckStringValid(src))
}

func TestControl_ForIn(t *testing.T) {
	cases := []string{
		"for x in lst { x; }",
		"for k, v in m { v; }",
	}
	for _, src := range cases {
		assert.True(t, NewSyntaxChecker().CheckStringValid(src), src)
	}
}

func TestControl_ForIn_RejectThreePart(t *testing.T) {
	// 三段式 for 应被拒绝
	assert.Error(t, NewSyntaxChecker().CheckStringErr("for (i = 0; i < 10; i++) {}"))
}

func TestControl_BreakContinueInFor(t *testing.T) {
	src := "for x in lst { if (x > 0) { break; }; x; }"
	assert.True(t, NewSyntaxChecker().CheckStringValid(src))
}

func TestSemantic_BreakOutsideFor(t *testing.T) {
	assert.Error(t, NewSyntaxChecker().CheckStringErr("break;"))
}

func TestSemantic_ContinueOutsideFor(t *testing.T) {
	assert.Error(t, NewSyntaxChecker().CheckStringErr("continue;"))
}

func TestRemoved_NoStruct(t *testing.T) {
	// struct/int/return/switch are no longer keywords (lexer rules removed).
	// "struct P { int name }" now parses as an identifier sequence — struct is
	// just a regular identifier. The test verifies this is no longer rejected.
	assert.NoError(t, NewSyntaxChecker().CheckStringErr("struct P { int name }"))
}

func TestRemoved_NoReturn(t *testing.T) {
	// "return" is no longer a keyword — it is a plain identifier, so
	// "return 1;" parses as a valid expression statement.
	assert.NoError(t, NewSyntaxChecker().CheckStringErr("return 1;"))
}

func TestRemoved_NoSwitch(t *testing.T) {
	// "switch" is no longer a keyword, but the ":" in "case 1:" is still
	// invalid syntax, so this is rejected for a different reason.
	assert.Error(t, NewSyntaxChecker().CheckStringErr("switch (x) { case 1: }"))
}

func TestRemoved_NoTypeAnnotation(t *testing.T) {
	// "int" is no longer a keyword — it is a plain identifier, so
	// "int a = 1;" parses as a valid statement (int treated as expression).
	assert.NoError(t, NewSyntaxChecker().CheckStringErr("int a = 1;"))
}

func TestSyntaxChecker_NoError(t *testing.T) {
	tree, err := NewSyntaxChecker().CheckString("var a = 1 + 2;")
	assert.NoError(t, err)
	assert.NotNil(t, tree)
}

func TestSyntaxChecker_WithError(t *testing.T) {
	_, err := NewSyntaxChecker().CheckString("var a = 1 + ;")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "syntax error")
}

func TestSyntaxChecker_MultipleErrors(t *testing.T) {
	_, errs := NewSyntaxChecker().CheckWithAllErrors(antlr.NewInputStream("var a = 1 + ; var b = 2 +"))
	assert.Greater(t, len(errs), 0)
}

func TestSyntaxErrorListener_CollectsMultiple(t *testing.T) {
	l := &SyntaxErrorListener{}
	l.SyntaxError(nil, nil, 1, 10, "e1", nil)
	l.SyntaxError(nil, nil, 2, 15, "e2", nil)
	assert.Equal(t, 2, len(l.GetAllErrors()))
	assert.Contains(t, l.GetFirstError().Error(), "e1")
}

func TestSyntaxChecker_IsValid(t *testing.T) {
	assert.True(t, NewSyntaxChecker().IsValid(antlr.NewInputStream("var x = 1;")))
	assert.False(t, NewSyntaxChecker().IsValid(antlr.NewInputStream("var x = 1 + ;")))
}

func TestSyntaxChecker_Reset(t *testing.T) {
	c := NewSyntaxChecker()
	_ = c.IsValid(antlr.NewInputStream("var x = 1 + ;"))
	c.Reset()
	assert.Equal(t, 0, c.GetErrorCount())
}

func TestSyntaxChecker_Concurrent(t *testing.T) {
	c := NewSyntaxChecker()
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			_ = c.IsValid(antlr.NewInputStream("var x = 1;"))
			done <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

// ErrAt 返回源码（占位，便于将来加错误定位断言）。
func ErrAt(src string) string { return src }

// CheckStringValid 是测试辅助：CheckString 无错返回 true。
func (c *SyntaxChecker) CheckStringValid(src string) bool {
	_, err := c.CheckString(src)
	return err == nil
}

// CheckStringErr 是测试辅助：返回 CheckString 的 error。
func (c *SyntaxChecker) CheckStringErr(src string) error {
	_, err := c.CheckString(src)
	return err
}
