// Code generated by sqlc. DO NOT EDIT.
// source: company.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createCompany = `-- name: CreateCompany :exec
INSERT INTO "company" (
    "description"
) VALUES (
    $1
) RETURNING id, description
`

func (q *Queries) CreateCompany(ctx context.Context, description sql.NullString) error {
	_, err := q.exec(ctx, q.createCompanyStmt, createCompany, description)
	return err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM "company"
WHERE "id" = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteCompanyStmt, deleteCompany, id)
	return err
}

const getCompany = `-- name: GetCompany :one
SELECT id, description FROM "company"
WHERE "id" = $1
`

func (q *Queries) GetCompany(ctx context.Context, id uuid.UUID) (Company, error) {
	row := q.queryRow(ctx, q.getCompanyStmt, getCompany, id)
	var i Company
	err := row.Scan(&i.ID, &i.Description)
	return i, err
}

const listCompanies = `-- name: ListCompanies :many
SELECT id, description FROM "company"
ORDER BY "id"
`

func (q *Queries) ListCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.query(ctx, q.listCompaniesStmt, listCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(&i.ID, &i.Description); err != nil {
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

const updateCompany = `-- name: UpdateCompany :exec
UPDATE "company" SET (
    "description"
) = (
    $1
) WHERE "id" = $2
RETURNING id, description
`

func (q *Queries) UpdateCompany(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.updateCompanyStmt, updateCompany, id)
	return err
}
