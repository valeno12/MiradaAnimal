-- name: CreateAnimalAtributo :exec
INSERT INTO animal_atributos (animal_id, atributo_id, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW());

-- name: GetAnimalAtributosByAnimalID :many
SELECT id, animal_id, atributo_id, created_at, updated_at, deleted_at
FROM animal_atributos
WHERE animal_id = $1 AND deleted_at IS NULL;

-- name: GetAnimalAtributosByAtributoID :many
SELECT id, animal_id, atributo_id, created_at, updated_at, deleted_at
FROM animal_atributos
WHERE atributo_id = $1 AND deleted_at IS NULL;

-- name: DeleteAnimalAtributo :exec
UPDATE animal_atributos
SET deleted_at = NOW()
WHERE id = $1;

-- name: UpdateAnimalAtributo :exec
UPDATE animal_atributos
SET animal_id = $2, atributo_id = $3, updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL;
