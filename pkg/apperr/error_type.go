package apperr

type ErrorType int

const (
	ErrNotFound ErrorType = iota + 1
	ErrUnauthorized
	ErrBadRequest
	ErrForbidden
)

var typeMap = map[ErrorType]string{
	ErrNotFound:     "NotFound",
	ErrBadRequest:   "BadRequest",
	ErrUnauthorized: "Unauthorized",
}

func MakeType(err AppError) string {
	if errType, ok := typeMap[err.Type()]; ok {
		return errType
	}

	return "UnknownType"
}
