package grpc

import (
	"bitbucket.org/HeilaSystems/serviceerror/errortypes"
	"google.golang.org/grpc/codes"
)

var grpcErrors = map[errortypes.ErrorType]codes.Code{
	errortypes.UnauthorizedError: codes.Unauthenticated,
	errortypes.Forbidden: codes.PermissionDenied,
	errortypes.BadRequestError: codes.Internal,
	errortypes.DbError: codes.Unavailable,
	errortypes.InternalServiceError: codes.Unavailable,
	errortypes.IoError: codes.Unavailable,
	errortypes.LogicError: codes.Internal,
	errortypes.NetworkError: codes.Unavailable,
	errortypes.NoContent: codes.Unknown,
	errortypes.ValidationError: codes.InvalidArgument,
	errortypes.MethodNotAllowed: codes.PermissionDenied,
}

func GetGrpcCode(et errortypes.ErrorType) codes.Code{
	return grpcErrors[et]
}
