CREATE TABLE "company" (
  "id" uuid PRIMARY KEY,
  "description" varchar(1024)
);

-- name: GetCompany :one
SELECT * FROM "company"
WHERE "id" = $1;

-- name: ListCompanies :many
SELECT * FROM "company"
ORDER BY "id";

-- name: CreateCompany :exec
INSERT INTO "company" (
    "description"
) VALUES (
    $1
) RETURNING *;

-- name: UpdateCompany :exec
UPDATE "company" SET (
    "description"
) = (
    $1
) WHERE "id" = $2
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM "company"
WHERE "id" = $1;
