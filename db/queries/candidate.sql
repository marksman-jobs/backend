CREATE TABLE "candidate" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL REFERENCES "user",
  "resume_id" char(64)
);

-- name: GetCandidate :one
SELECT * FROM "candidate"
WHERE "id" = $1;

-- name: ListCandidates :many
SELECT * FROM "candidate"
ORDER BY "id";

-- name: CreateCandidate :exec
INSERT INTO "candidate" (
    "user_id",
    "resume_id"
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateCandidate :exec
UPDATE "candidate" SET (
    "resume_id"
) = (
    $1
) WHERE "id" = $2
RETURNING *;

-- name: DeleteCandidate :exec
DELETE FROM "candidate"
WHERE "id" = $1;
