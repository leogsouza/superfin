-- name: ListUsers :many
SELECT * from users ORDER by id
LIMIT $1 OFFSET $2;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (
  email,
  password
) VALUES ($1, $2) RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users SET password = $1
where ID = $2 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;