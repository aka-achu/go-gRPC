package server

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
)

func (*Service) Sum(
	ctx context.Context,
	request *operation_pb.SumRequest,
) (
	*operation_pb.SumResponse,
	error,
) {
	return &operation_pb.SumResponse{
		SumResult: request.GetFirstNumber() + request.GetSecondNumber(),
	}, nil
}
