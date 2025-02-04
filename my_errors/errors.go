package my_errors

import "fmt"

// CustomError represents a structured error with an HTTP status code.
type CustomError struct {
	Code    int    `json:"code"`    // HTTP status code
	Message string `json:"message"` // Error message
}

// Implement the error interface for CustomError
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Helper functions to create common errors
func NewBadRequestError(message string) *CustomError {
	return &CustomError{
		Code:    400,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *CustomError {
	return &CustomError{
		Code:    401,
		Message: message,
	}
}

func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		Code:    404,
		Message: message,
	}
}

func NewConflictError(message string) *CustomError {
	return &CustomError{
		Code:    409,
		Message: message,
	}
}

func NewInternalError(message string) *CustomError {
	return &CustomError{
		Code:    500,
		Message: message,
	}
}

// Predefined custom errors
var (
	ErrNoUsersFound = &CustomError{
		Code:    404,
		Message: "No users found",
	}
)
