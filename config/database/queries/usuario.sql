-- name: CreateUsuario :exec
INSERT INTO usuario (nombre, email, password, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW());

-- name: GetUsuarioByID :one
SELECT id, nombre, email, password, created_at, updated_at
FROM usuario
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetUsuarios :many
SELECT id, nombre, email, password, created_at, updated_at
FROM usuario
WHERE deleted_at IS NULL;

-- name: UpdateUsuario :exec
UPDATE usuario
SET nombre = $2, email = $3, password = $4, updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteUsuario :exec
UPDATE usuario
SET deleted_at = NOW()
WHERE id = $1;
