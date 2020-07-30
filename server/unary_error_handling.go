package server

import (
	"context"
	"fmt"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
)

func (*OperationService) SquareRoot(
	ctx context.Context,
	request *operation_pb.SquareRootRequest,
) (
	*operation_pb.SquareRootResponse,
	error,
) {
	number := request.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalide argument. Trying to find square root of a negative number"),
			)
	}
	return &operation_pb.SquareRootResponse{
		Root: math.Sqrt(number),
	}, nil
}
