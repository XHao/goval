package goval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestProgramCompileRun 编译一次、对不同 context 反复求值（README 承诺的闭环）。
func TestProgramCompileRun(t *testing.T) {
	t.Run("compile_once_run_many", func(t *testing.T) {
		p, err := Compile("x + y", "x", "y")
		if assert.NoError(t, err) {
			v, err := p.Run(map[string]interface{}{"x": 1, "y": 2})
			if assert.NoError(t, err) {
				assert.Equal(t, int64(3), v)
			}
			v, err = p.Run(map[string]interface{}{"x": 10, "y": 20})
			if assert.NoError(t, err) {
				assert.Equal(t, int64(30), v)
			}
		}
	})
	t.Run("strict_assignment_over_context_keys", func(t *testing.T) {
		// context 键在编译期已知，裸赋值合法
		p, err := Compile("x = x * 2; x", "x")
		if assert.NoError(t, err) {
			v, err := p.Run(map[string]interface{}{"x": 5})
			if assert.NoError(t, err) {
				assert.Equal(t, int64(10), v)
			}
		}
	})
	t.Run("missing_key_errors_on_read", func(t *testing.T) {
		p, err := Compile("x + 1", "x")
		assert.NoError(t, err)
		_, err = p.Run(map[string]interface{}{})
		assert.Error(t, err)
	})
	t.Run("extra_keys_allowed", func(t *testing.T) {
		p, err := Compile("x + 1", "x")
		if assert.NoError(t, err) {
			v, err := p.Run(map[string]interface{}{"x": 1, "extra": 2})
			if assert.NoError(t, err) {
				assert.Equal(t, int64(2), v)
			}
		}
	})
	t.Run("compile_error_reported", func(t *testing.T) {
		_, err := Compile("y = 2", "x") // 严格赋值：y 未绑定
		assert.Error(t, err)
		_, err = Compile("1 +", "x")
		assert.Error(t, err)
	})
	t.Run("run_recovers_panics", func(t *testing.T) {
		p, err := Compile("1 / 0", "x")
		assert.NoError(t, err)
		_, err = p.Run(nil)
		assert.Error(t, err)
	})
	t.Run("context_type_validation", func(t *testing.T) {
		p, err := Compile("f(1)", "f")
		assert.NoError(t, err)
		_, err = p.Run(map[string]interface{}{"f": func(int) int { return 1 }})
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "f")
		}
	})
	t.Run("host_isolation_between_runs", func(t *testing.T) {
		// 脚本原地修改注入的 List：宿主数据不受影响，两次 Run 互不污染
		p, err := Compile("lst[0] = 99; lst[0]", "lst")
		assert.NoError(t, err)
		host := []interface{}{int64(1)}
		_, err = p.Run(map[string]interface{}{"lst": host})
		assert.NoError(t, err)
		assert.Equal(t, int64(1), host[0])
		v, err := p.Run(map[string]interface{}{"lst": host})
		if assert.NoError(t, err) {
			assert.Equal(t, int64(99), v)
		}
		assert.Equal(t, int64(1), host[0])
	})
	t.Run("evaluate_still_works", func(t *testing.T) {
		assertEvaluate(t, "x * 2", map[string]interface{}{"x": 3}, int64(6))
	})
}
