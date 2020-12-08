package types

const (
	DbErrorType                ReplyType = "Db"
	ValidationErrorType        ReplyType = "Validation"
	NetworkErrorType           ReplyType = "Network"
	IoErrorType                ReplyType = "Io"
	BadRequestErrorType        ReplyType = "BadRequest"
	NoContentErrorType         ReplyType = "NoContent"
	UnauthorizedErrorType      ReplyType = "Unauthorized"
	LogicUnauthorizedErrorType ReplyType = "LogicUnauthorized"
	ForbiddenErrorType         ReplyType = "Forbidden"
	InternalServiceErrorType   ReplyType = "Internal service error"
	MethodNotAllowedErrorType  ReplyType = "MethodNotAllowed"
	NoMatchErrorType           ReplyType = "NoMatch"
)

type ReplyType string
