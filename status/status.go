package status

import (
	"bitbucket.org/HeilaSystems/serviceerror/types"
)

var statusMap = map[types.ReplyType]Status{
	types.UnauthorizedErrorType:     UnauthorizedStatus,
	types.LogicUnauthorizedErrorType: UnauthorizedStatus,
	types.ForbiddenErrorType:        DeniedStatus,
	types.BadRequestErrorType:       ErrorStatus,
	types.DbErrorType:               ErrorStatus,
	types.InternalServiceErrorType:  ErrorStatus,
	types.IoErrorType:               ErrorStatus,
	types.NetworkErrorType:          ErrorStatus,
	types.NoContentErrorType:        NoMatchStatus,
	types.ValidationErrorType:       InvalidStatus,
	types.MethodNotAllowedErrorType: DeniedStatus,
	types.NoMatchErrorType:          NoMatchStatus,
}

type Status string

const (
	UnauthorizedStatus Status = "unauthorized"
	DeniedStatus       Status = "denied"
	NoMatchStatus      Status = "noMatch"
	ErrorStatus        Status = "error"
	InvalidStatus      Status = "invalid"
	SuccessStatus      Status = "success"
)

func GetStatus(et *types.ReplyType) Status {
	if et == nil {
		return SuccessStatus
	}
	return statusMap[*et]
}
