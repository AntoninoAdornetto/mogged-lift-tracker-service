// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: template.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createTemplate = `-- name: CreateTemplate :execresult
INSERT INTO template (
	name,
	lifts,
	created_by
) VALUES (
	?,
	?,
	UUID_TO_BIN(?)
)
`

type CreateTemplateParams struct {
	Name      string          `json:"name"`
	Lifts     json.RawMessage `json:"lifts"`
	CreatedBy string          `json:"created_by"`
}

func (q *Queries) CreateTemplate(ctx context.Context, arg CreateTemplateParams) (sql.Result, error) {
	return q.exec(ctx, q.createTemplateStmt, createTemplate, arg.Name, arg.Lifts, arg.CreatedBy)
}

const getTemplate = `-- name: GetTemplate :one
SELECT id, name, lifts, date_last_used, created_by FROM template
WHERE id = ? AND created_by = UUID_TO_BIN(?)
`

type GetTemplateParams struct {
	ID        int32  `json:"id"`
	CreatedBy string `json:"created_by"`
}

func (q *Queries) GetTemplate(ctx context.Context, arg GetTemplateParams) (Template, error) {
	row := q.queryRow(ctx, q.getTemplateStmt, getTemplate, arg.ID, arg.CreatedBy)
	var i Template
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lifts,
		&i.DateLastUsed,
		&i.CreatedBy,
	)
	return i, err
}

const listTemplates = `-- name: ListTemplates :many
SELECT id, name, lifts, date_last_used, created_by FROM template
WHERE created_by = UUID_TO_BIN(?)
`

func (q *Queries) ListTemplates(ctx context.Context, createdBy string) ([]Template, error) {
	rows, err := q.query(ctx, q.listTemplatesStmt, listTemplates, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Template
	for rows.Next() {
		var i Template
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Lifts,
			&i.DateLastUsed,
			&i.CreatedBy,
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
