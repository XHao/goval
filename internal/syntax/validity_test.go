package syntax

import "testing"

func TestValidity(t *testing.T) {
	t.Run("literals", func(t *testing.T) {
		for _, src := range []string{
			"42;", "0x1A;", "0755;", "0b1010;", "3.14;", "'a';",
			`"hello";`, "true;", "null;", "[1,2,3];", `{"k":1};`,
		} {
			assertValid(t, src)
		}
	})
	t.Run("var_decl", func(t *testing.T) {
		assertValid(t, "var x = 1;")
		assertValid(t, "var x = 1, y = 2;")
		assertValid(t, `var s = "hi";`)
	})
	t.Run("assignment_identifier", func(t *testing.T) {
		// 标识符左值赋值合法（语法层）
		assertValid(t, "x = 1;")
	})
	t.Run("operators", func(t *testing.T) {
		for _, src := range []string{
			"1 + 2 * 3;", "a && b || c;", "x == y ? 1 : 0;",
			"1 << 2;", "1 >> 2;", "!true;", "~0;", "1 & 2;", "1 | 2;", "1 ^ 2;",
			"2 in [1,2];", `"k" in m;`,
		} {
			assertValid(t, src)
		}
	})
	t.Run("lambda", func(t *testing.T) {
		assertValid(t, "(a, b) -> a + b;")
		assertValid(t, "x -> x + 1;")        // 单参无括号
		assertValid(t, "() -> 42;")          // 零参
		assertValid(t, "(x) -> { var t = x; t };") // 块体
	})
	t.Run("control_flow", func(t *testing.T) {
		assertValid(t, "if (x > 0) { var a = 1 };")
		assertValid(t, "if (x) { 1 } else { 2 };")
		assertValid(t, "if (x) { 1 } else if (y) { 2 } else { 3 };")
		assertValid(t, "for x in [1,2] { var a = x };")
		assertValid(t, `for k, v in m { var a = k };`)
		assertValid(t, "for x in [1] { break };")
		assertValid(t, "for x in [1] { continue };")
	})
	t.Run("access", func(t *testing.T) {
		assertValid(t, "a[0];")
		assertValid(t, `a["k"];`)
		assertValid(t, "a.b;")
		assertValid(t, "f(1);")
		assertValid(t, "a.b.c;")
	})
	t.Run("optional_semicolon", func(t *testing.T) {
		// 分号可选
		assertValid(t, "var x = 1\nx")
		assertValid(t, "var x = 1")
	})
	t.Run("comment", func(t *testing.T) {
		assertValid(t, "// line comment\nvar x = 1;")
		assertValid(t, "/* block */ var x = 1;")
	})

	// 被拒绝
	t.Run("field_write_rejected", func(t *testing.T) {
		assertInvalid(t, "p.name = 1;")
	})
	t.Run("index_write_rejected", func(t *testing.T) {
		assertInvalid(t, "lst[0] = 1;")
	})
	t.Run("removed_operators", func(t *testing.T) {
		assertInvalid(t, "x++;")   // ++ 已移除
		assertInvalid(t, "x--;")   // -- 已移除
		assertInvalid(t, "x += 1;") // 复合赋值已移除
		assertInvalid(t, "x -= 1;")
	})
	t.Run("removed_switch_case_syntax", func(t *testing.T) {
		// switch 已非关键字，但 case 1: 的 : 语法无效
		assertInvalid(t, "switch (x) { case 1: }")
	})
	t.Run("three_part_for_rejected", func(t *testing.T) {
		// 三段式 for 不支持
		assertInvalid(t, "for (i = 0; i < 10; i++) { 1 }")
	})

	// 已移除关键字现为普通标识符，应被接受
	t.Run("removed_keywords_now_identifiers", func(t *testing.T) {
		assertValid(t, "struct P { int name }") // struct/int 现为标识符
		assertValid(t, "return 1;")             // return 现为标识符
		assertValid(t, "int a = 1;")            // int 现为标识符
	})
}
