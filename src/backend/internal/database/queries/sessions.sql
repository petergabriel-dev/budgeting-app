-- name: CreateSession :one
INSERT INTO sessions (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetSessionByToken :one
SELECT s.*, u.id as user_id, u.email as user_email
FROM sessions s
JOIN users u ON s.user_id = u.id
WHERE s.token = $1 AND s.expires_at > NOW()
LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE token = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions
WHERE expires_at <= NOW();

-- name: DeleteUserSessions :exec
DELETE FROM sessions
WHERE user_id = $1;
