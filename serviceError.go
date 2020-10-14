package serviceerror

import (
	"bitbucket.org/HeilaSystems/serviceerror/commonError"
	"bitbucket.org/HeilaSystems/serviceerror/errortypes"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"strings"
	"text/template"
)

type ServiceError interface {
	WithError(error) ServiceError
	GetError() error
	WithReplyValues(ValuesMap) ServiceError
	GetReplyValues() ValuesMap
	WithLogMessage(string) ServiceError
	GetLogMessage() *string
	WithLogValues(ValuesMap) ServiceError
	GetLogValues() ValuesMap
	GetActionLog() string
	GetSource() string
	GetUserError() string
	GetErrorType() errortypes.ErrorType
}

type BaseServiceError struct {
	source string
	logMessage  *string
	actionLog   string
	logValues   map[string]interface{}
	err         error
	errorType   errortypes.ErrorType
	userMessage string
	extraData   map[string]interface{}
}
type BaseServiceErrorReply struct {
	Msg  string               `json:"msg"`
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

type ValuesMap map[string]interface{}

var tmplError string = "Templating error"

func (se *BaseServiceError) WithError(err error) ServiceError {
	if se.err != nil{
		se.err = errors.Wrap(se.err,err.Error())
	} else {
		se.err = err
	}
	return se
}

func (se *BaseServiceError) GetError() error {
	return se.err
}

func (se *BaseServiceError) WithReplyValues(extraData ValuesMap) ServiceError {
	se.extraData = extraData
	return se
}

func (se *BaseServiceError) GetReplyValues() ValuesMap {
	return se.extraData
}

func (se *BaseServiceError) WithLogMessage(logMessage string) ServiceError {
	se.logMessage = &logMessage
	return se
}


func (se *BaseServiceError) GetLogMessage() *string {
	if se.logValues != nil && se.logMessage != nil {
		tmpl, err := template.New("logMessage").Parse(*se.logMessage)
		if err != nil {
			return &tmplError
		}
		buf := &bytes.Buffer{}
		if err := tmpl.Execute(buf, se.logValues); err != nil {
			return &tmplError
		}
		s := buf.String()
		return &s
	}
	return se.logMessage
}

func (se *BaseServiceError) WithLogValues(logValues ValuesMap) ServiceError {
	se.logValues = logValues
	return se
}

func (se *BaseServiceError) GetLogValues() ValuesMap {
	return se.logValues
}

func (se *BaseServiceError) GetUserError() string {
	return se.userMessage
}

func (se *BaseServiceError) GetErrorType() errortypes.ErrorType {
	return se.errorType
}

func (se *BaseServiceError) GetActionLog() string {
	return se.actionLog
}
func (se *BaseServiceError) GetSource() string {
	return se.source
}

func newServiceError(errType errortypes.ErrorType, err error, userMessage string, runTimeCaller int) ServiceError {
	runTimeCaller += 1
	pc, fn, line, _ := runtime.Caller(runTimeCaller)
	sourceArr :=  strings.Split(fn,"/")
	if len(sourceArr)>=2 {
		sourceArr = sourceArr[len(sourceArr)-2:]
	}

	formattedAction := fmt.Sprintf("error in %s",runtime.FuncForPC(pc).Name() )
	source := fmt.Sprintf("%s:%d" , strings.Join(sourceArr,"/"), line )

	return &BaseServiceError{
		source:      source,
		actionLog:   formattedAction,
		err:         err,
		errorType:   errType,
		userMessage: userMessage,
	}
}

func NewDbError(err error) ServiceError {
	return newServiceError(errortypes.DbError, err, commonError.InternalServiceError, 1)
}

func NewForbiddenError() ServiceError {
	return newServiceError(errortypes.Forbidden, nil, commonError.Forbidden, 1)
}

func NewLogicError(userMessage string) ServiceError {
	return newServiceError(errortypes.LogicError, nil, userMessage, 1)
}

//action - what caused the error

func NewValidationError(userMessage string) ServiceError {
	return newServiceError(errortypes.ValidationError, nil, userMessage, 1)
}

func NewNetworkError(err error) ServiceError {
	return newServiceError(errortypes.NetworkError, err, commonError.InternalServiceError, 1)
}

func NewIoError(err error) ServiceError {
	return newServiceError(errortypes.IoError, err, commonError.InternalServiceError, 1)
}

func NewBadRequestError(err error) ServiceError {
	return newServiceError(errortypes.BadRequestError, err, commonError.InternalServiceError, 1)
}

func NewUnauthorizedError(err error) ServiceError {
	return newServiceError(errortypes.UnauthorizedError, err, commonError.InternalServiceError, 1)
}

func NewNoContentError(userMessage string) ServiceError {
	return newServiceError(errortypes.NoContent, nil, userMessage, 1)
}
func NewMethodNotAllowed(action, method string) ServiceError {
	return newServiceError(errortypes.MethodNotAllowed, fmt.Errorf(commonError.NoContent+action+"/"+method), commonError.NoContent, 1)
}

