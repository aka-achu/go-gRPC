package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"io"
	"log"
)

func PrimeFactors(c operation_pb.OperationServiceClient) {
	if stream, err := c.PrimeFactors(
		context.Background(),
		&operation_pb.PrimeFactorsRequest{
			Number: 1234567890,
	}); err != nil {
		log.Fatalf("Failed to make the request from prime factor service. -%v", err)
	} else {
		for {
			if response, err := stream.Recv(); err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Failed to receive data from the stream. -%v", err)
			} else {
				log.Printf("Response from server- PrimeFactor- %d", response.GetNumber())
			}
		}
	}
}