package entities

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrEmptyTodoTitle error = fmt.Errorf("title cannot be empty")
	ErrEmptyBoardUUID error = fmt.Errorf("board uuid cannot be empty")
)

type Todo struct {
	uuid        string
	title       string
	description string
	done        bool
	boardUUID   string
}

func NewTodo(
	title string,
	description string,
	boardUUID string) (*Todo, error) {

	// validations here
	if title == "" {
		return nil, ErrEmptyTodoTitle
	}

	if boardUUID == "" {
		return nil, ErrEmptyBoardUUID
	}

	return &Todo{
		uuid:        uuid.New().String(),
		title:       title,
		description: description,
		done:        false,
		boardUUID:   boardUUID,
	}, nil
}

func RestoreFromDB(
	uuid string,
	title string,
	description string,
	done bool) *Todo {
	return &Todo{
		uuid:        uuid,
		title:       title,
		description: description,
		done:        done,
	}
}

type TodoRepository interface {
	Find(ctx context.Context, uuid string) (*Todo, error)
	Create(ctx context.Context, todo *Todo) error
	Update(ctx context.Context, uuid string, todo *Todo) error
	Delete(ctx context.Context, uuid string) error
}

func (obj Todo) Uuid() string {
	return obj.uuid
}

func (obj Todo) Title() string {
	return obj.title
}

func (obj Todo) Description() string {
	return obj.description
}

func (obj Todo) Done() bool {
	return obj.done
}

func (obj Todo) BoardUUID() string {
	return obj.boardUUID
}
