package database

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"os"
	"time"

	"github.com/hermitpopcorn/rsaboard/types"
	_ "modernc.org/sqlite"
)

type SQLiteDatabase struct {
	context context.Context
	queries *Queries
}

//go:embed sql/schema.sql
var ddl string

// Opens a local SQLite database.
func OpenSQLiteDatabase(file string) (*SQLiteDatabase, DatabaseError) {
	ctx := context.Background()

	if file == "" {
		panic("No database file specified.")
	}

	if _, err := os.Stat(file); err != nil {
		os.Create(file)
	}

	db, err := sql.Open("sqlite", "file:"+file)
	if err != nil {
		return nil, makeDatabaseErrorFromGenericError(err)
	}

	if err := db.Ping(); err != nil {
		return nil, makeDatabaseErrorFromGenericError(err)
	}

	// Initialize tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, makeDatabaseErrorFromGenericError(err)
	}

	queries := New(db)

	return &SQLiteDatabase{
		queries: queries,
		context: ctx,
	}, nil
}

func (db *SQLiteDatabase) CreateMessage(code string, encryptedText string, email string, burnAfterMinutes int) (types.Message, DatabaseError) {
	message, err := db.queries.CreateMessage(db.context, CreateMessageParams{
		Code:                code,
		EncryptedText:       encryptedText,
		AuthorEmail:         email,
		ShouldBurn:          burnAfterMinutes >= 0,
		ShouldBurnInMinutes: sql.NullInt64{Int64: int64(burnAfterMinutes), Valid: burnAfterMinutes >= 0},
		CreatedAt:           time.Now(),
	})
	return message.Convert(), makeDatabaseErrorFromGenericError(err)
}

func (db *SQLiteDatabase) FindMessage(code string) (types.Message, DatabaseError) {
	message, err := db.queries.GetMessageByCode(db.context, code)
	if message.ID == 0 {
		return types.Message{}, &MessageNotFoundError{}
	}
	return message.Convert(), makeDatabaseErrorFromGenericError(err)
}

func (db *SQLiteDatabase) CheckCodeExists(code string) (bool, DatabaseError) {
	message, err := db.FindMessage(code)

	var notFoundError *MessageNotFoundError
	if errors.As(err, &notFoundError) {
		return false, nil
	}

	return (message.ID != 0), makeDatabaseErrorFromGenericError(err)
}

func (db *SQLiteDatabase) CreateAccessLog(message *types.Message, remoteAddress string) (types.AccessLog, DatabaseError) {
	accessLog, err := db.queries.CreateAccessLog(db.context, CreateAccessLogParams{
		MessageID:     message.ID,
		RemoteAddress: remoteAddress,
		AccessedAt:    time.Now(),
	})
	return accessLog.Convert(), makeDatabaseErrorFromGenericError(err)
}

func (db *SQLiteDatabase) SetMessageDeleteAt(message *types.Message, when time.Time) DatabaseError {
	err := db.queries.SetDeleteAtOnMessage(db.context, SetDeleteAtOnMessageParams{
		DeleteAt: sql.NullTime{
			Time:  when,
			Valid: true,
		},
		ID: message.ID,
	})

	return makeDatabaseErrorFromGenericError(err)
}

func (db *SQLiteDatabase) DeleteExpiredMessages(when time.Time) DatabaseError {
	err := db.queries.DeleteMessageWhereDeleteAtLessThanTime(db.context, sql.NullTime{
		Time:  when,
		Valid: true,
	})

	return makeDatabaseErrorFromGenericError(err)
}

func (message *Message) Convert() types.Message {
	shouldBurnInMinutes := message.ShouldBurnInMinutes.Int64
	if !message.ShouldBurn {
		shouldBurnInMinutes = 0
	}

	deleteAt := &message.DeleteAt.Time
	if !message.DeleteAt.Valid {
		deleteAt = nil
	}

	return types.Message{
		ID:                  message.ID,
		Code:                message.Code,
		EncryptedText:       message.EncryptedText,
		AuthorEmail:         message.AuthorEmail,
		ShouldBurn:          message.ShouldBurn,
		ShouldBurnInMinutes: shouldBurnInMinutes,
		DeleteAt:            deleteAt,
	}
}

func (accessLog *AccessLog) Convert() types.AccessLog {
	return types.AccessLog{
		MessageID:     accessLog.MessageID,
		RemoteAddress: accessLog.RemoteAddress,
		AccessedAt:    &accessLog.AccessedAt,
	}
}
