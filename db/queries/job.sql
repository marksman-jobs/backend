CREATE TABLE "job" (
  "id" uuid PRIMARY KEY,
  "title" varchar(96) NOT NULL,
  "type" job_type NOT NULL,
  "function" job_function NOT NULL,
  "currency" currency NOT NULL,
  "value" int NOT NULL,
  "number_of_applicants" int,
  "company_id" uuid REFERENCES "company",
  "recruiter_id" uuid REFERENCES "recruiter",
  "location_id" uuid NOT NULL REFERENCES "location"
);

-- name: GetJob :one
SELECT * FROM "job"
WHERE "id" = $1;

-- name: ListJobs :many
SELECT * FROM "job"
ORDER BY "id";

-- name: CreateJob :exec
INSERT INTO "job" (
  "title",
  "type",
  "function",
  "currency",
  "value",
  "number_of_applicants",
  "company_id",
  "recruiter_id",
  "location_id"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: UpdateJob :exec
UPDATE "job" SET (
  "title",
  "type",
  "function",
  "currency",
  "value",
  "number_of_applicants",
  "company_id",
  "recruiter_id",
  "location_id"
) = (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) WHERE "company_id" = $10
RETURNING *;

-- name: DeleteJob :exec
DELETE FROM "job"
WHERE "id" = $1;
