-- name: ListApplications :many
SELECT * FROM applications
WHERE status = $1
ORDER BY first_name;

-- name: CreateApplication :one
INSERT INTO applications (
  id, first_name, last_name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateApplication :exec
UPDATE applications
SET first_name = $1, last_name = $2, status = $3
WHERE id = $4;
