package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"io"
)

func PrimeFactors(c operation_pb.OperationServiceClient) {
	if stream, err := c.PrimeFactors(
		context.Background(),
		&operation_pb.PrimeFactorsRequest{
			Number: 120,
		}); err != nil {
		fatalLogger("Failed to make the request from prime factor service. -%v", err)
	} else {
		for {
			if response, err := stream.Recv(); err == io.EOF {
				break
			} else if err != nil {
				fatalLogger("Failed to receive data from the stream. -%v", err)
			} else {
				printer("Response from server- PrimeFactor- %d", response.GetNumber())
			}
		}
	}
}
