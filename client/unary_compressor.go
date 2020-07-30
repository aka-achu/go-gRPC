package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func SumWithCompressor(c operation_pb.OperationServiceClient) {
	if response, err := c.Sum(
		context.Background(),
		&operation_pb.SumRequest{
			FirstNumber:  10,
			SecondNumber: 20,
		},
		grpc.UseCompressor(gzip.Name),
	); err != nil {
		fatalLogger("Failed to make the Sum RPC request- %v", err)
	} else {
		printer("Response from server- Sum %d", response.SumResult)
	}
}