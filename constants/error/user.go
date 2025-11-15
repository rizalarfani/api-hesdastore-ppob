package error

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordInCorrect = errors.New("password incorrect")
	ErrApiKeyInvalid     = errors.New("Api Key Invalid")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrApiKeyInvalid,
}
