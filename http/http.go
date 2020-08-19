package http

import (
	"bitbucket.org/HeilaSystems/serviceerror/errortypes"
	"net/http"
)

var grpcErrors = map[errortypes.ErrorType]int{
	errortypes.UnauthorizedError: http.StatusUnauthorized,
	errortypes.Forbidden:  http.StatusForbidden,
	errortypes.BadRequestError:  http.StatusBadRequest,
	errortypes.DbError: http.StatusInternalServerError,
	errortypes.InternalServiceError:  http.StatusInternalServerError,
	errortypes.IoError: http.StatusInternalServerError,
	errortypes.LogicError: http.StatusBadRequest,
	errortypes.NetworkError: http.StatusBadGateway,
	errortypes.NoContent: http.StatusNoContent,
	errortypes.ValidationError: http.StatusUnprocessableEntity,
}

func GetHttpCode(et errortypes.ErrorType) int{
	return grpcErrors[et]
}