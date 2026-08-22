package goval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestV2EndToEnd 通过公开 API 验证 v2 语义端到端生效（Go 原生值进出）。
func TestV2EndToEnd(t *testing.T) {
	t.Run("person_object_with_methods", func(t *testing.T) {
		src := `
var Person = (name, age) -> {
    name: name,
    age: age,
    greet: () -> "hi " + name,
    birthday: () -> { age = age + 1; age },
    olderThan: (o) -> age > o.age
}
var p = Person("alice", 30)
var aged = p.birthday()
{ greet: p.greet(), age: aged, older: p.olderThan(Person("bob", 20)) }
`
		v, err := Evaluate(src, nil)
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{
			"greet": "hi alice",
			"age":   int64(31),
			"older": true,
		}, v)
	})
	t.Run("field_mutation_visible_in_result", func(t *testing.T) {
		v, err := Evaluate(`var p = {n: 1}; p.n = 2; p`, nil)
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"n": int64(2)}, v)
	})
	t.Run("loop_accumulation", func(t *testing.T) {
		v, err := Evaluate(`var s = 0; for x in [1, 2, 3] { s = s + x }; s`, nil)
		assert.NoError(t, err)
		assert.Equal(t, int64(6), v)
	})
	t.Run("context_variable_assignment", func(t *testing.T) {
		v, err := Evaluate(`x = x + 1; x`, map[string]interface{}{"x": 10})
		assert.NoError(t, err)
		assert.Equal(t, int64(11), v)
	})
	t.Run("strict_assignment_error", func(t *testing.T) {
		_, err := Evaluate(`typoVar = 1`, map[string]interface{}{"userId": 1})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "undefined variable 'typoVar'")
	})
	t.Run("capture_pragma_end_to_end", func(t *testing.T) {
		src := `
var level = "outer"
var P = capture { level: "field", read: () -> level }
P.read()
`
		v, err := Evaluate(src, nil)
		assert.NoError(t, err)
		assert.Equal(t, "outer", v)
	})
	t.Run("counter_private_state", func(t *testing.T) {
		src := `
var Counter = (start) -> {
    var count = start
    { inc: () -> { count = count + 1; count }, value: () -> count }
}
var c = Counter(0)
c.inc()
c.inc()
c.value()
`
		v, err := Evaluate(src, nil)
		assert.NoError(t, err)
		assert.Equal(t, int64(2), v)
	})
	t.Run("list_mutation_aliasing", func(t *testing.T) {
		v, err := Evaluate(`var l = [1, 2]; l[0] = 9; l`, nil)
		assert.NoError(t, err)
		assert.Equal(t, []interface{}{int64(9), int64(2)}, v)
	})
}
