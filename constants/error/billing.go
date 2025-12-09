package error

import "errors"

var (
	ErrBillNotAvailable = errors.New("Tagihan belum tersedia")
)

var BillingErrors = []error{
	ErrBillNotAvailable,
}
