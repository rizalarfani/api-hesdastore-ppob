package error

func ErrMapping(err error) bool {
	var (
		GeneralErrors     = GeneralErrors
		ProductErrors     = ProductErrors
		TransactionErrors = TransactionErrors
		UserErrors        = UserErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors...)
	allErrors = append(allErrors, ProductErrors...)
	allErrors = append(allErrors, TransactionErrors...)
	allErrors = append(allErrors, UserErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
