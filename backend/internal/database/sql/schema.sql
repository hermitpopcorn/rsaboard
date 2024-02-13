CREATE TABLE IF NOT EXISTS "message" (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"code" TEXT NOT NULL,
	"encrypted_text" TEXT NOT NULL,
	"author_email" TEXT NOT NULL,
	"should_burn" BOOLEAN NOT NULL DEFAULT 0,
	"should_burn_in_minutes" INTEGER,
	"created_at" DATETIME NOT NULL,
	"delete_at" DATETIME
);

CREATE TABLE IF NOT EXISTS "access_log" (
	"message_id" INTEGER NOT NULL,
	"remote_address" TEXT NOT NULL,
	"accessed_at" DATETIME NOT NULL
);
