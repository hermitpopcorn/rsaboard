package types

import "time"

type Message struct {
	ID                  int64
	Code                string
	EncryptedText       string
	AuthorEmail         string
	ShouldBurn          bool
	ShouldBurnInMinutes int64
	DeleteAt            *time.Time
}

type AccessLog struct {
	MessageID     int64
	RemoteAddress string
	AccessedAt    *time.Time
}
