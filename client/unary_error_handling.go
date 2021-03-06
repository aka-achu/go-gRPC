package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SquareRoot(c operation_pb.OperationServiceClient) {
	for _, number := range []float64{100, -100} {
		if response, err := c.SquareRoot(
			context.Background(),
			&operation_pb.SquareRootRequest{
				Number: number,
			},
		); err != nil {
			if responseError, stat := status.FromError(err); stat {
				if responseError.Code() == codes.InvalidArgument {
					logger("Invalid argument to the squre root service. Negative number. -%v", err)
				}
			} else {
				fatalLogger("Failed to make the request to the server. -%v", err)
			}
		} else {
			printer("Response from server. Square root-%f", response.GetRoot())
		}
	}
}
