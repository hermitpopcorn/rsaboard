-- name: GetMessageByCode :one
SELECT * FROM "message"
WHERE "code" = ? LIMIT 1;

-- name: CreateMessage :one
INSERT INTO "message" (
  "code", "encrypted_text", "author_email",  "should_burn", "should_burn_in_minutes", "created_at"
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: CreateAccessLog :one
INSERT INTO "access_log" (
  "message_id", "remote_address", "accessed_at"
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: SetDeleteAtOnMessage :exec
UPDATE "message"
SET "delete_at" = ?
WHERE "id" = ?;

-- name: DeleteMessage :exec
DELETE FROM "message"
WHERE "id" = ?;

-- name: DeleteMessageWhereDeleteAtLessThanTime :exec
DELETE FROM "message"
WHERE "delete_at" < ?;
