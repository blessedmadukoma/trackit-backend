-- name: CreateSession :one
INSERT INTO sessions (
 id,
 userid,
 email,
 refresh_token,
 user_agent,
 client_ip,
 is_blocked,
 expires_at
) VALUES (
 $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions 
WHERE id = $1 LIMIT 1;

-- -- name: GetSessionByAccessToken :one
-- SELECT * FROM sessions
-- WHERE email = $1
-- ORDER BY created_at DESC LIMIT 1;