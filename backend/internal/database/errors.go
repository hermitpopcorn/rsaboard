package database

type DatabaseError interface {
	Error() string
	Code() int
}

type GeneralDatabaseError struct {
	errorMessage string
}

func (e *GeneralDatabaseError) Error() string {
	return e.errorMessage
}
func (e *GeneralDatabaseError) Code() int {
	return 500
}

func makeDatabaseErrorFromGenericError(err error) DatabaseError {
	if err == nil {
		return nil
	}

	return &GeneralDatabaseError{
		errorMessage: err.Error(),
	}
}

type MessageNotFoundError struct{}

func (e *MessageNotFoundError) Error() string {
	return "No message was found."
}
func (e *MessageNotFoundError) Code() int {
	return 404
}
