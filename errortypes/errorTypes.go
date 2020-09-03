package errortypes

const (
	DbError              ErrorType = "Db"
	LogicError           ErrorType = "Logic"
	ValidationError      ErrorType = "Validation"
	NetworkError         ErrorType = "Network"
	IoError              ErrorType = "Io"
	BadRequestError      ErrorType = "BadRequest"
	NoContent            ErrorType = "NoContent"
	UnauthorizedError    ErrorType = "Unauthorized"
	Forbidden            ErrorType = "Forbidden"
	InternalServiceError ErrorType = "Internal service error"
	MethodNotAllowed     ErrorType = "MethodNotAllowed"
)

type ErrorType string
