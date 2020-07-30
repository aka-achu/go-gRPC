package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
)

func Sum(c operation_pb.OperationServiceClient) {
	if response, err := c.Sum(
		context.Background(),
		&operation_pb.SumRequest{
			FirstNumber:  10,
			SecondNumber: 20,
		}); err != nil {
		fatalLogger("Failed to make the Sum RPC request- %v", err)
	} else {
		printer("Response from server- Sum %d", response.SumResult)
	}
}
