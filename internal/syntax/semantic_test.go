package syntax

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

func TestSemantic(t *testing.T) {
	t.Run("break_outside_for", func(t *testing.T) {
		assertInvalid(t, "break;")
	})
	t.Run("continue_outside_for", func(t *testing.T) {
		assertInvalid(t, "continue;")
	})
	t.Run("break_inside_for_valid", func(t *testing.T) {
		assertValid(t, "for x in [1] { break };")
	})
	t.Run("continue_inside_for_valid", func(t *testing.T) {
		assertValid(t, "for x in [1] { continue };")
	})
	t.Run("syntax_checker_api_noerror", func(t *testing.T) {
		tree, err := NewSyntaxChecker().CheckString("var a = 1 + 2;")
		assert.NoError(t, err)
		assert.NotNil(t, tree)
	})
	t.Run("syntax_checker_api_witherror", func(t *testing.T) {
		_, err := NewSyntaxChecker().CheckString("var a = 1 + ;")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "syntax error")
	})
	t.Run("syntax_checker_multiple_errors", func(t *testing.T) {
		_, errs := NewSyntaxChecker().CheckWithAllErrors(antlr.NewInputStream("var a = 1 + ; var b = 2 +"))
		assert.Greater(t, len(errs), 0)
	})
	t.Run("syntax_checker_isvalid", func(t *testing.T) {
		assert.True(t, NewSyntaxChecker().IsValid(antlr.NewInputStream("var x = 1;")))
		assert.False(t, NewSyntaxChecker().IsValid(antlr.NewInputStream("var x = 1 + ;")))
	})
	t.Run("syntax_error_listener", func(t *testing.T) {
		l := &SyntaxErrorListener{}
		l.SyntaxError(nil, nil, 1, 10, "e1", nil)
		l.SyntaxError(nil, nil, 2, 15, "e2", nil)
		assert.Equal(t, 2, len(l.GetAllErrors()))
		assert.Contains(t, l.GetFirstError().Error(), "e1")
	})
}
