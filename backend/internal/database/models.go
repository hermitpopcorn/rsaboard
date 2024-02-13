// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"
)

type AccessLog struct {
	MessageID     int64
	RemoteAddress string
	AccessedAt    time.Time
}

type Message struct {
	ID                  int64
	Code                string
	EncryptedText       string
	AuthorEmail         string
	ShouldBurn          bool
	ShouldBurnInMinutes sql.NullInt64
	CreatedAt           time.Time
	DeleteAt            sql.NullTime
}
