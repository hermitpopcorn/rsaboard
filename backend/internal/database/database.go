// The functions in this package handles anything related to the database,
// be it querying for data or saving them.

package database

import (
	"time"

	"github.com/hermitpopcorn/rsaboard/types"
)

type Database interface {
	CreateMessage(code string, encryptedText string, email string, burnAfterMinutes int) (types.Message, DatabaseError)
	FindMessage(code string) (types.Message, DatabaseError)
	CheckCodeExists(code string) (bool, DatabaseError)
	CreateAccessLog(message *types.Message, remoteAddress string) (types.AccessLog, DatabaseError)
	SetMessageDeleteAt(message *types.Message, when time.Time) DatabaseError
	DeleteExpiredMessages(when time.Time) DatabaseError
}
