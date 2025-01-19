-- name: CreateAnimal :exec
INSERT INTO animal (nombre, especie, foto, sexo, edad_meses, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW());

-- name: GetAnimalByID :one
SELECT id, nombre, especie, foto, sexo, edad_meses, created_at, updated_at
FROM animal
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetAnimales :many
SELECT id, nombre, especie, foto, sexo, edad_meses, created_at, updated_at
FROM animal
WHERE deleted_at IS NULL;

-- name: UpdateAnimal :exec
UPDATE animal
SET nombre = $2, especie = $3, foto = $4, sexo = $5, edad_meses = $6, updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteAnimal :exec
UPDATE animal
SET deleted_at = NOW()
WHERE id = $1;
