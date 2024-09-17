package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Lazaro-Barros/boilerplate-golang/command/domain/entities"
	"github.com/Lazaro-Barros/boilerplate-golang/pkg/random"
)

func TestTodo(t *testing.T) {
	t.Run("Test NewTodo with valid data", func(t *testing.T) {
		boardUUID := "board-uuid"
		title := random.String(10)
		description := random.String(10)
		todo, err := entities.NewTodo(
			title,
			description,
			boardUUID)

		assert.NoError(t, err)
		assert.NotNil(t, todo)
		assert.Equal(t, title, todo.Title())
		assert.Equal(t, description, todo.Description())
		assert.Equal(t, boardUUID, todo.BoardUUID())
		assert.False(t, todo.Done())
		assert.NotEmpty(t, todo.Uuid())

	})

}
