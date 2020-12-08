package grpc

import (
	"bitbucket.org/HeilaSystems/serviceerror/errortypes"
	"google.golang.org/grpc/codes"
)

var grpcErrors = map[types.ErrorType]codes.Code{
	types.UnauthorizedErrorType:     codes.Unauthenticated,
	types.ForbiddenErrorType:        codes.PermissionDenied,
	types.BadRequestErrorType:       codes.Internal,
	types.DbErrorType:               codes.Unavailable,
	types.InternalServiceErrorType:  codes.Unavailable,
	types.IoErrorType:               codes.Unavailable,
	types.LogicErrorType:            codes.OK,
	types.NetworkErrorType:          codes.Unavailable,
	types.NoContentErrorType:        codes.OK,
	types.ValidationErrorType:       codes.InvalidArgument,
	types.MethodNotAllowedErrorType: codes.PermissionDenied,
}

func GetGrpcCode(et types.ErrorType) codes.Code{
	return grpcErrors[et]
}
