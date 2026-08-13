package goval

import "testing"

func TestEvaluate(t *testing.T) {
	t.Run("arithmetic", func(t *testing.T) {
		assertEvaluate(t, "1 + 2 * 3", nil, int64(7))
	})
	t.Run("context_inject", func(t *testing.T) {
		ctx := map[string]interface{}{"x": float64(10)}
		assertEvaluate(t, "x > 5", ctx, true)
	})
	t.Run("context_list", func(t *testing.T) {
		ctx := map[string]interface{}{"xs": []interface{}{int(1), int(2), int(3)}}
		assertEvaluate(t, "reduce(xs, 0, (a, x) -> a + x)", ctx, int64(6))
	})
	t.Run("context_map", func(t *testing.T) {
		ctx := map[string]interface{}{"m": map[string]interface{}{"name": "bob"}}
		assertEvaluate(t, "m.name", ctx, "bob")
	})
	t.Run("object_factory", func(t *testing.T) {
		src := `var Person = (name) -> { name: name, greet: () -> "hi " + name };
		        Person("bob").greet()`
		assertEvaluate(t, src, nil, "hi bob")
	})
	t.Run("reduce_builtin", func(t *testing.T) {
		assertEvaluate(t, "reduce([1,2,3], 0, (acc, x) -> acc + x)", nil, int64(6))
	})
	t.Run("returns_list", func(t *testing.T) {
		assertEvaluate(t, "range(1, 4)", nil, []interface{}{int64(1), int64(2), int64(3)})
	})
	t.Run("returns_map", func(t *testing.T) {
		assertEvaluate(t, `{"a": 1}`, nil, map[string]interface{}{"a": int64(1)})
	})
	t.Run("returns_null", func(t *testing.T) {
		assertEvaluate(t, "null", nil, nil)
	})
	t.Run("error_div_zero", func(t *testing.T) {
		// panic 被 recover 转为 error
		assertEvaluateError(t, "1 / 0", nil)
	})
	t.Run("error_compile", func(t *testing.T) {
		assertEvaluateError(t, "var x = 1; var x = 2", nil)
	})
}
