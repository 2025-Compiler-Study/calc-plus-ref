package calc4

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleSymbolTable_Register(t *testing.T) {
	symTable := NewSimpleSymbolTable()

	t.Run("Register new", func(t *testing.T) {
		// Register new variable
		err := symTable.Register("new")
		assert.NoError(t, err)
	})

	t.Run("Register duplicate", func(t *testing.T) {
		// Registering "exist" variable
		err := symTable.Register("exist")
		require.NoError(t, err)

		// Registering "exist" variable again should return an error
		err = symTable.Register("exist")
		assert.Error(t, err)
	})
}

func TestSimpleSymbolTable_SetVariable(t *testing.T) {
	symTable := NewSimpleSymbolTable()

	t.Run("Set variable", func(t *testing.T) {
		// Register new variable
		err := symTable.Register("ten")
		require.NoError(t, err)

		// Set the value to 10 should succeed
		err = symTable.SetVariable("ten", 10)
		assert.NoError(t, err)
	})

	t.Run("Set unregistered variable", func(t *testing.T) {
		// Set the unregistered variable to 10 should fail
		err := symTable.SetVariable("unregistered", 10)
		assert.Error(t, err)
	})
}

func TestSimpleSymbolTable_GetVariable(t *testing.T) {
	symTable := NewSimpleSymbolTable()

	t.Run("Default value is zero", func(t *testing.T) {
		// Register new variable
		err := symTable.Register("new")
		require.NoError(t, err)

		// The new registered variable should have value 0
		value, err := symTable.GetVariable("new")
		assert.NoError(t, err)
		assert.Equal(t, 0, value)
	})

	t.Run("Unregistered variable", func(t *testing.T) {
		_, err := symTable.GetVariable("unregistered")
		assert.Error(t, err)
	})

	t.Run("Changed value", func(t *testing.T) {
		// Register new variable
		err := symTable.Register("ten")
		require.NoError(t, err)

		// Set the value to 10
		err = symTable.SetVariable("ten", 10)
		require.NoError(t, err)

		// Get the value of "ten" variable
		value, err := symTable.GetVariable("ten")
		assert.NoError(t, err)
		assert.Equal(t, 10, value)
	})
}
