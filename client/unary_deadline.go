package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func Power(c operation_pb.OperationServiceClient, deadline time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()

	if response, err := c.Power(
		ctx,
		&operation_pb.PowerRequest{
			Base:     10,
			Exponent: 3,
		}); err != nil {
		if responseError, stat := status.FromError(err); stat {
			if responseError.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded for the request")
			}
		} else {
			log.Fatalf("Failed to make the request to the server. -%v", err)
		}
	} else {
		log.Printf("Response from server. Power-%f", response.GetResult())
	}

	defer cancel()
}