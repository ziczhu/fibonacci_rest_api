package api

// ErrorCode is a five-char code
type ErrorCode int

const (
	// 400s defines the codes of caller's error
	InvalidNumberErrorCode   = ErrorCode(40000)
	NegativeNumberErrorCode  = ErrorCode(40001)
	OverLimitNumberErrorCode = ErrorCode(40002)
)
