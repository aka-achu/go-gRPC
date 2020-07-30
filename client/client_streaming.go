package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
)

func ComputeAverage(c operation_pb.OperationServiceClient) {
	if stream, err := c.ComputeAverage(context.Background()); err != nil {
		fatalLogger("Failed to open a stream connection. -%v", err)
	} else {
		for _, number := range []int64{1, 2, 3, 4, 5, 6, 7, 8, 9} {
			if err := stream.Send(&operation_pb.ComputeAverageRequest{
				Number: number,
			}); err != nil {
				fatalLogger("Failed to send a number in the stream connection. -%v", err)
			}
		}
		if response, err := stream.CloseAndRecv(); err != nil {
			fatalLogger("Failed to receive average value from the server. -%v", err)
		} else {
			printer("Response from the server- Average %f", response.GetAverage())
		}
	}
}
