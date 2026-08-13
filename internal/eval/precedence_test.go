package eval

import "testing"

func TestPrecedence(t *testing.T) {
	t.Run("arithmetic_mul_over_add", func(t *testing.T) {
		assertEval(t, "2 + 3 * 4", int64(14))
		assertEval(t, "(2 + 3) * 4", int64(20))
	})
	t.Run("mul_div_mod_left_assoc", func(t *testing.T) {
		assertEval(t, "100 / 10 / 2", int64(5))
		assertEval(t, "100 % 7 % 3", int64(2)) // 100%7=2, 2%3=2
	})
	t.Run("shift_lower_than_additive", func(t *testing.T) {
		// 1 << 2 + 1 == 1 << (2+1) == 8
		assertEval(t, "1 << 2 + 1", int64(8))
	})
	t.Run("relational_higher_than_equality", func(t *testing.T) {
		assertEval(t, "1 + 1 == 2", true)
		assertEval(t, "1 < 2 == true", true)
	})
	t.Run("logic_not_and_or", func(t *testing.T) {
		// !false && true || false == ((!false) && true) || false == true
		assertEval(t, "!false && true || false", true)
		// ! 运算符高于 &&：!false && false == (!false) && false == false
		assertEval(t, "!false && false", false)
	})
	t.Run("bitwise_layering", func(t *testing.T) {
		// 优先级（低→高）：| < ^ < &，与 C 一致
		// 0xF0 & 0x0F | 0x10 == ((0xF0 & 0x0F) | 0x10) == (0 | 0x10) == 16
		assertEval(t, "0xF0 & 0x0F | 0x10", int64(16))
		// 0xFF & 0x0F ^ 0x10 == ((0xFF & 0x0F) ^ 0x10) == (0x0F ^ 0x10) == 0x1F == 31
		assertEval(t, "0xFF & 0x0F ^ 0x10", int64(31))
		// 用括号对照验证 & 高于 ^：先 ^ 后 & 的反例
		// 0xFF & (0x0F ^ 0x10) == 0xFF & 0x1F == 0x1F == 31 —— 与上面相同（巧合），
		// 改用 0xF0 & 0x0F ^ 0x01: ((0xF0&0x0F)^0x01) = (0^1) = 1;
		//   若 ^ 高于 & 则 0xF0 & (0x0F^0x01) = 0xF0 & 0x0E = 0x00 = 0
		assertEval(t, "0xF0 & 0x0F ^ 0x01", int64(1))
	})
	t.Run("ternary_right_assoc", func(t *testing.T) {
		// false ? 1 : true ? 2 : 3 == false ? 1 : (true ? 2 : 3) == 2
		assertEval(t, "false ? 1 : true ? 2 : 3", int64(2))
	})
	t.Run("assignment_lowest", func(t *testing.T) {
		// x = (3 > 2 ? 100 : 200); x —— 单赋值首次绑定
		assertEval(t, "var x = (3 > 2 ? 100 : 200); x", int64(100))
	})
}
