// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: usuario.sql

package models

import (
	"context"
	"database/sql"
)

const createUsuario = `-- name: CreateUsuario :exec
INSERT INTO usuario (nombre, email, password, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
`

func (q *Queries) CreateUsuario(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createUsuario)
	return err
}

const deleteUsuario = `-- name: DeleteUsuario :exec
UPDATE usuario
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteUsuario(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUsuario)
	return err
}

const getUsuarioByID = `-- name: GetUsuarioByID :one
SELECT id, nombre, email, password, created_at, updated_at
FROM usuario
WHERE id = $1 AND deleted_at IS NULL
`

type GetUsuarioByIDRow struct {
	ID        uint64
	Nombre    string
	Email     string
	Password  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (q *Queries) GetUsuarioByID(ctx context.Context) (GetUsuarioByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUsuarioByID)
	var i GetUsuarioByIDRow
	err := row.Scan(
		&i.ID,
		&i.Nombre,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsuarios = `-- name: GetUsuarios :many
SELECT id, nombre, email, password, created_at, updated_at
FROM usuario
WHERE deleted_at IS NULL
`

type GetUsuariosRow struct {
	ID        uint64
	Nombre    string
	Email     string
	Password  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (q *Queries) GetUsuarios(ctx context.Context) ([]GetUsuariosRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsuarios)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsuariosRow
	for rows.Next() {
		var i GetUsuariosRow
		if err := rows.Scan(
			&i.ID,
			&i.Nombre,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUsuario = `-- name: UpdateUsuario :exec
UPDATE usuario
SET nombre = $2, email = $3, password = $4, updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) UpdateUsuario(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateUsuario)
	return err
}
