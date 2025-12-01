package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrSQLError            = errors.New("database failed to execute query")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrValidatioin         = errors.New("validation error")
	ErrSecretKey           = errors.New("Belum mendapatkan Secret key")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrForbidden,
	ErrValidatioin,
	ErrSecretKey,
}
