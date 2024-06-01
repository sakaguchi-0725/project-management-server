package apperr

import "fmt"

type AppError interface {
	error
	Category() ErrorCategory
	Type() ErrorType
}

type appError struct {
	message  string
	category ErrorCategory
	errType  ErrorType
}

func (err *appError) Error() string {
	return err.message
}

func (err *appError) Category() ErrorCategory {
	return err.category
}

func (err *appError) Type() ErrorType {
	return err.errType
}

func NewAppError(errType ErrorType, errCategory ErrorCategory, format string, args ...any) error {
	return &appError{
		message:  fmt.Sprintf(format, args...),
		category: errCategory,
		errType:  errType,
	}
}
