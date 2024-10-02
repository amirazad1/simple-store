-- name: GetUser :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username DESC
LIMIT ?
OFFSET ?;

-- name: CreateUser :execresult
INSERT INTO users (
    username, password, full_name, mobile, password_changed_at
) VALUES (
             ?, ?, ?, ?, ?
         );

-- name: UpdateUser :execresult
UPDATE users
SET full_name=?, mobile=?
WHERE username = ?;

-- name: UpdateUserPassword :execresult
UPDATE users
SET password=?, password_changed_at=?
WHERE username = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = ?;