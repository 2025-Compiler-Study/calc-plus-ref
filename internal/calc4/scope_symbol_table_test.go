package calc4

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScopeSymbolTable_Lifecycle(t *testing.T) {
	t.Run("Global scope access", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		err := symTable.Register("globalVar")
		require.NoError(t, err)

		val, err := symTable.GetVariable("globalVar")
		assert.NoError(t, err)
		assert.Equal(t, 0, val)
	})

	t.Run("Inner scope can access outer variable", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		err := symTable.Register("outer")
		require.NoError(t, err)

		symTable.PushScope()
		val, err := symTable.GetVariable("outer")
		assert.NoError(t, err)
		assert.Equal(t, 0, val)
	})

	t.Run("Outer scope cannot access inner variable", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		symTable.PushScope()
		err := symTable.Register("inner")
		require.NoError(t, err)

		symTable.PopScope()
		_, err = symTable.GetVariable("inner")
		assert.Error(t, err)
	})
}

func TestScopeSymbolTable_Shadowing(t *testing.T) {
	t.Run("Re-register in inner scope", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		err := symTable.Register("x")
		require.NoError(t, err)
		err = symTable.SetVariable("x", 10)
		require.NoError(t, err)

		symTable.PushScope()
		err = symTable.Register("x")
		assert.NoError(t, err)
		val, err := symTable.GetVariable("x")
		assert.NoError(t, err)
		assert.Equal(t, 0, val)
	})

	t.Run("Inner re-registered value does not affect outer scope", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		err := symTable.Register("x")
		require.NoError(t, err)
		err = symTable.SetVariable("x", 10)
		require.NoError(t, err)

		symTable.PushScope()
		err = symTable.Register("x")
		require.NoError(t, err)
		val, err := symTable.GetVariable("x")
		require.NoError(t, err)
		require.Equal(t, 0, val)

		err = symTable.SetVariable("x", 20)
		assert.NoError(t, err)
		symTable.PopScope()
		val, err = symTable.GetVariable("x")
		assert.NoError(t, err)
		assert.Equal(t, 10, val)
	})
}

func TestScopeSymbolTable_SetVariable(t *testing.T) {
	t.Run("Update nearest Table variable", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		err := symTable.Register("a")
		require.NoError(t, err)

		symTable.PushScope()
		symTable.PushScope()
		err = symTable.SetVariable("a", 999)
		require.NoError(t, err)

		symTable.PopScope()
		symTable.PopScope()
		val, err := symTable.GetVariable("a")
		assert.Equal(t, 999, val)
	})

	t.Run("Set fails for unknown variable", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		symTable.PushScope()

		err := symTable.SetVariable("unknown", 123)
		assert.Error(t, err)
	})
}

func TestScopeSymbolTable_Redeclaration(t *testing.T) {
	t.Run("Redeclaration in same scope fails", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		symTable.PushScope()
		err := symTable.Register("a")
		require.NoError(t, err)

		err = symTable.Register("a")
		assert.Error(t, err)
	})

	t.Run("Redeclaration in different scope succeeds", func(t *testing.T) {
		symTable := NewScopeSymbolTable()
		err := symTable.Register("a")
		require.NoError(t, err)

		symTable.PushScope()
		err = symTable.Register("a")
		assert.NoError(t, err)
	})
}
