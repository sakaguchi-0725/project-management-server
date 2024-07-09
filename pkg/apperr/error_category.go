package apperr

type ErrorCategory int

func MakeCategory(err AppError) string {
	if category, ok := categoryMap[err.Category()]; ok {
		return category
	}

	return "UnkownCategory"
}

const (
	ErrCategoryCreateTaskInvalidParameter ErrorCategory = iota + 1
	ErrCategoryUpdateTaskInvalidParameter
	ErrCategoryDeleteTaskInvalidParameter
	ErrCategoryGetTaskByIdInvalidParameter
)

var categoryMap = map[ErrorCategory]string{
	ErrCategoryCreateTaskInvalidParameter:  "CreateTask/InvalidParameter",
	ErrCategoryUpdateTaskInvalidParameter:  "UpdateTask/InvalidParameter",
	ErrCategoryDeleteTaskInvalidParameter:  "DeleteTask/InvalidParameter",
	ErrCategoryGetTaskByIdInvalidParameter: "GetTaskById/InvalidParameter",
}
