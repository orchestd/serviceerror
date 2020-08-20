package errortypes

const (
	DbError              ErrorType = "Db"
	LogicError           ErrorType = "Logic"
	ValidationError      ErrorType = "Validation"
	NetworkError         ErrorType = "Network"
	IoError              ErrorType = "Io"
	BadRequestError      ErrorType = "BadRequest"
	NoContent            ErrorType = "noContent"
	UnauthorizedError    ErrorType = "Unauthorized"
	Forbidden            ErrorType = "Forbidden"
	InternalServiceError ErrorType = "Internal service error"
)

type ErrorType string
