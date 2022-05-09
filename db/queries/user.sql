CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "first_name" varchar(32) NOT NULL,
  "last_name" varchar(32) NOT NULL,
  "email" varchar(96) NOT NULL,
  "is_email_verified" boolean NOT NULL,
  "password_hash" char(72) NOT NULL,
  "pronouns" pronouns NOT NULL,
  "is_candidate" boolean NOT NULL,
  "is_recruiter" boolean NOT NULL
);

-- name: GetUser :one
SELECT * FROM "user"
WHERE "id" = $1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY "id";

-- name: CreateUser :exec
INSERT INTO "user" (
  "first_name",
  "last_name",
  "email",
  "is_email_verified",
  "password_hash",
  "pronouns",
  "is_candidate",
  "is_recruiter"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "user" SET (
  "first_name",
  "last_name",
  "email",
  "is_email_verified",
  "password_hash",
  "pronouns",
  "is_candidate",
  "is_recruiter"
) = (
    $1, $2, $3, $4, $5, $6, $7, $8
) WHERE "id" = $9
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE "id" = $1;
