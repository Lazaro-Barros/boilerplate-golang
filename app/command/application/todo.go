package application

import (
	"context"
	"fmt"

	"github.com/Lazaro-Barros/boilerplate-golang/command/domain/entities"
)

var (
	ErrTodoNotFound = fmt.Errorf("todo not found")
)

type TodoService struct {
	todoRepository entities.TodoRepository
}

func NewTodoService(todoRepository entities.TodoRepository) *TodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (s *TodoService) Create(ctx context.Context, in TodoIn) (string, error) {
	todo, err := entities.NewTodo(
		in.Title,
		in.Description,
		in.BoardUUID)

	if err != nil {
		return "", err
	}

	err = s.todoRepository.Create(ctx, todo)
	if err != nil {
		return "", err
	}

	return todo.Uuid(), nil
}

func (s *TodoService) Update(ctx context.Context, uuid string, in TodoIn) error {
	todo, err := s.todoRepository.Find(ctx, uuid)
	if err != nil {
		return err
	}
	if todo == nil {
		return ErrTodoNotFound
	}

	todo, err = entities.NewTodo(
		in.Title,
		in.Description,
		in.BoardUUID)
	if err != nil {
		return err
	}

	err = s.todoRepository.Update(ctx, uuid, todo)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) Delete(ctx context.Context, uuid string) error {
	return s.todoRepository.Delete(ctx, uuid)
}
