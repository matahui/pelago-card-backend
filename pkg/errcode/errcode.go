package errcode

const (
	// Success code
	Success = 0

	// System level error codes (1000xxx)
	ServerError       = 1000001 // Internal server error
	InvalidParams     = 1000002 // Invalid parameters
	NotFound          = 1000003 // Not found
	UnauthorizedError = 1000004 // Unauthorized
	TooManyRequests   = 1000005 // Too many requests

	// Merchant module error codes (1001xxx)
	MerchantNotFound     = 1001001 // Merchant not found
	MerchantExists       = 1001002 // Merchant already exists
	MerchantInactive     = 1001003 // Merchant inactive
	InvalidMerchantID    = 1001004 // Invalid merchant ID
	MerchantRegisterFail = 1001005 // Merchant registration failed
	MerchantUpdateFail   = 1001006 // Merchant update failed

	// Database error codes (1002xxx)
	DatabaseError     = 1002001 // Database error
	DataNotFound      = 1002002 // Data not found
	DataExists        = 1002003 // Data already exists
	TransactionFailed = 1002004 // Transaction failed
)
