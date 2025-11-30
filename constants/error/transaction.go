package error

import "errors"

var (
	ErrBalanceIsNotEnough  = errors.New("Saldo anda tidak mencukupi. Silahkan lakukan top up terlebih dahulu!")
	ErrBalanceIsZero       = errors.New("Saldo anda 0. Silahkan lakukan top up terlebih dahulu!")
	ErrProductIsFaulty     = errors.New("Product sedang ganguan")
	ErrProductIsAvalaible  = errors.New("Product tidak tersedia di HesdaStore!")
	ErrServiceNotAvailable = errors.New("Layanan sementara tidak tersedia. Silakan coba beberapa saat lagi!")
)

var TransactionErrors = []error{
	ErrBalanceIsNotEnough,
	ErrBalanceIsZero,
	ErrProductIsFaulty,
	ErrProductIsAvalaible,
	ErrServiceNotAvailable,
}
