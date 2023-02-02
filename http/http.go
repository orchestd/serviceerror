package http

import (
	"github.com/orchestd/serviceerror/types"
	"net/http"
)

var httpErrors = map[types.ReplyType]int{
	types.UnauthorizedErrorType:      http.StatusUnauthorized,
	types.LogicUnauthorizedErrorType: http.StatusOK,
	types.ForbiddenErrorType:         http.StatusForbidden,
	types.BadRequestErrorType:        http.StatusBadRequest,
	types.DbErrorType:                http.StatusInternalServerError,
	types.InternalServiceErrorType:   http.StatusInternalServerError,
	types.IoErrorType:                http.StatusInternalServerError,
	types.NetworkErrorType:           http.StatusBadGateway,
	types.NoContentErrorType:         http.StatusNoContent,
	types.ValidationErrorType:        http.StatusUnprocessableEntity,
	types.MethodNotAllowedErrorType:  http.StatusMethodNotAllowed,
	types.NoMatchErrorType:           http.StatusOK,
}

func GetHttpCode(et *types.ReplyType) int {
	if et == nil {
		return http.StatusOK
	}
	return httpErrors[*et]
}
