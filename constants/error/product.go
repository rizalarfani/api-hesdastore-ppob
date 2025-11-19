package error

import "errors"

var (
	ErrProductNotFound = errors.New("Product tidak ditemukan")
)

var ProductErrors = []error{
	ErrProductNotFound,
}
