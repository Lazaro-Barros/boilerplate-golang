-- name: GetTodo :one
SELECT uuid, title, description, done
FROM todo
WHERE uuid = $1;

-- name: ListTodos :many
SELECT uuid, title, description, done
FROM todo;

-- name: CreateTodo :one
INSERT INTO todo (uuid, title, description, done)
VALUES ($1, $2, $3, $4)
RETURNING uuid, title, description, done;

-- name: UpdateTodo :exec
UPDATE todo
SET title = $2, description = $3, done = $4
WHERE uuid = $1;

-- name: DeleteTodo :exec
DELETE FROM todo
WHERE uuid = $1;