package serviceerror

import (
	"bitbucket.org/HeilaSystems/serviceerror/commonError"
	"bitbucket.org/HeilaSystems/serviceerror/status"
	"bitbucket.org/HeilaSystems/serviceerror/types"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"strings"
	"text/template"
)

var SecurityError = errors.New("security error")

type ServiceReply interface {
	error
	WithError(error) ServiceReply
	GetError() error
	WithReplyValues(ValuesMap) ServiceReply
	GetReplyValues() ValuesMap
	WithLogMessage(string) ServiceReply
	GetLogMessage() *string
	WithLogValues(ValuesMap) ServiceReply
	GetLogValues() ValuesMap
	GetActionLog() string
	GetSource() string
	GetUserError() string
	GetErrorType() *types.ReplyType
}

const statusError = "error"
type BaseServiceError struct {
	source      string
	logMessage  *string
	actionLog   string
	logValues   ValuesMap
	err         error
	errorType   *types.ReplyType
	userMessage string
	extraData   ValuesMap
}
type Message struct {
	Id  string               `json:"id"`
	Values map[string]interface{} `json:"values"`
}

type BaseResponse struct {
	Status status.Status                 `json:"status"`
	Message  *Message `json:"message,omitempty"`
}

type Response struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"`
}

type ValuesMap map[string]interface{}

var tmplError string = "Templating error"

func (se *BaseServiceError) WithError(err error) ServiceReply {
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
func (se *BaseServiceError) Error() string{
	sr := Response{
		BaseResponse: BaseResponse{
			Status: statusError,
			Message: &Message{
				Id:    se.GetUserError(),
				Values: se.GetReplyValues(),
			} ,
		},
		Data:             nil,
	}
	seBytesArr  , _ := json.Marshal(sr)
	return string(seBytesArr)
}
func (se *BaseServiceError) Parse(err string) (Response,error){
	parsedSr := Response{}
	if err := json.Unmarshal([]byte(err) , &parsedSr);err != nil {
		return Response{},err
	}
	return parsedSr ,nil
}
func (se *BaseServiceError) WithReplyValues(extraData ValuesMap) ServiceReply {
	se.extraData = extraData
	return se
}

func (se *BaseServiceError) GetReplyValues() ValuesMap {
	return se.extraData
}

func (se *BaseServiceError) WithLogMessage(logMessage string) ServiceReply {
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

func (se *BaseServiceError) WithLogValues(logValues ValuesMap) ServiceReply {
	se.logValues = logValues
	return se
}

func (se *BaseServiceError) GetLogValues() ValuesMap {
	return se.logValues
}

func (se *BaseServiceError) GetUserError() string {
	return se.userMessage
}

func (se *BaseServiceError) GetErrorType() *types.ReplyType {
	return se.errorType
}

func (se *BaseServiceError) GetActionLog() string {
	return se.actionLog
}
func (se *BaseServiceError) GetSource() string {
	return se.source
}

func NewServiceError(errType *types.ReplyType, err error, userMessage string, runTimeCaller int) ServiceReply {
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

func NewDbError(err error) ServiceReply {
	et := types.DbErrorType
	return NewServiceError(&et, err, commonError.InternalServiceError, 1)
}

func NewForbiddenError() ServiceReply {
	et := types.ForbiddenErrorType
	return NewServiceError(&et, nil, commonError.Forbidden, 1)
}

//func NewLogicError(userMessage string) ServiceReply {
//	et := types.LogicErrorType
//	return NewServiceError(&et, nil, userMessage, 1)
//}

//action - what caused the error

func NewValidationError(userMessage string) ServiceReply {
	et := types.ValidationErrorType
	return NewServiceError(&et, nil, userMessage, 1)
}

func NewNetworkError(err error) ServiceReply {
	et := types.NetworkErrorType
	return NewServiceError(&et, err, commonError.InternalServiceError, 1)
}

func NewIoError(err error) ServiceReply {
	et := types.IoErrorType
	return NewServiceError(&et, err, commonError.InternalServiceError, 1)
}

func NewBadRequestError(err error) ServiceReply {
	et := types.BadRequestErrorType
	return NewServiceError(&et, err, commonError.InternalServiceError, 1)
}

func NewUnauthorizedError(err error) ServiceReply {
	et := types.UnauthorizedErrorType
	return NewServiceError(&et, err, commonError.InternalServiceError, 1)
}

func NewNoContentError(userMessage string) ServiceReply {
	et := types.NoContentErrorType
	return NewServiceError(&et, nil, userMessage, 1)
}

func NewMethodNotAllowed(action, method string) ServiceReply {
	et := types.MethodNotAllowedErrorType
	return NewServiceError(&et, fmt.Errorf(commonError.NoContent+action+"/"+method), commonError.NoContent, 1)
}

func NewMessage(userMessage string)ServiceReply {
	return NewServiceError(nil, nil, userMessage, 1)
}

func NewNoMatchReply(userMessage string) ServiceReply {
	et := types.NoMatchErrorType
	return NewServiceError(&et, nil, userMessage, 1)
}

func LogicUnauthorizedErrorType(userMessage string) ServiceReply {
	et := types.LogicUnauthorizedErrorType
	return NewServiceError(&et, nil, userMessage, 1)
}

