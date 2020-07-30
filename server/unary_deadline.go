package server

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"time"
)

func (*Service) Power(
	ctx context.Context,
	request *operation_pb.PowerRequest,
) (
	*operation_pb.PowerResponse,
	error,
) {
	time.Sleep(2 * time.Second)
	if ctx.Err() == context.DeadlineExceeded {
		logger("Deadline exceeds for the request. -%v", ctx.Err())
		return nil, status.Error(codes.DeadlineExceeded, "Deadline exceeds for the request")
	} else {
		return &operation_pb.PowerResponse{
			Result: math.Pow(request.GetBase(), request.GetExponent()),
		}, nil
	}
}
