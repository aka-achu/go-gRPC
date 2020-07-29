package client

import (
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc"
	"log"
)

func Initialize() {
	if clientConnection, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure()); err != nil {
		log.Fatalf("could not connect: %v", err)
	} else {
		defer clientConnection.Close()
		c := operation_pb.NewOperationServiceClient(clientConnection)
		Sum(c)
		PrimeFactors(c)
		ComputeAverage(c)
		FloorCeiling(c)
		SquareRoot(c)
	}
}
