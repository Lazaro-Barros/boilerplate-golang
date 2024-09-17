package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/Lazaro-Barros/boilerplate-golang/command/domain/entities"
	"github.com/Lazaro-Barros/boilerplate-golang/queries_sqlc/db"
)

type TodoRepository struct {
	db *db.Queries
}

func NewTodoRepository(db *db.Queries) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Find(ctx context.Context, id string) (*entities.Todo, error) {
	uuid, _ := uuid.Parse(id)
	todo, err := r.db.GetTodo(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return entities.RestoreFromDB(
		todo.Uuid.String(),
		todo.Title,
		todo.Description.String,
		todo.Done,
	), nil

}

func (r *TodoRepository) Create(ctx context.Context, todo *entities.Todo) error {
	_, err := r.db.CreateTodo(ctx, db.CreateTodoParams{
		Uuid:        uuid.MustParse(todo.Uuid()),
		Title:       todo.Title(),
		Description: sql.NullString{String: todo.Description(), Valid: true},
		Done:        todo.Done(),
	})
	return err
}

func (r *TodoRepository) Update(ctx context.Context, id string, todo *entities.Todo) error {
	uuid, _ := uuid.Parse(id)
	err := r.db.UpdateTodo(ctx, db.UpdateTodoParams{
		Uuid:        uuid,
		Title:       todo.Title(),
		Description: sql.NullString{String: todo.Description(), Valid: true},
		Done:        todo.Done(),
	})
	return err
}

func (r *TodoRepository) Delete(ctx context.Context, id string) error {
	uuid, _ := uuid.Parse(id)
	err := r.db.DeleteTodo(ctx, uuid)
	return err
}
