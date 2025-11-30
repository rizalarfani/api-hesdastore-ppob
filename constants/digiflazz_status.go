package constants

type TransactionStatus int
type TransactionStatusString string

const (
	PENDING TransactionStatus = 0
	SUCCESS TransactionStatus = 1
	FAILED  TransactionStatus = 2
	PROSESS TransactionStatus = 3
	INQUIRY TransactionStatus = 4

	PENDINGSTRING TransactionStatusString = "pending"
	SUCCESSSTRING TransactionStatusString = "success"
	FAILEDSTRING  TransactionStatusString = "failed"
	PROSESSSTRING TransactionStatusString = "prosess"
	INQUIRYTRING  TransactionStatusString = "inquiry"
)

var mapStatusStringToInt = map[TransactionStatusString]TransactionStatus{
	PENDINGSTRING: PENDING,
	SUCCESSSTRING: SUCCESS,
	FAILEDSTRING:  FAILED,
	PROSESSSTRING: PROSESS,
	INQUIRYTRING:  INQUIRY,
}

var mapStatusIntToString = map[TransactionStatus]TransactionStatusString{
	PENDING: PENDINGSTRING,
	SUCCESS: SUCCESSSTRING,
	FAILED:  FAILEDSTRING,
	PROSESS: PROSESSSTRING,
	INQUIRY: INQUIRYTRING,
}

func (p TransactionStatus) GetStatusString() TransactionStatusString {
	return mapStatusIntToString[p]
}

func (p TransactionStatusString) GetStatusInt() TransactionStatus {
	return mapStatusStringToInt[p]
}
