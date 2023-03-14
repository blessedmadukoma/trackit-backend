-- name: CreateUser :one
INSERT INTO "user" (
firstname, lastname, email, phone, password, user_type
) VALUES (
 $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM "user" 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "user" 
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE "user"
SET 
firstname = $2,
lastname = $3,
password = $4,
email = $5,
user_type = $6,
phone = $7,
updated_at = now()
WHERE id = $1
-- WHERE email = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;