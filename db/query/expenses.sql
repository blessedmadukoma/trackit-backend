-- name: CreateExpense :one
INSERT INTO "expenses" (
userid, email, amount, tag, description, date
) VALUES (
 $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetExpensesByID :one
SELECT * FROM "expenses" 
WHERE id = $1 LIMIT 1;

-- name: GetExpensesByUserEmail :one
SELECT * FROM "expenses" 
WHERE email = $1 LIMIT 1;

-- name: GetExpensesByUserID :one
SELECT * FROM "expenses" 
WHERE userid = $1 LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM "expenses"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateExpenses :one
UPDATE "expenses"
SET 
amount = $2,
description = $3,
date = $4,
tag = $5,
updated_at = now()
WHERE id = $1
-- WHERE email = $1
RETURNING *;

-- name: DeleteExpenses :exec
DELETE FROM "expenses"
WHERE id = $1;

-- -- name: GetCurrentExpensesByAccessToken :one
-- SELECT * FROM "expenses"
-- WHERE email = $1
-- LEFT JOIN expensess ON expensess.email = sessions.email
-- ORDER BY created_at DESC LIMIT 1;