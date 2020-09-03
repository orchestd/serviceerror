package serviceerror

import (
	"bitbucket.org/HeilaSystems/serviceerror/commonError"
	"bitbucket.org/HeilaSystems/serviceerror/errortypes"
	"fmt"
)

type ServiceError interface {
	GetInternalError() error
	GetUserError() string
	GetErrorType() errortypes.ErrorType
	AddExtraData(map[string]interface{}) ServiceError
	GetExtraData() map[string]interface{}
}

type BaseServiceError struct {
	action      string
	err         error
	errorType   errortypes.ErrorType
	userMessage string
	extraData   map[string]interface{}
}
type BaseServiceErrorReply struct {
	Msg  string `json:"msg"`
	Type errortypes.ErrorType `json:"type"`
}

type BaseServiceReply struct {
	Status string                 `json:"status"`
	Error  *BaseServiceErrorReply `json:"error,omitempty"`
}

type ServiceReply struct {
	BaseServiceReply
	Data interface{} `json:"data,omitempty"`
}

func (se *BaseServiceError) GetInternalError() error {
	internalError := se.action
	if se.err != nil && se.action != se.err.Error() {
		internalError += " " + se.err.Error()
	}
	if len(se.extraData) > 0 {
		for key, val := range se.extraData {
			internalError += "\r\n" + key + ": " + fmt.Sprint(val)
		}
	}
	return fmt.Errorf(internalError)
}

func (se *BaseServiceError) GetUserError() string {
	return se.userMessage
}

func (se *BaseServiceError) GetErrorType() errortypes.ErrorType {
	return se.errorType
}

func (se *BaseServiceError) AddExtraData(extraData map[string]interface{}) ServiceError {
	se.extraData = extraData
	return se
}
func (se *BaseServiceError) GetExtraData() map[string]interface{} {
	return se.extraData
}

func NewServiceError(action string, errType errortypes.ErrorType, err error, userMessage string) ServiceError {
	return &BaseServiceError{
		action:      action,
		err:         err,
		errorType:  errType,
		userMessage: userMessage,
	}

}

func NewDbError(action string, err error) ServiceError {
	return NewServiceError(action, errortypes.DbError, err, commonError.InternalServiceError)
}

func NewForbiddenError(action string, err error) ServiceError {
	return NewServiceError(action, errortypes.Forbidden, err, commonError.Forbidden)
}

func NewLogicError(action string, err error, userMessage string) ServiceError {
	return NewServiceError(action, errortypes.LogicError, err, userMessage)
}

//action - what caused the error

func NewValidationError(action, userMessage string) ServiceError {
	return NewServiceError(action, errortypes.ValidationError, fmt.Errorf(userMessage), userMessage)
}

func NewNetworkError(action string, err error) ServiceError {
	return NewServiceError(action, errortypes.NetworkError, err, commonError.InternalServiceError)
}

func NewIoError(action string, err error) ServiceError {
	return NewServiceError(action, errortypes.IoError, err, commonError.InternalServiceError)
}

func NewBadRequestError(action string, err error) ServiceError {
	return NewServiceError(action, errortypes.BadRequestError, err, commonError.InternalServiceError)
}

func NewUnauthorizedError(action string, err error) ServiceError {
	return NewServiceError(action,  errortypes.UnauthorizedError, err, commonError.InternalServiceError)
}

func NewNoContentError(action string, err error, userMessage string) ServiceError {
	return NewServiceError(action, errortypes.NoContent, err, userMessage)
}
func NewMethodNotAllowed(action, method string) ServiceError {
	return NewServiceError(action, errortypes.MethodNotAllowed, fmt.Errorf(commonError.NoContent+action+"/"+method), commonError.NoContent)
}
